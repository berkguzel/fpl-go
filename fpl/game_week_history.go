package fpl

import (
	"encoding/json"
	"errors"
	"strconv"
	
)

func (c *Client) ListWeeklyPoints(gameWeek int) ([]WeeklyInfo, error) {

	managerID, _ := Get()
	url := "https://fantasy.premierleague.com/api/entry/" + managerID + "/history/"

	
	response, err := c.NewRequest("GET", url)
	if err != nil{
		return nil, err
	}
	
	v := &Weekly{}

	_, err = c.Do(response, v)
	if err != nil{
		return nil, err
	}

	var wResponse []WeeklyInfo
	for index, value := range v.Current {

		if value.Event == gameWeek {

			m, err := json.Marshal(v.Current[index])
			if err != nil {
				return nil, err
			}

			w := &WeeklyInfo{}

			if err := json.Unmarshal(m, &w); err != nil {
				return nil, err
			}

			wResponse = append(wResponse, *w)
			return wResponse, nil
		}
	}
	return nil, errors.New("Could not find game week")
}

// List all weeks
func (c *Client) ListAllWeeks() ([]Weekly, error) {

	managerID, _ := Get()
	url := "https://fantasy.premierleague.com/api/entry/" + managerID + "/history/"

	response, err := c.NewRequest("GET", url)
	if err != nil{
		return nil, err
	}
	
	v := &Weekly{}

	_, err = c.Do(response, v)
	if err != nil{
		return nil, err
	}


	var allWeeks []Weekly
	allWeeks = append(allWeeks, *v)

	return allWeeks, nil
}

// individual performances of players weekly
func (c *Client) ListWeeklyPerformance(week int) ([]TeamWeekly, error) {

	managerID, _ := Get()

	url := "https://fantasy.premierleague.com/api/entry/" + managerID + "/event/" + strconv.Itoa(week) + "/picks/"

	response, err := c.NewRequest("GET", url)
	if err != nil{
		return nil, err
	}
	
	t := &TeamWeekly{}

	_, err = c.Do(response, t)
	if err != nil{
		return nil, err
	}

	var teamweekly []TeamWeekly
	teamweekly = append(teamweekly, *t)

	
	return teamweekly, nil
}
