package modules

import (
	"github.com/gen2brain/raylib-go/raylib"
	"strconv"
)

func (g *Game) Draw() {

	rl.BeginDrawing()
	// showing Menu.
	if g.Menu.Showed { // if Menu is showed.
		g.Menu.Draw()
	} else if g.GameOverMenu.Showed { // if GameOverMenu is showed.
		g.Player.Draw()
		g.GameOverMenu.Draw()
		rl.DrawText("SCORE: "+strconv.Itoa(int(g.Score)), (g.ScreenWidth-rl.MeasureText("SCORE: "+strconv.Itoa(int(g.Score)), 50))/2, g.ScreenHeight/2-50, 50, rl.White)
	} else if !g.GameOver { // if game is started.
		// Drawing game.
		g.Feed.Draw()
		g.Player.Draw()
		// Score.
		rl.DrawText("SCORE: "+strconv.Itoa(int(g.Score)), (g.ScreenWidth-rl.MeasureText("SCORE: "+strconv.Itoa(int(g.Score)), 20))/2, 0, 20, rl.White)
	} else if g.SettingsMenu.Showed { // if SettingsMenu is showed.
		g.SettingsMenu.Draw()
	}

	// Drawing FPS.
	if settings.FPSShowed {
		rl.DrawText(strconv.Itoa(int(rl.GetFPS())), 0, 0, 20, rl.White)
	}
	rl.ClearBackground(rl.Black)
	rl.EndDrawing()
}

func (Snake *Snake) Draw() {
	rl.DrawRectangleV(Snake.Position, rl.Vector2{10, 10}, Snake.Color)
	for _, cube := range Snake.Cubes {
		rl.DrawRectangleV(cube.Position, rl.Vector2{10, 10}, cube.Color)
	}
}

func (Feed *Feed) Draw() {
	rl.DrawRectangleV(Feed.Position, rl.Vector2{10, 10}, Feed.Color)
}

func (Menu *Menu) Draw() {
	for _, button := range Menu.Buttons {
		button.Draw()
	}
}

func (Button *Button) Draw() {
	rl.DrawRectangleV(Button.Position, Button.Size, Button.Color)
	rl.DrawText(Button.Text, int32(Button.Position.X)+(int32(Button.Size.X)-rl.MeasureText(Button.Text, Button.Font))/2, int32(Button.Position.Y)+((int32(Button.Size.Y)-Button.Font)/2), Button.Font, rl.White)
}
