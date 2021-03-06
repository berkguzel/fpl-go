package fpl

import (
	"fmt"
)

const (
	fixtureAddress = "https://fantasy.premierleague.com/api/fixtures/"
)

func(c *Client) GetFixture()(*Fixture, error){
	req, err := c.DoNewRequest("GET", fixtureAddress)
	if err != nil{
		fmt.Println(err)
	}
	
	response, respErr := c.Request(req)
	if respErr != nil{
		return nil, respErr
	}

	fmt.Println(response)

	return nil, nil
}