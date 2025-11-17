package src

import (
	"slices"
	"strconv"
)

func (s *Snake) Move(button byte) {
	//PrintBoard(Board, *s)
	legalKeyPresses := []byte{'w', 'a', 'd', 's', 'q', 0}
	isLegalButton := slices.Contains(legalKeyPresses, button)

	if !isLegalButton {
		return
	}

	if s.length == 0 {
		return
	}

	prevX, prevY := s.root.x, s.root.y
	prev := s.root
	// capture the sequence (cursor positioning) of the head _before_ it moves
	prevSeq := prev.seq
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
	if completedMove {
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
		PrintBoard(s)
	}

}

func (s *Snake) increaseSnakeLength() {
	head := s.GetHead()
	x, y := head.GetXY()
	y += 2
	s.AddFragment(x, y)
}

func (s *Snake) isValidMove() bool {

	

	for i := 0; i < GameBoard.partsLength; i++ {
		elm := GameBoard.parts[i]
		if !elm.obstacle {
			continue
		}

		if s.root.x < 0 || s.root.x > Size-1{
			IsGameOver = true
			return false
		}

		if s.root.y < 0 || s.root.y > Size-1{
			IsGameOver = true
			return false
		}
	}
	if (s.root.tail != nil){
		if s.root.x == s.root.tail.x && s.root.y == s.root.tail.y{
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
