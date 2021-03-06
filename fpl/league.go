package fpl

import (
	"fmt"
	"errors"
	"encoding/json"

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
	errMars := json.Unmarshal(response, &leagueInfo)
	if errMars != nil{
		fmt.Println(errMars)
	}
	standings := leagueInfo.Standings.Results

	league := []LeagueResponse{}
	var l LeagueResponse
	
	for _, v := range standings {
		resp, _ := json.Marshal(v)
		errMars = json.Unmarshal(resp, &l)
		if errMars != nil{
			fmt.Println(errMars)
		}
		league = append(league,l)

	}

	return league, errors.New("Could not find league info")
}