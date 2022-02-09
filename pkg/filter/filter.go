package filter

// simple aliases.
type (
	Letter   = string
	Word     = []Letter
	Wordlist = []Word
	Alphabet = []Letter
)

// A predicate that tell  if a word matches.
type predicate = func(Word) bool

// All the predicate is in and relation.
type predicates []predicate

func (p Pattern) Filter(wl Wordlist) Wordlist {
	var ret Wordlist
	t := mkTester(p)
	for _, w := range wl {
		if t.test(w) {
			ret = append(ret, w)
		}
	}
	return ret
}

func mkTester(p Pattern) predicates {
	var ret predicates

	for _, m := range p {
		if m.Marker == Green {
			ret = append(ret, greenMatcher(m))
		}
		if m.Marker == Orange {
			ret = append(ret, orangeMatcher(m.Letter))
		}
		if m.Marker == Gray {
			ret = append(ret, grayMatcher(m.Letter))
		}
	}

	return ret
}

func greenMatcher(m Mark) predicate {
	return func(w Word) bool {
		return w[m.Position] == m.Letter
	}
}

func orangeMatcher(l Letter) predicate {
	return containsLetter(l)
}

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

func not(p predicate) predicate {
	return func(w Word) bool { return !p(w) }
}

func (ps predicates) test(w Word) bool {
	ret := true
	for _, p := range ps {
		ret = ret && p(w)
		if !ret {
			return false
		}
	}
	return ret
}
