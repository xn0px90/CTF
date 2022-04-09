package level

import (
	_ "embed"
	"errors"
	"fmt"
	"math/rand"
	"pppordle/game"
	"unicode"
)

//go:embed assets/level2_emojis.txt
var level2ValidEmojis string

//go:embed assets/flag1.txt
var flag1 string

var (
	level2Emojis   []rune
	level2EmojiMap = make(map[rune]struct{})
)

func Level2() *Level {
	return &Level{
		Number:       2,
		GenerateGame: level2GenerateGame,
	}
}

func level2GenerateGame() *game.Game {
	loadLevel2Emojis()
	g := game.Game{
		Validator:       level2Validator,
		Level:           2,
		Candidates:      level2Emojis,
		CompleteMessage: fmt.Sprintf("🏁 %s 🏁", flag1),
	}

	word, attempts := emojiGenerator()

	g.Word = word
	g.Guesses = attempts

	return &g
}

func emojiGenerator() (word []rune, attempts int) {
	for i := 0; i < 5; i++ {
		word = append(word, unicode.ToUpper(level2Emojis[rand.Intn(len(level2Emojis))]))
	}
	return word, 6
}

func level2Validator(game *game.Game, guess []rune) error {
	if !ContainsAll(level2EmojiMap, guess) {
		return errors.New("Not in emoji list")
	}

	return nil

}

func loadLevel2Emojis() {
	level2Emojis = []rune(level2ValidEmojis)
	for _, emoji := range level2Emojis {
		level2EmojiMap[emoji] = struct{}{}
	}
}
