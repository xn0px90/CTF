package level

import (
	"bufio"
	"bytes"
	_ "embed"
	"errors"
	"math/rand"
	"unicode"

	"pppordle/game"
)

//go:embed assets/level1_wordlist.txt
var level1WordlistBytes []byte

var (
	level1Wordlist     [][]rune
	level1Candidates   []rune
	level1CandidateMap = make(map[rune]struct{})
	level1WordMap      = make(map[string]struct{})
)

func Level1() *Level {
	loadLevel1Wordlist()
	return &Level{
		Number:       1,
		GenerateGame: level1GenerateGame,
	}
}

func loadLevel1Wordlist() {
	scanner := bufio.NewScanner(bytes.NewReader(level1WordlistBytes))
	for scanner.Scan() {
		newWord := []rune(scanner.Text())
		for i := range newWord {
			newWord[i] = unicode.ToUpper(newWord[i])
		}
		level1Wordlist = append(level1Wordlist, newWord)
		level1WordMap[string(newWord)] = struct{}{}
	}

	for i := 'A'; i <= 'Z'; i++ {
		level1Candidates = append(level1Candidates, i)
		level1CandidateMap[i] = struct{}{}
	}
}

func level1GenerateGame() *game.Game {
	g := game.Game{
		Validator:       level1Validator,
		Level:           1,
		Candidates:      level1Candidates,
		CompleteMessage: "Nice!",
	}

	word, attempts := wordGenerator()

	g.Word = word
	g.Guesses = attempts

	return &g
}

func wordGenerator() (word []rune, attempts int) {
	return level1Wordlist[rand.Intn(len(level1Wordlist))], 6
}

func level1Validator(game *game.Game, guess []rune) error {
	if !ContainsAll(level1CandidateMap, guess) {
		return errors.New("Not in character list")
	}

	if !ContainsAll(level1WordMap, []string{string(guess)}) {
		return errors.New("Not in word list")
	}

	return nil
}
