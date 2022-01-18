package hanabi

import (
	"math/rand"
	"time"
)

// 卡牌颜色
type CardColor string

const (
	RED     CardColor = "红"
	GREEN   CardColor = "绿"
	WHITE   CardColor = "白"
	YELLOW  CardColor = "黄"
	BLUE    CardColor = "蓝"
	UNKNOWN CardColor = "未知"
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

// 卡牌是否包含所有信息
func (c Card) HasFullInfo() bool {
	return c.color != UNKNOWN && c.number > 0
}