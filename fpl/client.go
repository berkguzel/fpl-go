package fpl

import (
	"net/http"
	"net/url"
	"fmt"
	"io/ioutil"
	"log"


)

type Client struct {
	BaseURL *url.URL
	UserAgent string 
	httpClient http.Client 
}



func (c *Client) Do(method string, url string)([]byte, error){

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("User-Agent", "")

	client := &c.httpClient
	resp, err := client.Do(req)
	if err != nil{
		log.Fatal(err)
	}

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		log.Fatal(err)
	}
	defer resp.Body.Close()
	b := []byte(response)

	return b, nil
}