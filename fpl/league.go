package fpl

import (
	"fmt"
	"errors"
	"encoding/json"
	//ms "github.com/mitchellh/mapstructure" 

)

const (
	leagueAddres = "https://fantasy.premierleague.com/api/leagues-classic/1859418/standings/"
)
func (c *Client) ListStandings()(*LeagueInfo, error){

	req, err := c.DoNewRequest("GET", leagueAddres)
	if err != nil{
		fmt.Println(err)
	}
	
	response, respErr := c.Request(req)
	if respErr != nil{
		return nil, respErr
	}

	league := &LeagueInfo{}
	json.Unmarshal(response, &league)
	standings := league.Standings.Results

	var leaguex LeagueResponse
	hadi, _ := json.Marshal(standings[1])
	fmt.Println(string(hadi))
	json.Unmarshal(hadi, &leaguex)
	fmt.Println(leaguex)
	
	


	return nil, errors.New("Could not find league info")
}