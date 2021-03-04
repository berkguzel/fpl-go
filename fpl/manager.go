package fpl

import (
	"encoding/json"
	"fmt"
)

const (
	managerAddress = "https://fantasy.premierleague.com/api/entry/7270639/"
)
func (c *Client) Manager()(*Manager, error) {
	req, err := c.DoNewRequest("GET", managerAddress)
	if err != nil{
		fmt.Println(err)
	}
	
	response, respErr := c.Request(req)
	if respErr != nil{
		return nil, respErr
	}


	manager := &Manager{}
	err = json.Unmarshal(response, &manager)
	if err != nil{
		return nil, err
	}

	return manager, nil
}