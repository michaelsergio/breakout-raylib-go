package main

import (
	"fmt"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func DrawTextBar(txt string, x int32) {
	rl.DrawText(txt, x, 0, 20, rl.LightGray)
}

func DrawTextBarColor(txt string, x int32, color rl.Color) {
	rl.DrawText(txt, x, 0, 20, color)
}

func DrawTextBarRight(txt string, x int32) {
	rl.DrawText(txt, int32(WINDOW_W)-x, 0, 20, rl.LightGray)
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
	scoreColor := rl.LightGray
	if game.Score > game.MaxScore {
		scoreColor = rl.Gold
	}

	DrawTextBar("Lives", 10)
	DrawTextBar(strconv.Itoa(game.Lives), 70)
	DrawTextBar("Level", 110)
	DrawTextBar(strconv.Itoa(game.Level), 170)
	DrawTextBar("Score", 200)
	DrawTextBarColor(strconv.Itoa(game.Score), 270, scoreColor)

	DrawTextBarRight("Speed", 120)
	DrawTextBarRight(strconv.FormatFloat(speed, 'f', 1, 32), 50)
	DrawTextBarRight("Max Score", 320)
	DrawTextBarRight(strconv.Itoa(game.MaxScore), 200)
}

func Render(game *Game) {
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
	if game.Mode == Start {
		rl.DrawText("Press Enter to start", 200, 200, 40, rl.LightGray)
	} else if game.Mode == WaitBall {
		rl.DrawText("Press Enter to release ball", 120, 200, 40, rl.LightGray)
	} else if game.Mode == GameOver {
		rl.DrawText(strconv.Itoa(game.Score), 180, 0, 20, rl.LightGray)
		var text = fmt.Sprintf("Your Score is %d.", game.Score)
		rl.DrawText(text, 200, 200, 40, rl.LightGray)
		rl.DrawText("Press Enter to start a new game", 120, 260, 32, rl.LightGray)
	} else if game.Mode == GameOverNewMaxScore {
		rl.DrawText("NEW MAX SCORE!", int32(WINDOW_W)/3, 180, 32, rl.LightGray)
		rl.DrawText(strconv.Itoa(game.Score), int32(WINDOW_W)/2-40, 220, 42, rl.LightGray)
		rl.DrawText("Press Enter to start a new game", 120, 280, 32, rl.LightGray)
	}

	rl.EndDrawing()
}
