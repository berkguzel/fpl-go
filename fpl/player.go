package fpl

import (
	"encoding/json"
	"errors"
	"strconv"
)

// all information of football player
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

// football player's informations for next fixtures
func (c *Client) ListPlayerFixture(playerID int) ([]PlayerFixture, error) {

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

// football player's informations for previous fixtures
func (c *Client) ListPlayerHistory(playerID int) ([]PlayerHistory, error) {

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

// football player's informations for previous seasons
func (c *Client) ListPlayerHistoryPast(playerID int) ([]PlayerHistoryPast, error) {

	player, err := c.GetPlayer(playerID)
	if err != nil {
		return nil, err
	}

	pf, err := json.Marshal(player.HistoryPast)
	if err != nil {
		return nil, err
	}

	var playerHistoryPast []PlayerHistoryPast
	if err := json.Unmarshal(pf, &playerHistoryPast); err != nil {
		return nil, err
	}

	return playerHistoryPast, nil
}

// learn code of player
func (c *Client) GetCodeOfPlayer(name string) (int, error) {

	players, err := c.GetInfoOfPlayers()
	if err != nil {
		return 0, err
	}

	for _, v := range players {
		if v.FirstName == name {
			return v.ID, nil
		}
	}

	return 0, errors.New("Could not find footballer")

}

// detailed informations of footballers
func (c *Client) GetInfoOfPlayers() ([]PlayerDetailedInfo, error) {

	general, _ := c.GetGeneral()

	var e []PlayerDetailedInfo

	m, err := json.Marshal(general.Elements)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(m, &e); err != nil {
		return nil, err
	}

	return e, nil

}
