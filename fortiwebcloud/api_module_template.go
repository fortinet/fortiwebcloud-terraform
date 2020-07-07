package fortiwebcloud

import (
	"encoding/json"
	"errors"
)

type Template struct {
	TemplateName string
}

type TemplateClient struct {
	d *Template
	r *Request
}

func NewTemplateClient(c *CloudWafClient, d interface{}) *TemplateClient {

	var tpl TemplateClient

	data := d.(*Template)

	request := NewRequest(c, "GET", "template", nil, nil)

	tpl.d = data
	tpl.r = request
	return &tpl
}
func (tc *TemplateClient) Send() error {
	return tc.r.Send()
}

func (tc *TemplateClient) ReadData() (interface{}, error) {
	body, err := tc.r.ReadData()
	if err != nil {
		return nil, err
	}
	var mbody map[string]interface{}
	err = json.Unmarshal(body.([]byte), &mbody)

	result, ok := mbody["result"]
	tlist, succ := result.([]interface{})

	if !ok || !succ {
		return nil, errors.New("couldn't find the template")
	}
	for _, e := range tlist {
		eMap, _ := e.(map[string]interface{})
		if eMap["name"].(string) == tc.d.TemplateName {
			return eMap, nil
		}
	}
	return nil, errors.New("Not found the template by name")

}
