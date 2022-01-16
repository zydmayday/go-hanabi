package hanabi

import "fmt"

type Player struct {
	Name  string
	board Board
	cards Cards
}

// 参与游戏
func (p *Player) Join(b *Board) {
	fmt.Printf("玩家(%v)加入游戏\n", p.Name)
	b.players = append(b.players, p)
	p.board = *b
}

// 玩家的动作
func (p Player) Play() {

}

func (p Player) Print() {
	fmt.Printf("玩家(%v)的手牌: %v\n", p.Name, p.cards)
}
