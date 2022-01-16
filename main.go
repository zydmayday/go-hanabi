package main

import (
	"fmt"
	"zydmayday/hanabi/hanabi"
)

func main() {
	board := hanabi.Board{}
	board.Init()
	board.Print()

	for i := 0; i < 4; i++ {
		p := hanabi.Player{Name: fmt.Sprintf("%s%d", "player-", i)}
		p.Join(&board)
	}

	board.DealAll()
	board.Print()

	board.Start()
}
