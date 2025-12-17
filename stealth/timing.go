package stealth

import (
	"math/rand"
	"time"
)

// Think simulates human thinking time between actions
func Think(minMs int, maxMs int) {
	if maxMs <= minMs {
		time.Sleep(time.Duration(minMs) * time.Millisecond)
		return
	}

	delay := rand.Intn(maxMs-minMs) + minMs
	time.Sleep(time.Duration(delay) * time.Millisecond)
}
