package fpl

import (
	"encoding/json"
)

const (
	fixtureAddress = "https://fantasy.premierleague.com/api/fixtures/"
)

func(c *Client) GetFixture()(*Fixture, error){
	
	response, _ := c.Do("GET", fixtureAddress)
	fixture := &Fixture{}
	err := json.Unmarshal(response, &fixture)
	if err != nil{
		return nil, err
	}
	return fixture, nil
}