package fpl

import (
	"net/http"
	"net/url"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

type Client struct {
	BaseURL *url.URL
	UserAgent string 
	httpClient http.Client 
	response string
}


func DoNewRequest(method string, url string) (*http.Request, error){

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("User-Agent", "")

	return req, nil
	
}

func (c *Client) Request(resp *http.Response) (string, error){

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		fmt.Println(err)
	}
	return string(response), nil

}

func UnmarshallJSON(b []byte,v interface{}) (interface{}, error) {
	
	//data := make(map[string]interface{})


	if err := json.Unmarshal(b, v); err != nil {
		return nil, err
	}
	//data["a"] = v

	return v, nil
}
/*
func main(){

	
	req, err := http.NewRequest("GET", "https://fantasy.premierleague.com/api/entry/7270639/history/", nil)
	if err != nil{
		fmt.Println(err)
	}

	req.Header.Add("User-Agent", "")

	var client = &http.Client{
		Timeout: time.Second * 10,
	}

	resp, err := client.Do(req)
	if err != nil{
		fmt.Println(err)
	}

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		fmt.Println(err)
	}

	fmt.Println(string(response))
	
	Get()

}
*/