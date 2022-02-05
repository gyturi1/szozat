package lib

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

const FixedLetterMarker = "*"

func ParseGuess(args []string) (Word, Word, error) {
	var w []Letter = make([]Letter, 0)
	var fixedWord []Letter = make([]Letter, 0)
	if len(args) == 0 {
		return Word{}, Word{}, nil
	}
	if len(args) != WordLength {
		return Word{}, Word{}, fmt.Errorf("not valid length: %s required %d", args, WordLength)
	}
	for _, v := range args {
		p := strings.Trim(v, " ")
		fixed := false
		if strings.HasPrefix(v, FixedLetterMarker) {
			p = strings.TrimPrefix(v, FixedLetterMarker)
			fixed = true
		}

		l, err := MkLetter(p)
		if err != nil {
			return InvalidWord, InvalidWord, errors.Wrap(err, "parse words")
		}
		w = append(w, l)
		if fixed {
			fixedWord = append(fixedWord, l)
		} else {
			fixedWord = append(fixedWord, EmptyLetterMarker)
		}
	}

	return ToWord(w), ToWord(fixedWord), nil
}

func ParseRemainingLetters(args []string) ([]Letter, error) {
	var letters []Letter = make([]Letter, 0)
	if len(args) == 0 {
		return letters, nil
	}
	for _, v := range args {
		v = strings.Trim(v, " ")
		if len(v) == 0 {
			continue
		}
		l, err := MkLetter(v)
		if err != nil {
			return letters, errors.Wrap(err, "parse available")
		}
		if l != EmptyLetterMarker {
			letters = append(letters, l)
		}
	}
	return letters, nil
}
