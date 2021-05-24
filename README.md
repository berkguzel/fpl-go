# fpl-go

This package presents a client for [Fantasy Premier League](https://fantasy.premierleague.com/) API.


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
# FEATURES
- Standings
- Manager specific information
- Fixture
- Player specific information
- Weekly events



#### You can find an idea about how to use this client in this [repo](github.com/berkguzel/fpl-discord-bot).