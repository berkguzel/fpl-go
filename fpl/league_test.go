package fpl

import (
	"testing"
)

func TestLeague(t *testing.T) {

	c := NewClient(nil)

	got, _ := c.League("")
	if got != nil {
		t.Errorf("Got: %+v", got)
	}

}

func TestGetStandings(t *testing.T) {

	c := NewClient(nil)

	got, _ := c.GetStandings("")
	if got != nil {
		t.Errorf("Got: %+v", got)
	}

}

func TestGetNewEntrries(t *testing.T) {

	c := NewClient(nil)

	got, _ := c.GetNewEntries("")
	if got != nil {
		t.Errorf("Got: %+v", got)
	}

}
