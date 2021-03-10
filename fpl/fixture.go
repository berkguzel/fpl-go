package fpl

const (
	fixtureAddress = "https://fantasy.premierleague.com/api/fixtures/"
)

func (c *Client) GetFixture() (*Fixture, error) {

	fixture := &Fixture{}

	response, err := c.NewRequest("GET", fixtureAddress)
	if err != nil{
		return nil, err
	}
	
	_, err = c.Do(response, fixture)
	if err != nil{
		return nil, err
	}
	
	return fixture, nil
}
