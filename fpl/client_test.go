package fpl

import (
	"testing"
)

func TestNewRequest(t *testing.T) {

	c := NewClient(nil)

	got, err := c.NewRequest("GET", "https://fantasy.premierleague.com/api")
	if err != nil {
		t.Errorf("GOT: %v, There is no accessible url", got)
	}
}

func TestDo(t *testing.T) {

	c := NewClient(nil)

	response, err := c.NewRequest("GET", "https://fantasy.premierleague.com/api")
	if err != nil{
		t.Error(err)
	}

	got, _ := c.Do(response, nil)
	if got == nil {
		t.Errorf("Returned nil")
	}
}