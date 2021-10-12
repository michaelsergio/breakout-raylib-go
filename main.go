package main

import (
	"strconv"
)
import "github.com/gen2brain/raylib-go/raylib"

func clamp(x, min, max float32) float32 {
	if x > max {
		return max
	}
	if x < min {
		return min
	}
	return x
}

var windowW = float32(800)
var windowH = float32(450)
var BrickScore = 200
var incSpeed = float32(1.025)
var	ballRadius = float32(5.0)

func resetBall(ballPos, ballVel *rl.Vector2) {
	ballPos.X = float32(5)
	ballPos.Y = float32(200)
	ballVel.X = float32(1)
	ballVel.Y = float32(5)
}

type Brick struct {
	Exists bool
	Rec rl.Rectangle
	Color rl.Color
}

var bricksPerRow = 13
var brickRows = 8
var totalBricks = bricksPerRow * brickRows
var brickWidth = windowW / float32(bricksPerRow)
var brickHeight = float32(15)
var brickPaddingX = float32(2)
var brickPaddingY = float32(3)
var brickInitialY = float32(20)

var colors = []rl.Color{
	rl.Pink,
	rl.Red,
	rl.Orange,
	rl.Gold,
	rl.Green,
	rl.Blue,
	rl.Purple,
	rl.SkyBlue,
}

func resetBricks(bricks []Brick) {
	for i := 0; i < len(bricks); i++ {
		x := i % bricksPerRow
		y := i / bricksPerRow
		bricks[i].Exists = true
		bricks[i].Rec.X = float32(x) * (brickWidth + brickPaddingX)
		bricks[i].Rec.Y = brickInitialY + (float32(y) * (brickHeight + brickPaddingY))
		bricks[i].Rec.Width = float32(brickWidth)
		bricks[i].Rec.Height = float32(brickHeight)
		bricks[i].Color = colors[y % len(colors)]
	}
}

func resetPaddle(paddle *rl.Rectangle, velX *float32) {
	paddle.X = float32(windowW) / 2 - float32(80)
	paddle.Y = float32(windowH) - float32(10) - float32(15)
	paddle.Width = float32(80)
	paddle.Height = float32(15)
	*velX = float32(6)
}


type Game struct {
	Death int
	Score int
	BallPos rl.Vector2
	BallVel rl.Vector2
	PaddleRect rl.Rectangle
	PaddleVelX float32
}

func main() {
	rl.InitWindow(int32(windowW), int32(windowH), "Breakout")
	rl.SetTargetFPS(60)

	game := Game{} 

	resetBall(&game.BallPos, &game.BallVel)
	resetPaddle(&game.PaddleRect, &game.PaddleVelX)

	var bricks = make([]Brick, totalBricks)
	resetBricks(bricks);

	for !rl.WindowShouldClose() {

		// input 
		if rl.IsKeyDown(rl.KeyRight) {
			game.PaddleRect.X += game.PaddleVelX
		} else if rl.IsKeyDown(rl.KeyLeft) {
			game.PaddleRect.X -= game.PaddleVelX
		}
		game.PaddleRect.X = clamp(game.PaddleRect.X, 0.0, windowW)


		// update
		game.BallPos.X = game.BallPos.X + game.BallVel.X;
		game.BallPos.Y = game.BallPos.Y + game.BallVel.Y;

		if game.BallPos.X > windowW || game.BallPos.X < 0 {
			game.BallPos.X = clamp(game.BallPos.X, 0, windowW)
			game.BallVel.X *= -1
		}

		if game.BallPos.Y < 0 {
			game.BallPos.Y = clamp(game.BallPos.Y, 0, windowH)
			game.BallVel.Y *= -1
		}
		// Reset to start. Add Death 
		if game.BallPos.Y > windowH {
			resetBall(&game.BallPos, &game.BallVel)
			game.Death++
		}

		if rl.CheckCollisionCircleRec(game.BallPos, ballRadius, game.PaddleRect) {
			var paddleMid = game.PaddleRect.X + (game.PaddleRect.Width / 2.0)
			var dVelX = collideXVel(paddleMid, game.BallPos.X, game.BallVel.X)
			var changeV = rl.Vector2 {
				X: dVelX,
				Y: float32(-1.0),
			}
			game.BallVel = rl.Vector2Multiply(game.BallVel, changeV)
			game.BallVel = rl.Vector2Scale(game.BallVel, incSpeed)
		}

		for i := 0; i < len(bricks); i++ {
			if bricks[i].Exists {
				if rl.CheckCollisionCircleRec(game.BallPos, ballRadius, bricks[i].Rec) {
					bricks[i].Exists = false
					game.Score += BrickScore

					var mid = bricks[i].Rec.X + (bricks[i].Rec.Width / 2.0)
					var dVelX = collideXVel(mid, game.BallPos.X, game.BallVel.X)
					var changeV = rl.Vector2 {
						X: dVelX,
						Y: float32(-1.0),
					}
					game.BallVel = rl.Vector2Multiply(game.BallVel, changeV)
					game.BallVel = rl.Vector2Scale(game.BallVel, incSpeed)
				}
			}
		}


		// render
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

		// Paddle
		rl.DrawRectangleRec(game.PaddleRect, rl.LightGray)

		// Bricks
		for _, brick := range bricks {
			if brick.Exists {
				rl.DrawRectangleRec(brick.Rec, brick.Color)
			}
		}

		// Ball
		rl.DrawCircleV(game.BallPos, ballRadius, rl.Blue)
	}
	rl.CloseWindow()
}

func collideXVel(rectMid, ballPosX, ballVelX float32) float32 {
	isLeftSideOfPaddle := ballPosX < rectMid 
	if isLeftSideOfPaddle {
		if ballVelX > 0 { /* make sure it is negative */
			return -1
		}
	} else {
		if ballVelX < 0 { /* make sure it is positive */
			return -1
		}
	}
	return 1
}

