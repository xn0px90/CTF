package level

import (
	_ "embed"
	"errors"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"

	"pppordle/game"
)

//go:embed assets/level3_cursed.txt
var level3Word string

func Level3() *Level {
	var word []rune

	// Reference: https://stackoverflow.com/a/26722698
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	normalized, _, _ := transform.String(t, level3Word)

	for _, l := range []rune(normalized) {
		word = append(word, unicode.ToUpper(l))
	}

	var level3Candidates []rune
	for _, r := range word {
		level3Candidates = append(level3Candidates, r)
	}

	g := &game.Game{
		Validator:       level3Validator,
		Level:           3,
		Word:            word,
		Guesses:         1,
		Candidates:      level3Candidates,
		CompleteMessage: "A̷͚̘͚̤̠̼̝͊́̌̋͊͒̚m̶̥̳̃͂̃͊ͅa̶̧̞̺͚̱̔ź̶͉͔̄̋̉̏̍͘ị̶̧̨̱̠̣͈̲̾͌̀̒͂̾̈́͜ñ̴̨͓̖̞̮̙̩̉͆̒̑́͜ǵ̸͈̤̆!̷̨̼͑̈́̄̏͂!̸̢͎͚͔̌͋̈́̂̈́̊͠͠!̴̧͈̫̤̯̓̃̓͗̓͐͛́̕",
	}

	return &Level{
		Number: 3,
		GenerateGame: func() *game.Game {
			return g
		},
	}
}

func level3Validator(game *game.Game, guess []rune) error {
	if string(guess) != level3Word {
		return errors.New("Not cursed enough")
	}

	return nil
}
