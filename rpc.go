package watcher

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type rpcClient struct {
	addr string
	c    *http.Client
}

func (c *rpcClient) url(endpoint string) *url.URL {
	return &url.URL{
		Scheme: "http",
		Host:   c.addr,
		Path:   endpoint,
	}
}

func (c *rpcClient) Get(endpoint string, output interface{}) error {
	resp, err := c.c.Get(c.url(endpoint).String())
	if err != nil {
		return err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, output)
}
