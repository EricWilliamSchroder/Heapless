package src

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

const TerminalOffset = 4

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
            seq := "\033[" + strconv.Itoa(y+TerminalOffset) + ";" + strconv.Itoa(x+TerminalOffset) + "H"
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

func drawBox(board [Size][Size]int, offsetY, offsetX int) {
    width := len(board[0])
    height := len(board)

    for i := 0; i < width+3; i++ {
        seq := "\033[" + strconv.Itoa(offsetY+1) + ";" + strconv.Itoa(offsetX+i) + "H"
        os.Stdout.Write([]byte(seq))
        os.Stdout.Write([]byte("─"))
    }

    for y := 0; y < height+1; y++ {
        seqLeft := "\033[" + strconv.Itoa(offsetY+1+y) + ";" + strconv.Itoa(offsetX) + "H"
        os.Stdout.Write([]byte(seqLeft))
        os.Stdout.Write([]byte("│"))

        seqRight := "\033[" + strconv.Itoa(offsetY+1+y) + ";" + strconv.Itoa(offsetX+width+3) + "H"
        os.Stdout.Write([]byte(seqRight))
        os.Stdout.Write([]byte("│"))
    }

    for i := 0; i < width+4; i++ {
        seq := "\033[" + strconv.Itoa(offsetY+height+2) + ";" + strconv.Itoa(offsetX+i) + "H"

        os.Stdout.Write([]byte(seq))

        if i == 0 {
            os.Stdout.Write([]byte("│"))
        } else if i == width+3 {
            os.Stdout.Write([]byte("│"))
        } else {
            os.Stdout.Write([]byte("─"))
        }
    }


}

func PrintBoard(board [Size][Size]int, snake Snake) {
    drawBox(board, 2, 2)
    for y := 0; y < len(board); y++ {
        for x := range len(board[0]) {
            drawn := snake.drawSnake(x, y)
            if !drawn {
                // Skriv bakgrund
                seq := "\033[" + strconv.Itoa(y+TerminalOffset) + ";" + strconv.Itoa(x+TerminalOffset) + "H"
                os.Stdout.Write([]byte(seq))
                os.Stdout.Write([]byte("\033[31m  \033[0m"))
            }
        }
    }
}


func (part *Part) Debug() {
	seq := "\033[" + strconv.Itoa(20) + ";" + strconv.Itoa(20) + "H"
	os.Stdout.Write([]byte(seq))
	fmt.Printf("\033[20;20H%d", part.y)
}
