package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func ProcessInput(game *Game) {
	// Paddle Movement
	if rl.IsKeyDown(rl.KeyRight) {
		game.PaddleRect.X += game.PaddleVelX
	} else if rl.IsKeyDown(rl.KeyLeft) {
		game.PaddleRect.X -= game.PaddleVelX
	}
	game.PaddleRect.X = rl.Clamp(game.PaddleRect.X, 0.0, WINDOW_W-game.PaddleRect.Width)
}

func processStartInput(game *Game) {
	if rl.IsKeyReleased(rl.KeyEnter) {
		transitionToPlayingMode(game)
	}
}

func processWaitBallInput(game *Game) {
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

func transitionToWaitBall(game *Game) {
	game.Mode = WaitBall
	holdBall(&game.BallPos, &game.BallVel)
}

func transitionToLevelWon(game *Game) {
	game.Mode = WaitBall
	holdBall(&game.BallPos, &game.BallVel)
	resetBricks(game.Bricks)
	game.Level++
}

func transitionToGameOverNewMaxScoreMode(game *Game) {
	game.Mode = GameOverNewMaxScore
	game.MaxScore = game.Score
	game.SavedGames.MaxScore = game.Score
	WriteMaxScoreFile(SAVE_GAME_FILE_PATH, game.SavedGames)
}

func transitionToGameOverMode(game *Game) {
	game.Mode = GameOver
	holdBall(&game.BallPos, &game.BallVel)

	var isNewMaxScore = game.Score > game.MaxScore
	if isNewMaxScore {
		transitionToGameOverNewMaxScoreMode(game)
	}
}

func transitionToNewGame(game *Game) {
	holdBall(&game.BallPos, &game.BallVel)
	resetPaddle(&game.PaddleRect, &game.PaddleVelX)
	resetBricks(game.Bricks)
	game.Score = 0
	game.Lives = STARTING_LIVES
	game.Level = 1
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

	rl.InitAudioDevice()
	defer rl.CloseAudioDevice()

	pongSound := rl.LoadSound(AUDIO_FILE_PATH)
	defer rl.UnloadSound(pongSound)
	game.PongSound = pongSound

	rl.InitWindow(int32(WINDOW_W), int32(WINDOW_H), "Breakout")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	transitionToNewGame(&game)

	for !rl.WindowShouldClose() {
		if game.Mode == Start {
			processStartInput(&game)
		} else if game.Mode == WaitBall {
			processWaitBallInput(&game)
		} else if game.Mode == GameOver || game.Mode == GameOverNewMaxScore {
			processGameOverInput(&game)
		}
		// Go this in start and playing mode
		if game.Mode != GameOver {
			ProcessInput(&game)
		}
		Update(&game)
		Render(&game)
	}
}
