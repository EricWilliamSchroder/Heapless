package main


import (
	"Heapless/src"
)


var parts [src.MaxSnakeLength]src.Part // global array, guaranteed not on heap
var snake src.Snake                // global Snake, stays stack/global

func main(){
	snake.InitWithParts(&parts)
	snake.CreateSnake(5, 5)
	snake.AddPart(5, 6)
	snake.AddPart(5, 7)
	snake.PrintSnake()
}