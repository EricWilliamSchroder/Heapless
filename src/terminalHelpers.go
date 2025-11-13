package src

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

func Clear() {
	os.Stdout.Write([]byte("\033[2J"))
	os.Stdout.Write([]byte("\033[H"))
	os.Stdout.Write([]byte("\033[?25l"))
}

func Reset() {
	exec.Command("stty", "sane").Run()
	os.Stdout.Write([]byte("\033[?25h\033[0m\033[2J\033[H"))

}

func (snake *Snake) drawSnake(x, y int) bool{
    drawn := false
    for i := 0; i < snake.Length(); i++ {
        p := &(snake.GetParts())[i]
        if p.x == x && p.y == y {
            // Sätt ANSI-position innan du ritar
            seq := "\033[" + strconv.Itoa(y+1) + ";" + strconv.Itoa(x+1) + "H"
            os.Stdout.Write([]byte(seq))
            if i == 0 {
                os.Stdout.Write([]byte("\033[32m0\033[0m")) // green head
            } else {
                os.Stdout.Write([]byte("\033[34m█\033[0m")) // blue body
            }
            drawn = true
            break
        }
    }
    return drawn
}

func PrintBoard(board [10][10]int, snake Snake) {
    for y := 0; y < len(board); y++ {
        for x := range len(board[0]) {
            drawn := snake.drawSnake(x, y)
            if !drawn {
                // Skriv bakgrund
                seq := "\033[" + strconv.Itoa(y+1) + ";" + strconv.Itoa(x+1) + "H"
                os.Stdout.Write([]byte(seq))
                os.Stdout.Write([]byte("\033[31m#\033[0m"))
            }
        }
    }
}




func (part *Part) ReplaceWithHash(x, y int){
	seq := "\033[" + strconv.Itoa(x) + ";" + strconv.Itoa(y) + "H"
	os.Stdout.Write([]byte(seq))
	os.Stdout.Write([]byte("#"))

}

func (part *Part) Debug() {
	seq := "\033[" + strconv.Itoa(20) + ";" + strconv.Itoa(20) + "H"
	os.Stdout.Write([]byte(seq))
	fmt.Printf("\033[20;20H%d", part.y)
}
