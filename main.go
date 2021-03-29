package main

import (
	"flag"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	Fullscreen = flag.Bool("fullscreen", false, "Starts in fullscreen mode")

	Height     int
	Width      int
	MatrixFont rl.Font

	BgColor        = rl.Color{R: 10, G: 10, B: 0, A: 100}
	HeadGlyphColor = rl.Color{R: 155, G: 255, B: 155, A: 255}

	MaxTailSize = 15
	GlyphSize   = 20

	FrameRate = int32(45)
)

func main() {
	flag.Parse()

	rl.InitWindow(1280, 768, "Matrix Code")

	rl.SetTargetFPS(FrameRate)
	MatrixFont = rl.LoadFont("matrix-code-nfi.ttf")

	if *Fullscreen {
		Width = rl.GetMonitorWidth(0)
		Height = rl.GetMonitorHeight(0)
		rl.ToggleFullscreen()
	} else {
		Width = 800
		Height = 800
	}

	rl.SetWindowSize(Width, Height)
	rl.SetWindowPosition(10, 30)

	// Create initial matrix streams
	streams := []*MatrixStream{}
	for i := 0; i < 30; i++ {
		s := NewMatrixStream()
		streams = append(streams, s)
	}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		// A/Z - Change tail size
		// TODO: It's not working as expected, leaving some glyphs showing forever...
		if rl.IsKeyDown(rl.KeyA) {
			MaxTailSize++
		} else if rl.IsKeyDown(rl.KeyZ) {
			if MaxTailSize > 2 {
				MaxTailSize--
			}
		}

		// S/X - Change glyph size
		if rl.IsKeyDown(rl.KeyS) {
			GlyphSize++
		} else if rl.IsKeyDown(rl.KeyX) {
			if GlyphSize > 2 {
				GlyphSize--
			}
		}

		// D/C - Change framerate
		if rl.IsKeyDown(rl.KeyD) {
			FrameRate++
			rl.SetTargetFPS(FrameRate)
		} else if rl.IsKeyDown(rl.KeyC) {
			if FrameRate > 2 {
				FrameRate--
				rl.SetTargetFPS(FrameRate)
			}
		}

		// F/V - Change number os matrix streams
		if rl.IsKeyDown(rl.KeyF) {
			s := NewMatrixStream()
			streams = append(streams, s)
		} else if rl.IsKeyDown(rl.KeyV) {
			if len(streams) > 1 {
				streams = streams[:len(streams)-1]
			}
		}

		// Draws a black translucent background
		rl.DrawRectangle(0, 0, int32(Width), int32(Height), BgColor)

		// Draw the streams
		for _, s := range streams {
			s.Draw()
		}

		rl.EndDrawing()
	}

	rl.UnloadFont(MatrixFont)
	rl.ClearDroppedFiles()
	rl.CloseWindow()
}
