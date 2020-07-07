package fortiwebcloud

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
)

type ServerTest struct {
	BackendIP   string `json:"backend_ip,omitempty"`
	BackendType string `json:"backend_type,omitempty"`
	Domain      string `json:"domain_name,omitempty"`
}

type ServerTestClient struct {
	d *ServerTest
	r *Request
}

type DnsLookup struct {
	Domain string `json:"domain,omitempty"`
}

type DnsLookupClient struct {
	d *DnsLookup
	r *Request
}

type IPRegion struct {
	Domain      string     `json:"domain_name,omitempty"`
	EpIP        string     `json:"ep_ip,omitempty"`
	ExtraDomain []string   `json:"extra_domains,omitempty"`
	CustomPort  CustomPort `json:"custom_port,omitempty"`
}

type IPRegionClient struct {
	d *IPRegion
	r *Request
}

func NewServerTestClient(c *CloudWafClient, d interface{}) *ServerTestClient {

	var server ServerTestClient
	data := d.(*ServerTest)
	request := NewRequest(c, "GET", "misc/backend-ip-test", nil, nil)
	request.FillUrlParams("backend_ip", data.BackendIP)
	request.FillUrlParams("backend_type", data.BackendType)
	request.FillUrlParams("domain_name", data.Domain)
	server.d = data
	server.r = request
	return &server
}
func (sc *ServerTestClient) Send() error {
	return sc.r.Send()
}

func (sc *ServerTestClient) ReadData() (interface{}, error) {
	body, err := sc.r.ReadData()
	if err != nil {
		return nil, err
	}

	var mbody map[string]interface{}
	err = json.Unmarshal(body.([]byte), &mbody)

	status, ok := mbody["network_connectivity"]
	if !ok {
		return nil, errors.New("Response the network connectivity check failed")
	}
	st := status.(float64)
	if st != 1 {
		return nil, errors.New("Response the status check failed")
	}

	return mbody, nil

}

func NewDnsLookupClient(c *CloudWafClient, d interface{}) *DnsLookupClient {

	var dnsClient DnsLookupClient
	data := d.(*DnsLookup)

	locJSON, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	bytes := bytes.NewBuffer(locJSON)

	request := NewRequest(c, "POST", "misc/dns-lookup", nil, bytes)

	dnsClient.d = data
	dnsClient.r = request
	return &dnsClient
}
func (dc *DnsLookupClient) Send() error {
	return dc.r.Send()
}

func (dc *DnsLookupClient) ReadData() (interface{}, error) {
	body, err := dc.r.ReadData()
	if err != nil {
		return nil, err
	}
	var mbody map[string]interface{}
	err = json.Unmarshal(body.([]byte), &mbody)

	address := mbody["A"]
	return address, err

}

func NewIPRegionClient(c *CloudWafClient, d interface{}) *IPRegionClient {

	var regionClient IPRegionClient
	data := d.(*IPRegion)

	locJSON, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	bytes := bytes.NewBuffer(locJSON)

	request := NewRequest(c, "POST", "misc/check-ip-region", nil, bytes)

	regionClient.d = data
	regionClient.r = request
	return &regionClient
}
func (rc *IPRegionClient) Send() error {
	return rc.r.Send()
}

func (rc *IPRegionClient) ReadData() (interface{}, error) {
	body, err := rc.r.ReadData()
	if err != nil {
		return nil, err
	}
	var mbody map[string]interface{}
	err = json.Unmarshal(body.([]byte), &mbody)

	return mbody, err

}
