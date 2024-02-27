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

	// list the standings in the league
	standings, err := c.GetStandings("league id")
	if err != nil{
		fmt.Println(err)
	}

	// list the info about the manager
	m, err := c.GetManager("manager id")
	if err != nil {
		fmt.Println(err)
	}

}


```



# FEATURES
- League
    - Standings
    - New entries
    - Team info  
- Manager specific information
    - Points of the manager for all game weeks
    - Info about the clssic leagues and cups  
- Fixture
    - General fixture
    - Weekly fixture
- Player specific information
	- Code of player in FPL
	- Next fixtures
	- Information for previous seasons and next fixtures
	- Detailed info
- Weekly events
    - Points of the manager and player for game weeks
    - Info about FPL players by week
- General
    - Info about current EPL clubs and game weeks
    - Game setttings
    - Phases of EPL season
  


#### You can find an idea about how to use this client in this [repo](https://github.com/berkguzel/fpl-discord-bot).
