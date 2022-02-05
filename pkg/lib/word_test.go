package lib

import (
	"reflect"
	"testing"
)

func TestMkLetter(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		want    Letter
		wantErr bool
	}{
		{"T1", "", InvalidLetter, true},
		{"T2", "dt", InvalidLetter, true},
		{"T2", "DT", InvalidLetter, true},
		{"T3", "cs", Letter("cs"), false},
		{"T3", "ly", Letter("ly"), false},
		{"T3", "lY", Letter("ly"), false},
		{"T3", "LY", Letter("ly"), false},
		{"T3", "a", Letter("a"), false},
		{"T3", "A", Letter("a"), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MkLetter(tt.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("MkLetter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MkLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hasEmptySlot(t *testing.T) {
	tests := []struct {
		name string
		w    Word
		want bool
	}{
		{"T1", [5]Letter{"_", "", "", "", ""}, true},
		{"T2", [5]Letter{"_", "a", "a", "cs", "a"}, true},
		{"T3", [5]Letter{"", "", "", "", ""}, false},
		{"T4", [5]Letter{"a", "a", "a", "a", "a"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasEmptySlot(tt.w); got != tt.want {
				t.Errorf("hasEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWord_NotEmptyLetters(t *testing.T) {
	tests := []struct {
		name string
		w    *Word
		want []Letter
	}{
		{"T1", &Word{"a", "b", "c", "d", "zs"}, []Letter{"a", "b", "c", "d", "zs"}},
		{"T2", &Word{EmptyLetterMarker, "b", "c", "d", "zs"}, []Letter{"b", "c", "d", "zs"}},
		{"T3", &Word{EmptyLetterMarker, "a", EmptyLetterMarker, EmptyLetterMarker, EmptyLetterMarker}, []Letter{"a"}},
		{"T3", &EmptyWord, []Letter{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.w.NotEmptyLetters(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Word.NotEmptyLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMkWord(t *testing.T) {
	type args struct {
		s1 string
		s2 string
		s3 string
		s4 string
		s5 string
	}
	tests := []struct {
		name    string
		args    args
		want    Word
		wantErr bool
	}{
		{
			name:    "T1",
			args:    args{"", "", "", "", ""},
			want:    InvalidWord,
			wantErr: true,
		},
		{
			name:    "T2",
			args:    args{"adfad", "lyy", "", "_", "_"},
			want:    InvalidWord,
			wantErr: true,
		},
		{
			name:    "T3",
			args:    args{"a", "ly", "_", "_", "_"},
			want:    Word{"a", "ly", "_", "_", "_"},
			wantErr: false,
		},
		{
			name:    "T4",
			args:    args{"a", "ly", "k", "b", "D"},
			want:    Word{"a", "ly", "k", "b", "d"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MkWord(tt.args.s1, tt.args.s2, tt.args.s3, tt.args.s4, tt.args.s5)
			if (err != nil) != tt.wantErr {
				t.Errorf("MkWord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MkWord() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWord_Match(t *testing.T) {
	tests := []struct {
		name string
		w    *Word
		patt Word
		want bool
	}{
		{"T1", safeMkWord("a", "b", "c", "d", "e"), *safeMkWord("a", "b", "c", "d", "e"), true},
		{"T2", safeMkWord("a", "b", "c", "d", "e"), *safeMkWord("a", "b", "c", "_", "e"), true},
		{"T3", safeMkWord("a", "ly", "c", "d", "e"), *safeMkWord("_", "ly", "c", "_", "e"), true},
		{"T4", safeMkWord("a", "ly", "c", "d", "e"), *safeMkWord("_", "_", "_", "_", "_"), true},
		{"T5", safeMkWord("a", "ly", "c", "d", "e"), *safeMkWord("_", "_", "_", "_", "k"), false},
		{"T5", safeMkWord("a", "ly", "c", "d", "e"), *safeMkWord("a", "_", "_", "_", "k"), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.w.Match(tt.patt); got != tt.want {
				t.Errorf("Word.Match() = %v, want %v", got, tt.want)
			}
		})
	}
}

func safeMkWord(s1, s2, s3, s4, s5 string) *Word {
	return &Word{Letter(s1), Letter(s2), Letter(s3), Letter(s4), Letter(s5)}
}
