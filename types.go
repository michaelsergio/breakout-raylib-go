package main

import "github.com/gen2brain/raylib-go/raylib"

type Brick struct {
	Exists bool
	Rec rl.Rectangle
	Color rl.Color
}

type Game struct {
	Death int
	Score int
	BallPos rl.Vector2
	BallVel rl.Vector2
	PaddleRect rl.Rectangle
	PaddleVelX float32
	Bricks []Brick
}