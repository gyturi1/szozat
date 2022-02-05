package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/gyturi1/szozat/pkg/lib"
	"github.com/gyturi1/szozat/pkg/wordmap"
)

//go run cmd/main.go -g "_ _ cs *ö k"  -l "h f ü ny ly t"
func main() {
	initialWord, fixedWordPattern, letters := parseArgs()
	ws := lib.Gen(initialWord, fixedWordPattern, letters)
	ws = wordmap.Filter(ws)
	printResult(ws)
}

func parseArgs() (lib.Word, lib.Word, []lib.Letter) {
	guess := flag.String("g", "_ _ _ _ _", `the last guess example: " _ ö *cs _ k": _ denotes missing letter; * means letter in the rigth position; other are included but not the rigth position`)
	availableLetters := flag.String("l", "", "space separated list of letters available, the letter from the guess is appended here")
	ver := flag.Bool("v", false, "prints the version info")
	ex := flag.Bool("e", false, "examples")

	flag.Parse()
	if *ver {
		printVersion()
		os.Exit(0)
	}
	if *ex {
		printExamples()
		os.Exit(0)
	}

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

var commit, version, date string

func printVersion() {
	fmt.Printf("Current build version: %s, commit:%s, date: %s \n", version, commit, date)
}

func printExamples() {
	fmt.Println(`You have the first 3 letter all in rigth position and 10 remaining letter: -g "*k *o *cs _ _" -l " i s t ny dzs gy w q x ú" `)
	fmt.Println(`You have the first 3 letter none of them in rigth position and 14 remaining letter: -g "p a j _ _" -l " i s t ny dzs gy w q v ú cs ű í x" `)
	fmt.Println(`Only the letters from guess and provided after the -f flag are used to construct the new possible words`)
}
