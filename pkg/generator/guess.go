package generator

import (
	"fmt"
	"regexp"

	"github.com/pkg/errors"
)

type marker string

const Green marker = "*"
const Orange marker = "?"
const Gray marker = "#"

var ValidMarkers []string = []string{string(Gray), string(Orange), string(Green)}

type Guess struct {
	Word    word
	Markers [WordLength]marker
}

var parser string = fmt.Sprintf(`([\\%s\\%s\\%s]{0,1}\p{Latin}{0,3})`, string(Green), string(Orange), string(Gray))
var re = regexp.MustCompile(parser)

//Parse a guess which must be WordLength letter long, and all letter prefixed with a marker.
func Parse(s string) (Guess, error) {
	submatches := re.FindAllStringSubmatch(s, WordLength)

	if len(submatches) != WordLength {
		return Guess{}, fmt.Errorf("invalid guess length got: %v", submatches)
	}

	var ls []letter
	var ms []marker
	for _, m := range submatches {
		m, l, err := parseSubMatch(m[0])
		if err != nil {
			return Guess{}, errors.Wrap(err, "parse submatch")
		}
		ls = append(ls, l)
		ms = append(ms, m)
	}
	w := (*word)(ls)
	m := (*[WordLength]marker)(ms)
	return Guess{Word: *w, Markers: *m}, nil
}

func parseSubMatch(m string) (marker, letter, error) {
	if len(m) < 2 {
		return "", "", fmt.Errorf("both marker and letter are mandatory")
	}
	marker := m[:1]
	rest := m[1:]

	retM, err := parseMarker(marker)
	if err != nil {
		return "", "", errors.Wrap(err, "parse marker")
	}

	retL, err := mkLetter(rest)
	if err != nil {
		return "", "", errors.Wrap(err, "parse letter")
	}

	return retM, retL, nil
}

//parseMarker parses a marker string
func parseMarker(s string) (marker, error) {
	rs := []rune(s)
	if len(rs) == 0 {
		return "", fmt.Errorf("no marker")
	}
	if len(rs) > 1 {
		return "", fmt.Errorf("invalid marker: %v, valid markers: %v", rs, ValidMarkers)
	}
	r := rs[0]
	switch {
	case marker(r) == Green:
		return Green, nil
	case marker(r) == Gray:
		return Gray, nil
	case marker(r) == Orange:
		return Orange, nil
	default:
		return "", fmt.Errorf("unknown marker: %v valid markers: %v", r, ValidMarkers)
	}
}

//template is a special word containing empty letters
type template struct {
	word
}

func (t template) toWord() (word, error) {
	if hasEmptySlot(t) {
		return word{}, fmt.Errorf("template has empty slot")
	}
	return t.word, nil
}

func hasEmptySlot(t template) bool {
	for _, l := range t.word {
		if l == emptyLetter {
			return true
		}
	}
	return false
}

//mkTemplate keep green letters and turns the rest into Empty marker
func mkTemplate(gs []Guess) template {
	var ret template
	for _, g := range gs {
		for i := 0; i < WordLength; i++ {
			if g.Markers[i] == Green {
				ret.word[i] = g.Word[i]
			}
		}
	}
	for i := 0; i < WordLength; i++ {
		if ret.word[i] == "" {
			ret.word[i] = emptyLetter
		}
	}
	return ret
}

//matchingGreens returns true if all green letters match in the word
func (g Guess) matchingGreens(w word) bool {
	for i := 0; i < WordLength; i++ {
		if g.Markers[i] == Green && g.Word[i] != w[i] {
			return false
		}
	}
	return true
}

//overlapOrange return true if there is any matching orange letter in the word
func (g Guess) overlapOrange(w word) bool {
	for i := 0; i < WordLength; i++ {
		if g.Markers[i] == Orange && g.Word[i] == w[i] {
			return true
		}
	}
	return false
}

//grays returns all gray marked letters from the guess
func (g Guess) grays() letterSet {
	var ret letterSet = make(letterSet, 0)
	for i := 0; i < len(g.Markers); i++ {
		if g.Markers[i] == Gray {
			ret[g.Word[i]] = empty
		}
	}
	return ret
}
