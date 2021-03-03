package fpl

import (
	"fmt"
	"encoding/json"
	"errors"
	"strings"

	//"github.com/fatih/structs"
)
type Person struct {
	Name  string
	Last string
}

const (
	addr = "https://fantasy.premierleague.com/api/entry/7270639/history/"
)

func (c *Client) ListWeeklyPoints(gameWeek int) (*WeeklyResponse, error){

	req, err := c.DoNewRequest("GET", addr)
	if err != nil{
		fmt.Println(err)
	}
	
	response, respErr := c.Request(req)
	if respErr != nil{
		return nil, respErr
	}

	v := &Weekly{}
	if err := json.Unmarshal(response, &v); err != nil {
		fmt.Println(err)
    }

	
	for index, value := range  v.Current{

		if value.Event == gameWeek {
			
			m, err := json.Marshal(v.Current[index])
			if err != nil{
				fmt.Println(err)
			}
			fmt.Printf("%T", v.Current[index])

			w := &WeeklyResponse{}
			
			dec := json.NewDecoder(strings.NewReader(string(m)))
			erros := dec.Decode(&w)
			if erros != nil{
				fmt.Println(erros)
			}
			return w, nil		
		} 		
	}
	return nil, errors.New("Could not find game week")
}

func (c *Client) ListWeeklyPointsx(gameWeek int) (*WeeklyResponse, error){

	req, err := c.DoNewRequest("GET", addr)
	if err != nil{
		fmt.Println(err)
	}
	
	response, respErr := c.Request(req)
	if respErr != nil{
		return nil, respErr
	}

	v := &Weekly{}
	if err := json.Unmarshal(response, &v); err != nil {
		fmt.Println(err)
    }

	var person map[string]interface{}	
	json.Unmarshal([]byte(response), &person)

	for _, v := range person {
		switch v := v.(type) {
		case []interface{}:
            for _, ival := range v {
				
                fmt.Printf("%T", ival)
            }
		}


    }

	
	
	return nil, errors.New("Could not find game week")
}


func (c *Client) ListAllWeeks() ([]*Weekly, error) {
	
	req, err := c.DoNewRequest("GET", addr)
	if err != nil{
		fmt.Println(err)
	}
	
	response, respErr := c.Request(req)
	if err != nil{
		fmt.Println(respErr)
	}

	w := &Weekly{}
	json.Unmarshal(response, &w)

	var arrayOfWeeks []*Weekly
	arrayOfWeeks = append(arrayOfWeeks, w)

	return  arrayOfWeeks, nil
}
