package main

import (
	"Heapless/src"
	"os"
	"os/signal"
)

var snake src.Snake // global Snake, stays stack/global
var board [10][10]int

var cleanupDone = make(chan struct{})

func checkForInterrupt(signalChan chan os.Signal) {
	<-signalChan
	// Call reset and signal the game loop to stop. Do not exit immediately here
	// so that other cleanup can occur in the main goroutine if needed.
	src.Reset()

	// notify gameLoop to stop
	close(cleanupDone)
}

func gameLoop() {
	src.Clear()
	src.PrintBoard(board, snake)
	keypresses := src.StartKeyboardReader()
	// Print the current head (if any) using the snake API so coordinates
	// are correct and any pointer/value semantics are preserved.
	if h := snake.GetHead(); h != nil {
		h.PrintPart()
	}

	for {
		select {
		case <-cleanupDone:
			return
		case value, ok := <-keypresses:
			if !ok {
				return
			}

			switch value {
			case 'w':
				snake.MoveUp()
				src.PrintBoard(board, snake)
				break
			case 's':
				snake.MoveDown()
				src.PrintBoard(board, snake)
				break

			case 'd':
				snake.MoveRight()
				src.PrintBoard(board, snake)
				break
			case 'a':
				snake.MoveLeft()
				src.PrintBoard(board, snake)
				break

			case 'q':

				head := snake.GetHead()
				x, y := head.GetXY()
				y++ // öka y för att växa nedåt
				snake.AddPart(x, y)

				if snake.Length() > 0 {
					snake.PrintSnake()
				}
				break

			case 27: // ESC
				return
			}
		}
	}
}

func main() {
	snake.InitWithParts()
	snake.CreateSnake(5, 0)

	// run the game loop in the main goroutine so the program doesn't exit
	// immediately. Starting it as a goroutine allows main to return which
	// terminates the process before the loop runs.

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	go checkForInterrupt(signalChan)
	gameLoop()

	src.Reset()

	// // snake.PrintSnake()

	// parts := snake.GetParts()
	//src.Clear()

	//src.PrintBoard(board)
	// for i := 0; i < snake.Length(); i++ {
	// 	p := parts[i]
	// 	p.PrintPart()
	// }
	// src.Reset()
}
