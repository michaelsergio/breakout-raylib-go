package main

import rl "github.com/gen2brain/raylib-go/raylib"

func startBall(ballPos, ballVel *rl.Vector2) {
	ballVel.X = float32(1)
	ballVel.Y = float32(5)
}

func holdBall(ballPos, ballVel *rl.Vector2) {
	startY := BOARD_INITIAL_Y + BRICK_INITIAL_Y +
		float32(ROWS_OF_BRICKS)*(BRICK_HEIGHT+BRICK_PAD_Y) +
		BALL_START_OFFSET_Y
	ballPos.X = float32(5)
	ballPos.Y = float32(startY)
	ballVel.X = float32(0)
	ballVel.Y = float32(0)
}

func resetBricks(bricks []Brick) {
	for i := 0; i < len(bricks); i++ {
		x := i % BRICKS_PER_ROW
		y := i / BRICKS_PER_ROW
		bricks[i].Exists = true
		bricks[i].Rec.X = BOARD_INITIAL_X + BOARD_INITIAL_X + float32(x)*(BRICK_WIDTH+BRICK_PAD_X)
		bricks[i].Rec.Y = BOARD_INITIAL_Y + BRICK_INITIAL_Y + (float32(y) * (BRICK_HEIGHT + BRICK_PAD_Y))
		bricks[i].Rec.Width = float32(BRICK_WIDTH)
		bricks[i].Rec.Height = float32(BRICK_HEIGHT)
		bricks[i].Color = COLORS[y%len(COLORS)]
	}
}

func resetPaddle(paddle *rl.Rectangle, velX *float32) {
	paddle.X = float32(WINDOW_W)/2 - float32(80)
	paddle.Y = float32(WINDOW_H) - float32(10) - float32(15)
	paddle.Width = float32(80)
	paddle.Height = float32(15)
	*velX = float32(6)
}
