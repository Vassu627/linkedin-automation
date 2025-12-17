package stealth

import (
	"math/rand"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

// HumanMove moves the mouse in a human-like way
func HumanMove(page *rod.Page, x float64, y float64) {
	// start slightly away from target
	startX := x - float64(rand.Intn(50)+20)
	startY := y - float64(rand.Intn(50)+20)

	page.Mouse.MoveTo(proto.Point{X: startX, Y: startY})
	Think(100, 300)

	// move closer (slight overshoot)
	page.Mouse.MoveTo(proto.Point{X: x + 5, Y: y + 5})
	Think(50, 150)

	// micro correction
	page.Mouse.MoveTo(proto.Point{X: x, Y: y})
	Think(50, 120)
}
