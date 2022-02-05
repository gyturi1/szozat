package lib

import (
	"fmt"
	"strings"
)

const WordLength = 5
const EmptyLetterMarker Letter = "_"
const InvalidLetter Letter = Letter("")

var InvalidWord Word = [WordLength]Letter{"", "", "", "", ""}
var EmptyWord Word = [WordLength]Letter{EmptyLetterMarker, EmptyLetterMarker, EmptyLetterMarker, EmptyLetterMarker, EmptyLetterMarker}

type Letter string
type Word [WordLength]Letter

var hugarian_alphabet [45]Letter = [45]Letter{
	EmptyLetterMarker,
	"a", "á", "b", "c",
	"cs", "d", "dz", "dzs",
	"e", "é", "f", "g",
	"gy", "h", "i", "í",
	"j", "k", "l", "ly",
	"m", "n", "ny", "o",
	"ó", "ö", "ő", "p",
	"q", "r", "s", "sz",
	"t", "ty", "u", "ú",
	"ü", "ű", "v", "w",
	"x", "y", "z", "zs",
}

func (w Word) String() string {
	return strings.Join(asString(w[:]), "")
}

func asString(ls []Letter) []string {
	var ret []string
	for _, l := range ls {
		ret = append(ret, string(l))
	}
	return ret
}

func Unique(intSlice []Word) []Word {
	keys := make(map[Word]bool)
	list := []Word{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func ToWord(s []Letter) Word {
	return *(*[5]Letter)(s)
}

func MkLetter(s string) (Letter, error) {
	p := strings.ToLower(s)
	for _, l := range hugarian_alphabet {
		if string(l) == p {
			return Letter(p), nil
		}
	}
	return InvalidLetter, fmt.Errorf("not a valid Letter %s", s)
}

func HasEmptySlot(w Word) bool {
	for _, l := range w {
		if l == EmptyLetterMarker {
			return true
		}
	}
	return false
}

func (w *Word) NotEmptyLetters() []Letter {
	var ret []Letter = make([]Letter, 0)
	for _, l := range w {
		if l != EmptyLetterMarker {
			ret = append(ret, l)
		}
	}
	return ret
}

func (w *Word) Match(p Word) bool {
	ret := true
	for i, l := range p {
		if l == EmptyLetterMarker {
			continue
		}
		if w[i] != l {
			return false
		}
	}
	return ret
}

func MkWord(s1, s2, s3, s4, s5 string) (Word, error) {
	var err error
	l1, err := MkLetter(s1)
	if err != nil {
		return InvalidWord, err
	}
	l2, err := MkLetter(s2)
	if err != nil {
		return InvalidWord, err
	}
	l3, err := MkLetter(s3)
	if err != nil {
		return InvalidWord, err
	}
	l4, err := MkLetter(s4)
	if err != nil {
		return InvalidWord, err
	}
	l5, err := MkLetter(s5)
	if err != nil {
		return InvalidWord, err
	}
	return Word{l1, l2, l3, l4, l5}, nil
}
