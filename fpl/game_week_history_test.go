package fpl

import (
	"testing" 

)

func TestListWeeklyPoints(t *testing.T){

	c := &Client{}

	got, _ := c.ListWeeklyPoints(0)
	if got != nil {
		t.Errorf("Wanted nil but got: %+v", got)
	}
}

func TestListAllWeeks(t *testing.T){

	c := &Client{}

	got, err := c.ListAllWeeks()
	if err != nil{
		t.Fatal(err)
	}

	if got == nil {
		t.Errorf("Could do not be succeed, Got: %+v", got)
	}

}

func TestListWeeklyPerformance(t *testing.T){

	c := &Client{}

	got, _ := c.ListWeeklyPerformance(-1)
	if got != nil {
		t.Errorf("Wanted nil but got: %+v", got)
	}

}