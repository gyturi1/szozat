package filter

import (
	"fmt"
	"strconv"
	"strings"
)

// prases the gray letters given from an encoded string of space separated letters list.
func ParseGrayLetters(s string) ([]Marker, error) {
	var ret []Marker
	letters := strings.Split(s, " ")
	for _, l := range letters {
		if len(l) == 0 {
			continue
		}
		if !validLetter(l) {
			return ret, fmt.Errorf("not a valid letter: %s", l)
		}
		ret = append(ret, Marker{Letter: l, M: Gray})
	}
	return ret, nil
}

// parse the orange letters from an encoded string of space separated letters list.
func ParseOrangeLetters(s string) ([]Marker, error) {
	var ret []Marker
	letters := strings.Split(s, " ")
	for _, l := range letters {
		if len(l) == 0 {
			continue
		}
		if !validLetter(l) {
			return ret, fmt.Errorf("not a valid letter: %s", l)
		}
		ret = append(ret, Marker{Letter: l, M: Orange})
	}
	return ret, nil
}

// parse the green letters from an encoded string of pair of letter:index list.
func ParseGreenLetters(s string) ([]Marker, error) {
	var ret []Marker
	pairs := strings.Split(s, " ")
	for _, p := range pairs {
		if len(p) == 0 {
			continue
		}
		letterAndIndex := strings.Split(p, ":")
		if len(letterAndIndex) != 2 {
			return ret, fmt.Errorf("invalid letter and/or index: %s please specify green letter as letter:index, example: ly:2", p)
		}
		l := letterAndIndex[0]
		if !validLetter(l) {
			return ret, fmt.Errorf("not a valid letter: %s", l)
		}
		i := letterAndIndex[1]
		ix, err := strconv.Atoi(i)
		if err != nil {
			return ret, fmt.Errorf("invalid index not a number: %s", i)
		}
		if ix < 1 || ix > 5 {
			return ret, fmt.Errorf("invalid index: %s index must be between 1 and 5 inclusive", i)
		}
		ret = append(ret, Marker{Letter: l, Position: ix - 1, M: Green})
	}
	return ret, nil
}
