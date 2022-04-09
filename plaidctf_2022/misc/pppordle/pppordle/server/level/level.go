package level

import (
	"math/rand"
	"time"

	"pppordle/game"
)

func init() {
	rand.Seed(time.Now().UnixMilli())
}

type Level struct {
	Number       int
	GenerateGame func() *game.Game
}

func ContainsAll[K comparable](m map[K]struct{}, candidates []K) bool {
	for _, c := range candidates {
		if _, ok := m[c]; !ok {
			return false
		}
	} 

	return true
}
