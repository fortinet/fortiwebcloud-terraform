package fortiwebcloud

import (
	"bytes"
	"encoding/json"
	"log"
)

type CustomPort struct {
	Http  int `json:"http,omitempty"`
	Https int `json:"https,omitempty"`
}
type AppCreate struct {
	AppName    string `json:"app_name"`
	DomainName string `json:"domain_name"`

	ExtraDomains   []string   `json:"extra_domains"`
	BlockMode      int        `json:"block_mode"`
	ServerAddress  string     `json:"server_address"`
	CustomPort     CustomPort `json:"custom_port"`
	Service        [2]string  `json:"service"`
	TemplateEnable int        `json:"template_enable,omitempty"`
	Availability   float64    `json:"head_availability"`
	StatusCode     float64    `json:"head_status_code"`
	CDNStatus      int        `json:"cdn_status"`
	ServerCountry  string     `json:"server_country"`
	Region         string     `json:"region"`
	TemplateId     string     `json:"template_id,omitempty"`
	ServerType     string     `json:"server_type"`
	Platform       string     `json:"platform"`
	IsGlobaCdn     int        `json:"is_global_cdn"`
	Continent      string     `json:"continent,omitempty"`
}
type AppCreateChk struct {
	EPId string
}
type AppQuery struct {
	AppName string
}
type AppDel struct {
	EPId string
}
type AppCreateClient struct {
	d *AppCreate
	r *Request
}
type AppCreateChkClient struct {
	d *AppCreateChk
	r *Request
}
type AppQueryClient struct {
	d *AppQuery
	r *Request
}
type AppDelClient struct {
	d *AppDel
	r *Request
}

func NewAppCreateClient(c *CloudWafClient, d interface{}) *AppCreateClient {

	var app AppCreateClient

	data := d.(*AppCreate)
	locJSON, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	bytes := bytes.NewBuffer(locJSON)
	request := NewRequest(c, "POST", "application", nil, bytes)

	app.d = data
	app.r = request
	return &app
}
func (cc *AppCreateClient) Send() error {
	return cc.r.Send()
}

func (cc *AppCreateClient) ReadData() (interface{}, error) {
	body, err := cc.r.ReadData()
	if err != nil {
		return nil, err
	}
	var mbody map[string]interface{}
	err = json.Unmarshal(body.([]byte), &mbody)
	if err == nil {
		return mbody["domain_info"], err
	}

	return mbody, err
}

func NewAppQueryClient(c *CloudWafClient, d interface{}) *AppQueryClient {
	var app AppQueryClient

	data := d.(*AppQuery)

	request := NewRequest(c, "GET", "application", nil, nil)
	request.FillUrlParams("partial", "true")

	app.d = data
	app.r = request
	return &app
}
func (qc *AppQueryClient) Send() error {
	return qc.r.Send()
}

func (qc *AppQueryClient) ReadData() (interface{}, error) {
	body, err := qc.r.ReadData()
	if err != nil {
		return nil, err
	}

	var mbody []interface{}
	err = json.Unmarshal(body.([]byte), &mbody)

	for _, e := range mbody {
		eMap := e.(map[string]interface{})
		output("elem app_name:" + eMap["app_name"].(string) + "<>" + qc.d.AppName)
		if eMap["app_name"].(string) == qc.d.AppName {
			return e, nil
		}
	}
	return nil, err
}

func NewAppDelClient(c *CloudWafClient, d interface{}) *AppDelClient {
	var app AppDelClient

	data := d.(*AppDel)
	output("the ep id is >>:" + data.EPId)
	request := NewRequest(c, "DELETE", "application/"+data.EPId, nil, nil)

	app.d = data
	app.r = request
	return &app
}
func (qc *AppDelClient) Send() error {
	return qc.r.Send()
}

func (qc *AppDelClient) ReadData() (interface{}, error) {
	_, err := qc.r.ReadData()
	if err != nil {
		return nil, err
	}
	return nil, nil
}
