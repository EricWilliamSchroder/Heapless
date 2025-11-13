package src

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

const terminalOffset = 4

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
            seq := "\033[" + strconv.Itoa(y+terminalOffset) + ";" + strconv.Itoa(x+terminalOffset) + "H"
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

func PrintBoard(board [Size][Size]int, snake Snake) {
    for y := 0; y < len(board); y++ {
        for x := range len(board[0]) {
            drawn := snake.drawSnake(x, y)
            if !drawn {
                // Skriv bakgrund
                seq := "\033[" + strconv.Itoa(y+terminalOffset) + ";" + strconv.Itoa(x+terminalOffset) + "H"
                os.Stdout.Write([]byte(seq))
                os.Stdout.Write([]byte("\033[31m#\033[0m"))
            }
        }
    }
}


func (part *Part) Debug() {
	seq := "\033[" + strconv.Itoa(20) + ";" + strconv.Itoa(20) + "H"
	os.Stdout.Write([]byte(seq))
	fmt.Printf("\033[20;20H%d", part.y)
}
