package src

import (
	"time"
)



func GameLoop(snake *Snake, cleanupDone chan struct{}) {

	Clear()
	keyPresses := StartKeyBoardReader()
	ticker := time.NewTicker(UpdateSpeed)
	defer ticker.Stop()

	PrintBoard(snake)

	for !IsGameOver{
		snake.onPowerUp()
		
		if (GameBoard.fruitsLength == 0){
			snake.AddFruits(3)
		}

		select {
		case <-cleanupDone:
			return

		case value, ok := <-keyPresses:
			if !ok {
				return
			}
			// Only update direction — no movement here
			snake.Move(value)

		case <-ticker.C:
			// Move every tick using the current direction
			snake.Move(0)
		}
		
	}

	drawBox(snake, 0)
	snake.drawSnake()
	drawBox(snake, 1)
	

	time.Sleep(2*time.Second)
}
