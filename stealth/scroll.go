package stealth

import (
	"math/rand"

	"github.com/go-rod/rod"
)

// RandomScroll simulates natural human scrolling behavior
func RandomScroll(page *rod.Page) {
	steps := rand.Intn(3) + 3 // 3–5 scroll actions

	for i := 0; i < steps; i++ {
		// scroll down smoothly
		page.Mouse.Scroll(0, float64(rand.Intn(300)+100), rand.Intn(5)+3)
		Think(400, 900)
	}

	// occasional scroll up (human correction)
	if rand.Float32() < 0.3 {
		page.Mouse.Scroll(0, -float64(rand.Intn(200)), rand.Intn(4)+2)
		Think(300, 600)
	}
}
