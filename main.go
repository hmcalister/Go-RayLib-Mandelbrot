package main

import (
	"flag"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	WINDOW_WIDTH  int32 = 800
	WINDOW_HEIGHT int32 = 450
)

func main() {
	rlLogLevelFlag := flag.String("rlLogLevel", "none", "Set the raylib log level. Valid values are: fatal, error, warning, info, debug, trace, none.")
	slogLevelFlag := flag.String("slogLevel", "none", "Set the slog level. Valid values are: fatal, error, warning, info, debug, trace, none.")
	slogFormatFlag := flag.String("slogFormat", "pretty", "Set the slog format. Valid values are: text, pretty, json.")
	flag.Parse()
	setupLogging(*slogLevelFlag, *slogFormatFlag, *rlLogLevelFlag)

	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)


		rl.EndDrawing()
	}
}
