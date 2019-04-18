package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	"math/rand"
	"strconv"
)

// GUI.
type Button struct {
	position rl.Vector2
	text     string
	size     rl.Vector2
	color    rl.Color
	font     int32
}
type Menu struct {
	buttons    []Button
	showed     bool
	background rl.Color
}

// entities.
type Cube struct {
	position     rl.Vector2
	color        rl.Color
	pastPosition rl.Vector2
}
type Snake struct {
	living       bool
	color        rl.Color
	position     rl.Vector2
	nextPosition rl.Vector2
	pastPosition rl.Vector2
	cubes        []Cube
	speed        int32
}
type Feed struct {
	position rl.Vector2
	color    rl.Color
	power    int32
}

type Game struct {
	// menues.
	menu         Menu
	gameOverMenu Menu

	screenWidth  int32
	screenHeight int32
	nameWindow   string
	showedFPS    bool
	background   rl.Color

	gameOver  bool
	time      int32
	player    Snake
	feed      Feed
	score     int32
	bestScore int32
}

var game Game           // game
var places []rl.Vector2 // all places.
func main() {

	game.init()
	rl.InitWindow(game.screenWidth, game.screenHeight, game.nameWindow)

	rl.SetTargetFPS(20)

	for !rl.WindowShouldClose() {
		game.update()
	}
	rl.CloseWindow()
}

func (g *Game) init() {
	// window initilization.
	g.screenWidth = 500
	g.screenHeight = 500
	g.nameWindow = "Snake"
	g.showedFPS = true
	g.background = rl.Lime

	// game initilization.
	g.gameOver = false
	// fill places(variable).
	for x := 0; int32(x) < g.screenWidth; x += 10 {
		for y := 0; int32(y) < g.screenHeight; y += 10 {
			places = append(places, rl.Vector2{float32(x), float32(y)})
		}
	}
	// snake initilization.
	g.player.color = rl.Pink
	g.player.position = rl.Vector2{float32(g.screenWidth / 2), float32(g.screenHeight / 2)}
	g.player.speed = 1
	// feed initilization.
	g.feed.power = 10
	g.feed.color = rl.Lime
	// menu initilization.
	var bExit = Button{rl.Vector2{0, 200}, "Exit", rl.Vector2{float32(g.screenWidth), 100}, rl.Red, 50}
	var bStart = Button{rl.Vector2{0, 0}, "Play", rl.Vector2{float32(g.screenWidth), 100}, rl.Green, 50}
	var bSettings = Button{rl.Vector2{0, 100}, "Settings", rl.Vector2{float32(g.screenWidth), 100}, rl.Blue, 50}
	g.menu = Menu{
		buttons:    []Button{bStart, bSettings, bExit},
		showed:     true,
		background: rl.Green,
	}
	// gameOverMenu initilization.
	var bMenu = Button{rl.Vector2{float32((g.screenWidth - (g.screenWidth - 20))) / 2, float32((g.screenHeight + 50)) / 2}, "Menu", rl.Vector2{float32(g.screenWidth - 20), 100}, rl.Orange, 50}
	g.gameOverMenu = Menu{
		buttons:    []Button{bMenu},
		showed:     false,
		background: g.background,
	}
}

var addPosition rl.Vector2

func (g *Game) update() {

	if !g.menu.showed && !g.gameOver {
		// game started.
		p := &g.player
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
		// move head.
		p.pastPosition = p.position
		p.nextPosition = rl.Vector2{addPosition.X + p.position.X, addPosition.Y + p.position.Y}
		// teleportation head.
		switch {
		case p.nextPosition.X == -10:
			p.nextPosition.X = float32(g.screenWidth - 10)
		case p.nextPosition.X == float32(g.screenWidth):
			p.nextPosition.X = 0
		case p.nextPosition.Y == -10:
			p.nextPosition.Y = float32(g.screenHeight - 10)
		case p.nextPosition.Y == float32(g.screenHeight):
			p.nextPosition.Y = 0
		}
		p.position = p.nextPosition

		// move body.

		if len(p.cubes) != 0 {
			// first cube.
			cube := &p.cubes[0]
			cube.pastPosition = cube.position
			cube.position = p.pastPosition
			// next cubes.
			if len(p.cubes) > 1 {
				for i := range p.cubes[1:] {
					p.cubes[i+1].pastPosition = p.cubes[i+1].position
					p.cubes[i+1].position = p.cubes[i].pastPosition
					// checking a collision.
					if p.position == p.cubes[i+1].position {
						p.die()
					}
				}
			}

		}

		// feedChecker.
		if g.feed.position == p.position {
			for i := 1; int32(i) <= g.feed.power; i++ {
				if len(p.cubes) == 0 {
					p.cubes = append(p.cubes, Cube{p.position, rl.Red, p.position})
				} else {
					p.cubes = append(p.cubes, Cube{p.cubes[len(p.cubes)-1].pastPosition, rl.Red, p.cubes[len(p.cubes)-1].pastPosition})
				}
			}
			g.feed.rePlace()
			g.score += g.feed.power
		}
	} else if g.menu.showed {
		switch {
		case g.menu.buttons[0].isClicked(): // start.
			g.start()
		case g.menu.buttons[2].isClicked(): // exit.
			rl.CloseWindow()
		}
	} else if g.gameOverMenu.showed {
		switch {
		case g.gameOverMenu.buttons[0].isClicked():
			g.gameOverMenu.showed = false
			g.menu.showed = true
			// returning snake to start.
			g.player.cubes = []Cube{}
			g.player.position = rl.Vector2{float32(g.screenWidth / 2), float32(g.screenHeight / 2)}
		}
	}
	g.draw()
}

func (g *Game) draw() {

	rl.BeginDrawing()
	// showing menu.
	if g.menu.showed {
		g.menu.draw()
	} else if g.gameOver {
		g.player.draw()
		g.gameOverMenu.draw()
		rl.DrawText("SCORE: "+strconv.Itoa(int(g.score)), (game.screenWidth-rl.MeasureText("SCORE: "+strconv.Itoa(int(g.score)), 50))/2, game.screenHeight/2-50, 50, rl.White)
	} else if !g.gameOver {
		// drawing game.
		g.feed.draw()
		g.player.draw()
		// score.

		rl.DrawText("SCORE: "+strconv.Itoa(int(g.score)), (game.screenWidth-rl.MeasureText("SCORE: "+strconv.Itoa(int(g.score)), 20))/2, 0, 20, rl.White)
	}

	// drawing FPS.
	if g.showedFPS {
		rl.DrawText(strconv.Itoa(int(rl.GetFPS())), 0, 0, 20, rl.White)
	}
	rl.ClearBackground(rl.Black)
	rl.EndDrawing()
}

func (menu *Menu) draw() {
	for _, button := range menu.buttons {
		button.draw()
	}
}

func (button *Button) draw() {
	rl.DrawRectangleV(button.position, button.size, button.color)
	rl.DrawText(button.text, int32(button.position.X)+(int32(button.size.X)-rl.MeasureText(button.text, button.font))/2, int32(button.position.Y)+((int32(button.size.Y)-button.font)/2), button.font, rl.White)
}

func (snake *Snake) die() {
	snake.living = false
	game.gameOver = true
	if game.score > game.bestScore {
		game.bestScore = game.score
	}
	game.gameOverMenu.showed = true
}

func (game *Game) start() {
	game.gameOver = false
	game.menu.showed = false
	game.feed.rePlace()
	game.score = 0
}
func (player *Snake) draw() {
	rl.DrawRectangleV(player.position, rl.Vector2{10, 10}, player.color)
	for _, cube := range player.cubes {
		rl.DrawRectangleV(cube.position, rl.Vector2{10, 10}, cube.color)
	}
}
func (feed *Feed) draw() {
	rl.DrawRectangleV(feed.position, rl.Vector2{10, 10}, feed.color)
}

func (feed *Feed) rePlace() {
	var freePlaces []rl.Vector2
	freePlaces = places
	// note: places is variable being declared above main function.
	var elemsToDel []int
	for i, position := range freePlaces {
		for _, cube := range game.player.cubes {
			if cube.position == position {

				elemsToDel = append(elemsToDel, i)
			}
		}
		if game.player.position == position {

			elemsToDel = append(elemsToDel, i)
		}
	}
	for p, i := range elemsToDel {
		freePlaces = append(freePlaces[:i-p], freePlaces[i-p+1:]...)
	}
	feed.position = freePlaces[rand.Intn(len(freePlaces))]
}

func (button *Button) isClicked() bool {
	MousePosition := rl.GetMousePosition()
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) && ((MousePosition.X >= button.position.X) && (MousePosition.Y >= button.position.Y) && ((MousePosition.X <= button.position.X+button.size.X) && (MousePosition.Y <= button.position.Y+button.size.Y))) {
		return true
	}
	return false
}
