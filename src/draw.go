package src

import (
	"os"
	"strconv"
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
	for _, el := range fragments {
		os.Stdout.Write(el.GetSeq())
		os.Stdout.Write(el.GetValue())
	}

}



func drawBox(snake *Snake, promptType int) {
	parts := GameBoard.GetParts()
	os.Stdout.Write(GameBoard.seq)
	os.Stdout.Write([]byte("                "))
	if (IsGameOver && promptType == 0){
		GameBoard.CenterText("GAME OVER!!!")
	} else if (!IsGameOver && promptType == 0){
		GameBoard.CenterText("Score: " + strconv.Itoa(snake.length))
	} else {
		GameBoard.CenterText("Game over!! Your core: " + strconv.Itoa(snake.length))
	}
	if (promptType != 1){

		for _, el := range parts {
			os.Stdout.Write(el.GetSeq())
			os.Stdout.Write(el.GetValue())
		}
	}

	os.Stdout.Write(GameBoard.seq)
	os.Stdout.Write(GameBoard.prompt)
}

func PrintBoard(snake *Snake) {
	drawBox(snake, 0)
	snake.drawSnake()

}
