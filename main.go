package main

import (
	"flag"
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	Fullscreen = flag.Bool("fullscreen", false, "Starts in fullscreen mode")

	Height     int
	Width      int
	MatrixFont rl.Font

	BgColor        = rl.Color{R: 10, G: 10, B: 0, A: 100}
	HeadGlyphColor = rl.Color{R: 185, G: 255, B: 185, A: 255}

	MaxTailSize = 50
	GlyphSize   = 15

	FrameRate = int32(60)

	streams []*MatrixStream

	BlurSize = float32(8.0)

	ShowDebug = false
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

	// Create a RenderTexture2D to be used for render to texture
	target := rl.LoadRenderTexture(int32(Width), int32(Height))
	shader := rl.LoadShader("", "shaders/blur.fs")

	// Create initial matrix streams
	streams = []*MatrixStream{}
	for i := 0; i < 100; i++ {
		s := NewMatrixStream()
		streams = append(streams, s)
	}

	shaderWindowSizeLoc := rl.GetShaderLocation(shader, "windowSize")
	rl.SetShaderValue(shader, shaderWindowSizeLoc, []float32{float32(Width), float32(Height)}, 1)

	shaderBlurSizeLoc := rl.GetShaderLocation(shader, "blurSize")

	for !rl.WindowShouldClose() {
		handleUserInput()

		rl.BeginDrawing()

		rl.BeginTextureMode(target)
		rl.ClearBackground(rl.Black)

		// rl.ClearBackground(rl.Blank)

		// // Draw the streams
		for _, s := range streams {
			s.Draw()
		}

		rl.EndTextureMode()

		rl.SetShaderValue(shader, shaderBlurSizeLoc, []float32{BlurSize}, 0)

		rl.BeginShaderMode(shader)

		rl.DrawTextureRec(target.Texture, rl.NewRectangle(0, 0, float32(target.Texture.Width), float32(-target.Texture.Height)), rl.NewVector2(0, 0), rl.White)

		rl.EndShaderMode()

		rl.DrawTextureRec(target.Texture, rl.NewRectangle(0, 0, float32(target.Texture.Width), float32(-target.Texture.Height)), rl.NewVector2(0, 0), rl.Color{255, 255, 255, 128})

		// Help GUI
		if ShowDebug {
			rl.DrawText(fmt.Sprintf("%f FPS", rl.GetFPS()), 10, 10, 10, rl.White)
			rl.DrawText(fmt.Sprintf("A/Z - Change glyph lifetime = %d", MaxTailSize), 10, 30, 10, rl.White)
			rl.DrawText(fmt.Sprintf("S/X - Change glyph size = %d", GlyphSize), 10, 50, 10, rl.White)
			rl.DrawText(fmt.Sprintf("D/C - Change framerate = %d", FrameRate), 10, 70, 10, rl.White)
			rl.DrawText(fmt.Sprintf("F/V - Change qt of streams = %d", len(streams)), 10, 90, 10, rl.White)
			rl.DrawText(fmt.Sprintf("G/B - Change blur size = %f", BlurSize), 10, 110, 10, rl.White)
		} else if rl.GetTime() < 8 {
			rl.DrawText("Press F1 to show commands", 10, 10, 18, rl.White)
		}

		rl.EndDrawing()
	}

	rl.UnloadFont(MatrixFont)
	rl.UnloadShader(shader)
	rl.UnloadRenderTexture(target)
	rl.ClearDroppedFiles()
	rl.CloseWindow()
}

func handleUserInput() {
	if rl.IsKeyPressed(rl.KeyF1) {
		ShowDebug = !ShowDebug
	}

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

	// F/V - Change number of matrix streams
	if rl.IsKeyDown(rl.KeyF) {
		s := NewMatrixStream()
		streams = append(streams, s)
	} else if rl.IsKeyDown(rl.KeyV) {
		if len(streams) > 1 {
			streams = streams[:len(streams)-1]
		}
	}

	// G/B - Change blur size
	if rl.IsKeyDown(rl.KeyG) {
		BlurSize += 0.5
	} else if rl.IsKeyDown(rl.KeyB) {
		if BlurSize > 0 {
			BlurSize -= 0.5
		}
	}
}
