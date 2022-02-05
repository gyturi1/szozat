package main

import (
	"flag"
	"fmt"
	"sort"
	"strings"

	"github.com/gyturi1/szozat/pkg/lib"
	"github.com/gyturi1/szozat/pkg/wordmap"
)

//go run cmd/main.go -g "_ _ cs ö k"  -l "h f ü ny ly t"
func main() {
	initialWord, fixedWordPattern, letters := parseArgs()
	ws := lib.Gen(initialWord, fixedWordPattern, letters)
	ws = wordmap.Filter(ws)
	printResult(ws)
}

func parseArgs() (lib.Word, lib.Word, []lib.Letter) {
	guess := flag.String("g", "_ _ _ _ _", "the last guess ex: _ ö *cs _ k; _ denotes missing letter, * means letter in the rigth position, other are included but not the rigth position")
	availableLetters := flag.String("l", "", "space separated list of letters available")

	flag.Parse()

	return parseFlags(*guess, *availableLetters)

}

func parseFlags(guess, availableLetters string) (lib.Word, lib.Word, []lib.Letter) {
	initialWord, fixedWordPattern, err := lib.ParseGuess(strings.Split(guess, " "))
	if err != nil {
		panic(fmt.Errorf("flag -g: %w", err))

	}

	letters, err := lib.ParseRemainingLetters(strings.Split(availableLetters, " "))
	if err != nil {
		panic(fmt.Errorf("flag -l: %w", err))
	}
	return initialWord, fixedWordPattern, letters
}

func printResult(ws []lib.Word) {
	var converted []string
	for _, w := range ws {
		converted = append(converted, w.String())
	}
	sort.Strings(converted)

	for _, v := range converted {
		fmt.Println(v)
	}

	fmt.Println(len(ws))
}
