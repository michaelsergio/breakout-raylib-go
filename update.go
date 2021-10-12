package main

import rl "github.com/gen2brain/raylib-go/raylib"

func Update(game *Game) {
	// Update the ball velocity 1 tick
	game.BallPos.X = game.BallPos.X + game.BallVel.X
	game.BallPos.Y = game.BallPos.Y + game.BallVel.Y

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
		game.Lives--
		if game.Lives <= 0 {
			transitionToGameOverMode(game)
		} else {
			transitionToWaitBall(game)
		}
	}

	// Check for paddle collision
	if rl.CheckCollisionCircleRec(game.BallPos, BALL_RADIUS, game.PaddleRect) {
		playNoiseBar(game)
		var paddleMid = game.PaddleRect.X + (game.PaddleRect.Width / 2.0)
		var dVelX = collideXVel(paddleMid, game.BallPos.X, game.BallVel.X)
		var changeV = rl.Vector2{
			X: dVelX,
			Y: float32(-1.0),
		}
		game.BallVel = rl.Vector2Multiply(game.BallVel, changeV)
		game.BallVel = rl.Vector2Add(game.BallVel,
			rl.Vector2Scale(
				rl.Vector2Normalize(game.BallVel),
				BALL_INC_SPEED))
	}

	// Check for brick collision
	for i := 0; i < len(game.Bricks); i++ {
		if game.Bricks[i].Exists {
			if rl.CheckCollisionCircleRec(game.BallPos, BALL_RADIUS, game.Bricks[i].Rec) {
				playNoiseBrick(game)

				game.Bricks[i].Exists = false
				game.Score += BRICK_SCORE

				var mid = game.Bricks[i].Rec.X + (game.Bricks[i].Rec.Width / 2.0)
				var dVelX = collideXVel(mid, game.BallPos.X, game.BallVel.X)
				var changeV = rl.Vector2{
					X: dVelX,
					Y: float32(-1.0),
				}
				game.BallVel = rl.Vector2Multiply(game.BallVel, changeV)
				game.BallVel = rl.Vector2Add(game.BallVel,
					rl.Vector2Scale(
						rl.Vector2Normalize(game.BallVel),
						BALL_INC_SPEED))
			}
		}
	}
}

// Return 1.0 or -1.0 depending on if the which side of the rect makes contact with the ball
// The return value should be multipled by the x of the ball Velocity to change its movement.
func collideXVel(mid, ballPosX, ballVelX float32) float32 {
	isLeftSideOfRec := ballPosX < mid
	if isLeftSideOfRec {
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

func playNoiseBrick(game *Game) {
	playNoiseBar(game)
}

func playNoiseBar(game *Game) {
	rl.PlaySound(game.PongSound)
}
