package generator

import (
	"fmt"
)

const WordLength = 5

type (
	letter string
	word   [WordLength]letter
)

func (w word) String() string {
	var ret string
	for _, l := range w {
		ret += string(l)
	}
	return ret
}

func mkLetter(s string) (letter, error) {
	if _, ok := hunAlphabet[s]; !ok {
		return "", fmt.Errorf("invalid letter: %s", s)
	}
	return letter(s), nil
}
