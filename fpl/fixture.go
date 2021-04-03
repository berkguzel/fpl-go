package fpl

import (
	"fmt"
	"strconv"
)
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

func (c *Client) GetWeeklyFixture(gameWeek int) (*WeeklyFixture, error){

	url := fmt.Sprintf(fixtureAddress +"?event=" + strconv.Itoa(gameWeek) )

	wf := &WeeklyFixture{}

	response, err := c.NewRequest("GET", url)
	if err != nil {
		return nil, err
	}

	_, err = c.Do(response, wf)
	if err != nil {
		return nil, err
	}

	return wf, nil
}