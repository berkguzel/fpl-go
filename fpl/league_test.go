package fpl

import (
	"testing"
)

func TestLeague(t *testing.T) {

	c := NewClient(nil)

	got, err := c.League()
	if err != nil {
		t.Error(err)
	}

	if got == nil {
		t.Errorf("Could do not be succeed, Got: %+v", got)
	}

}

func TestGetStandings(t *testing.T) {

	c := NewClient(nil)

	got, err := c.GetStandings()
	if err != nil {
		t.Error(err)
	}

	if got == nil {
		t.Errorf("Could do not be succeed, Got: %+v", got)
	}

}

func TestGetNewEntrries(t *testing.T) {

	c := NewClient(nil)

	got, err := c.GetNewEntries()
	if err != nil {
		t.Error(err)
	}

	if got == nil {
		t.Errorf("Could do not be succeed, Got: %+v", got)
	}

}