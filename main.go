package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	CHARS = "abcdefghijklmnopqrstuvwxyz0123456789"
)

var (
	HEIGHT  int32
	WIDTH   int32
	BGCOLOR = rl.Color{R: 10, G: 10, B: 0, A: 100}
	FONT    rl.Font
)

func main() {
	rl.InitWindow(0, 0, "Matrix Code")

	WIDTH = int32(rl.GetMonitorWidth(0))
	HEIGHT = int32(rl.GetMonitorHeight(0))

	rl.ToggleFullscreen()
	rl.SetTargetFPS(60)

	FONT = rl.LoadFont("matrix-code-nfi.ttf")

	x := float32(190)

	// If we clear at least once, it looks better, less flicker on the background...
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)
	rl.EndDrawing()

	// glyphs := make([]*MatrixGlyph, 0, 100)
	// for i := 0; i < 100; i++ {
	// 	randomIdx := rand.Intn(len(CHARS))
	// 	randomChar := string(CHARS[randomIdx])
	// 	g := &MatrixGlyph{
	// 		Char:   randomChar,
	// 		X:      rand.Float32() * float32(WIDTH),
	// 		Y:      rand.Float32() * float32(HEIGHT),
	// 		Health: 100,
	// 		IsHead: rand.Intn(2) == 0,
	// 	}

	// 	glyphs = append(glyphs, g)
	// }

	stream := &MatrixStream{
		Glyphs: []*MatrixGlyph{
			{
				Char:   "a",
				X:      10,
				Y:      40,
				Health: 100,
				IsHead: true,
			},
			{
				Char:   "b",
				X:      10,
				Y:      20,
				Health: 100,
				IsHead: false,
			},
		},
	}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.DrawRectangle(0, 0, WIDTH, HEIGHT, BGCOLOR)

		stream.Draw()
		// for _, g := range glyphs {
		// 	g.Draw()
		// }

		rl.EndDrawing()
		x++
	}

	rl.UnloadFont(FONT)
	rl.ClearDroppedFiles()
	rl.CloseWindow()
}
