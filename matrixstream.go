package main

import (
	"math/rand"
)

// MatrixStream represents a column of Matrix Glyphs
type MatrixStream struct {
	Head     *MatrixGlyph
	Tail     []*MatrixGlyph
	TailSize int
}

// NewMatrixStream returns a new column of glyphs, at a random screen location.
func NewMatrixStream() *MatrixStream {
	s := &MatrixStream{
		Head:     randomGlyph(),
		Tail:     make([]*MatrixGlyph, MaxTailSize+3),
		TailSize: randomTailSize(),
	}

	return s
}

// Draw the stream and all its glyphs, updates their position, creates new glyphs copies of the Head
func (s *MatrixStream) Draw() {
	// Draws the Head, the brightest glyph at the front of the stream
	s.Head.Draw()

	// Draws the Tail
	for _, g := range s.Tail {
		if g == nil || g.Health <= 0 {
			// The glyph is dead
			continue
		}

		g.Draw()

		// Update glyphs health based on tail size
		g.Health -= float32(100/s.TailSize + (rand.Intn(10) - 5))

		if g.Health <= 0 {
			// The glyph died
			g.Health = 0
		}

		if rand.Intn(15) == 0 {
			// Looking at the movie reference, the glyphs left behing change characters over time
			g.Char = randomChar()
		}
	}

	if rand.Intn(5) > 3 {
		// The Head leaves a copy of the character behind
		newTail := &MatrixGlyph{
			Char:   s.Head.Char,
			X:      s.Head.X,
			Y:      s.Head.Y,
			IsHead: false,
			Health: 100,
		}

		// Update Head with random char and new Y
		s.Head.Char = randomChar()
		s.Head.Y += float32(GlyphSize)

		if s.Head.Y > float32(Height) {
			// The Head went offscreen
			gridWidth := int(Width) / GlyphSize
			gridHeight := int(Height) / GlyphSize
			x := float32(rand.Intn(gridWidth) * GlyphSize)
			y := float32(rand.Intn(gridHeight) * GlyphSize)
			s.Head.X = x
			s.Head.Y = -y

			s.TailSize = randomTailSize()
		}

		// Find a tail position with a dead or empty glyph and substitute with the new tail
		for i, g := range s.Tail {
			if g != nil && g.Health > 0 {
				continue
			}

			s.Tail[i] = newTail
			break
		}
	}
}

const (
	CHARS = "abcdefghijklmnopqrstuvwxyz0123456789"
)

// randomChar returns a random alphanumeric char
func randomChar() string {
	randomIdx := rand.Intn(len(CHARS))
	randomChar := string(CHARS[randomIdx])

	return randomChar
}

// randomGlyph returns a new glyph for the Head, with a random character and position.
func randomGlyph() *MatrixGlyph {
	gridWidth := int(Width) / GlyphSize
	gridHeight := int(Height) / GlyphSize

	x := float32(rand.Intn(gridWidth) * GlyphSize)
	y := float32(rand.Intn(gridHeight) * GlyphSize)

	g := &MatrixGlyph{
		Char:   randomChar(),
		X:      x,
		Y:      y,
		IsHead: true,
	}

	return g
}

// randomTailSize returns a random size for the tails, from 3 to MaxTailSize
func randomTailSize() int {
	return rand.Intn(MaxTailSize) + 3
}
