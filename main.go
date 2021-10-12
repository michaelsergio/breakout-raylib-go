package main

import "github.com/gen2brain/raylib-go/raylib"

func ProcessInput(game *Game) {
	// Paddle Movement
	if rl.IsKeyDown(rl.KeyRight) {
		game.PaddleRect.X += game.PaddleVelX
	} else if rl.IsKeyDown(rl.KeyLeft) {
		game.PaddleRect.X -= game.PaddleVelX
	}
	game.PaddleRect.X = clamp(game.PaddleRect.X, 0.0, WINDOW_W)
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