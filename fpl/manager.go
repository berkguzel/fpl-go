package fpl

import (
	"encoding/json"
)


func (c *Client) Manager() (*Manager, error){
	
	manager := &Manager{}
	managerID, _ := Get()
	url := "https://fantasy.premierleague.com/api/entry/" + managerID + "/"

	res, _ := c.Do("GET", url)

	err := json.Unmarshal(res, manager)
	if err != nil{
		return nil, err
	}
	
	return manager, nil
}

func (c *Client) LeagueClassic() ([]ManagerLeaguesClassic, error){
	
	manager, err := c.Manager()
	if err != nil{
		return nil, err
	}
	
	b, err := json.Marshal(manager.Leagues.Classic)
	if err != nil{
		return nil, err
	}
	
	var classic  []ManagerLeaguesClassic
	err = json.Unmarshal(b, &classic)
	if err != nil {
		return nil, err
	}

	return classic, nil
}

func (c *Client) LeagueCup() ([]ManagerLeaguesCup, error){

	manager, err := c.Manager()
	if err != nil{
		return nil, err
	}
	
	b, err := json.Marshal(manager.Leagues.Cup)
	if err != nil{
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