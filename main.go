package main


import (
	"Heapless/src"
)

func main(){
	var snake *src.Snake;
	var part src.Part;

	part.X = 3
	part.Y = 4

	src.AddToSnake(&part, &snake)
	src.AddToSnake(&part, &snake)
}