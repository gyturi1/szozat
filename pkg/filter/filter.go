package filter

import (
	"github.com/rs/zerolog/log"
)

// simple aliases.
type (
	Letter   = string
	Word     = []Letter
	Wordlist = []Word
	Alphabet = []Letter
)

// A predicate that tests whether a word matches.
type predicate = func(Word) bool

// Predicate slice.
type Predicates []predicate

// Filter the provided wordlist for the Markers.
func Filter(wl Wordlist, ps Predicates) Wordlist {
	log.Info().Str("method", "Filter()").Msg("")
	var ret Wordlist
	for _, w := range wl {
		if ps.all(w) {
			ret = append(ret, w)
		}
	}
	return ret
}

// all runs all the predicates for the given word and pipes them with AND operator starting with true, so it could short circuits.
func (ps Predicates) all(w Word) bool {
	ret := true
	for _, p := range ps {
		ret = ret && p(w)
		if !ret {
			return false
		}
	}
	return ret
}

func containsLetter(l Letter) predicate {
	return func(w Word) bool {
		for _, s := range w {
			if s == l {
				return true
			}
		}
		return false
	}
}

// not negates the given predicate.
func not(p predicate) predicate {
	return func(w Word) bool { return !p(w) }
}
