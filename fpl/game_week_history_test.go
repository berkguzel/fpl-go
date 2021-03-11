package fpl

import (
	"testing"
)

func TestListWeeklyPoints(t *testing.T) {

	c := NewClient(nil)

	got, _ := c.ListWeeklyPoints(0, "")
	if got != nil {
		t.Errorf("Got: %+v", got)
	}
}

func TestListAllWeeks(t *testing.T) {

	c := NewClient(nil)

	got, _ := c.ListAllWeeks("")
	if got != nil {
		t.Errorf("Got: %+v", got)
	}

}

func TestListWeeklyPerformance(t *testing.T) {

	c := NewClient(nil)

	got, _ := c.ListWeeklyPerformance(-1, "")
	if got != nil {
		t.Errorf("Got: %+v", got)
	}
}
