package hanabi

import "fmt"

// 卡片颜色种类
var colors []CardColor = []CardColor{RED, GREEN, WHITE, YELLOW, BLUE}

// 游戏面板
type Board struct {
	// 所有的卡片，在玩家抽取后会逐渐减少，当没有卡片时则重新洗牌
	cards Cards
	// 提示指示牌
	noteTokensNum int
	// 生命指示牌
	stormTokensNum int
	// 用户完成的hanabi
	hanabi Hanabi
	// 参加玩家
	players []*Player
	// 已存在的提示信息, {玩家名: { 卡牌index: 卡牌信息 } }
	hints map[string]map[int]Card
}

type Hanabi map[CardColor]int

// 初始化游戏
func (b *Board) Init() {
	fmt.Println("Start init board...")
	b.noteTokensNum = 8
	b.stormTokensNum = 3
	b.initCards()
	b.cards.Shuffle()
	b.initHint()
	b.initHanabi()
	fmt.Println("End init board...")
}

// 初始化所有卡片
func (b *Board) initCards() {
	nums := []int{3, 2, 2, 2, 1}
	for i, num := range nums {
		for num > 0 {
			for _, c := range colors {
				card := Card{i + 1, c}
				b.cards = append(b.cards, card)
			}
			num--
		}
	}
}

// 初始化hanabi
func (b *Board) initHanabi() {
	b.hanabi = map[CardColor]int{}
	for _, c := range colors {
		b.hanabi[c] = 0
	}
}

// 初始化hint
func (b *Board) initHint() {
	b.hints = map[string]map[int]Card{}
}

// 牌堆是否有剩余牌
func (b Board) hasCardLeft() bool {
	return len(b.cards) > 0
}

// 发牌
func (b *Board) deal(p *Player) {
	// 选取牌堆最上方一张牌，发给玩家
	var c Card = b.cards[0]
	fmt.Printf("给玩家(%v)发牌：%v\n", p.Name, c)
	p.cards = append(p.cards, c)
	b.cards = b.cards[1:]
}

// 发牌
func (b *Board) dealWithIdx(p *Player, idx int) {
	// 选取牌堆最上方一张牌，发给玩家
	var c Card = b.cards[0]
	fmt.Printf("给玩家(%v)发牌：%v\n", p.Name, c)
	p.cards[idx] = c
	b.cards = b.cards[1:]
}

// 给所有玩家发牌，默认一人发4张
func (b *Board) DealAll() {
	for i := 0; i < 4; i++ {
		for _, p := range b.players {
			b.deal(p)
		}
	}
}

// 开始游戏
func (b *Board) Start() {
	fmt.Println("正式开始游戏")
	/*
		每个玩家轮流执行动作
		玩家有三个选择
		  1 提示其他玩家的手牌信息
			2 打牌
			3 弃牌
		对于动作1
		  1.1 从提示指示牌中扣除一枚，如果没有提示指示牌，则无法进行此操作
			1.2 指定一名其他玩家，提示其任意一张手牌的花色或者数字
		对于动作2
		  2.1 选择一张手牌打出
			2.2 board判断打出的手牌是否有效
			2.2.1 如果有效则加入hanabi，增加一枚提示指示牌（上限为8枚）
			2.2.2 如果无效则扣除一枚生命指示牌
			2.3 判断游戏终止条件
			2.3.1 如果无生命指示牌，游戏失败
			2.3.2 如果hanabi完成，则游戏成功
			2.3.3 否则，给玩家补牌
		对于动作3
			3.1 选择一张手牌丢弃
			3.2 增加一枚提示指示牌（上限为8枚）
			3.3 给玩家补牌
	*/
	round := 0
	for !b.isOver(){
		for _, p := range b.players {
			round++
			fmt.Println("-------------------------------")
			fmt.Printf("----------- 第%v回合 -----------\n", round)
			fmt.Println("-------------------------------")
			if b.isOver() {
				break
			}
			if !b.hasCardLeft() {
				b.initCards()
			}
			p.Play()
			b.Print()
		}
	}
	fmt.Printf("游戏结束! 一共经历%v回合\n", round)
	b.Print()
}

// 增加提示指示物
func (b *Board) increaseNoteTokenSum() {
	if b.noteTokensNum < 8 {
		b.noteTokensNum++
	}
}

func (b Board) HasNoteToken() bool {
	return b.noteTokensNum > 0
}

// 玩家打出手牌，board检查并处理
func (b *Board) ReceiveCard(p *Player, idx int) {
	/*
		1. 判断该手牌是否有效，如果有效，则
			1.1 加入hanabi
			1.2 增加提示指示物
		2. 如果无效，则
			2.1 扣除生命指示物
		3 移除hints相关信息
		4 给玩家发牌
	*/
	c := p.cards[idx]
	if b.CanAddToHanabi(c) {
		b.hanabi[c.color]++
		b.increaseNoteTokenSum()
	} else {
		b.stormTokensNum--
	}
	delete(b.hints[p.Name], idx)
	b.dealWithIdx(p, idx)
}

// board收到提示
func (b *Board) ReceiveHint(name string, idx int, card Card) {
	_, ok := b.hints[name]
	if !ok {
		b.hints[name] = map[int]Card{}
	}
	b.hints[name][idx] = card
	b.noteTokensNum--
}

// 该牌是否可以加入最终的hanabi
func (b Board) CanAddToHanabi(card Card) bool {
	h, ok := b.hanabi[card.color]
	if ok {
		return h == card.number-1
	}
	return false
}

// 游戏终止条件
func (b Board) isOver() bool {
	return b.stormTokensNum <= 0
}

func (b Board) Print() {
	fmt.Println()
	fmt.Println("==============================================================")
	fmt.Printf("牌堆数量：%v, 提示牌数量：%v, 生命牌数量：%v，\n当前hanabi：%v\n已知提示信息：%v\n", len(b.cards), b.noteTokensNum, b.stormTokensNum, b.hanabi, b.hints)
	for _, p := range b.players {
		p.Print()
	}
	fmt.Println("==============================================================")
	fmt.Println()
}
