package fpl

import (
	"testing"
)

func TestListStandings(t *testing.T){

	c := &Client{}

	got, err := c.ListStandings()
	if err != nil{
		t.Error(err)
	}

	if got == nil{
		t.Errorf("Could do not be succeed, Got: %+v", got)
	}

}