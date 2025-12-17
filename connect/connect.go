package connect

import (
	"fmt"

	"linkedin-automation/storage"
	"linkedin-automation/search"
	"linkedin-automation/stealth"
)

// Connector handles connection request logic
type Connector struct {
	Limiter *stealth.RateLimiter
	Store   *storage.Store
}


// NewConnector creates a new connector
func NewConnector(limit int, store *storage.Store) *Connector {
	return &Connector{
		Limiter: stealth.NewRateLimiter(limit),
		Store:   store,
	}
}


// SendRequests simulates sending connection requests
func (c *Connector) SendRequests(profiles []search.Result) {
	for _, profile := range profiles {

		// duplicate prevention
		if c.Store.State.SentConnections[profile.ProfileURL] {
			fmt.Println("Already sent request to:", profile.ProfileURL)
			continue
		}

		// rate limit enforcement
		if !c.Limiter.Allow() {
			fmt.Println("Daily connection limit reached. Stopping.")
			return
		}

		// simulate visiting profile
		fmt.Println("Visiting profile:", profile.ProfileURL)
		stealth.Think(1000, 2500)

		// simulate sending request
		fmt.Println("Sending connection request to:", profile.ProfileURL)
		stealth.Think(800, 1500)

		c.Store.MarkConnectionSent(profile.ProfileURL)

	}
}
