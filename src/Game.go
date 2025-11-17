package src

import (
	"time"
)



func GameLoop(snake *Snake, cleanupDone chan struct{}) {
	const updateSpeed = 1000 * time.Millisecond // move every 150ms
	Clear()
	keyPresses := StartKeyBoardReader()
	ticker := time.NewTicker(updateSpeed)
	defer ticker.Stop()

	PrintBoard(snake)

	for !IsGameOver {
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
