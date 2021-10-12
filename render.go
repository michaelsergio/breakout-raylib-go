package main

import (
	"strconv"
	"fmt"
)
import "github.com/gen2brain/raylib-go/raylib"


func DrawTextBar(txt string, x int32) {
	rl.DrawText(txt, x, 0, 20, rl.LightGray)
}

func DrawTextBarRight(txt string, x int32) {
	rl.DrawText(txt, int32(WINDOW_W) - x, 0, 20, rl.LightGray)
}

func DrawBricks(bricks []Brick) {
	for _, brick := range bricks {
		if brick.Exists {
			rl.DrawRectangleRec(brick.Rec, brick.Color)
		}
	}
}

func DrawUI(game *Game) {
	var speed = float64(rl.Vector2Length(game.BallVel))

	DrawTextBar("Lives", 10)
	DrawTextBar(strconv.Itoa(game.Lives), 80)
	DrawTextBar("Score", 110)
	DrawTextBar(strconv.Itoa(game.Score), 180)

	DrawTextBarRight("Speed", 120)
	DrawTextBarRight(strconv.FormatFloat(speed, 'f', 1, 32), 50)
	DrawTextBarRight("Max Score", 320)
	DrawTextBarRight(strconv.Itoa(game.MaxScore), 200)
}

func Render (game *Game) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)


	// Always Draw the main UI, regardless of GameMode
	
	DrawUI(game)
	DrawBricks(game.Bricks)
	
	// Paddle
	rl.DrawRectangleRec(game.PaddleRect, rl.LightGray)

	// Ball
	rl.DrawCircleV(game.BallPos, BALL_RADIUS, rl.Blue)

	// Game Mode specific Drawing
	if (game.Mode == Start) {
		rl.DrawText("Press Enter to start", 200, 200, 40, rl.LightGray)
	} else if (game.Mode == GameOver) {
		rl.DrawText(strconv.Itoa(game.Score), 180, 0, 20, rl.LightGray)
		var text = fmt.Sprintf("Your Score is %d.", game.Score)
		rl.DrawText(text, 200, 200, 40, rl.LightGray)
		rl.DrawText("Press Enter to start a new game", 120, 260, 32, rl.LightGray)
	}

	rl.EndDrawing()
}
