package main

import (
	"fmt"
)
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

func processStartInput(game *Game) {
	if rl.IsKeyReleased(rl.KeyEnter) {
		transitionToPlayingMode(game)
	}
}
func processGameOverInput(game *Game) {
	if rl.IsKeyReleased(rl.KeyEnter) {
		transitionToNewGame(game)
	}
}


func transitionToPlayingMode(game *Game) {
		game.Mode = Playing
		resetBall(&game.BallPos, &game.BallVel)
}
func transitionToGameOverMode(game *Game) {
		game.Mode = GameOver
		holdBall(&game.BallPos, &game.BallVel)

		var isNewMaxScore = game.Score > game.MaxScore
		if isNewMaxScore {
			game.MaxScore = game.Score
			game.SavedGames.MaxScore = game.Score
			WriteMaxScoreFile(SAVE_GAME_FILE_PATH, game.SavedGames)
		}
}

func transitionToNewGame(game *Game) {
	holdBall(&game.BallPos, &game.BallVel)
	resetPaddle(&game.PaddleRect, &game.PaddleVelX)
	resetBricks(game.Bricks);
	game.Score = 0
	game.Lives = STARTING_LIVES
	game.Mode = Start
}


func main() {
	totalBricks := BRICKS_PER_ROW * ROWS_OF_BRICKS
	game := Game{
		Bricks: make([]Brick, totalBricks),
	} 

	gameSave, err := ReadMaxScoreFile(SAVE_GAME_FILE_PATH)
	if err != nil {
		fmt.Printf("Failed to read file: %v\n", err)
	} else {
		game.MaxScore = gameSave.MaxScore
	}

	rl.InitWindow(int32(WINDOW_W), int32(WINDOW_H), "Breakout")
	rl.SetTargetFPS(60)

	transitionToNewGame(&game)

	for !rl.WindowShouldClose() {
		if game.Mode == Start {
			processStartInput(&game)
		} else if game.Mode == GameOver {
			processGameOverInput(&game)
		}
		// Go this in start and playing mode
		if game.Mode != GameOver {
			ProcessInput(&game)
		}
		Update(&game)
		Render(&game)
	}
	rl.CloseWindow()
}