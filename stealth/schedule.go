package stealth

import "time"

// AllowedNow checks if automation is allowed in human-like hours
func AllowedNow() bool {
	hour := time.Now().Hour()

	// simulate business hours: 9 AM – 6 PM
	return hour >= 9 && hour <= 18
}
