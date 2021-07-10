package main

import (

	github.com/berkguzel/fpl-go
)

// This part shows how to see information about manager
func main() {
	c := Authentication()

	id, err := ManagerID(leagueID, teamName)
	if err != nil {
		return "", errors.New("Could not find the manager id")
	}

	m, err := c.GetManager(strconv.Itoa(id))
	if err != nil {
		return "", errors.New("Could not find the manager")
	}
}