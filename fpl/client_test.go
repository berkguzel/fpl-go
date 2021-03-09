package fpl

import (
	"testing"
)

func TestDo(t *testing.T) {

	c := &Client{}

	got, err := c.Do("GET", "https://fantasy.premierleague.com/api")
	if err != nil {
		t.Errorf("GOT: %v, There is no accessible url", got)
	}
}
