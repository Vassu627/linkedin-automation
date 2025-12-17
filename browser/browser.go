package browser

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

// New creates and connects a visible, stealthy browser instance
func New() *rod.Browser {
	rand.Seed(time.Now().UnixNano())

	// random viewport (human-like screens)
	width := 1200 + rand.Intn(400)
	height := 700 + rand.Intn(300)

	l := launcher.New().
		Leakless(false). // Windows antivirus fix
		Headless(false).
		Set("disable-blink-features", "AutomationControlled").
		Set("window-size", fmt.Sprintf("%d,%d", width, height))

	url := l.MustLaunch()

	browser := rod.New().ControlURL(url).MustConnect()

	// hide navigator.webdriver
	browser.MustPage().MustEval(`
		Object.defineProperty(navigator, 'webdriver', {
			get: () => undefined
		})
	`)

	return browser
}
