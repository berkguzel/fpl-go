package fpl

import (
	"encoding/json"
)

func (c *Client) ListTransfers() ([]TransferHistory, error) {

	managerID, _ := Get()
	url := "https://fantasy.premierleague.com/api/entry/" + managerID + "/transfers/"

	response, err := c.Do("GET", url)
	if err != nil {
		return nil, err
	}

	var transfer []TransferHistory

	err = json.Unmarshal(response, &transfer)
	if err != nil {
		return nil, err
	}

	return transfer, nil
}
