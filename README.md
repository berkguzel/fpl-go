# fpl-go

This package presents a client for [Fantasy Premier League](https://fantasy.premierleague.com/) API.

fpl-go is still in progress.


```go
package main

import (
	"fmt"

	"github.com/berkguzel/fpl-go/fpl"
)

func main(){

	c := fpl.NewClient(nil)

	standings, err := c.GetStandings("league id")
	if err != nil{
		fmt.Println(err)
	}

	fmt.Println(standings)
}


```
