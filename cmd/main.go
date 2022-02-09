package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/gyturi1/szozat/pkg/filter"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type params struct {
	isAll      bool
	isDownlaod bool
}

func main() {
	params, guessStrings := parseArgs()
	p := parseGuesses(guessStrings)
	wl := wordList(params.isDownlaod)
	printResult(p.Filter(wl), params.isAll)
}

// returns the wrodlist from: embedded, cache, or download it if requested.
func wordList(download bool) [][]string {
	log.Info().Msg("Getting wordlist")
	wl, et, err := filter.Embedded()
	if err != nil {
		panic(err)
	}
	log.Debug().Msg("Loaded from embedded")

	cachedWordList, e2 := filter.LatestCached()
	if len(cachedWordList) > 0 {
		wl = cachedWordList
		et = e2
		log.Debug().Str("etag", et).Msg("Loaded from cache")
	}

	if download {
		d, err := filter.Download(et)
		if err != nil {
			panic(err)
		}
		if len(d) > 0 {
			wl = d
		}
		log.Debug().Msg("Downloaded")
	}
	log.Debug().Int("wl.length", len(wl)).Msg("")
	return wl
}

// Parse the guesses given as command line argument.
func parseGuesses(s []string) filter.Markers {
	log.Info().Msg("Parsing guesses")
	p, err := filter.ParseAll(s)
	if err != nil {
		panic(err)
	}
	return p
}

// parse the command line arguments, an returns the flags as params, and the guesses as a string slice.
func parseArgs() (params, []string) {
	v := flag.Bool("v", false, "prints the version info")
	e := flag.Bool("e", false, "examples")
	a := flag.Bool("a", false, "print all results")
	d := flag.Bool("d", false, "download word list if new available")
	l := flag.Int("l", 7, "log level can be: https://pkg.go.dev/github.com/rs/zerolog#DebugLevel")
	flag.Parse()

	level := zerolog.Level(*l)
	zerolog.SetGlobalLevel(level)

	ret := params{isAll: *a, isDownlaod: *d}
	args := flag.Args()

	log.Debug().Str("flags", fmt.Sprintf("%v", ret)).Msg("ParseArgs")
	log.Debug().Str("guesses", fmt.Sprintf("%v", args)).Msg("ParseArgs")

	if *v {
		printVersion()
		os.Exit(0)
	}
	if *e {
		printExamples()
		os.Exit(0)
	}

	return ret, args
}

const maxresult = 20

func printResult(wl filter.Wordlist, all bool) {
	log.Info().Msg("Printing results")
	var ws []string
	for _, w := range wl {
		ws = append(ws, strings.Join(w, ""))
	}
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

	if len(ws) == 0 {
		fmt.Println("No result.")
		fmt.Println("Maybe old wordlist. Use -d to update the wordlist")
	}
}

var commit, version, date string

func printVersion() {
	fmt.Printf("Current build version: %s, commit:%s, date: %s \n", version, commit, date)
}

var markersInfo = map[string]string{
	string(filter.Gray):   "The letter is Gray",
	string(filter.Orange): "The letter is Orange",
	string(filter.Green):  "The letter is Green",
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
	fmt.Printf("Suppose you made the guess 'kocsis', and k is green cs is orange the rest is gray: \"%sk%so%scs%si%ss\"\n", filter.Green, filter.Gray, filter.Orange, filter.Gray, filter.Gray)
	fmt.Println("")
	fmt.Printf("Multiple guesses is separeted with space on the command line: szozat \"guess1\" \"guess2\"\n")
	fmt.Println("")
}
