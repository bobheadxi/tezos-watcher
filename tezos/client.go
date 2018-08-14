package tezos

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Client is used for requests to the Tezos RPC API
type Client struct {
	addr string
	c    *http.Client
}

// New instantiates a new instance of the Tezos RPC API client and checks if
// given node was bootstrapped
func New(addr string) (*Client, error) {
	c := &Client{addr, http.DefaultClient}
	return c, c.Get("/monitor/bootstrapped", nil)
}

// Get executes a GET request on given endpoint, and unmarshals output into
// the given struct
func (c *Client) Get(endpoint string, output interface{}) error {
	resp, err := c.c.Get(c.url(endpoint).String())
	if err != nil {
		return err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if output == nil {
		return nil
	}
	return json.Unmarshal(b, output)
}

func (c *Client) url(endpoint string) *url.URL {
	return &url.URL{
		Scheme: "http",
		Host:   c.addr,
		Path:   endpoint,
	}
}
