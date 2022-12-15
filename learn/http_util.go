package main

import (
	"bytes"
	//"github.com/polarismesh/polaris/test/http"
	"net/http"
)

//func main() {
//	//clientHTTP := http.NewClient(httpserverAddress, httpserverVersion)
//	clientHttp := NewClient()
//	clientHttp.SendRequest("Post", "https://www.baidu.com", nil)
//}

func NewClient() *Client { //address, version string
	return &Client{
		//Address: address,
		//Version: version,
		Worker: &http.Client{},
	}
}

// Client HTTP客户端
type Client struct {
	Address string
	Version string
	Worker  *http.Client
}

// SendRequest 发送请求 HTTP Post/Put
func (c *Client) SendRequest(method string, url string, body *bytes.Buffer) (*http.Response, error) {
	var request *http.Request
	var err error
	if body == nil {
		request, err = http.NewRequest(method, url, nil)
	} else {
		request, err = http.NewRequest(method, url, body)
	}
	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Request-Id", "test")
	response, err := c.Worker.Do(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}
