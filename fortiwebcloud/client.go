package fortiwebcloud

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Config struct {
	HostName string
	UserName string
	PassWord string
	Token    string
}

type FortiwebCloudClient struct {
	//to sdk client
	CloudClient *CloudWafClient
}

func (c *Config) CreateClient() (interface{}, error) {
	var client FortiwebCloudClient

	if err := c.CheckConfig(); err != nil {
		return nil, fmt.Errorf("input param check failed, for %v", err)
	}

	createTLSClient(&client, c)
	return &client, nil
}

func (c *Config) CheckConfig() error {
	if c.HostName == "" {
		return errors.New("empty key fields")
	}
	if c.Token == "" {
		if c.UserName == "" || c.PassWord == "" {
			return errors.New("invalid authorization fields")
		}
	}
	return nil
}

func createTLSClient(client *FortiwebCloudClient, c *Config) error {
	config := &tls.Config{InsecureSkipVerify: true}

	tr := &http.Transport{
		TLSClientConfig: config,
	}

	httpc := &http.Client{
		Transport: tr,
		Timeout:   time.Second * 600,
	}
	client.CloudClient = NewClient(c, httpc)

	sessionEnvString := getEnvSession()
	if sessionEnvString != "" {
		client.CloudClient.SetToken(sessionEnvString)
		return nil
	}
	OutPut("start login:" + sessionEnvString)
	token := client.CloudClient.Login()
	if token == "" {
		return fmt.Errorf("FortiwebCloud Login failed")
	}
	setEnvSession(token)
	return nil
}
func getEnvSession() (session string) {

	sessionString := os.Getenv("FORTIWEB_CLOUD_SESSION")
	sessionTimeStamp := os.Getenv("FORTIWEB_CLOUD_SESSION_TIMESTAMP")
	if sessionString == "" || sessionTimeStamp == "" {

		return ""
	}

	lasttime, err := time.Parse(time.RFC3339, sessionTimeStamp)
	if err != nil {

		return ""
	}

	now := time.Now()
	subM := now.Sub(lasttime)

	if subM.Minutes() > 60 {

		return ""
	}

	return sessionString
}

func setEnvSession(session string) {
	now := time.Now()
	os.Setenv("FORTIWEB_CLOUD_SESSION", session)
	os.Setenv("FORTIWEB_CLOUD_SESSION_TIMESTAMP", now.Format(time.RFC3339))

}
