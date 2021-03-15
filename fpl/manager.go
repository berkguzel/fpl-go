package fpl

import (
	"encoding/json"
)

// information about manager
func (c *Client) Manager(managerID string) (*Manager, error) {

	url := "https://fantasy.premierleague.com/api/entry/" + managerID + "/"

	manager := &Manager{}

	response, err := c.NewRequest("GET", url)
	if err != nil {
		return nil, err
	}

	_, err = c.Do(response, manager)
	if err != nil {
		return nil, err
	}

	return manager, nil
}

// information about classic leagues the manager joined
func (c *Client) LeagueClassic(managerID string) ([]ManagerLeaguesClassic, error) {

	manager, err := c.Manager(managerID)
	if err != nil {
		return nil, err
	}

	b, err := json.Marshal(manager.Leagues.Classic)
	if err != nil {
		return nil, err
	}

	var classic []ManagerLeaguesClassic
	err = json.Unmarshal(b, &classic)
	if err != nil {
		return nil, err
	}

	return classic, nil
}

// information about cups the manager joined
func (c *Client) LeagueCup(managerID string) ([]ManagerLeaguesCup, error) {

	manager, err := c.Manager(managerID)
	if err != nil {
		return nil, err
	}

	b, err := json.Marshal(manager.Leagues.Cup)
	if err != nil {
		return nil, err
	}

	cup := &ManagerLeaguesCup{}
	err = json.Unmarshal(b, &cup)
	if err != nil {
		return nil, err
	}

	var leagueCup []ManagerLeaguesCup
	leagueCup = append(leagueCup, *cup)

	return leagueCup, nil
}
