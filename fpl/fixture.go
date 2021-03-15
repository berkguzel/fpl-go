package fpl

const (
	fixtureAddress = "https://fantasy.premierleague.com/api/fixtures/"
)

// returns fixture endpoint
func (c *Client) GetFixture() ([]Fixture, error) {

	f := &Fixture{}

	response, err := c.NewRequest("GET", fixtureAddress)
	if err != nil {
		return nil, err
	}

	_, err = c.Do(response, f)
	if err != nil {
		return nil, err
	}

	var fixture []Fixture
	fixture = append(fixture, *f)

	return fixture, nil
}
