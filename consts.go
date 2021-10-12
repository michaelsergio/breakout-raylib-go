package main

import rl "github.com/gen2brain/raylib-go/raylib"

var SAVE_GAME_FILE_PATH = "./saved_game.txt"
var AUDIO_FILE_PATH = "./pong.wav"
var WINDOW_W = float32(800)
var WINDOW_H = float32(450)
var BRICK_SCORE = 200

var STARTING_LIVES = 3

var BALL_RADIUS = float32(5.0)
var BALL_INC_SPEED = float32(0.1)
var BALL_START_OFFSET_Y = float32(5)

var BOARD_INITIAL_X = float32(0)
var BOARD_INITIAL_Y = float32(20)

var BRICK_HEIGHT = float32(15)
var BRICK_PAD_X = float32(2)
var BRICK_PAD_Y = float32(3)
var BRICK_INITIAL_X = float32(0)
var BRICK_INITIAL_Y = float32(40)

var BRICKS_PER_ROW = 13
var ROWS_OF_BRICKS = 8
var BRICK_WIDTH = WINDOW_W/float32(BRICKS_PER_ROW) - BRICK_PAD_X

var COLORS_NEON = []rl.Color{
	rl.Pink,
	rl.Red,
	rl.Orange,
	rl.Gold,
	rl.Green,
	rl.Blue,
	rl.Purple,
	rl.SkyBlue,
}

// Classic Palette
var COLORS_CLASSIC = []rl.Color{
	rl.Red,
	rl.Red,
	rl.Orange,
	rl.Orange,
	rl.Green,
	rl.Green,
	rl.Gold,
	rl.Gold,
}

var COLORS_ARKANOID = []rl.Color{
	rl.LightGray,
	rl.Red,
	rl.Gold,
	rl.Gold,
	rl.Blue,
	rl.Blue,
	rl.Magenta,
	rl.Magenta,
	rl.Green,
}

// var COLORS = COLORS_CLASSIC
var COLORS = COLORS_ARKANOID
