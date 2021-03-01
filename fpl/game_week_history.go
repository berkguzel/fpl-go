package fpl

import (
	"fmt"
	"encoding/json"
	
	"github.com/fatih/structs"
)


func (c *Client) Get(gameWeek int) (*WeeklyResponse, error){

	req, err := DoNewRequest("GET", "https://fantasy.premierleague.com/api/entry/7270639/history/")
	if err != nil{
		fmt.Println(err)
	}
	
	client := &c.httpClient
	resp, err := client.Do(req)
	if err != nil{
		return nil, err
	}

	x, respErr := c.Request(resp)
	if err != nil{
		fmt.Println(respErr)
	}

	v := &Weekly{}
	b := []byte(x)
	

	_, JSONErr := UnmarshallJSON(b, &v)
	if JSONErr != nil{
		return nil, JSONErr
	}

	curr := v.Current[1]
	m, _ := json.Marshal(curr)
	wee := &WeeklyResponse{}
	_ = json.Unmarshal(m, &wee)
	

	
	return wee,   nil
}


func (c *Client) ListWeeklyPoints() map[string]interface{}{
	
	weeklyResponse, _ := c.Get(5)
	m := structs.Map(weeklyResponse)

	return m
}