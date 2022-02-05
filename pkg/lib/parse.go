package lib

import (
	"strings"

	"github.com/pkg/errors"
)

const FixedLetterMarker = "*"

func ParseWords(args []string) (Word, Word, error) {
	var w []Letter
	var fixedWord []Letter
	for _, v := range args {
		p := v
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
	for _, v := range args {
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
