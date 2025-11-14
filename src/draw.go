package src

import (
	"os"
)

func (p *Fragment) getTypeOfTail() string {
	head := p.head

	if head.x > p.x {
		return "→"
	} else if head.x < p.x {
		return "←"
	} else if head.y > p.y {
		return "↓"
	}
	return "↑"
}

func (snake *Snake) drawSnake() {
    fragments := snake.GetFragments()
    for _, el := range fragments{
		os.Stdout.Write(el.GetSeq())
		os.Stdout.Write(el.GetValue())
	}

}

func drawBox(board Board) {
	parts := board.GetParts()
    for _, el := range parts{
		os.Stdout.Write(el.GetSeq())
		os.Stdout.Write(el.GetValue())
	}
}


func PrintBoard(board Board, snake *Snake) {
    drawBox(board)
    snake.drawSnake()
    //snake.Debug()
	
}


func (snake *Snake) Debug(){
    println("\nCoords: ", snake.root.x, snake.root.y)
    println("\nValid: " , snake.isValidMove())
}
