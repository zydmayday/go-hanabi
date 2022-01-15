package hanabi

import "fmt"

// 游戏面板
type Board struct {
	// 所有的卡片，在玩家抽取后会逐渐减少，当没有卡片时则重新洗牌
	cards Cards
	// 提示指示牌
	noteTokensNum int
	// 生命指示牌
	stormTokensNum int
}

// 初始化游戏
func (b *Board) Init() {
	fmt.Println("Start init board...")
	b.noteTokensNum = 8
	b.stormTokensNum = 3
	b.initCards()
	b.cards.Shuffle()
	fmt.Println("End init board...")
}

// 初始化所有卡片
func (b *Board) initCards() {
	nums := []int{3, 2, 2, 2, 1}
	colors := []CardColor{RED, GREEN, WHITE, YELLOW, BLUE}
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

func (b Board) Print() {
	fmt.Printf("%+v\n", b)
}
