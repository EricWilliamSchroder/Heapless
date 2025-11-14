package src

import "slices"

func (s *Snake) Move(button byte, board [Size][Size]int){
	//PrintBoard(board, *s)
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


func (s *Snake) increaseSnakeLength(){
	head := s.GetHead()
	x, y := head.GetXY()
	y++ 
	s.AddPart(x, y)
}

func (s *Snake) isValidMove() bool{
	side := Size

	if (s.root.y >= side+1 || s.root.y < -1){return false}
	if (s.root.x >= side+1 || s.root.x < -1){return false}

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