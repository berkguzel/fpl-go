package fpl

import (
	"encoding/json"
)

const (
	transferHistoryAddress  = "https://fantasy.premierleague.com/api/entry/7270639/transfers/"
)
func(c *Client) ListTransfers()([]TransferHistory, error) {
	req, err := c.DoNewRequest("GET", transferHistoryAddress)
	if err != nil{
		return nil, err
	}
	
	response, respErr := c.Request(req)
	if respErr != nil{
		return nil, respErr
	}

	var transfer []TransferHistory

	if err := json.Unmarshal(response, &transfer); err != nil {
		return nil, err
    }

	return transfer, nil
}