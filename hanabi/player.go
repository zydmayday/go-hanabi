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
	/*
	1. 判断是否有手牌可以打出
		1.1 查看board.hints是否有自己相关的提示，比对当前的board.hanabi是否有合理的牌可以打出
		1.2 如果有，则打出，执行完毕
		1.3 如果没有，则跳转到2
	2. 判断是否有提示指示符
		2.1 如果没有，则跳转到3
		2.2 如果有，则查看board.hints，提示还没有的*有效*信息，如果没有可提示信息，则跳转到3（一般不会没有，所以这部分暂时不考虑），执行完毕
	3. 丢牌相关
		3.1 查看board.hints，看是否有可以丢弃的手牌
		3.2 如果有，则丢弃，执行完毕
		3.3 如果没有，则跳转到4
	4. 随机执行阶段，10%打出，90%弃牌
	*/
}

func (p Player) Print() {
	fmt.Printf("玩家(%v)的手牌: %v\n", p.Name, p.cards)
}
