package main

import (
	"./modules"
	"github.com/gen2brain/raylib-go/raylib"
)

var (
	game modules.Game // game.
)

func main() {

	game.Init()
	rl.InitWindow(game.ScreenWidth, game.ScreenHeight, game.NameWindow)

	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {
		game.Update()

	}
	rl.CloseWindow()
}
