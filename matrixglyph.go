package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	HEADGLYPHCOLOR = rl.Color{R: 155, G: 255, B: 155, A: 255}
)

type MatrixGlyph struct {
	Char   string
	X      float32
	Y      float32
	IsHead bool
	Health float32
}

func (g *MatrixGlyph) Draw() {
	dx := rand.Float32() * 3
	dy := rand.Float32() * 3

	color := HEADGLYPHCOLOR
	if !g.IsHead {
		f := uint8(60.0 / 100.0 * g.Health)
		if f > 255 {
			f = 255
		}

		h := uint8(255.0 / 100.0 * g.Health)
		if h > 255 {
			h = 255
		}

		color = rl.Color{R: f, G: h, B: f, A: 255}
		// text.setFillColor(sf::Color(f, h, f, 255));
	}
	rl.DrawTextEx(FONT, g.Char, rl.Vector2{X: g.X + dx, Y: g.Y + dy}, float32(20), float32(0), color)

	if g.Health > 0 {
		g.Health--
	}
}
