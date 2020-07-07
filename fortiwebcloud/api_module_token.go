package fortiwebcloud

import (
	"bytes"
	"encoding/json"
	"log"
)

type Token struct {
	UserName string `json:"username,omitempty"`
	PassWord string `json:"password,omitempty"`
	Group    string `json:"group,omitempty"`
}
type TokenClient struct {
	d *Token
	r *Request
}

func NewTokenClient(c *CloudWafClient, d interface{}) *TokenClient {

	var tokenClient TokenClient
	data := d.(*Token)
	locJSON, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	bytes := bytes.NewBuffer(locJSON)
	request := NewRequest(c, "POST", "token", nil, bytes)

	tokenClient.d = data
	tokenClient.r = request
	return &tokenClient
}
func (tc *TokenClient) Send() error {
	return tc.r.Send()
}

func (tc *TokenClient) ReadData() (interface{}, error) {
	body, err := tc.r.ReadData()
	if err != nil {
		return nil, err
	}
	var tokenBody map[string]interface{}
	err = json.Unmarshal(body.([]byte), &tokenBody)

	token := tokenBody["token"]
	return token, err
}
