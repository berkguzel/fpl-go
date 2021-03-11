package fpl

import (
	"encoding/json"
	"strconv"
)

func (c *Client) GetPlayer(playerID int) (*Player, error) {

	id := strconv.Itoa(playerID)
	url := "https://fantasy.premierleague.com/api/element-summary/" + id + "/"

	response, err := c.NewRequest("GET", url)
	if err != nil {
		return nil, err
	}

	player := &Player{}

	_, err = c.Do(response, player)
	if err != nil {
		return nil, err
	}

	return player, nil
}

func (c *Client) PlayerFixture(playerID int) ([]PlayerFixture, error) {

	player, err := c.GetPlayer(playerID)
	if err != nil {
		return nil, err
	}

	pf, err := json.Marshal(player.Fixtures)
	if err != nil {
		return nil, err
	}

	var playerFixture []PlayerFixture
	if err := json.Unmarshal(pf, &playerFixture); err != nil {
		return nil, err
	}

	return playerFixture, nil
}

func (c *Client) PlayerHistory(playerID int) ([]PlayerHistory, error) {

	player, err := c.GetPlayer(playerID)
	if err != nil {
		return nil, err
	}

	pf, err := json.Marshal(player.History)
	if err != nil {
		return nil, err
	}

	var playerHistory []PlayerHistory
	if err := json.Unmarshal(pf, &playerHistory); err != nil {
		return nil, err
	}

	return playerHistory, nil
}

func (c *Client) PlayerHistoryPast(playerID int) ([]PlayerHistory, error) {

	player, err := c.GetPlayer(playerID)
	if err != nil {
		return nil, err
	}

	pf, err := json.Marshal(player.HistoryPast)
	if err != nil {
		return nil, err
	}

	var playerHistory []PlayerHistory
	if err := json.Unmarshal(pf, &playerHistory); err != nil {
		return nil, err
	}

	return playerHistory, nil
}
