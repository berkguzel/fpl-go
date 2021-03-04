package main

import (
	"fmt"
	"github.com/berkguzel/fpl-go/fpl"
)
func main(){

	c := fpl.Client{}
	//v, _ := c.ListStandings()
	//fmt.Println(v)
	//array, _ :=  c.ListWeeklyPointsx(8)
	array, _ := c.Manager()
	
	fmt.Println(array)
	
}