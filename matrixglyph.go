package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// MatrixGlyph represents a Matrix code character
type MatrixGlyph struct {
	Char   string
	X      float32
	Y      float32
	IsHead bool
	Health float32
}

// Draw the glyph
func (g *MatrixGlyph) Draw() {
	if !g.IsHead && g.Health <= 0 {
		// The glyph is dead
		return
	}

	// Bright glyph if Head
	color := HeadGlyphColor

	if !g.IsHead {
		// Calculate glyph color if tail, based on the glyph's health
		f := 60.0 / 100.0 * g.Health
		if f > 255 {
			f = 255
		} else if f < 0 {
			f = 0
		}

		h := 255.0 / 100.0 * g.Health
		if h > 255 {
			h = 255
		} else if h < 0 {
			h = 0
		}

		color = rl.Color{R: uint8(f), G: uint8(h), B: uint8(f), A: 255}
	}

	rl.DrawTextEx(MatrixFont,
		g.Char,
		rl.Vector2{X: g.X, Y: g.Y},
		float32(GlyphSize),
		float32(0),
		color)
}
