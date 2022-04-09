package level

import (
	_ "embed"
	"errors"
	"pppordle/game"
	"reflect"
	"unicode"
)

//go:embed assets/flag2.txt
var flag2 string

var (
	level4CandidateMap = make(map[rune]struct{})
)

func Level4() *Level {
	var flag []rune

	for _, l := range []rune(flag2) {
		flag = append(flag, unicode.ToUpper(l))
	}

	var level4Candidates []rune
	for i := '!'; i <= '~'; i++ {
		level4Candidates = append(level4Candidates, i)
		level4CandidateMap[i] = struct{}{}
	}

	g := &game.Game{
		Validator:       flagValidator,
		Level:           4,
		Word:            flag,
		Guesses:         6,
		Candidates:      level4Candidates,
		CompleteMessage: "Congrats!",
	}

	return &Level{
		Number: 4,
		GenerateGame: func() *game.Game {
			return g
		},
	}
}

func flagValidator(game *game.Game, guess []rune) error {
	var prefix = []rune{'P', 'C', 'T', 'F', '{'}
	var suffix = []rune{'}'}

	if !reflect.DeepEqual(prefix, guess[:len(prefix)]) ||
		!reflect.DeepEqual(suffix, guess[len(guess)-len(suffix):]) {
		return errors.New("Invalid flag")
	}

	if !ContainsAll(level4CandidateMap, guess) {
		return errors.New("Not in character list")
	}

	return nil
}
