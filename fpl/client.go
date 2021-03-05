package fpl

import (
	"net/http"
	"net/url"
	"fmt"
	"io/ioutil"

)

type Client struct {
	BaseURL *url.URL
	UserAgent string 
	httpClient http.Client 
	response string
}


func (c *Client)DoNewRequest(method string, url string) (*http.Request, error){

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("User-Agent", "")

	return req, nil
}

func (c *Client) Request(req *http.Request) ([]byte, error) {

	client := &c.httpClient
	resp, err := client.Do(req)
	if err != nil{
		return nil, err
	}

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		return nil, err
	}
	defer resp.Body.Close()
	b := []byte(response)

	return b, nil
}

