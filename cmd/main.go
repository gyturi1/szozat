package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"

	"github.com/gyturi1/szozat/pkg/generator"
	"github.com/gyturi1/szozat/pkg/wordmap"
)

func main() {
	all, guessStrings := parseArgs()
	gs := parseGuesses(guessStrings...)
	i := generator.Input{
		Guesses:   gs,
		ValidWord: validWords(),
	}
	ws, err := generator.Generate(i)
	if err != nil {
		panic(err)
	}
	printResult(ws, all)
}

func parseGuesses(ss ...string) []generator.Guess {
	var ret []generator.Guess
	for _, s := range ss {
		g, err := generator.Parse(s)
		if err != nil {
			panic(err)
		}
		ret = append(ret, g)
	}
	return ret
}

func parseArgs() (bool, []string) {
	v := flag.Bool("v", false, "prints the version info")
	e := flag.Bool("e", false, "examples")
	a := flag.Bool("a", false, "print all results")

	flag.Parse()
	if *v {
		printVersion()
		os.Exit(0)
	}
	if *e {
		printExamples()
		os.Exit(0)
	}

	return *a, flag.Args()

}

const maxresult = 20

func printResult(ws []string, all bool) {
	sort.Strings(ws)

	c := len(ws)
	if !all && len(ws) > maxresult {
		c = maxresult
	}
	for _, w := range ws[:c] {
		fmt.Println(w)
	}

	fmt.Printf("(%d/%d)\n", c, len(ws))
	if len(ws) > c {
		fmt.Println("Use -a flag to see all results")
	}
}

var commit, version, date string

func printVersion() {
	fmt.Printf("Current build version: %s, commit:%s, date: %s \n", version, commit, date)
}

var markersInfo map[string]string = map[string]string{
	string(generator.Gray):   "The letter is Gray",
	string(generator.Orange): "The letter is Orange",
	string(generator.Green):  "The letter is Green",
}

func printExamples() {
	fmt.Println("")
	fmt.Printf("szozat [options] [guess...]\n")
	fmt.Println("")
	fmt.Printf("Guesses are separated with space. A guess is encoded, see examples below.")
	fmt.Println("")
	fmt.Printf("I a guess prefix each letter with one of the markers, meaning of the markers:\n")
	fmt.Printf("\t %v\n", markersInfo)
	fmt.Println("")
	fmt.Printf("Suppose you made the guess 'kocsis', and k is green cs is orange the rest is gray: *k#o?cs#i#s\n")
	fmt.Println("")
	fmt.Printf("Multiple guesses is separeted with space on the command line: szozat *k-o?cs-i-s *k#a#r?cs#ú\n")
	fmt.Println("")
}

func validWords() generator.ValidWord {
	m, err := wordmap.Read(downloadWordMap())

	if err == nil && len(m) > 0 {
		return func(s string) bool {
			_, ok := m[s]
			return ok
		}
	}
	return wordmap.Contains
}

//this will download the wordlist from github/mdanka/szozat/main/src/constants/hungarian-word-letter-list.json
func downloadWordMap() []byte {
	resp, err := http.Get("https://raw.githubusercontent.com/mdanka/szozat/main/src/constants/hungarian-word-letter-list.json")
	if err != nil {
		fmt.Println("Tried downloading fresh word list, but failed, using what i have")
		return nil
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Tried reading fresh word list, but failed, using what i have")
		return nil
	}
	return b
}
