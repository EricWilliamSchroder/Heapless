package main

import (
	"Heapless/src"
	"os"
	"os/signal"
	
)

var snake src.Snake // global Snake, stays stack/global

var board src.Board

var cleanupDone = make(chan struct{})

func checkForInterrupt(signalChan chan os.Signal) {
	<-signalChan
	// Call reset and signal the game loop to stop. Do not exit immediately here
	// so that other cleanup can occur in the main goroutine if needed.
	src.Reset()

	// notify gameLoop to stop
	close(cleanupDone)
}

func main() {
	snake.InitWithFragments()
	snake.CreateSnake(5, 0)

	// run the game loop in the main goroutine so the program doesn't exit
	// immediately. Starting it as a goroutine allows main to return which
	// terminates the process before the loop runs.

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	go checkForInterrupt(signalChan)
	board = src.CreateBoard()
	src.GameLoop(&snake, board, cleanupDone)



	src.Reset()

}
