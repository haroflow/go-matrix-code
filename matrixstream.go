package main

type MatrixStream struct {
	Glyphs []*MatrixGlyph
}

func (s *MatrixStream) Draw() {
	for _, g := range s.Glyphs {
		g.Draw()
	}
}
