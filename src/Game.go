package src

import (
	"time"
)

const Size int = 10

func GameLoop(snake *Snake, board [Size][Size]int, cleanupDone chan struct{}) {
    const updateSpeed = 1000 * time.Millisecond // move every 150ms
	Clear()
    keyPresses := StartKeyboardReader()
    ticker := time.NewTicker(updateSpeed)
    defer ticker.Stop()

    PrintBoard(board, *snake)

    for {
        select {
        case <-cleanupDone:
            return

        case value, ok := <-keyPresses:
            if !ok {
                return
            }
            // Only update direction — no movement here
            snake.Move(value, board)

        case <-ticker.C:
            // Move every tick using the current direction
            snake.Move(0, board)
        }
    }
}
