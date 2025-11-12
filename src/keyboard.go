package src


import (
	"os"
	"os/exec"
)

func GetKeyboardPress(c chan byte) {
	// Put terminal in raw mode
	cmd := exec.Command("stty", "cbreak", "-echo")
	cmd.Stdin = os.Stdin
	_ = cmd.Run()

	defer func() {
		exec.Command("stty", "-cbreak", "echo").Run()
	}()

	var buf [1]byte
	
	_, err := os.Stdin.Read(buf[:])
	if err != nil {
		return;
	}

	b := buf[0]

	c <- b
	
	
}