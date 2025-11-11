package src


import "os"

func Clear(){
	os.Stdout.Write([]byte("\033[2J"))
	os.Stdout.Write([]byte("\033[?25l"))
}