package main

import (
	"fmt"
	"github.com/berkguzel/fpl-go/fpl"
)
func main(){

	c := fpl.Client{}
	array := c.ListWeeklyPoints()
	fmt.Println(array.Rank)
	for i, v := range array {
		fmt.Println(i, v)
	}

	

}