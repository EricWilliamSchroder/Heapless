package src

import (
	"time"
)

func (s *Snake) IsGameOver() bool {
	side := Size

	if s.root.y >= side || s.root.y <= 0 {
		return true
	}
	if s.root.x >= side || s.root.x <= 0 {
		return true
	}

	return false

}

func GameLoop(snake *Snake, cleanupDone chan struct{}) {
	const updateSpeed = 1 * time.Millisecond // move every 150ms
	Clear()
	keyPresses := StartKeyBoardReader()
	ticker := time.NewTicker(updateSpeed)
	defer ticker.Stop()

	PrintBoard(snake)

	for {
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
}
