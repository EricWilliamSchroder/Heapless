package src

import (
	"math"
	"math/rand"
	"strconv"
)

// rent krasst så skulle man kunna lägga in obstecles här det hade funkat utan problem
type Part struct {
	obstacle   bool
	x, y       int
	seq, value []byte
}

type Fruit struct {
	x, y       int
	seq, value []byte
}

type Board struct {
	parts        [Size * Size]Part
	fruits       [Size * Size]Fruit
	prompt       []byte
	seq          []byte
	partsLength  int
	fruitsLength int
}

func (s *Snake) isGoodSpot(x, y int) bool {
	// Only check the active fragments (up to s.length). Iterating the entire
	// fixed-size array includes zero-value fragments which can wrongly block
	// placement or detection.
	if s.Fragments == nil || s.length == 0 {
		return true
	}
	for i := 0; i < s.length; i++ {
		f := (*s.Fragments)[i]
		if f.x == x && f.y == y {
			return false
		}
		if math.Abs(float64(f.x-x)) <= 1 && math.Abs(float64(f.y-y)) <= 1 {
			return false
		}
	}
	return true
}

// N is the number of fruits randomly placed
func (s *Snake) AddFruits(N int) {
	var seenCords [Size * Size][2]int
	seenIdx := 0

	fruits := []string{"🍓", "🍆", "🍒", "🍉", "🥒"}
	GameBoard.fruitsLength = 0

	for GameBoard.fruitsLength < N {

		xCord := rand.Intn(Size-3) + 1
		yCord := rand.Intn(Size-3) + 1
		cords := [2]int{xCord, yCord}

		if !s.isGoodSpot(xCord, yCord) {
			continue
		}
		if hasBeenSeen(cords, &seenCords) {
			continue
		}

		seenCords[seenIdx] = cords
		seenIdx++

		seq := []byte("\033[" + strconv.Itoa(yCord+YOffset) +
			";" + strconv.Itoa(xCord+XOffset) + "H")
		value := []byte("\033[33m" + fruits[rand.Intn(len(fruits))] + "\033[0m")

		GameBoard.fruits[GameBoard.fruitsLength] = Fruit{
			x:     xCord,
			y:     yCord,
			seq:   seq,
			value: value,
		}

		GameBoard.fruitsLength++
	}
}

func hasBeenSeen(newCords [2]int, seen *[Size * Size][2]int) bool {
	// 'seen' is a pointer to an array; iterate over the dereferenced array.
	for i := 0; i < len(*seen); i++ {
		if (*seen)[i] == newCords {
			return true
		}
	}
	return false
}

func (b *Board) CenterText(prompt string) {
	midX := (Size-len(prompt)) / 2
	seq := []byte("\033[" + strconv.Itoa(YOffset-2) +
		";" + strconv.Itoa(midX+XOffset) + "H")
	value := []byte("\033[34m" + prompt + "\033[0m")

	b.prompt = value
	b.seq = seq
}

func CreateBoard() Board {
	var board Board
	board.partsLength = 0
	var val string
	var obs bool

	text := "Score 0"
	board.CenterText(text)

	for y := 0; y < Size; y++ {
		for x := 0; x < Size; x++ {

			if x == 0 || x == Size-1 {
				val = "│"
				obs = true
			} else if y == 0 || y == Size-1 {
				val = "─"
				obs = true
			} else{
				val = "#"
				obs = false
			}

			seq := []byte("\033[" + strconv.Itoa(y+YOffset) +
				";" + strconv.Itoa(x+XOffset) + "H")
			value := []byte("\033[36m" + val + "\033[0m")

			board.parts[board.partsLength] = Part{
				x:        x,
				y:        y,
				seq:      seq,
				value:    value,
				obstacle: obs,
			}

			board.partsLength++
		}
	}

	return board // returneras by value = stack
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
