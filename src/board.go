package src

import (
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
	fragments := s.Fragments
	for _, f := range fragments{
		if (!(x + 5 < f.x || y - 5 > f.x)){return false}
		if (!(y + 5 < f.y || y - 5 > f.y)){return false}
	}
	return true
}

// N is the number of fruits randomly placed
func (s *Snake) AddFruits(N int) {
    var seenCords [Size * Size][2]int
    seenIdx := 0

    fruits := []string{"🍑", "🍆", "🍒", "🍉", "🥒"}

    GameBoard.fruitsLength = 0

    for GameBoard.fruitsLength < N {

        xCord := rand.Intn(Size-2) + 1
        yCord := rand.Intn(Size-2) + 1
        cords := [2]int{xCord, yCord}

		if (!s.isGoodSpot(xCord, yCord)) {continue}
        if hasBeenSeen(cords, &seenCords) {continue}

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

func hasBeenSeen(newCords [2]int, seen *[Size*Size][2]int) bool {
    for i := range len(seen) {
        if seen[i] == newCords {return true}
    }
    return false
}


func (b *Board) CenterText(prompt string) {
	midX := (Size/2) + len(prompt)
	seq := []byte("\033[" + strconv.Itoa(YOffset-1) +
		";" + strconv.Itoa(midX) + "H")
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
			} else {
				val = " "
				obs = false
			}

			seq := []byte("\033[" + strconv.Itoa(y+YOffset) +
				";" + strconv.Itoa(x+XOffset) + "H")
			value := []byte("\033[31m" + val + "\033[0m")

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


