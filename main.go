package main

import (
	"flag"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	WINDOW_WIDTH  int32 = 800
	WINDOW_HEIGHT int32 = 600
)

type drawParameters struct {
	centerX float64
	centerY float64
	zoom    float64
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

	param := drawParameters{
		-0.75,
		0.0,
		200.0,
	}
	drawTexture := createMandelbrotTexture(param)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		// TODO: Handle user input and redraw image

		rl.DrawTexture(drawTexture, 0, 0, rl.White)

		rl.EndDrawing()
	}
}
