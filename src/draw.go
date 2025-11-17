package src

import (
	"os"
	"slices"
	"strconv"
)

var prevLength int

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
		if (prevLength != snake.length){
			GameBoard.CenterText("Score: " + strconv.Itoa(snake.length))	
			prevLength = snake.length
		}
	} else {
		GameBoard.CenterText("Game over!! Your Score: " + strconv.Itoa(snake.length))
	}
	if (promptType != 1){

		for _, el := range parts {
			if slices.Contains(el.value, byte('#')) && snake.length > 6{
				el.value = []byte("\033[36m" + " " + "\033[0m")
			}
			os.Stdout.Write(el.GetSeq())
			os.Stdout.Write(el.GetValue())
		}
	}

	os.Stdout.Write(GameBoard.seq)
	os.Stdout.Write(GameBoard.prompt)
}

func drawFruits(){
	fruits := GameBoard.fruits

	for _, el := range fruits{
		os.Stdout.Write(el.seq)
		os.Stdout.Write(el.value)
	}
}

func PrintBoard(snake *Snake) {
	drawBox(snake, 0)
	drawFruits()
	snake.drawSnake()

}
