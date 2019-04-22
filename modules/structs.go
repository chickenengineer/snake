package modules

import (
	"github.com/gen2brain/raylib-go/raylib"
)

type (
	Button struct {
		Position rl.Vector2
		Text     string
		Size     rl.Vector2
		Color    rl.Color
		Font     int32
	}

	Menu struct {
		Buttons    []Button
		Showed     bool
		Background rl.Color
	}

	// entities.
	Cube struct {
		Position     rl.Vector2
		Color        rl.Color
		PastPosition rl.Vector2
	}

	Snake struct {
		Living       bool
		Color        rl.Color
		Position     rl.Vector2
		NextPosition rl.Vector2
		PastPosition rl.Vector2
		Cubes        []Cube
	}

	Feed struct {
		Position rl.Vector2
		Color    rl.Color
		Power    int32
	}

	Game struct {
		// menues.
		Menu         Menu
		GameOverMenu Menu
		SettingsMenu Menu

		ScreenWidth  int32
		ScreenHeight int32
		NameWindow   string
		Background   rl.Color

		GameOver  bool
		Time      int32
		Player    Snake
		Feed      Feed
		Score     int32
		BestScore int32
	}
)
