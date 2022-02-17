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
	guesses    filter.Markers
	greens     filter.Markers
	grays      filter.Markers
	oranges    filter.Markers
}

func main() {
	params := parseArgs()
	wl := wordList(params.isDownlaod)
	var p filter.Predicates
	if len(params.guesses) > 0 {
		p = append(p, params.guesses.ToPredicates()...)
	}
	if len(params.greens) > 0 {
		p = append(p, params.greens.ToPredicates()...)
	}
	if len(params.grays) > 0 {
		p = append(p, params.grays.ToPredicates()...)
	}
	if len(params.oranges) > 0 {
		p = append(p, params.oranges.ToPredicates()...)
	}

	printResult(filter.Filter(wl, p), params.isAll)
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

// parse the command line arguments, an returns the flags as params, and the guesses as a string slice.
func parseArgs() params {
	v := flag.Bool("v", false, "prints the version info")
	e := flag.Bool("e", false, "examples")
	a := flag.Bool("a", false, "print all results")
	d := flag.Bool("d", false, "download word list if new available")
	l := flag.Int("l", 7, "log level can be: https://pkg.go.dev/github.com/rs/zerolog#DebugLevel")
	green := flag.String("g", "", "the green letters encoded as space separated letter:index")
	gray := flag.String("b", "", "the gray letters a space separated letter list")
	orange := flag.String("o", "", "the orange letter space separated letter list")
	flag.Parse()

	level := zerolog.Level(*l)
	zerolog.SetGlobalLevel(level)

	if *v {
		printVersion()
		os.Exit(0)
	}
	if *e {
		printExamples()
		os.Exit(0)
	}

	guesses, err := filter.ParseAll(flag.Args())
	if err != nil {
		panic(err)
	}
	greens, err := filter.ParseGreenLetters(*green)
	if err != nil {
		panic(err)
	}
	oranges, err := filter.ParseOrangeLetters(*orange)
	if err != nil {
		panic(err)
	}
	grays, err := filter.ParseGrayLetters(*gray)
	if err != nil {
		panic(err)
	}
	ret := params{isAll: *a, isDownlaod: *d, guesses: guesses, greens: greens, grays: grays, oranges: oranges}

	log.Debug().Str("flags", fmt.Sprintf("%v", ret)).Msg("ParseArgs")

	return ret
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
