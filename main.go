package main

import (
	"strconv"
)
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


func resetBall(ballPos, ballVel *rl.Vector2) {
	ballPos.X = float32(5)
	ballPos.Y = float32(200)
	ballVel.X = float32(1)
	ballVel.Y = float32(5)
}

func resetBricks(bricks []Brick) {
	for i := 0; i < len(bricks); i++ {
		x := i % BRICKS_PER_ROW
		y := i / BRICKS_PER_ROW
		bricks[i].Exists = true
		bricks[i].Rec.X = BRICK_INITIAL_X + float32(x) * (BRICK_WIDTH + BRICK_PAD_X)
		bricks[i].Rec.Y = BRICK_INITIAL_Y + (float32(y) * (BRICK_HEIGHT + BRICK_PAD_Y))
		bricks[i].Rec.Width = float32(BRICK_WIDTH)
		bricks[i].Rec.Height = float32(BRICK_HEIGHT)
		bricks[i].Color = colors[y % len(colors)]
	}
}

func resetPaddle(paddle *rl.Rectangle, velX *float32) {
	paddle.X = float32(WINDOW_W) / 2 - float32(80)
	paddle.Y = float32(WINDOW_H) - float32(10) - float32(15)
	paddle.Width = float32(80)
	paddle.Height = float32(15)
	*velX = float32(6)
}

func ProcessInput(game *Game) {
	// Paddle Movement
	if rl.IsKeyDown(rl.KeyRight) {
		game.PaddleRect.X += game.PaddleVelX
	} else if rl.IsKeyDown(rl.KeyLeft) {
		game.PaddleRect.X -= game.PaddleVelX
	}
	game.PaddleRect.X = clamp(game.PaddleRect.X, 0.0, WINDOW_W)
}


func Update(game *Game) {

	// Update the ball velocity 1 tick
	game.BallPos.X = game.BallPos.X + game.BallVel.X;
	game.BallPos.Y = game.BallPos.Y + game.BallVel.Y;

	// Bounce off side walls
	if game.BallPos.X > WINDOW_W || game.BallPos.X < 0 {
		game.BallPos.X = clamp(game.BallPos.X, 0, WINDOW_W)
		game.BallVel.X *= -1
	}

	// Check for death
	if game.BallPos.Y < 0 {
		game.BallPos.Y = clamp(game.BallPos.Y, 0, WINDOW_H)
		game.BallVel.Y *= -1
	}

	// Reset to start. Add Death 
	if game.BallPos.Y > WINDOW_H {
		resetBall(&game.BallPos, &game.BallVel)
		game.Death++
	}

	// Check for paddle collision
	if rl.CheckCollisionCircleRec(game.BallPos, BALL_RADIUS, game.PaddleRect) {
		var paddleMid = game.PaddleRect.X + (game.PaddleRect.Width / 2.0)
		var dVelX = collideXVel(paddleMid, game.BallPos.X, game.BallVel.X)
		var changeV = rl.Vector2 {
			X: dVelX,
			Y: float32(-1.0),
		}
		game.BallVel = rl.Vector2Multiply(game.BallVel, changeV)
		// game.BallVel = rl.Vector2Scale(game.BallVel, BALL_INC_SPEED)
		game.BallVel = rl.Vector2Add(game.BallVel,
			rl.Vector2Scale( 
				rl.Vector2Normalize(game.BallVel),
				BALL_INC_SPEED))
	}

	// Check for brick collision
	for i := 0; i < len(game.Bricks); i++ {
		if game.Bricks[i].Exists {
			if rl.CheckCollisionCircleRec(game.BallPos, BALL_RADIUS, game.Bricks[i].Rec) {
				game.Bricks[i].Exists = false
				game.Score += BRICK_SCORE

				var mid = game.Bricks[i].Rec.X + (game.Bricks[i].Rec.Width / 2.0)
				var dVelX = collideXVel(mid, game.BallPos.X, game.BallVel.X)
				var changeV = rl.Vector2 {
					X: dVelX,
					Y: float32(-1.0),
				}
				game.BallVel = rl.Vector2Multiply(game.BallVel, changeV)
				// game.BallVel = rl.Vector2Scale(game.BallVel, BALL_INC_SPEED)
				game.BallVel = rl.Vector2Add(game.BallVel,
					rl.Vector2Scale( 
						rl.Vector2Normalize(game.BallVel),
						BALL_INC_SPEED))
					}
				}
	}
}

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

func clamp(x, min, max float32) float32 {
	if x > max {
		return max
	}
	if x < min {
		return min
	}
	return x
}

func main() {
	rl.InitWindow(int32(WINDOW_W), int32(WINDOW_H), "Breakout")
	rl.SetTargetFPS(60)

	totalBricks := BRICKS_PER_ROW * ROWS_OF_BRICKS
	game := Game{} 
	game.Bricks = make([]Brick, totalBricks)

	resetBall(&game.BallPos, &game.BallVel)
	resetPaddle(&game.PaddleRect, &game.PaddleVelX)
	resetBricks(game.Bricks);

	for !rl.WindowShouldClose() {
		ProcessInput(&game)
		Update(&game)
		Render(&game)
	}
	rl.CloseWindow()
}