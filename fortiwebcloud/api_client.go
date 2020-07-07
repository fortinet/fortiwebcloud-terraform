package fortiwebcloud

import (
	"log"
	"net/http"
)

type CloudWafClient struct { // FmgSDKClient
	Host    string
	User    string
	Passwd  string
	Debug   string
	HTTPCon *http.Client //Client
	Init    bool
	Token   string
}

// NewClient is for creating new client
// Input:
//   @ip: ipaddress of fortiwebcloud
//   @user: username of fortiwebcloud
//   @passwd: passwd of fortiwebcloud
//   @client: http client used
// Output:
//   @CloudWafClient: client
func NewClient(c *Config, client *http.Client) *CloudWafClient {

	if c.Token != "" {
		token := "Basic " + c.Token

		return &CloudWafClient{
			Host:    c.HostName,
			HTTPCon: client,
			Debug:   "OFF",
			Init:    true,
			Token:   token,
		}
	} else {

	}
	return &CloudWafClient{
		Host:    c.HostName,
		User:    c.UserName,
		Passwd:  c.PassWord,
		HTTPCon: client,
		Debug:   "OFF",
		Init:    false,
		Token:   "",
	}
}

// Login is for logging in
// Output:
//   @session: login session
//   @err: error details if failure, and nil if success
func (c *CloudWafClient) SetToken(token string) {
	c.Token = token
	c.Init = true
}
func (c *CloudWafClient) Login() string {
	if !c.Init {
		token := Token{UserName: c.User, PassWord: c.Passwd}
		tokenClient := NewTokenClient(c, &token)
		tokenClient.Send()
		strToken, err := tokenClient.ReadData()
		if err == nil {
			c.Init = true
			c.Token = strToken.(string)
			return c.Token
		}
		return ""
	}
	return c.Token
}

// Logout is for logging out
// Input:
//   @s: login session
// Output:
//   @err: error details if failure, and nil if success
func (c *CloudWafClient) Logout(s string) (err error) {

	return nil
	// should be logout
}

// Trace is for debugging
// Input:
//   @s: function name to be traced
// Output:
//   @func()
func (f *CloudWafClient) Trace(s string) func() {
	if f.Debug == "ON" {
		log.Printf("[TRACEDEBUG] -> Enter %s <-", s)
		return func() { log.Printf("[TRACEDEBUG]    -> Leave %s <-", s) }
	}

	return func() {}
}

func (c *CloudWafClient) SetInit() {
	c.Init = true
}
