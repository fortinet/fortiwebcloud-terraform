package fortiwebcloud

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"time"
)

// Describes the request to FortiwebCloud
type Request struct {
	Client       *CloudWafClient
	HTTPRequest  *http.Request
	HTTPResponse *http.Response
	Path         string
	Params       interface{}
	Data         *bytes.Buffer
}

type UploadFiles struct {
	FileName string
	FilePath string
}

func NewRequestWithFiles(c *CloudWafClient, method string, path string, params interface{}, data *bytes.Buffer, mfiles []UploadFiles) (*Request, error) {
	var h *http.Request
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	dd := []byte(data.String())
	var mdata map[string]interface{}
	tojson := json.Unmarshal(dd, &mdata)
	if tojson != nil {
		output(fmt.Sprintf("Error: %s\n", tojson))
	}
	for key, val := range mdata {
		switch val.(type) {
		case map[string]interface{}:
			jsonString, _ := json.Marshal(val)
			_ = writer.WriteField(key, fmt.Sprintf("%s", jsonString))
			break
		default:
			_ = writer.WriteField(key, fmt.Sprintf("%v", val))
			break
		}
	}

	var i int = 0
	for _, mfile := range mfiles {
		i++
		file, err := os.Open(mfile.FilePath)
		defer file.Close()
		if err != nil {
			output(fmt.Sprintf("Error open file %s: %v", mfile.FilePath, err))
			return nil, err
		}
		part, err := writer.CreateFormFile(mfile.FileName, mfile.FilePath)
		if err != nil {
			output(fmt.Sprintf("Error reading file: %v", err))
		}
		_, err = io.Copy(part, file)
	}
	if err := writer.Close(); err != nil {
		output(fmt.Sprintf("Error closing multipart writer: %v", err))
	}
	if data == nil {
		h, _ = http.NewRequest(method, "", nil)
	} else {
		h, _ = http.NewRequest(method, "", body)
		output("send body data>>>:" + body.String())
	}
	h.Header.Set("Content-Type", writer.FormDataContentType())

	r := &Request{
		Client:      c,
		Path:        path,
		HTTPRequest: h,
		Params:      params,
		Data:        data,
	}
	return r, nil
}

// New creates request object with http method, path, params and data,
// It will save the http request, path, etc. for the next operations
// such as sending data, getting response, etc.
// It returns the created request object to the client.
func NewRequest(c *CloudWafClient, method string, path string, params interface{}, data *bytes.Buffer) *Request {
	var h *http.Request

	if data == nil {
		h, _ = http.NewRequest(method, "", nil)
	} else {
		h, _ = http.NewRequest(method, "", data)
		output("send body data>>:" + data.String())
	}
	h.Header.Set("Content-Type", "application/json")

	r := &Request{
		Client:      c,
		Path:        path,
		HTTPRequest: h,
		Params:      params,
		Data:        data,
	}
	return r
}

// Store URL request params

func (r *Request) FillUrlParams(key string, val string) {
	if r.HTTPRequest.Form == nil {
		r.HTTPRequest.Form = make(map[string][]string)
	}
	r.HTTPRequest.Form.Set(key, val)
}

func (r *Request) Send() error {

	r.HTTPRequest.Header.Set("Authorization", r.Client.Token)

	u := buildURL(r)

	var err error
	r.HTTPRequest.URL, err = url.Parse(u)
	output("send data to url>>:" + r.HTTPRequest.URL.String())
	if err != nil {
		log.Print(err)
		return err
	}

	retry := 0
	for {
		//Send

		rsp, errdo := r.Client.HTTPCon.Do(r.HTTPRequest)
		r.HTTPResponse = rsp
		output("response status>>:" + strconv.Itoa(rsp.StatusCode))
		if errdo != nil {

			if retry > 3 {
				err = fmt.Errorf("Error found: %s", errdo)
				return err
			}
			time.Sleep(time.Duration(1) * time.Second)
			log.Printf("Error found: %s, will resend again %s, %d", errdo, u, retry)

			retry++

		} else {

			return errdo
		}
	}

	return nil

}
func (r *Request) ReadData() (interface{}, error) {
	if r.HTTPResponse == nil {
		err := fmt.Errorf("cannot get response")
		return nil, err
	}

	body, err := ioutil.ReadAll(r.HTTPResponse.Body)
	defer r.HTTPResponse.Body.Close()
	output("read body data>>:" + string(body))
	if err != nil || body == nil {
		output("readdata error ####:" + string(err.Error()))
		err = fmt.Errorf("cannot get response body %s", err)
		return nil, err
	}
	if r.HTTPResponse.StatusCode >= 400 {

		return nil, fmt.Errorf("check failed for: %s", string(body))
	}

	return body, nil

}

func buildURL(r *Request) string {
	host := "https://" + r.Client.Host + "/"
	path := path.Join("v2", r.Path)
	u := host + path
	para := ""
	for k, _ := range r.HTTPRequest.Form {
		if para != "" {
			para += "&"
		}
		para += k
		para += "="
		para += r.HTTPRequest.Form.Get(k)

	}
	if para != "" {
		u += "?"
		u += para
	}
	return u
}
