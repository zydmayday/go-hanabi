package hanabi

import (
	"fmt"
	"math/rand"
)

type Player struct {
	Name  string
	board *Board
	cards Cards
}

// 参与游戏
func (p *Player) Join(b *Board) {
	fmt.Printf("玩家(%v)加入游戏\n", p.Name)
	b.players = append(b.players, p)
	p.board = b
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
	if idx := p.canPlayCard(); idx != -1 {
		p.PlayCard(idx)
	} else if p.board.HasNoteToken() {
		name, idx, card := p.giveHint()
		p.board.ReceiveHint(name, idx, card)
	} else {
		// 随机出牌
		p.PlayCard(rand.Intn(4))
	}
}

// 出牌
func (p Player) PlayCard(idx int) {
	fmt.Printf("玩家(%v)打出第%v张手牌\n", p.Name, idx+1)
	p.board.ReceiveCard(&p, idx)
}

// 是否可以出牌
func (p Player) canPlayCard() int {
	cards, ok := p.board.hints[p.Name]
	if ok {
		for idx, card := range cards {
			if card.HasFullInfo() && p.board.CanAddToHanabi(card) {
				return idx
			}
		}
	}
	return -1
}

// 给出提示
func (p Player) giveHint() (name string, idx int, card Card) {
	// 这里做一个简化，每次都提示完整的卡牌信息
	// 提示先做成随机
	// TODO: 之后把逻辑弄复杂
	for {
		playerIdx := rand.Intn(4)
		if p.board.players[playerIdx].Name != p.Name {
			cardIdx := rand.Intn(4)
			targetPlayer := p.board.players[playerIdx]
			fmt.Printf("玩家(%v)给出提示： 提示玩家(%v)的第%v张手牌为%v\n", p.Name, targetPlayer.Name, cardIdx+1, card)
			return targetPlayer.Name, cardIdx, targetPlayer.cards[cardIdx]
		}
	}
}

func (p Player) Print() {
	fmt.Printf("玩家(%v)的手牌: %v\n", p.Name, p.cards)
}
