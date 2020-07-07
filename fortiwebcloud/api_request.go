package fortiwebcloud

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
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

	r.HTTPRequest.Header.Set("Content-Type", "application/json")

	r.HTTPRequest.Header.Set("Authorization", r.Client.Token)

	u := buildURL(r)

	var err error
	r.HTTPRequest.URL, err = url.Parse(u)
	output("send data to url>>:" + r.HTTPRequest.URL.String())
	if err != nil {
		log.Fatal(err)
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
	path := path.Join("v1", r.Path)
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
