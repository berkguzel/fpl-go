package fpl

import (
	"os"
)

func Get() (string, string) {
	managerID := os.Getenv("MANAGER_ID")
	leagueID := os.Getenv("LEAGUE_ID")

	return managerID, leagueID
}
