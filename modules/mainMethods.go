package modules

import (
	"github.com/gen2brain/raylib-go/raylib"
	"math/rand"
)

var (
	game     *Game        // game (alias is "g" in methods).
	places   []rl.Vector2 // all places.
	Settings settings
)

// Init initilization. just once for execution of program.
func (g *Game) Init() {
	game = g
	// Settings initilization.
	Settings = settings{}
	// window initilization.
	g.ScreenWidth = 500
	g.ScreenHeight = 500
	g.NameWindow = "Snake"
	g.Background = rl.Lime

	// game initilization.
	g.GameOver = true
	// fill places(variable).
	for x := 0; int32(x) < g.ScreenWidth; x += 10 {
		for y := 0; int32(y) < g.ScreenHeight; y += 10 {
			places = append(places, rl.Vector2{float32(x), float32(y)})
		}
	}
	// snake initilization.
	g.Player.Color = rl.Pink
	g.Player.Position = rl.Vector2{float32(g.ScreenWidth / 2), float32(g.ScreenHeight / 2)}
	// Feed initilization.
	g.Feed.Power = 10
	g.Feed.Color = rl.Lime
	// Menu initilization.
	var bExit = Button{rl.Vector2{0, 200}, "Exit", rl.Vector2{float32(g.ScreenWidth), 100}, rl.Red, 50}
	var bStart = Button{rl.Vector2{0, 0}, "Play", rl.Vector2{float32(g.ScreenWidth), 100}, rl.Green, 50}
	var bSettings = Button{rl.Vector2{0, 100}, "Settings", rl.Vector2{float32(g.ScreenWidth), 100}, rl.Blue, 50}
	g.Menu = Menu{
		Buttons:    []Button{bStart, bSettings, bExit},
		Showed:     true,
		Background: rl.Green,
	}
	// GameOverMenu initilization.
	var bMenu = Button{rl.Vector2{float32((g.ScreenWidth - (g.ScreenWidth - 20))) / 2, float32((g.ScreenHeight + 50)) / 2}, "Menu", rl.Vector2{float32(g.ScreenWidth - 20), 100}, rl.Orange, 50}
	g.GameOverMenu = Menu{
		Buttons:    []Button{bMenu},
		Showed:     false,
		Background: g.Background,
	}
	// SettingsMenu initilization.
	var bMode = Button{rl.Vector2{0, 0}, "Mode: Normal", rl.Vector2{float32(g.ScreenWidth), 100}, rl.Gray, 50}
	var bDifficult = Button{rl.Vector2{0, 100}, "Difficult: Easy", rl.Vector2{float32(g.ScreenWidth), 100}, rl.DarkGray, 50}
	var bShowedFPS = Button{rl.Vector2{0, 200}, "FPS Showed: NO", rl.Vector2{float32(g.ScreenWidth), 100}, rl.Brown, 50}
	var bCancel = Button{rl.Vector2{0, 300}, "Cancel", rl.Vector2{float32(g.ScreenWidth), 100}, rl.Gold, 50}
	g.SettingsMenu = Menu{
		Buttons:    []Button{bMode, bDifficult, bShowedFPS, bCancel},
		Showed:     false,
		Background: g.Background,
	}
}

var (
	addPosition    rl.Vector2
	latestSaveTime float32
)

func (g *Game) Update() {

	if !g.GameOver { // if game is Started.
		// keyChecker.
		switch {
		case rl.IsKeyPressed(rl.KeyW) && (addPosition != rl.Vector2{0, 10}):
			addPosition = rl.Vector2{0, -10}
		case rl.IsKeyPressed(rl.KeyA) && (addPosition != rl.Vector2{10, 0}):
			addPosition = rl.Vector2{-10, 0}
		case rl.IsKeyPressed(rl.KeyS) && (addPosition != rl.Vector2{0, -10}):
			addPosition = rl.Vector2{0, 10}
		case rl.IsKeyPressed(rl.KeyD) && (addPosition != rl.Vector2{-10, 0}):
			addPosition = rl.Vector2{10, 0}
		}

		if latestSaveTime < (rl.GetTime() - 0.05 + 0.01*float32(Settings.difficult)) {
			g.Player.Move(g)
		}
	} else if g.Menu.Showed { // if Menu is Showed.
		switch {
		case g.Menu.Buttons[0].IsClicked(): // Start.
			g.Start()
		case g.Menu.Buttons[1].IsClicked(): // Settings.
			g.Menu.Showed = false
			g.SettingsMenu.Showed = true
		case g.Menu.Buttons[2].IsClicked(): // exit.
			rl.CloseWindow()
		}
	} else if g.GameOverMenu.Showed { // if GameOverMenu is Showed.
		switch {
		case g.GameOverMenu.Buttons[0].IsClicked():
			g.GameOverMenu.Showed = false
			g.Menu.Showed = true
			// returning snake to Start.
			g.Player.Cubes = []Cube{}
			g.Player.Position = rl.Vector2{float32(g.ScreenWidth / 2), float32(g.ScreenHeight / 2)}
		}
	} else if g.SettingsMenu.Showed { // if SettingsMenu is Showed.
		switch {
		case g.SettingsMenu.Buttons[1].IsClicked(): // difficult.
			switch Settings.difficult {
			case 0:
				Settings.difficult++
				g.SettingsMenu.Buttons[1].Text = "Difficult: Normal"
			case 1:
				Settings.difficult++
				g.SettingsMenu.Buttons[1].Text = "Difficult: Hard"
				g.SettingsMenu.Buttons[1].Color = rl.Red
			case 2:
				Settings.difficult = 0
				g.SettingsMenu.Buttons[1].Text = "Difficult: Easy"
				g.SettingsMenu.Buttons[1].Color = rl.DarkGray
			}
		case g.SettingsMenu.Buttons[2].IsClicked(): // Settings.
			if Settings.FPSShowed {
				Settings.FPSShowed = false

				g.SettingsMenu.Buttons[2].Text = "FPS Showed: NO"
			} else {
				Settings.FPSShowed = true

				g.SettingsMenu.Buttons[2].Text = "FPS Showed: YES"
			}
		case g.SettingsMenu.Buttons[3].IsClicked(): // cancel.
			g.Menu.Showed = true
			g.SettingsMenu.Showed = false
		}
	}
	g.Draw()
}

func (snake *Snake) Die() {
	snake.Living = false
	game.GameOver = true
	if game.Score > game.BestScore {
		game.BestScore = game.Score
	}
	game.GameOverMenu.Showed = true
}

func (snake *Snake) Move(g *Game) {
	latestSaveTime = rl.GetTime()
	p := &g.Player

	// move head.
	p.PastPosition = p.Position
	p.NextPosition = rl.Vector2{addPosition.X + p.Position.X, addPosition.Y + p.Position.Y}
	// teleportation head.
	switch {
	case p.NextPosition.X == -10:
		p.NextPosition.X = float32(g.ScreenWidth - 10)
	case p.NextPosition.X == float32(g.ScreenWidth):
		p.NextPosition.X = 0
	case p.NextPosition.Y == -10:
		p.NextPosition.Y = float32(g.ScreenHeight - 10)
	case p.NextPosition.Y == float32(g.ScreenHeight):
		p.NextPosition.Y = 0
	}
	p.Position = p.NextPosition

	// move body.

	if len(p.Cubes) != 0 {
		// first cube.
		cube := &p.Cubes[0]
		cube.PastPosition = cube.Position
		cube.Position = p.PastPosition
		// next Cubes.
		if len(p.Cubes) > 1 {
			for i := range p.Cubes[1:] {
				p.Cubes[i+1].PastPosition = p.Cubes[i+1].Position
				p.Cubes[i+1].Position = p.Cubes[i].PastPosition
				// checking a collision.
				if p.Position == p.Cubes[i+1].Position {
					p.Die()
				}
			}
		}

	}

	// FeedChecker.
	if g.Feed.Position == p.Position {
		for i := 1; int32(i) <= g.Feed.Power; i++ {
			if len(p.Cubes) == 0 {
				p.Cubes = append(p.Cubes, Cube{p.Position, rl.Red, p.Position})
			} else {
				p.Cubes = append(p.Cubes, Cube{p.Cubes[len(p.Cubes)-1].PastPosition, rl.Red, p.Cubes[len(p.Cubes)-1].PastPosition})
			}
		}
		g.Feed.RePlace()
		g.Score += g.Feed.Power
	}
}

func (game *Game) Start() {
	game.GameOver = false
	game.Menu.Showed = false
	game.Feed.RePlace()
	game.Score = 0
	game.Feed.Power = int32(10 + 10*Settings.difficult)
}

func (feed *Feed) RePlace() {
	var filledPlaces = append(getPosOfCubes(game.Player.Cubes), game.Player.Position) // positions of cubes + position of head.
	var emptyPlaces []rl.Vector2
	// note: places is variable being declared above methods.
	for _, place := range places {
		if !filled(place, filledPlaces) {
			emptyPlaces = append(emptyPlaces, place)
		}
	}
	feed.Position = emptyPlaces[rand.Intn(len(emptyPlaces))]
}

func (button *Button) IsClicked() bool {
	mousePosition := rl.GetMousePosition()
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) && ((mousePosition.X >= button.Position.X) && (mousePosition.Y >= button.Position.Y) && ((mousePosition.X <= button.Position.X+button.Size.X) && (mousePosition.Y <= button.Position.Y+button.Size.Y))) {
		return true
	}
	return false
}

func filled(place rl.Vector2, filledPlaces []rl.Vector2) bool {
	for _, filledPlace := range filledPlaces {
		if place == filledPlace {
			return true
		}
	}
	return false
}

func getPosOfCubes(cubes []Cube) []rl.Vector2 {
	var positions []rl.Vector2
	for _, cube := range cubes {
		positions = append(positions, cube.Position)
	}
	return positions
}
