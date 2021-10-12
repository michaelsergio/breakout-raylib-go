package main

import (
	"strconv"
)
import "github.com/gen2brain/raylib-go/raylib"


func Render (game *Game) {
	var speed = float64(rl.Vector2Length(game.BallVel))
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)
	rl.DrawText("Lives", 10, 0, 20, rl.LightGray)
	rl.DrawText(strconv.Itoa(game.Death), 80, 0, 20, rl.LightGray)
	rl.DrawText("Score", 110, 0, 20, rl.LightGray)
	rl.DrawText(strconv.Itoa(game.Score), 180, 0, 20, rl.LightGray)
	rl.DrawText("Speed", 220, 0, 20, rl.LightGray)
	rl.DrawText(strconv.FormatFloat(speed, 'f', 1, 32), 300, 0, 20, rl.LightGray)
	rl.EndDrawing()

	// Bricks
	for _, brick := range game.Bricks {
		if brick.Exists {
			rl.DrawRectangleRec(brick.Rec, brick.Color)
		}
	}

	// Paddle
	rl.DrawRectangleRec(game.PaddleRect, rl.LightGray)

	// Ball
	rl.DrawCircleV(game.BallPos, BALL_RADIUS, rl.Blue)
}
