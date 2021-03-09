package fpl

import (
	"encoding/json"
)

func (c *Client) GetGeneral() (*General, error) {

	url := "https://fantasy.premierleague.com/api/bootstrap-static/"

	response, _ := c.Do("GET", url)
	general := &General{}

	if err := json.Unmarshal(response, &general); err != nil {
		return nil, err
	}

	return general, nil
}

func (c *Client) ListTeamInfo() ([]TeamResponse, error) {

	general, _ := c.GetGeneral()

	var t []TeamResponse

	m, err := json.Marshal(general.Teams)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(m, &t); err != nil {
		return nil, err
	}

	return t, nil
}

func (c *Client) ListEventInfo() ([]EventsResponse, error) {

	general, _ := c.GetGeneral()

	var e []EventsResponse

	m, err := json.Marshal(general.Events)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(m, &e); err != nil {
		return nil, err
	}

	return e, nil

}

func (c *Client) ListPhasesInfo() ([]PhasesResponse, error) {

	general, _ := c.GetGeneral()

	var p []PhasesResponse

	m, err := json.Marshal(general.Phases)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(m, &p); err != nil {
		return nil, err
	}

	return p, nil

}

func (c *Client) ListElementsInfo() ([]ElementsResponse, error) {

	general, _ := c.GetGeneral()

	var e []ElementsResponse

	m, err := json.Marshal(general.Elements)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(m, &e); err != nil {
		return nil, err
	}

	return e, nil

}

func (c *Client) ListElementStatsInfo() ([]ElementStatsResponse, error) {

	general, _ := c.GetGeneral()

	var es []ElementStatsResponse

	m, err := json.Marshal(general.ElementStats)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(m, &es); err != nil {
		return nil, err
	}

	return es, nil

}

func (c *Client) ListElementTypesInfo() ([]ElementTypesResponse, error) {

	general, _ := c.GetGeneral()

	var et []ElementTypesResponse

	m, err := json.Marshal(general.ElementTypes)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(m, &et); err != nil {
		return nil, err
	}

	return et, nil

}

func (c *Client) ListGameSettings() ([]GameSettingsResponse, error) {

	general, _ := c.GetGeneral()

	g := &GameSettingsResponse{}

	m, err := json.Marshal(general.GameSettings)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(m, &g); err != nil {
		return nil, err
	}

	var gameResponse []GameSettingsResponse
	gameResponse = append(gameResponse, *g)

	return gameResponse, nil

}
