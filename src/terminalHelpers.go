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

func (part *Part) Debug() {
	seq := "\033[" + strconv.Itoa(20) + ";" + strconv.Itoa(20) + "H"
	os.Stdout.Write([]byte(seq))
	fmt.Printf("\033[20;20H%d", part.y)
}
