package fpl

import (
	"net/http"
	"encoding/json"
)

type Client struct {
	httpClient *http.Client
	UserAgent string
}

func (c *Client) NewRequest(method string, path string) (*http.Request, error) {

    req, err := http.NewRequest(method, path, nil)
    if err != nil {
        return nil, err
    }

    req.Header.Set("User-Agent", c.UserAgent)
    return req, nil
}

func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
    resp, err := c.httpClient.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    err = json.NewDecoder(resp.Body).Decode(v)
    return resp, err
}

func NewClient(httpClient *http.Client) *Client{
	if httpClient == nil {
		httpClient = &http.Client{}
	}

	c := &Client{
		httpClient : httpClient,

	}

	return c

}