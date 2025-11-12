package main

import (
	"Heapless/src"
	//"math"
	
)



var snake src.Snake               // global Snake, stays stack/global
var board [10][10]int


func main(){
	snake.InitWithParts()
	snake.CreateSnake(5, 0)
	for i:= 1; i < 10; i++{
		snake.AddPart(5, i)

	}

	// snake.PrintSnake()

	
	parts := snake.GetParts()
	src.Clear()

	src.PrintBoard(board)
	for i := 0; i < snake.Length(); i++ {
		p := parts[i]
		p.PrintPart()
	}
	src.Reset()
}