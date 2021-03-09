package fpl

import (
	"encoding/json"

)

func (c *Client) ListStandings()([]LeagueResponse, error){

	_, leagueID := Get()

	url := "https://fantasy.premierleague.com/api/leagues-classic/" + leagueID + "/standings/"
	

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

	return league, nil
}