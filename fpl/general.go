package fpl

import (
	"fmt"
	"encoding/json"

)

const (
	generalAddress = "https://fantasy.premierleague.com/api/bootstrap-static/"
)

func (c *Client) GetGeneral() (*General, error){

	req, err := c.DoNewRequest("GET", generalAddress)
	if err != nil{
		fmt.Println(err)
	}
	
	response, respErr := c.Request(req)
	if respErr != nil{
		fmt.Println(respErr)
	}

	general  := &General{}

	if err := json.Unmarshal(response, &general); err != nil {
		fmt.Println(err)
    }

	return general, nil
}

func (c *Client) ListTeamInfo()([]TeamResponse, error){

	general, _ := c.GetGeneral()
	
	var t []TeamResponse
	
	m, err := json.Marshal(general.Teams)
	if err != nil{
		fmt.Println(err)
	}
	
	if err := json.Unmarshal(m, &t); err != nil {
		fmt.Println(err)
    }

	return t, nil
}

func (c *Client) ListEventInfo() ([]EventsResponse, error) {

	general, _ := c.GetGeneral()
	
	var e []EventsResponse
	
	m, err := json.Marshal(general.Events)
	if err != nil{
		fmt.Println(err)
	}
	
	if err := json.Unmarshal(m, &e); err != nil {
		fmt.Println(err)
    }

	return e, nil

}

func (c *Client) ListPhasesInfo() ([]PhasesResponse, error) {

	general, _ := c.GetGeneral()
	
	var p []PhasesResponse
	
	m, err := json.Marshal(general.Phases)
	if err != nil{
		fmt.Println(err)
	}
	
	if err := json.Unmarshal(m, &p); err != nil {
		fmt.Println(err)
    }

	return p, nil

}

func (c *Client) ListElementsInfo() ([]ElementsResponse, error) {

	general, _ := c.GetGeneral()
	
	var e []ElementsResponse
	
	m, err := json.Marshal(general.Elements)
	if err != nil{
		fmt.Println(err)
	}
	
	if err := json.Unmarshal(m, &e); err != nil {
		fmt.Println(err)
    }

	return e, nil

}

func (c *Client) ListElementStatsInfo() ([]ElementStatsResponse, error) {

	general, _ := c.GetGeneral()
	
	var es []ElementStatsResponse
	
	m, err := json.Marshal(general.ElementStats)
	if err != nil{
		fmt.Println(err)
	}
	
	if err := json.Unmarshal(m, &es); err != nil {
		fmt.Println(err)
    }

	return es, nil

}

func(c *Client) ListElementTypesInfo() ([]ElementTypesResponse, error) {

	general, _ := c.GetGeneral()
	
	var et []ElementTypesResponse
	
	m, err := json.Marshal(general.ElementTypes)
	if err != nil{
		fmt.Println(err)
	}
	
	if err := json.Unmarshal(m, &et); err != nil {
		fmt.Println(err)
    }

	return et, nil

}

func(c *Client) ListGameSettings() ([]GameSettingsResponse, error) {

	general, _ := c.GetGeneral()
	
	var g []GameSettingsResponse
	
	m, err := json.Marshal(general.GameSettings)
	if err != nil{
		fmt.Println(err)
	}
	
	if err := json.Unmarshal(m, &g); err != nil {
		fmt.Println(err)
    }

	return g, nil

}