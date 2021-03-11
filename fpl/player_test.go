package fpl

import (
	"testing"
)

func TestGetPlayer(t *testing.T) {

	c := NewClient(nil)

	got, err := c.GetPlayer(1)
	if err != nil {
		t.Errorf("Could not access to data")
	}

	if got == nil {
		t.Errorf("Somethings went wrong")
	}

}

func TestPlayerFixture(t *testing.T) {

	c := NewClient(nil)

	got, err := c.PlayerFixture(1)
	if err != nil {
		t.Errorf("Could not access to data")
	}

	if got == nil {
		t.Errorf("Somethings went wrong")
	}

}

func TestPlayerHistory(t *testing.T) {

	c := NewClient(nil)

	got, err := c.PlayerHistory(1)
	if err != nil {
		t.Errorf("Could not access to data")
	}
	if got == nil {
		t.Errorf("Somethings went wrong")
	}

}

func TestPlayerHistoryPast(t *testing.T) {

	c := NewClient(nil)

	got, err := c.PlayerHistoryPast(1)
	if err != nil {
		t.Errorf("Could not access to data")
	}

	if got == nil {
		t.Errorf("Somethings went wrong")
	}

}
