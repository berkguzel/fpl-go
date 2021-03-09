package fpl

import (
	"testing" 

)

func TestGetFixture( t *testing.T){
	c := &Client{}
	
	got, err := c.GetFixture()
	if err != nil{
		t.Fatal(err)
	}
	if got == nil {
		t.Errorf("GOT: %+v, Could not find fixture", got)
	}
}