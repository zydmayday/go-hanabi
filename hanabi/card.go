package hanabi

import (
	"math/rand"
	"time"
)

// 卡牌颜色
type CardColor string

const (
	RED     = "红"
	GREEN   = "绿"
	WHITE   = "白"
	YELLOW  = "黄"
	BLUE    = "蓝"
	UNKNOWN = "未知"
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
