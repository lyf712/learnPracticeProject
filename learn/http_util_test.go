package main

import (
	"testing"
)

func TestHttp(t *testing.T) {
	clientHttp := NewClient()
	println(clientHttp.SendRequest("Post", "https://www.baidu.com", nil))
}
