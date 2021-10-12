package main

import "github.com/gen2brain/raylib-go/raylib"

var WINDOW_W = float32(800)
var WINDOW_H = float32(450)
var BRICK_SCORE = 200

var	BALL_RADIUS = float32(5.0)
var BALL_INC_SPEED = float32(0.1)

var BRICK_HEIGHT = float32(15)
var BRICK_PAD_X = float32(2)
var BRICK_PAD_Y = float32(3)
var BRICK_INITIAL_X = float32(0)
var BRICK_INITIAL_Y = float32(20)

var BRICKS_PER_ROW = 13
var ROWS_OF_BRICKS = 8
var BRICK_WIDTH = WINDOW_W / float32(BRICKS_PER_ROW) - BRICK_PAD_X

var COLORS = []rl.Color{
	rl.Pink,
	rl.Red,
	rl.Orange,
	rl.Gold,
	rl.Green,
	rl.Blue,
	rl.Purple,
	rl.SkyBlue,
}
