package filter

var hunAlphabet = Alphabet{
	"a", "á", "b", "c", "cs",
	"d", "dz", "dzs", "e", "é",
	"f", "g", "gy", "h", "i",
	"í", "j", "k", "l", "ly",
	"m", "n", "ny", "o", "ó",
	"ö", "ő", "p", "q", "r",
	"s", "sz", "t", "ty", "u",
	"ú", "ü", "ű", "v", "w",
	"x", "y", "z", "zs",
}

func validLetter(l Letter) bool {
	for _, s := range hunAlphabet {
		if s == l {
			return true
		}
	}
	return false
}
