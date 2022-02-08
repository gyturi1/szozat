package generator

import "github.com/pkg/errors"

type ValidWord = func(string) bool

type Input struct {
	Guesses []Guess
	ValidWord
}

func Generate(i Input) ([]string, error) {
	var gens []word
	t := mkTemplate(i.Guesses)
	grays := grays(i.Guesses)

	r := hun_alphabet.asSet().removeAll(grays)
	all := generateAll([]template{t}, r.asSlice())
	for _, a := range all {
		w, err := a.toWord()
		if err != nil {
			return nil, errors.Wrap(err, "generate")
		}
		gens = append(gens, w)
	}

	gens = filterByPreviousGuesses(gens, i.Guesses)
	gens = filterByValidWords(gens, i.ValidWord)

	var ret []string
	for _, g := range gens {
		ret = append(ret, g.String())
	}
	return ret, nil
}

//generateAll generates all the possible words recursively adding all letters to the empty slots in template.
func generateAll(startingPoint []template, availableLetters []letter) []template {
	if len(startingPoint) == 0 || len(availableLetters) == 0 {
		return startingPoint
	}
	var ret []template
	finished, unfinished := split(startingPoint)
	ret = append(ret, finished...)
	for _, u := range unfinished {
		r := putAllLettersToFirstEmptySlot(availableLetters, u)
		ret = append(ret, generateAll(r, availableLetters)...)
	}
	return ret
}

func putAllLettersToFirstEmptySlot(ls []letter, w template) []template {
	if len(ls) == 0 {
		return []template{w}
	}
	var ret []template
	for _, l := range ls {
		ret = append(ret, putLetterToFirstEmptySlot(w, l))
	}
	return ret
}

func putLetterToFirstEmptySlot(t template, l letter) template {
	var w word
	replaced := false
	for i := 0; i < WordLength; i++ {
		w[i] = t.word[i]
		if !replaced && t.word[i] == emptyLetter {
			w[i] = l
			replaced = true
		}
	}
	return template{word: w}
}

func filterByValidWords(ws []word, v ValidWord) []word {
	var ret []word = make([]word, 0)
	for _, w := range ws {
		if v(w.String()) {
			ret = append(ret, w)
		}
	}
	return ret
}

func filterByPreviousGuesses(ws []word, pgs []Guess) []word {
	if len(pgs) == 0 {
		return ws
	}
	var ret []word = make([]word, 0)
	for _, w := range ws {
		var satisfies = true
		for _, p := range pgs {
			satisfies = satisfies && (p.matchingGreens(w) && !p.overlapOrange(w))
			if !satisfies {
				break
			}
		}
		if satisfies {
			ret = append(ret, w)
		}
	}
	return ret
}

func grays(gs []Guess) letterSet {
	var ret letterSet = make(letterSet)
	for _, g := range gs {
		ret = ret.addAll(g.grays())
	}
	return ret
}

func split(ws []template) (finshed []template, unfinished []template) {
	unfinished = filter(ws, hasEmptySlot)
	finshed = filter(ws, not(hasEmptySlot))
	return
}

type predicate = func(template) bool

func not(p predicate) predicate {
	return func(w template) bool {
		return !p(w)
	}
}

func filter(ws []template, p predicate) []template {
	var ret []template
	for _, w := range ws {
		if p(w) {
			ret = append(ret, w)
		}
	}
	return ret
}
