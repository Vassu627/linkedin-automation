package stealth

import (
	"math/rand"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
)

// HumanType simulates realistic human typing
func HumanType(el *rod.Element, text string) {
	for _, ch := range text {

		// type character
		el.Input(string(ch))

		// occasional typo + correction
		if rand.Float32() < 0.05 {
			ka, err := el.KeyActions()
			if err == nil {
				ka.Press(input.Backspace).Do()
			}
			el.Input(string(ch))
		}

		// random typing delay
		Think(80, 200)
	}
}
