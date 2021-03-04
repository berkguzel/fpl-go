package fpl

import (
	"fmt"
	"errors"
	"encoding/json"
	//ms "github.com/mitchellh/mapstructure" 

)

const (
	leagueAddress = "https://fantasy.premierleague.com/api/leagues-classic/1859418/standings/"
)

func (c *Client) ListStandings()([]LeagueResponse, error){

	req, err := c.DoNewRequest("GET", leagueAddress)
	if err != nil{
		fmt.Println(err)
	}
	
	response, respErr := c.Request(req)
	if respErr != nil{
		return nil, respErr
	}

	leagueInfo := &LeagueInfo{}
	json.Unmarshal(response, &leagueInfo)
	standings := leagueInfo.Standings.Results

	league := []LeagueResponse{}
	var l LeagueResponse
	
	for _, v := range standings {
		resp, _ := json.Marshal(v)
		json.Unmarshal(resp, &l)
		league = append(league,l)

	}

	return league, errors.New("Could not find league info")
}