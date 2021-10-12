package main

import "github.com/gen2brain/raylib-go/raylib"

type Brick struct {
	Exists bool
	Rec rl.Rectangle
	Color rl.Color
}

type GameMode int
const ( 
	Start GameMode = iota
	Playing
	GameOver
)

type Game struct {
	Lives int
	Score int
	BallPos rl.Vector2
	BallVel rl.Vector2
	PaddleRect rl.Rectangle
	PaddleVelX float32
	Bricks []Brick
	Mode GameMode
	MaxScore int
	SavedGames SavedGames
}

type SavedGames struct {
	MaxScore int
}