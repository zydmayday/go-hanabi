package hanabi

import (
	"math/rand"
	"time"
)

// 卡牌颜色
type CardColor int

const (
	RED    = iota
	GREEN  = iota
	WHITE  = iota
	YELLOW = iota
	BLUE   = iota
)

type Card struct {
	number int
	color  CardColor
}

type Cards []Card

// 洗牌
func (c Cards) Shuffle() Cards {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(c), func(i, j int) {
		c[i], c[j] = c[j], c[i]
	})
	return c
}
