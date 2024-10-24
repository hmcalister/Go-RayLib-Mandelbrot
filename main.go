package main

import (
	"flag"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	WINDOW_WIDTH  int32   = 800
	WINDOW_HEIGHT int32   = 600
	ZOOM_FACTOR   float64 = 1.25
)

type drawParameters struct {
	centerX       float64
	centerY       float64
	zoom          float64
	maxIterations int
}

func (params drawParameters) convertPixelToComplex(pixelX, pixelY int32) ComplexNumber {
	return ComplexNumber{
		params.centerX + (float64(pixelX)-float64(WINDOW_WIDTH)/2)/params.zoom,
		params.centerY + (float64(pixelY)-float64(WINDOW_HEIGHT)/2)/params.zoom,
	}
}

func main() {
	rlLogLevelFlag := flag.String("rlLogLevel", "none", "Set the raylib log level. Valid values are: fatal, error, warning, info, debug, trace, none.")
	slogLevelFlag := flag.String("slogLevel", "none", "Set the slog level. Valid values are: fatal, error, warning, info, debug, trace, none.")
	slogFormatFlag := flag.String("slogFormat", "pretty", "Set the slog format. Valid values are: text, pretty, json.")
	flag.Parse()
	setupLogging(*slogLevelFlag, *slogFormatFlag, *rlLogLevelFlag)

	rl.InitWindow(WINDOW_WIDTH, WINDOW_HEIGHT, "Mandelbrot Set")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	params := drawParameters{
		-0.75,
		0.0,
		200.0,
		100,
	}
	drawTexture := createMandelbrotTexture(params)

	leftMouseButtonPressed := false
	lastMouseX := int32(0)
	lastMouseY := int32(0)
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		mouseX := rl.GetMouseX()
		mouseY := rl.GetMouseY()
		scroll := rl.GetMouseWheelMove()

		// Handle mouse scrolling zoom
		if scroll != 0 {
			mouseComplex := params.convertPixelToComplex(mouseX, mouseY)
			if scroll > 0 {
				params.zoom *= ZOOM_FACTOR
			} else {
				params.zoom /= ZOOM_FACTOR
			}
			params.centerX = mouseComplex.Re - (float64(mouseX)-float64(WINDOW_WIDTH)/2.0)/params.zoom
			params.centerY = mouseComplex.Im - (float64(mouseY)-float64(WINDOW_HEIGHT)/2.0)/params.zoom
			drawTexture = createMandelbrotTexture(params)
		}

		// Handle mouse panning
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			rl.DrawText("MOUSE", 100, 100, 20, rl.White)
			leftMouseButtonPressed = true
			lastMouseX = mouseX
			lastMouseY = mouseY
		}
		if rl.IsMouseButtonReleased(rl.MouseLeftButton) {
			leftMouseButtonPressed = false
		}
		if leftMouseButtonPressed {
			deltaX := float64(lastMouseX - mouseX)
			deltaY := float64(lastMouseY - mouseY)
			params.centerX += deltaX / params.zoom
			params.centerY += deltaY / params.zoom
			lastMouseX = mouseX
			lastMouseY = mouseY
			drawTexture = createMandelbrotTexture(params)
		}

		rl.DrawTexture(drawTexture, 0, 0, rl.White)

		rl.EndDrawing()
	}
}
