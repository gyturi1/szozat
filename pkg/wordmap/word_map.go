package wordmap

import (
	"bufio"
	"embed"

	"github.com/gyturi1/szozat/pkg/lib"
)

type e struct{}
type wordmap map[string]e

//go:embed words.txt
var f embed.FS

//193676 is the line count of the words.txt
const wordmapsize = 193_676

var m wordmap

func Filter(ws []lib.Word) []lib.Word {
	var ret []lib.Word = make([]lib.Word, 0)
	for _, w := range ws {
		if containsWord(w) {
			ret = append(ret, w)
		}
	}
	return ret
}

func containsWord(w lib.Word) bool {
	return contains(w.String())
}

func contains(w string) bool {
	_, ok := m[w]
	return ok
}

func init() {
	if m == nil {
		m = make(wordmap, wordmapsize)
		f, err := f.Open("words.txt")
		if err != nil {
			panic(err)
		}
		defer f.Close()

		sc := bufio.NewScanner(f)

		for sc.Scan() {
			m[string(sc.Bytes())] = e{}
		}

		if sc.Err() != nil {
			panic(sc.Err())
		}
	}
}
