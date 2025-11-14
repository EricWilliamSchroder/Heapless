package src

const MaxSnakeLength = Size*Size
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
	
	if s.length > 0 {
		prevTail := &(*s.parts)[s.length-1]
		prevTail.tail = newPart
	}
	
	s.length++
}

// Getters
func (p *Part) GetXY() (int, int){
	return (*p).x, (*p).y
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