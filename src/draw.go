package src

import(
	"os"
	"strconv"
)

func (p *Part) getTypeOfTail() string {
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


func (snake *Snake) drawSnake(x, y int) bool{
    drawn := false
    for i := 0; i < snake.Length(); i++ {
        p := &(snake.GetParts())[i]
        if p.x == x && p.y == y {
            seq := "\033[" + strconv.Itoa(y+TerminalOffset) + ";" + 
                    strconv.Itoa(x+TerminalOffset) + "H"

            os.Stdout.Write([]byte(seq))
            if i == 0 {
                os.Stdout.Write([]byte("\033[32m0\033[0m")) // green head
            } else if p.tail != nil{
                os.Stdout.Write([]byte("\033[34m8\033[0m")) // blue body
            } else {
                format := "\033[31m" + p.getTypeOfTail() +" \033[0m"
                os.Stdout.Write([]byte(format))
            }
            drawn = true
            break
        }
    }
    return drawn
}


func drawBox(board [Size][Size]int, offsetY, offsetX int, info string, snakeX, snakeY int) {
    width := len(board[0])
    height := len(board)

    // Räkna om snake-position till terminalposition
    termSnakeX := offsetX + 1 + snakeX      // +1 för ramen
    termSnakeY := offsetY + 1 + snakeY      // +1 för ramen

    // Header
    seq := "\033[" + strconv.Itoa(offsetY) + ";" + strconv.Itoa(offsetX+(Size/2)-2) + "H"
    os.Stdout.Write([]byte(seq))
    os.Stdout.Write([]byte("                    "))
    os.Stdout.Write([]byte(seq))
    os.Stdout.Write([]byte(info))

    // ─── TOP ───
    for i := 0; i < width+2; i++ {
        y := offsetY + 1
        x := offsetX + i

        seq := "\033[" + strconv.Itoa(y) + ";" + strconv.Itoa(x) + "H"
        os.Stdout.Write([]byte(seq))

        if x == termSnakeX && y == termSnakeY {
            os.Stdout.Write([]byte("0"))
        } else {
            os.Stdout.Write([]byte("─"))
        }
    }

    // │ SIDES │
    for y := 0; y < height; y++ {
        realY := offsetY + 1 + y

        // vänster
        seqLeft := "\033[" + strconv.Itoa(realY) + ";" + strconv.Itoa(offsetX) + "H"
        os.Stdout.Write([]byte(seqLeft))
        if offsetX == termSnakeX && realY == termSnakeY {
            os.Stdout.Write([]byte("0"))
        } else {
            os.Stdout.Write([]byte("│"))
        }

        // höger
        seqRight := "\033[" + strconv.Itoa(realY) + ";" + strconv.Itoa(offsetX+width+1) + "H"
        os.Stdout.Write([]byte(seqRight))
        if offsetX+width+1 == termSnakeX && realY == termSnakeY {
            os.Stdout.Write([]byte("0"))
        } else {
            os.Stdout.Write([]byte("│"))
        }
    }

    // ─── BOTTOM ───
    for i := 0; i < width+2; i++ {
        y := offsetY + height + 1
        x := offsetX + i

        seq := "\033[" + strconv.Itoa(y) + ";" + strconv.Itoa(x) + "H"
        os.Stdout.Write([]byte(seq))

        if x == termSnakeX && y == termSnakeY {
            os.Stdout.Write([]byte("0"))
        } else if i == 0 || i == width+1 {
            os.Stdout.Write([]byte("│"))
        } else {
            os.Stdout.Write([]byte("─"))
        }
    }
}

func (snake *Snake) drawSnakeMov(board [Size][Size]int){
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

func PrintBoard(board [Size][Size]int, snake Snake) {
    var prompt string;
    if (snake.IsGameOver()){
        prompt = "GAME OVER!!!"
    } else {
        prompt = "Score: " + strconv.Itoa(snake.length-1)
    }

    snake.drawSnakeMov(board)
    drawBox(board, TerminalOffset/2, TerminalOffset/2, prompt, snake.root.x, snake.root.y)
}
