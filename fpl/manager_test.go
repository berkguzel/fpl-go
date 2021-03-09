package fpl

import (
	"testing"
)

func TestManager(t *testing.T) {

	c := &Client{}

	got, err := c.Manager()
	if err != nil {
		t.Error(err)
	}
	if got == nil {
		t.Errorf("Could do not be succed, got %+v", got)
	}
}

func TestLeagueClassic(t *testing.T) {

	c := &Client{}

	got, err := c.LeagueClassic()
	if err != nil {
		t.Error(err)
	}
	if got == nil {
		t.Errorf("Could do not be succed, got %+v", got)
	}
}

func TestLeagueCup(t *testing.T) {

	c := &Client{}

	got, err := c.LeagueCup()
	if err != nil {
		t.Error(err)
	}
	if got == nil {
		t.Errorf("Could do not be succed, got %+v", got)
	}
}
