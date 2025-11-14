package src

import "strconv"

type Part struct {
	obstacle bool
	x, y     int
	seq      []byte // for showing the value later
	value    []byte
}

type Board struct {
	parts  [Size * Size]Part
	length int
}

func CreateBoard() Board {
	var Board Board
	Board.length = 0
	var val string
	var obs bool

	for y := 0; y < Size; y++ {
		for x := 0; x < Size; x++ {

			if x == 0 || x == Size-1 {
				val = "│"
				obs = true
			} else if y == 0 || y == Size-1 {
				val = "─"
				obs = true
			} else {
				val = " "
				obs = false
			}

			seq := []byte("\033[" + strconv.Itoa(y+YOffset) +
				";" + strconv.Itoa(x+XOffset) + "H")
			value := []byte("\033[31m" + val + "\033[0m")

			Board.parts[Board.length] = Part{
				x:        x,
				y:        y,
				seq:      seq,
				value:    value,
				obstacle: obs,
			}

			Board.length++
		}
	}

	return Board // returneras by value = stack
}

func (part *Part) GetSeq() []byte {
	return part.seq
}

func (part *Part) GetValue() []byte {
	return part.value
}

func (Board Board) GetParts() [Size * Size]Part {
	return Board.parts
}
