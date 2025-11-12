package src

import (
	"os"
	"strconv"
)

func Clear(){
	os.Stdout.Write([]byte("\033[2J"))
	os.Stdout.Write([]byte("\033[H"))
	os.Stdout.Write([]byte("\033[?25l"))
}

func Reset(){
	os.Stdout.Write([]byte("\033[2J"))
	os.Stdout.Write([]byte("\033[H"))
	os.Stdout.Write([]byte("\033[?25h"))
}

func PrintBoard(board [10][10]int){
	for i:= 0; i < len(board); i++{
		for range len(board[0]) {
			os.Stdout.Write([]byte("#")) 
		}
		os.Stdout.Write([]byte("\n"))
	}

}

func (part *Part) PrintPart(){

	// fmt.Println((*part).x, (*part).y)
	seq := "\033[" + strconv.Itoa((*part).y) + ";" + strconv.Itoa((*part).x) + "H"
	os.Stdout.Write([]byte(seq))
	os.Stdout.Write([]byte("█")) 
}