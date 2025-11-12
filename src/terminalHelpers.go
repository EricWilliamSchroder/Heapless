package src

import (
	"fmt"
	"os"
	"strconv"
)

func Clear() {
	os.Stdout.Write([]byte("\033[2J"))
	os.Stdout.Write([]byte("\033[H"))
	os.Stdout.Write([]byte("\033[?25l"))
}

func Reset() {
	os.Stdout.Write([]byte("\033[2J"))
	os.Stdout.Write([]byte("\033[H"))
	os.Stdout.Write([]byte("\033[?25h"))
}

func PrintBoard(board [10][10]int, snake Snake) {

    for y := 0; y < len(board); y++ {
        for x := 0; x < len(board[0]); x++ {
            drawn := false
            for i := 0; i < snake.Length(); i++ {
                p := &(*snake.GetParts())[i]
                if p.x == x && p.y == y {
                    // Sätt ANSI-position innan du ritar
                    seq := "\033[" + strconv.Itoa(y+1) + ";" + strconv.Itoa(x+1) + "H"
                    os.Stdout.Write([]byte(seq))
                    if i == 0 {
                        os.Stdout.Write([]byte("0")) // huvud
                    } else {
                        os.Stdout.Write([]byte("█")) // kropp
                    }
                    drawn = true
                    break
                }
            }
            if !drawn {
                // Skriv bakgrund
                seq := "\033[" + strconv.Itoa(y+1) + ";" + strconv.Itoa(x+1) + "H"
                os.Stdout.Write([]byte(seq))
                os.Stdout.Write([]byte("#"))
            }
        }
    }
}



func (part *Part) PrintPart() {
	// Terminal coordinates are 1-based. Add 1 to both x and y so that
	// logical board coordinates (starting at 0) map correctly to the terminal.
	seq := "\033[" + strconv.Itoa((*part).y+1) + ";" + strconv.Itoa((*part).x+1) + "H"
	os.Stdout.Write([]byte(seq))
	if (*part).index < 1 {
		os.Stdout.Write([]byte("0"))
	} else {
		os.Stdout.Write([]byte("█"))
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
