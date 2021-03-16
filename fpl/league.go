package fpl

import (
	"encoding/json"
	"errors"
)

// league endpoint according to given league id
func (c *Client) GetLeague(leagueID string) (*LeagueInfo, error) {

	url := "https://fantasy.premierleague.com/api/leagues-classic/" + leagueID + "/standings/"

	response, err := c.NewRequest("GET", url)
	if err != nil {
		return nil, err
	}

	l := &LeagueInfo{}

	_, err = c.Do(response, l)
	if err != nil {
		return nil, err
	}

	return l, nil
}

// standings according to given league id
func (c *Client) GetStandings(leagueID string) ([]StandingsResponse, error) {

	s, err := c.GetLeague(leagueID)
	if err != nil {
		return nil, err
	}

	standings := s.Standings.Results

	var league []StandingsResponse
	l := &StandingsResponse{}

	for _, v := range standings {
		resp, _ := json.Marshal(v)
		if err = json.Unmarshal(resp, &l); err != nil {
			return nil, err
		}
		league = append(league, *l)

	}

	return league, nil

}

// new entries of the league
func (c *Client) GetNewEntries(leagueID string) ([]NewEntriesResponse, error) {

	s, err := c.GetLeague(leagueID)
	if err != nil {
		return nil, err
	}

	newEntries := s.NewEntries

	ent, err := json.Marshal(newEntries)
	if err != nil {
		return nil, err
	}

	e := &NewEntriesResponse{}
	var newEntriesResponse []NewEntriesResponse
	if err := json.Unmarshal(ent, &e); err != nil {
		return nil, err
	}
	newEntriesResponse = append(newEntriesResponse, *e)

	return newEntriesResponse, nil
}

// information about the team according to league
func (c *Client) GetTeamInfoInLeague(leagueID string, teamName string) (*StandingsResponse, error) {

	league, err := c.GetStandings(leagueID)
	if err != nil {
		return nil, err
	}

	for _, v := range league {
		if v.EntryName == teamName {
			return &v, nil
		}
	}

	return nil, errors.New("Could not find team in league")
}
