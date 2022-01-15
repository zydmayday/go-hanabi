package main

import (
	"zydmayday/hanabi/hanabi"
)

func main() {
	board := hanabi.Board{}
	board.Init()
	board.Print()
}
