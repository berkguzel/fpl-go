package fpl

import (
	"errors"
	"encoding/json"

)

func (c *Client) ListStandings()([]LeagueResponse, error){

	_, teamID := Get()

	url := "https://fantasy.premierleague.com/api/leagues-classic/" + teamID + "/standings/"
	

	response, _ := c.Do("GET", url)

	leagueInfo := &LeagueInfo{}
	err := json.Unmarshal(response, &leagueInfo)
	if err != nil{
		return nil, err
	}
	standings := leagueInfo.Standings.Results

	league := []LeagueResponse{}
	var l LeagueResponse
	
	for _, v := range standings {
		resp, _ := json.Marshal(v)
		err = json.Unmarshal(resp, &l)
		if err != nil{
			return nil, err
		}
		league = append(league,l)

	}

	return league, errors.New("Could not find league info")
}