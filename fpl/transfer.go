package fpl


func (c *Client) ListTransfers() ([]TransferHistory, error) {

	managerID, _ := Get()
	url := "https://fantasy.premierleague.com/api/entry/" + managerID + "/transfers/"

	response, err := c.NewRequest("GET", url)
	if err != nil{
		return nil, err
	}
	
	var transfer []TransferHistory
	
	_, err = c.Do(response, &transfer)
	if err != nil{
		return nil, err
	}

	return transfer, nil
}
