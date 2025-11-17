package src

import (
	"time"
)



func GameLoop(snake *Snake, cleanupDone chan struct{}) {

	Clear()
	keyPresses := StartKeyBoardReader()
	ticker := time.NewTicker(tickSpeed)
	defer ticker.Stop()

	PrintBoard(snake)

	for !IsGameOver {
		snake.IsGameOverCompletely()
		snake.onPowerUp()
		if GameBoard.fruitsLength == 0 {
			snake.AddFruits(3)
		}

		select {
		case <-cleanupDone:
			return

		case key := <-keyPresses:
			snake.UpdateDirection(key)   // endast riktningsbyte

		case <-ticker.C:
			snake.Move()   // faktiskt movement
		}
	}


	drawBox(snake, 0)
	snake.drawSnake()
	drawBox(snake, 1)
	

	time.Sleep(2*time.Second)
}
