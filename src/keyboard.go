package src

import (
	"os"
	"os/exec"
)

func StartKeyBoardReader() <-chan byte {
	ch := make(chan byte)

	// Sätt terminalen i raw mode **innan** goroutinen startas
	cmd := exec.Command("stty", "cbreak", "-echo")
	cmd.Stdin = os.Stdin
	_ = cmd.Run()

	go func() {
		defer func() {
			exec.Command("stty", "-cbreak", "echo").Run()
		}()

		var buf [1]byte
		for {
			_, err := os.Stdin.Read(buf[:])
			if err != nil {
				close(ch)
				return
			}
			ch <- buf[0]
		}
	}()

	return ch
}
