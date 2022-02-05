package lib

//Gen will generate all the possible permutaion. The initial word is the already found letters.
//fixedWordPattern: from the initial word the letters in the good position
func Gen(initialWord Word, fixedWordPattern Word, availableLetters []Letter) []Word {
	iws := []Word{initialWord}
	a := availableLetters
	a = append(a, initialWord.NotEmptyLetters()...)
	ws := Unique(generateAll(iws, a))
	return filterByPattern(ws, fixedWordPattern)
}

func filterByPattern(ws []Word, pattern Word) []Word {
	var ret []Word = make([]Word, 0)
	for _, w := range ws {
		if w.Match(pattern) {
			ret = append(ret, w)
		}
	}
	return ret
}

func generateAll(initialWord []Word, availableLetters []Letter) []Word {
	if len(initialWord) == 0 || len(availableLetters) == 0 {
		return initialWord
	}
	var ret []Word
	finished, unfinished := split(initialWord)
	ret = append(ret, finished...)
	for _, u := range unfinished {
		r := putAllLettersToFirstEmptySlot(availableLetters, u)
		ret = append(ret, generateAll(r, availableLetters)...)
	}
	return ret
}

func split(ws []Word) (finshed []Word, unfinished []Word) {
	unfinished = filter(ws, HasEmptySlot)
	finshed = filter(ws, not(HasEmptySlot))
	return
}

type predicate = func(Word) bool

func not(p predicate) predicate {
	return func(w Word) bool {
		return !p(w)
	}
}

func filter(ws []Word, p predicate) []Word {
	var ret []Word = make([]Word, 0)
	for _, w := range ws {
		if p(w) {
			ret = append(ret, w)
		}
	}
	return ret
}

func putAllLettersToFirstEmptySlot(ls []Letter, w Word) []Word {
	if len(ls) == 0 {
		return []Word{w}
	}
	var ret []Word
	for _, l := range ls {
		ret = append(ret, putLetterToFirstEmptySlot(w, l))
	}
	return ret
}

func putLetterToFirstEmptySlot(w Word, l Letter) Word {
	var ret []Letter
	replaced := false
	for _, p := range w {
		if !replaced && p == EmptyLetterMarker {
			ret = append(ret, l)
			replaced = true
		} else {
			ret = append(ret, p)

		}
	}
	return ToWord(ret)
}
