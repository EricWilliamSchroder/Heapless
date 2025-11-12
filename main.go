package main

import (
	"Heapless/src"
)





var snake src.Snake               // global Snake, stays stack/global
var board [10][10]int


func gameLoop(){
	// TODO: fix deadlock

	c := make(chan byte)
	go src.GetKeyboardPress(c)
	var value byte
	for {
		value = <- c
		if (value == 'w'){
			part := snake.GetParts()[snake.Length()]
			x ,y  := part.GetXY()
			snake.AddPart(x, y+1)
			parts := snake.GetParts()
			parts[snake.Length()].PrintPart()
			
		} else if (value == 27){
			break
		}
	}
}

func main(){
	snake.InitWithParts()
	snake.CreateSnake(5, 0)
	src.Clear()
	src.PrintBoard(board)
	gameLoop()

	src.Reset()

	// // snake.PrintSnake()

	
	// parts := snake.GetParts()
	// src.Clear()

	// src.PrintBoard(board)
	// for i := 0; i < snake.Length(); i++ {
	// 	p := parts[i]
	// 	p.PrintPart()
	// }
	// src.Reset()
}