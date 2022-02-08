package generator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func mkGuesses(t *testing.T, ss ...string) []Guess {
	t.Helper()
	var ret []Guess
	for _, s := range ss {
		g, err := Parse(s)
		if err != nil {
			t.Fatal(err)
		}
		ret = append(ret, g)
	}
	return ret
}

func Test_grays(t *testing.T) {
	type args struct {
		gs []Guess
	}
	tests := []struct {
		name string
		args args
		want letterSet
	}{
		{
			name: "Empty",
			args: args{gs: []Guess{}},
			want: make(letterSet),
		},
		{
			name: "No gray",
			args: args{gs: mkGuesses(t, "?a?a?a?a?a")},
			want: make(letterSet),
		},
		{
			name: "One gray",
			args: args{gs: mkGuesses(t, "?a#a?a?a?a")},
			want: mkLetterSet("a"),
		},
		{
			name: "Multiple gray",
			args: args{gs: mkGuesses(t, "?a#a?a#b?a")},
			want: mkLetterSet("a", "b"),
		},
		{
			name: "Multiple gray in multi guess",
			args: args{gs: mkGuesses(t, "?a#a?a#b?a", "#cs?ly#dzs*n*gy")},
			want: mkLetterSet("a", "b", "cs", "dzs"),
		},
		{
			name: "Multiple gray in multi guess repeating grays",
			args: args{gs: mkGuesses(t, "?a#a?a#b?a", "#cs?ly#dzs*n#a")},
			want: mkLetterSet("a", "b", "cs", "dzs"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := grays(tt.args.gs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("grays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_filterByPreviousGuesses(t *testing.T) {
	type args struct {
		ws  []word
		pgs []Guess
	}
	tests := []struct {
		name string
		args args
		want []word
	}{
		{
			name: "Empty",
			args: args{ws: []word{}, pgs: mkGuesses(t)},
			want: []word{},
		},
		{
			name: "Empty words",
			args: args{ws: []word{}, pgs: mkGuesses(t, "*a*b*l*a*k")},
			want: []word{},
		},
		{
			name: "Empty guesses",
			args: args{ws: []word{{"a", "b", "l", "a", "k"}}, pgs: mkGuesses(t)},
			want: []word{{"a", "b", "l", "a", "k"}},
		},
		{
			name: "Green matches",
			args: args{ws: []word{{"a", "b", "l", "a", "k"}}, pgs: mkGuesses(t, "*a*b*l*a*k")},
			want: []word{{"a", "b", "l", "a", "k"}},
		},
		{
			name: "Orange matches",
			args: args{
				ws:  []word{{"a", "b", "l", "a", "k"}, {"b", "a", "n", "d", "a"}},
				pgs: mkGuesses(t, "?a?b#r#o#sz"),
			},
			want: []word{{"b", "a", "n", "d", "a"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := filterByPreviousGuesses(tt.args.ws, tt.args.pgs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filterByPreviousGuesses() = %v, want %v", got, tt.want)
			}
		})
	}
}

var (
	constTrue  ValidWord = func(s string) bool { return true }
	constFalse ValidWord = func(s string) bool { return false }
	validWords ValidWord = func(s string) bool {
		var m map[string]struct{} = map[string]struct{}{
			"kutyus": empty,
			"abált":  empty,
		}
		_, ok := m[s]
		return ok
	}
)

func Test_filterByValidWords(t *testing.T) {
	type args struct {
		ws []word
		v  ValidWord
	}
	tests := []struct {
		name string
		args args
		want []word
	}{
		{
			name: "Empty",
			args: args{ws: []word{}, v: nil},
			want: []word{},
		},
		{
			name: "All accepted",
			args: args{ws: []word{{"k", "o", "m", "o", "ly"}}, v: constTrue},
			want: []word{{"k", "o", "m", "o", "ly"}},
		},
		{
			name: "All rejected",
			args: args{ws: []word{{"k", "o", "m", "o", "ly"}}, v: constFalse},
			want: []word{},
		},
		{
			name: "valid words",
			args: args{ws: []word{{"k", "o", "m", "o", "ly"}, {"a", "b", "á", "l", "t"}}, v: validWords},
			want: []word{{"a", "b", "á", "l", "t"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := filterByValidWords(tt.args.ws, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filterByValidWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_putLetterToFirstEmptySlot(t *testing.T) {
	type args struct {
		t template
		l letter
	}
	tests := []struct {
		name string
		args args
		want template
	}{
		{
			name: "No empty slot",
			args: args{t: template{word: word{"f", "a", "r", "k", "a"}}, l: "dzs"},
			want: template{word: word{"f", "a", "r", "k", "a"}},
		},
		{
			name: "One empty slot",
			args: args{t: template{word: word{emptyLetter, "a", "r", "k", "a"}}, l: "dzs"},
			want: template{word: word{"dzs", "a", "r", "k", "a"}},
		},
		{
			name: "Two empty slot",
			args: args{t: template{word: word{emptyLetter, emptyLetter, "r", "k", "a"}}, l: "dzs"},
			want: template{word: word{"dzs", emptyLetter, "r", "k", "a"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := putLetterToFirstEmptySlot(tt.args.t, tt.args.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("putLetterToFirstEmptySlot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_putAllLettersToFirstEmptySlot(t *testing.T) {
	type args struct {
		ls []letter
		w  template
	}
	tests := []struct {
		name string
		args args
		want []template
	}{
		{
			name: "Emtpy",
			args: args{ls: []letter{}, w: template{word: word{}}},
			want: []template{{word: word{}}},
		},
		{
			name: "Emtpy letters",
			args: args{ls: []letter{}, w: template{word: word{emptyLetter, emptyLetter, emptyLetter, emptyLetter, emptyLetter}}},
			want: []template{{word: word{emptyLetter, emptyLetter, emptyLetter, emptyLetter, emptyLetter}}},
		},
		{
			name: "One letter to fill",
			args: args{ls: []letter{"ly"}, w: template{word: word{emptyLetter, emptyLetter, emptyLetter, emptyLetter, emptyLetter}}},
			want: []template{{word: word{"ly", emptyLetter, emptyLetter, emptyLetter, emptyLetter}}},
		},
		{
			name: "two letter to fill",
			args: args{ls: []letter{"ly", "ö"}, w: template{word: word{emptyLetter, emptyLetter, emptyLetter, emptyLetter, emptyLetter}}},
			want: []template{
				{word: word{"ly", emptyLetter, emptyLetter, emptyLetter, emptyLetter}},
				{word: word{"ö", emptyLetter, emptyLetter, emptyLetter, emptyLetter}},
			},
		},
		{
			name: "multiple letter to fill",
			args: args{ls: []letter{"ly", "ö", "cs"}, w: template{word: word{emptyLetter, "k", emptyLetter, emptyLetter, emptyLetter}}},
			want: []template{
				{word: word{"ly", "k", emptyLetter, emptyLetter, emptyLetter}},
				{word: word{"cs", "k", emptyLetter, emptyLetter, emptyLetter}},
				{word: word{"ö", "k", emptyLetter, emptyLetter, emptyLetter}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := putAllLettersToFirstEmptySlot(tt.args.ls, tt.args.w)
			assert.ElementsMatch(t, got, tt.want)
		})
	}
}

func Test_split(t *testing.T) {
	type args struct {
		ws []template
	}
	tests := []struct {
		name           string
		args           args
		wantFinshed    []template
		wantUnfinished []template
	}{
		{
			name: "Empty",
			args: args{
				ws: []template{},
			},
			wantFinshed:    nil,
			wantUnfinished: nil,
		},
		{
			name: "Only finished",
			args: args{
				ws: []template{
					{word: word{"r", "u", "h", "á", "s"}},
				},
			},
			wantFinshed: []template{
				{word: word{"r", "u", "h", "á", "s"}},
			},
			wantUnfinished: nil,
		},
		{
			name: "Only unfinished",
			args: args{
				ws: []template{
					{word: word{emptyLetter, "u", "h", "á", "s"}},
				},
			},
			wantFinshed: nil,
			wantUnfinished: []template{
				{word: word{emptyLetter, "u", "h", "á", "s"}},
			},
		},
		{
			name: "Mixed gnished unfinished I",
			args: args{
				ws: []template{
					{word: word{emptyLetter, "u", "h", "á", "s"}},
					{word: word{"j", "u", "h", "o", "k"}},
				},
			},
			wantFinshed: []template{
				{word: word{"j", "u", "h", "o", "k"}},
			},
			wantUnfinished: []template{
				{word: word{emptyLetter, "u", "h", "á", "s"}},
			},
		},
		{
			name: "Mixed gnished unfinished II",
			args: args{
				ws: []template{
					{word: word{emptyLetter, "u", "h", "á", "s"}},
					{word: word{emptyLetter, "o", "p", "á", "r"}},
					{word: word{"j", "u", "h", "o", "k"}},
					{word: word{"p", "u", "s", "k", "a"}},
				},
			},
			wantFinshed: []template{
				{word: word{"p", "u", "s", "k", "a"}},
				{word: word{"j", "u", "h", "o", "k"}},
			},
			wantUnfinished: []template{
				{word: word{emptyLetter, "o", "p", "á", "r"}},
				{word: word{emptyLetter, "u", "h", "á", "s"}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFinshed, gotUnfinished := split(tt.args.ws)
			assert.ElementsMatch(t, gotFinshed, tt.wantFinshed)
			assert.ElementsMatch(t, gotUnfinished, tt.wantUnfinished)
		})
	}
}

func Test_generateAll(t *testing.T) {
	type args struct {
		initialWord      []template
		availableLetters []letter
	}
	tests := []struct {
		name string
		args args
		want []template
	}{
		{
			name: "Emtpy",
			args: args{initialWord: []template{{word: word{}}}, availableLetters: []letter{}},
			want: []template{{word: word{}}},
		},
		{
			name: "One possible gen",
			args: args{
				initialWord:      []template{{word: word{emptyLetter, "u", "p", "a", "k"}}},
				availableLetters: []letter{"k"},
			},
			want: []template{
				{word: word{"k", "u", "p", "a", "k"}},
			},
		},
		{
			name: "multiple possible gen I",
			args: args{
				initialWord:      []template{{word: word{"s", "u", "t", emptyLetter, "ó"}}},
				availableLetters: []letter{"ty", "dzs", "ö"},
			},
			want: []template{
				{word: word{"s", "u", "t", "ty", "ó"}},
				{word: word{"s", "u", "t", "ö", "ó"}},
				{word: word{"s", "u", "t", "dzs", "ó"}},
			},
		},
		{
			name: "multiple possible gen II",
			args: args{
				initialWord:      []template{{word: word{"s", emptyLetter, "t", emptyLetter, "ó"}}},
				availableLetters: []letter{"ty", "u"},
			},
			want: []template{
				{word: word{"s", "ty", "t", "ty", "ó"}},
				{word: word{"s", "u", "t", "u", "ó"}},
				{word: word{"s", "u", "t", "ty", "ó"}},
				{word: word{"s", "ty", "t", "u", "ó"}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateAll(tt.args.initialWord, tt.args.availableLetters)
			assert.ElementsMatch(t, got, tt.want)
		})
	}
}

func TestGenerate(t *testing.T) {
	type args struct {
		i Input
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name:    "One letter missing",
			args:    args{i: Input{Guesses: mkGuesses(t, "#k*ú*t*o*r"), ValidWord: func(s string) bool { return s == "bútor" }}},
			want:    []string{"bútor"},
			wantErr: false,
		},
		{
			name:    "two letter missing",
			args:    args{i: Input{Guesses: mkGuesses(t, "#k#o*t*o*r"), ValidWord: func(s string) bool { return s == "bútor" }}},
			want:    []string{"bútor"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Generate(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Generate() = %v, want %v", got, tt.want)
			}
		})
	}
}
