package src

import "strconv"


var direction byte = 's'

// Snake represents the full snake in the game.
// It holds a reference to the root Fragment of the snake (usually the head)
// and can be extended to store additional game-related data such as score or length.
type Snake struct {
	// root is the first segment of the snake, typically the head.
	root      *Fragment
	Fragments *[MaxSnakeLength]Fragment
	length    int
}

// Fragment is a basic linked list
// Fragment represents a single segment of the snake's body.
// Each Fragment knows its coordinates and optionally links to the head and tail.
// This allows easy traversal and management of the snake body.
type Fragment struct {
	// x and y are the coordinates of this snake Fragment on the game grid.
	x, y  int
	index int
	head  *Fragment
	tail  *Fragment
	seq []byte
	value []byte 
}

func (s *Snake) InitWithFragments() {
	s.Fragments = &Fragments
	s.length = 0
	s.root = nil
}

func (s *Snake) CreateSnake(initX, initY int) {
	(*s.Fragments)[0] = Fragment{x: initX, y: initY, head: &(*s.Fragments)[0], index: 0}
	s.root = &(*s.Fragments)[0]
	
	s.root.seq = []byte("\033[" + strconv.Itoa(initY+YOffset) + 
						";" + strconv.Itoa(initX+XOffset) + "H")
	s.root.value = []byte("\033[31m" + "0" + "\033[0m")
	s.length = 1
}

func (s *Snake) AddFragment(x, y int) {
	if s.length >= len(*s.Fragments) {
		return
	}

	newFragment := &(*s.Fragments)[s.length]
	newFragment.x = x
	newFragment.y = y
	newFragment.index = s.length
	newFragment.head = s.root
	newFragment.tail = nil


	newFragment.seq = []byte("\033[" + strconv.Itoa(y+YOffset) + 
						";" + strconv.Itoa(x+XOffset) + "H")
	newFragment.value = []byte("\033[33m" + "v" + "\033[0m")

	if s.length > 0 {
		prevTail := &(*s.Fragments)[s.length-1]
		prevTail.tail = newFragment
		if (prevTail.index != 0){
			prevTail.value = []byte("\033[32m" + "8" + "\033[0m")
		} else {
			prevTail.value = []byte("\033[31m" + "0" + "\033[0m")
		}
	}

	s.length++
}

// Getters
func (p *Fragment) GetXY() (int, int) {
	return (*p).x, (*p).y
}

func (s *Snake) GetFragments() *[MaxSnakeLength]Fragment {
	return s.Fragments
}

func (s *Snake) GetHead() *Fragment {
	if s.length == 0 || s.Fragments == nil {
		return nil
	}
	return &(*s.Fragments)[s.length-1]
}

func (p *Fragment) GetTail() *Fragment {
	return p.tail
}

func (s *Snake) Length() int {
	return s.length
}

func (f *Fragment) GetValue() []byte {
	return f.value
}


func (f *Fragment) GetSeq() []byte {
	return f.seq
}