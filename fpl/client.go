package fpl

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

type Client struct {
	BaseURL    *url.URL
	Header     string
	httpClient http.Client
}

func (c *Client) Do(method string, url string) ([]byte, error) {

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "")

	client := &c.httpClient
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b := []byte(response)

	return b, nil
}
