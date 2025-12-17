package main

import (
	"fmt"

	"linkedin-automation/connect"
	"linkedin-automation/search"
	"linkedin-automation/storage"
	"linedin-automation/logger"
)

func main() {
	logger.Info("Starting LinkedIn Automation PoC - Persistence Demo")

	store := storage.NewStore("state.json")

	results := search.SearchProfiles(
		search.SearchQuery{
			JobTitle: "Engineer",
			Company:  "Tech",
			Keywords: "Automation",
		},
		1,
	)

	connector := connect.NewConnector(3, store)
	connector.SendRequests(results)

	fmt.Println("Run completed. Restart to verify persistence.")
}
