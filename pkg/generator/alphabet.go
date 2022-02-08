package generator

type Alphabet map[string]struct{}
type letterSet map[letter]struct{}

const emptyLetter letter = "_"

//use for item in letter set
var empty struct{}

var hun_alphabet Alphabet = Alphabet{
	"a":   {},
	"á":   {},
	"b":   {},
	"c":   {},
	"cs":  {},
	"d":   {},
	"dz":  {},
	"dzs": {},
	"e":   {},
	"é":   {},
	"f":   {},
	"g":   {},
	"gy":  {},
	"h":   {},
	"i":   {},
	"í":   {},
	"j":   {},
	"k":   {},
	"l":   {},
	"ly":  {},
	"m":   {},
	"n":   {},
	"ny":  {},
	"o":   {},
	"ó":   {},
	"ö":   {},
	"ő":   {},
	"p":   {},
	"q":   {},
	"r":   {},
	"s":   {},
	"sz":  {},
	"t":   {},
	"ty":  {},
	"u":   {},
	"ú":   {},
	"ü":   {},
	"ű":   {},
	"v":   {},
	"w":   {},
	"x":   {},
	"y":   {},
	"z":   {},
	"zs":  {},
}

func mkLetterSet(s ...string) letterSet {
	var ret letterSet = make(letterSet)
	for _, i := range s {
		ret[letter(i)] = empty
	}
	return ret
}

func (a Alphabet) asSet() letterSet {
	ls := make(letterSet)
	for k := range a {
		ls[letter(k)] = empty
	}
	return ls
}

//addAll creates a copy of the original letterSet and adds all the element provided as argument to this method.
//Both the original and the passed in sets are left intact.
func (ls letterSet) addAll(s letterSet) letterSet {
	var ret letterSet = make(letterSet)
	for l := range ls {
		ret[l] = empty
	}
	for l2 := range s {
		ret[l2] = empty
	}
	return ret
}

//removeAll copies from the original set what is not present in the provided set.
//Both the original and the passed in sets are left intact.
func (ls1 letterSet) removeAll(ls2 letterSet) letterSet {
	var ret letterSet = make(letterSet)
	for l1 := range ls1 {
		if _, cont := ls2[l1]; !cont {
			ret[l1] = empty
		}
	}
	return ret
}

func (ls letterSet) asSlice() []letter {
	var ret []letter
	for l := range ls {
		ret = append(ret, l)
	}
	return ret
}
