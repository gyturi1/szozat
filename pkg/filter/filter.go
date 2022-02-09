package filter

import (
	"fmt"

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
type predicates []predicate

// Filter the provided wordlist for the Markers.
func (p Markers) Filter(wl Wordlist) Wordlist {
	log.Info().Str("method", "Filter()").Msg("")
	var ret Wordlist
	ps := p.predicates()
	for _, w := range wl {
		if ps.all(w) {
			ret = append(ret, w)
		}
	}
	return ret
}

// builds the predicate list from the Markers.
func (p Markers) predicates() predicates {
	log.Info().Str("method", "predicates()").Msg("")
	log.Debug().Str("markers", fmt.Sprintf("%v", p)).Msg("")
	var ret predicates

	for _, m := range p {
		if m.M == Green {
			log.Debug().Str("mark", fmt.Sprintf("%v", m)).Msg("Adding green predicate")
			ret = append(ret, greenMatcher(m))
		}
		if m.M == Orange {
			log.Debug().Str("letter", m.Letter).Msg("Adding orange predicate")
			ret = append(ret, orangeMatcher(m.Letter))
		}
		if m.M == Gray {
			log.Debug().Str("letter", m.Letter).Msg("Adding gray predicate")
			ret = append(ret, grayMatcher(m.Letter))
		}
	}

	return ret
}

// all runs all the predicates for the given word and pipes them with AND operator starting with true, so it could short circuits.
func (ps predicates) all(w Word) bool {
	ret := true
	for _, p := range ps {
		ret = ret && p(w)
		if !ret {
			return false
		}
	}
	return ret
}

// builds a predicate for testing that the word has a letter in the given position.
func greenMatcher(m Marker) predicate {
	return func(w Word) bool {
		return w[m.Position] == m.Letter
	}
}

// builds a predicate for testing if the word contains the letter.
func orangeMatcher(l Letter) predicate {
	return containsLetter(l)
}

// builds a predicate for testing if the word does not contain the letter.
func grayMatcher(l Letter) predicate {
	return not(containsLetter(l))
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
