package src

import (
	"os"
	"slices"
	"strconv"
	"math"
)


const MaxSnakeLength = 100
var parts [MaxSnakeLength]Part // global array, guaranteed not on heap
var direction byte = 's'


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
	index int
	head *Part
	tail *Part
}

func (s *Snake) GetParts() *[MaxSnakeLength]Part {
	return s.parts
}

func (s *Snake) GetHead() *Part {
	if s.length == 0 || s.parts == nil {
		return nil
	}
	return &(*s.parts)[s.length-1]
}


func (p *Part) GetTail() *Part{
	return p.tail
}

func (s *Snake) Length()int{
	return s.length
}

func (s *Snake) InitWithParts(){
	s.parts = &parts
	s.length = 0
	s.root = nil
}

func (s *Snake) CreateSnake(initX, initY int){
	(*s.parts)[0] = Part{x : initX, y : initY, head : &(*s.parts)[0], index: 0}
	s.root = &(*s.parts)[0]
	s.length = 1
}

func (p *Part) GetXY() (int, int){
	return (*p).x, (*p).y
}


func (s *Snake) AddPart(x, y int) {
	

	if s.length >= len(*s.parts) {
		return
	}

	newPart := &(*s.parts)[s.length]
	newPart.x = x
	newPart.y = y
	newPart.index = s.length
	newPart.head = s.root
	newPart.tail = nil

	// sätt tail på förra sista delen
	if s.length > 0 {
		prevTail := &(*s.parts)[s.length-1]
		prevTail.tail = newPart
	}

	s.length++
}

func (s *Snake) Move(button byte, board [Size][Size]int){
	legalKeyPresses := []byte{'w', 'a', 'd', 's', 'q', 0}
	isLegalButton := slices.Contains(legalKeyPresses, button)
	
	if (!isLegalButton){return}

	if s.length == 0 {return}

	prevX, prevY := s.root.x, s.root.y
	completedMove := false

	if button == legalKeyPresses[4] {
		s.increaseSnakeLength()
		button = direction
	} 
	if button != 0 && direction != button {
		direction = button
		button = 0
	}

	switch direction {
		case 'w':
			completedMove = s.moveUp()
		case 's':
			completedMove = s.moveDown()
		case 'd':
			completedMove = s.moveRight()
		case 'a':
			completedMove = s.moveLeft()
		case 27:
			return

	}

	
	if (completedMove){
		for i := 1; i < s.length; i++ {
			p := &s.parts[i]
			p.x, p.y, prevX, prevY = prevX, prevY, p.x, p.y
		}
	}
		
	PrintBoard(board, *s)
}


// TODO: Fix smarter adding of part right now the parts just adds to the positive Y axis

func (s *Snake) increaseSnakeLength(){
	head := s.GetHead()
	x, y := head.GetXY()
	y++ 
	s.AddPart(x, y)
}

func (s *Snake) isValidMove() bool{
	sideF := math.Sqrt(MaxSnakeLength)
	side := int(sideF)

	if (s.root.y >= side || s.root.y < 0){return false}
	if (s.root.x >= side || s.root.x < 0){return false}

	return true
	
}

func (s *Snake) moveUp() bool {
	s.root.y-- // flytta huvudet upp
	if (!s.isValidMove()){
		// redo the move
		s.root.y++
		return false
	}
	return true


}

func (s *Snake) moveDown() bool {
	s.root.y++ 
	if (!s.isValidMove()){
		// redo the move
		s.root.y--
		return false
	}
	return true


}

func (s *Snake) moveLeft() bool {
	s.root.x-- 
	if (!s.isValidMove()){
		// redo the move
		s.root.x++
		return false
	}
	return true


}

func (s *Snake) moveRight() bool {
	s.root.x++ 
	if (!s.isValidMove()){
		// redo the move
		s.root.x--
		return false
	}
	return true

}

func (s *Snake) PrintSnake() {
	for i := 0; i < s.length; i++ {
		p := &(*s.parts)[i]
		seq := "\033[" + strconv.Itoa((*p).y+1) + ";" + strconv.Itoa((*p).x+1) + "H"
		os.Stdout.Write([]byte(seq))
		if (*p).index < 1 {
			os.Stdout.Write([]byte("0"))
		} else {
			os.Stdout.Write([]byte("█"))
		} 
	}
}

