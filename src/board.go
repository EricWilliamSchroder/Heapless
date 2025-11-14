package src

import "strconv"


const XOffset int = 1
const YOffset int = 2

type Part struct{
	x, y int
	seq []byte  // for showing the value later
	value []byte
}

type Board struct {
	parts [Size * Size]Part
	length int
}

func CreateBoard() Board {
    var board Board
    board.length = 0
    var val string

    for y := 0; y < Size; y++ {
        for x := 0; x < Size; x++ {

            if x == 0 || x == Size-1 {
                val = "│"
            } else if y == 0 || y == Size-1 {
                val = "─"
            } else {
                val = " "
            }
			
            seq := []byte("\033[" + strconv.Itoa(y+YOffset) + 
						";" + strconv.Itoa(x+XOffset) + "H")
            value := []byte("\033[31m" + val + "\033[0m")

            board.parts[board.length] = Part{
                x:     x+1,
                y:     y+2,
                seq:   seq,
                value: value,
            }

            board.length++
        }
    }

    return board // returneras by value = stack
}



func (part *Part) GetSeq() []byte{
	return part.seq
}

func (part *Part) GetValue() []byte{
	return part.value
}

func (board Board) GetParts() [Size*Size]Part{
	return board.parts
}