package src

import (
	"strconv"
)

func (s *Snake) UpdateDirection(key byte) {
	switch key {
	case 'w', 'a', 's', 'd':
		// byt inte till motsatt riktning rakt av
		direction = key
	case 'e':
		tickSpeed += 500
	case 'q':
		tickSpeed -= 100
	}
}

func (s *Snake) Move() {
	//PrintBoard(Board, *s)
	var completedMove bool

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
	if completedMove {
		s.shiftFragments()
		PrintBoard(s)
	}

}

func (s *Snake) shiftFragments() {
	prevX, prevY := s.root.x, s.root.y
	prev := s.root
	// capture the sequence (cursor positioning) of the head _before_ it moves
	prevSeq := prev.seq
	// update head seq to its new coordinates
	prev.seq = []byte("\033[" +
		strconv.Itoa(prev.y+YOffset) +
		";" + strconv.Itoa(prev.x+XOffset) + "H")

	for i := 1; i < s.length; i++ {
		p := &s.Fragments[i]
		// shift coordinates (keep previous coordinates in prevX/prevY)
		p.x, p.y, prevX, prevY = prevX, prevY, p.x, p.y
		// shift seq the same way: give p the previous seq and carry p's old seq forward
		p.seq, prevSeq = prevSeq, p.seq
		prev = p
	}

}

func (s *Snake) increaseSnakeLength() {
	head := s.GetHead()
	x, y := head.GetXY()
	y += 2
	s.AddFragment(x, y)
}



func (s *Snake) onPowerUp() {
	for i := 0; i < GameBoard.fruitsLength; i++ {
		elm := GameBoard.fruits[i]

		// Protect against nil root
		if s.root == nil {
			continue
		}
		if s.root.x == elm.x && s.root.y == elm.y {
			s.increaseSnakeLength()
			lastIdx := GameBoard.fruitsLength - 1
			if i != lastIdx {
				GameBoard.fruits[i] = GameBoard.fruits[lastIdx]
			}
			GameBoard.fruits[lastIdx] = Fruit{}
			GameBoard.fruitsLength--
			i--
		}
	}
}

func (s *Snake)IsGameOverCompletely(){
	for i := 0; i < GameBoard.partsLength; i++ {
		elm := GameBoard.parts[i]
		if !elm.obstacle {
			continue
		}

		if s.root.x <= 0 || s.root.x > Size-2 {
			IsGameOver = true
		}

		if s.root.y <= 0 || s.root.y > Size-2 {
			IsGameOver = true
		}
	}
}

func (s *Snake) isValidMove() bool {
	for i := 0; i < GameBoard.partsLength; i++ {
		elm := GameBoard.parts[i]
		if !elm.obstacle {
			continue
		}

		if s.root.x < 0 || s.root.x > Size-1 {
			return false
		}

		if s.root.y < 0 || s.root.y > Size-1 {
			return false
		}
	}

	// dont walk on yourself
    for i := 0; i < s.length; i++ {
        part := &s.Fragments[i]
        if part == s.root { // skip head
            continue
        }
        if s.root.x == part.x && s.root.y == part.y {
            IsGameOver = true
            return false
        }
    }

	return true
}

func (s *Snake) moveUp() bool {
	s.root.y-- // move head up
	if !s.isValidMove() {
		// redo the move
		s.root.y++
		return false
	}
	return true

}

func (s *Snake) moveDown() bool {
	s.root.y++ // move head down
	if !s.isValidMove() {
		// redo the move
		s.root.y--
		return false
	}
	return true

}

func (s *Snake) moveLeft() bool {
	s.root.x-- // move head left
	if !s.isValidMove() {
		// redo the move
		s.root.x++
		return false
	}
	return true

}

func (s *Snake) moveRight() bool {
	s.root.x++ // move head right
	if !s.isValidMove() {
		// redo the move
		s.root.x--
		return false
	}
	return true

}
