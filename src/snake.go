package src

import (
	"fmt"
)


const MaxSnakeLength = 100



// Snake represents the full snake in the game.
// It holds a reference to the root part of the snake (usually the head)
// and can be extended to store additional game-related data such as score or length.
type Snake struct {
	// root is the first segment of the snake, typically the head.
	root *Part
	parts *[MaxSnakeLength]Part
	length int

}


// Part is a basic linked list
// Part represents a single segment of the snake's body.
// Each Part knows its coordinates and optionally links to the head and tail.
// This allows easy traversal and management of the snake body.
type Part struct {
	// x and y are the coordinates of this snake part on the game grid.
	x, y int
	head *Part
	tail *Part
}

func (s *Snake) InitWithParts(preParts *[]Part){
	s.parts	= preParts
	s.length = 0
	s.root = nil
}

func (s *Snake) CreateSnake(initX, initY int){
	(*s.parts)[0] = Part{x : initX, y : initY, head : &(*s.parts)[0]}
	s.root = &(*s.parts)[0]
	s.length = 1
}


func (s *Snake) AddPart(x, y int){
	if (s.length >= len(*s.parts)){
		// TODO: Add victory as the snake cannot be longer
		return
	}

	newPart := &(*s.parts)[s.length]
	newPart.x = x
	newPart.y = y
	newPart.head = s.root
	newPart.tail = nil

	previousTail := &(*s.parts)[s.length - 1]
	previousTail.tail = newPart
	s.length++
}

func (s *Snake) PrintSnake() {
	for i := 0; i < s.length; i++ {
		p := &(*s.parts)[i]
		fmt.Printf("Part %d: x=%d, y=%d\n", i, p.x, p.y)
	}
}

