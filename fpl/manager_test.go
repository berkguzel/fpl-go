package fpl

import (
	"testing"
)

func TestManager(t *testing.T) {

	c := NewClient(nil)

	got, _ := c.GetManager("")
	if got != nil {
		t.Errorf("Got: %+v", got)
	}
}

func TestLeagueClassic(t *testing.T) {

	c := NewClient(nil)

	got, _ := c.LeagueClassic("")
	if got != nil {
		t.Errorf("Got: %+v", got)
	}
}

func TestLeagueCup(t *testing.T) {

	c := NewClient(nil)

	got, _ := c.LeagueCup("")
	if got != nil {
		t.Errorf("Got: %+v", got)
	}
}
