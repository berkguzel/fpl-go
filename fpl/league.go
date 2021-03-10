package fpl

import (
	"encoding/json"


)


func (c *Client) League() (*LeagueInfo, error) {

	_, leagueID := Get()

	url := "https://fantasy.premierleague.com/api/leagues-classic/" + leagueID + "/standings/"


	response, err := c.NewRequest("GET", url)
	if err != nil{
		return nil, err
	}
	
	l := &LeagueInfo{}

	_, err = c.Do(response, l)
	if err != nil{
		return nil, err
	}

	return l, nil
}

func (c *Client) GetStandings()([]StandingsResponse, error){

	s, err := c.League()
	if err != nil{
		return nil, err
	}

	standings := s.Standings.Results

	var league []StandingsResponse
	l := &StandingsResponse{}

	for _, v := range standings {
		resp, _ := json.Marshal(v)
		if err = json.Unmarshal(resp, &l); err != nil{
			return nil, err
		}
		league = append(league, *l)

	}

	return league, nil

}

func (c *Client) GetNewEntries() ([]NewEntriesResponse, error){

	s, err := c.League()
	if err != nil{
		return nil, err
	}

	newEntries := s.NewEntries

	ent, err := json.Marshal(newEntries)
	if err != nil{
		return nil, err
	}

	e := &NewEntriesResponse{} 
	var newEntriesResponse []NewEntriesResponse
	if err := json.Unmarshal(ent, &e); err !=nil{
		return nil, err
	}
	newEntriesResponse = append(newEntriesResponse, *e)

	return newEntriesResponse, nil
}