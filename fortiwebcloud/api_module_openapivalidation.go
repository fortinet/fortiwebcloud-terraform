package fortiwebcloud

import (
	"bytes"
	"encoding/json"
	"log"
)

type OpenapiValidationCreate struct {
	EPId                   string
	Template_status        bool       `json:"template"`
	OpenapiValidationCfg   SchemaFile `json:"configs"`
}

type SchemaFile struct {
	Status                 bool       `json:"status"`
	Action                 string     `json:"action"`
	SchemaFiles            []OFiles   `json:"file_list"`
}

type OFiles struct {
	OpenapiFile string `json:"name"`
	Seq         int    `json:"idx"`
}

type OpenapiValidationQueryClient struct {
	d *OpenapiValidationQuery
	r *Request
}

type OpenapiValidationCreateClient struct {
	d *OpenapiValidationCreate
	r *Request
}

type OpenapiValidationQuery struct {
	EPId string
}

func NewOpenapiValidationCreateClient(c *CloudWafClient, d interface{}, uploadFiles []UploadFiles) (*OpenapiValidationCreateClient, error) {
	var OpenapiValidationCreateClient OpenapiValidationCreateClient
	data := d.(*OpenapiValidationCreate)
	locJSON, err := json.Marshal(data)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	bytes := bytes.NewBuffer(locJSON)
	request, err := NewRequestWithFiles(c, "PUT", "application/"+data.EPId+"/api_protection", nil, bytes, uploadFiles)
	output("URL >>: " + "application/" + data.EPId + "/api_protection")
	if err != nil {
		log.Print(err)
		return nil, err
	}
	OpenapiValidationCreateClient.d = data
	OpenapiValidationCreateClient.r = request
	return &OpenapiValidationCreateClient, nil
}

func (tc *OpenapiValidationCreateClient) Send() error {
	return tc.r.Send()
}

func (tc *OpenapiValidationCreateClient) ReadData() (interface{}, error) {
	body, err := tc.r.ReadData()
	if err != nil {
		log.Print(err)
		return nil, err
	}
	var openapiValidationBody map[string]interface{}
	err = json.Unmarshal(body.([]byte), &openapiValidationBody)
	return openapiValidationBody, err
}

func NewOpenapiValidationQueryClient(c *CloudWafClient, d interface{}) *OpenapiValidationQueryClient {
	var openapiValidationQueryClient OpenapiValidationQueryClient

	data := d.(*OpenapiValidationQuery)
	request := NewRequest(c, "GET", "application/"+data.EPId+"/api_protection", nil, nil)

	openapiValidationQueryClient.d = data
	openapiValidationQueryClient.r = request
	return &openapiValidationQueryClient
}

func (tc *OpenapiValidationQueryClient) Send() error {
	return tc.r.Send()
}

func (tc *OpenapiValidationQueryClient) ReadData() (interface{}, error) {
	body, err := tc.r.ReadData()
	if err != nil {
		return nil, err
	}
	var openapiValidationBody map[string]interface{}
	err = json.Unmarshal(body.([]byte), &openapiValidationBody)
	if err != nil {
		return body, err
	}
	return openapiValidationBody, nil
}
