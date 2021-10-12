package main

import rl "github.com/gen2brain/raylib-go/raylib"

func Update(game *Game) {
	// Update the ball velocity 1 tick
	game.BallPos.X = game.BallPos.X + game.BallVel.X
	game.BallPos.Y = game.BallPos.Y + game.BallVel.Y

	// Bounce off side walls
	if game.BallPos.X > WINDOW_W || game.BallPos.X < 0 {
		game.BallPos.X = rl.Clamp(game.BallPos.X, 0, WINDOW_W)
		game.BallVel.X *= -1
	}

	// Bounce off top wall
	if game.BallPos.Y < 0+BOARD_INITIAL_Y {
		game.BallPos.Y = rl.Clamp(game.BallPos.Y, 0+BOARD_INITIAL_Y, WINDOW_H)
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
		norm := rl.Vector2Normalize(rl.Vector2{X: game.PaddleRect.X, Y: game.PaddleRect.Y})
		reflect := Vector2Reflect(game.BallVel, norm)
		game.BallVel.X = reflect.X
		game.BallVel.Y = reflect.Y

		game.BallVel = rl.Vector2Add(game.BallVel,
			rl.Vector2Scale(
				rl.Vector2Normalize(game.BallVel),
				BALL_INC_SPEED))
	}

	// Check for brick collision
	// Also check for win while looping through bricks
	didWin := true
	for i := 0; i < len(game.Bricks); i++ {
		if game.Bricks[i].Exists {
			didWin = false
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
	if didWin {
		transitionToLevelWon(game)
	}
}

// Code taken from raylib source directly
func Vector2Reflect(v, normal rl.Vector2) rl.Vector2 {
	result := rl.Vector2{}
	dotProduct := rl.Vector2DotProduct(v, normal)
	result.X = v.X - (2.0*normal.X)*dotProduct
	result.Y = v.Y - (2.0*normal.Y)*dotProduct
	return result
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

func playNoiseBrick(game *Game) {
	playNoiseBar(game)
}

func playNoiseBar(game *Game) {
	rl.PlaySound(game.PongSound)
}
