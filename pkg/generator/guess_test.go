package generator

import (
	"reflect"
	"testing"
)

func Test_matchingGreens(t *testing.T) {
	type args struct {
		w word
		g Guess
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Empty",
			args: args{
				w: word{},
				g: Guess{},
			},
			want: true,
		},
		{
			name: "Empty guess",
			args: args{
				w: word{"a", "b", "l", "a", "k"},
				g: Guess{},
			},
			want: true,
		},
		{
			name: "Guess has no green",
			args: args{
				w: word{"a", "b", "l", "a", "k"},
				g: Guess{Word: word{"k", "o", "ny", "í", "t"}, Markers: [WordLength]marker{Gray, Gray, Gray, Gray, Gray}},
			},
			want: true,
		},
		{
			name: "Guess has one green word matches",
			args: args{
				w: word{"a", "b", "l", "a", "k"},
				g: Guess{Word: word{"a", "l", "s", "ó", "n"}, Markers: [WordLength]marker{Green, Gray, Gray, Gray, Gray}},
			},
			want: true,
		},
		{
			name: "Guess has two green word matches",
			args: args{
				w: word{"a", "b", "l", "a", "k"},
				g: Guess{Word: word{"a", "b", "b", "ó", "l"}, Markers: [WordLength]marker{Green, Green, Gray, Gray, Gray}},
			},
			want: true,
		},
		{
			name: "Guess has two green word does not match",
			args: args{
				w: word{"a", "k", "k", "o", "r"},
				g: Guess{Word: word{"a", "b", "b", "ó", "l"}, Markers: [WordLength]marker{Green, Green, Gray, Gray, Gray}},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.g.matchingGreens(tt.args.w); got != tt.want {
				t.Errorf("matchingGreens() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_overlapOrange(t *testing.T) {
	type args struct {
		w word
		g Guess
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Empty",
			args: args{
				w: word{},
				g: Guess{},
			},
			want: false,
		},
		{
			name: "Empty guess",
			args: args{
				w: word{"a", "k", "k", "o", "r"},
				g: Guess{},
			},
			want: false,
		},
		{
			name: "Guess all gray",
			args: args{
				w: word{"a", "k", "k", "o", "r"},
				g: Guess{Word: word{"m", "i", "d", "e", "n"}, Markers: [WordLength]marker{Gray, Gray, Gray, Gray, Gray}},
			},
			want: false,
		},
		{
			name: "Guess one orange word not overlap",
			args: args{
				w: word{"a", "k", "k", "o", "r"},
				g: Guess{Word: word{"m", "i", "d", "e", "n"}, Markers: [WordLength]marker{Orange, Gray, Gray, Gray, Gray}},
			},
			want: false,
		},
		{
			name: "Guess two orange word not overlap",
			args: args{
				w: word{"a", "k", "k", "o", "r"},
				g: Guess{Word: word{"m", "i", "d", "e", "n"}, Markers: [WordLength]marker{Orange, Gray, Orange, Gray, Gray}},
			},
			want: false,
		},
		{
			name: "Guess two orange word overlap",
			args: args{
				w: word{"m", "o", "d", "j", "a"},
				g: Guess{Word: word{"m", "i", "d", "e", "n"}, Markers: [WordLength]marker{Orange, Gray, Orange, Gray, Gray}},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.g.overlapOrange(tt.args.w); got != tt.want {
				t.Errorf("overlapOrange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mkTemplate(t *testing.T) {
	tests := []struct {
		name string
		gs   []Guess
		want template
	}{
		{
			name: "No guesses",
			gs:   []Guess{},
			want: template{word: word{emptyLetter, emptyLetter, emptyLetter, emptyLetter, emptyLetter}},
		},
		{
			name: "No green in guesses",
			gs: []Guess{
				{Word: word{"b", "i", "z", "o", "ny"}, Markers: [WordLength]marker{Gray, Orange, Gray, Gray, Gray}},
				{Word: word{"k", "o", "m", "o", "ly"}, Markers: [WordLength]marker{Gray, Orange, Gray, Gray, Gray}},
			},
			want: template{word: word{emptyLetter, emptyLetter, emptyLetter, emptyLetter, emptyLetter}},
		},
		{
			name: "One green in guesses",
			gs: []Guess{
				{Word: word{"b", "i", "z", "o", "ny"}, Markers: [WordLength]marker{Gray, Orange, Gray, Gray, Gray}},
				{Word: word{"k", "o", "m", "o", "ly"}, Markers: [WordLength]marker{Green, Orange, Gray, Gray, Gray}},
			},
			want: template{word: word{"k", emptyLetter, emptyLetter, emptyLetter, emptyLetter}},
		},
		{
			name: "Multiple green in guesses",
			gs: []Guess{
				{Word: word{"b", "i", "z", "o", "ny"}, Markers: [WordLength]marker{Gray, Orange, Green, Gray, Gray}},
				{Word: word{"k", "o", "m", "o", "ly"}, Markers: [WordLength]marker{Green, Orange, Gray, Gray, Gray}},
			},
			want: template{word: word{"k", emptyLetter, "z", emptyLetter, emptyLetter}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mkTemplate(tt.gs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mkTemplate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseMarker(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    marker
		wantErr bool
	}{
		{
			name:    "Empty",
			args:    args{s: ""},
			want:    "",
			wantErr: true,
		},
		{
			name:    "Invalid marker I",
			args:    args{s: ":"},
			want:    "",
			wantErr: true,
		},
		{
			name:    "Invalid marker II",
			args:    args{s: "adf"},
			want:    "",
			wantErr: true,
		},
		{
			name:    "Green",
			args:    args{s: "*"},
			want:    Green,
			wantErr: false,
		},
		{
			name:    "Orange",
			args:    args{s: "?"},
			want:    Orange,
			wantErr: false,
		},
		{
			name:    "Gray",
			args:    args{s: "#"},
			want:    Gray,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseMarker(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseMarker() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parseMarker() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseSubMatch(t *testing.T) {
	type args struct {
		m string
	}
	tests := []struct {
		name    string
		args    args
		want    marker
		want1   letter
		wantErr bool
	}{
		{
			name:    "Empty",
			args:    args{m: ""},
			want:    "",
			want1:   "",
			wantErr: true,
		},
		{
			name:    "Invalid maker and letter",
			args:    args{m: ":iy"},
			want:    "",
			want1:   "",
			wantErr: true,
		},
		{
			name:    "Invalid letter",
			args:    args{m: "*iy"},
			want:    "",
			want1:   "",
			wantErr: true,
		},
		{
			name:    "Invalid marker",
			args:    args{m: ":i"},
			want:    "",
			want1:   "",
			wantErr: true,
		},
		{
			name:    "Green M",
			args:    args{m: "*m"},
			want:    Green,
			want1:   "m",
			wantErr: false,
		},
		{
			name:    "Orange dzs",
			args:    args{m: "*dzs"},
			want:    Green,
			want1:   "dzs",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := parseSubMatch(tt.args.m)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseSubMatch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parseSubMatch() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("parseSubMatch() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestParse(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    Guess
		wantErr bool
	}{
		{
			name:    "Empty",
			args:    args{s: ""},
			want:    Guess{},
			wantErr: true,
		},
		{
			name:    "Not WordLength",
			args:    args{s: "*g"},
			want:    Guess{},
			wantErr: true,
		},
		{
			name:    "Invalid marker",
			args:    args{s: ":l:á:n:dzs:a"},
			want:    Guess{},
			wantErr: true,
		},
		{
			name:    "Invalid letter",
			args:    args{s: "#l#my*n?dzs#a"},
			want:    Guess{},
			wantErr: true,
		},
		{
			name:    "Valid I",
			args:    args{s: "#l#á*n*dzs#a"},
			want:    Guess{Word: word{"l", "á", "n", "dzs", "a"}, Markers: [WordLength]marker{Gray, Gray, Green, Green, Gray}},
			wantErr: false,
		},
		{
			name:    "Valid II",
			args:    args{s: "?ny?gy?ö?dzs?a"},
			want:    Guess{Word: word{"ny", "gy", "ö", "dzs", "a"}, Markers: [WordLength]marker{Orange, Orange, Orange, Orange, Orange}},
			wantErr: false,
		},
		{
			name:    "Valid III",
			args:    args{s: "#ny*gy?ö*dzs#a"},
			want:    Guess{Word: word{"ny", "gy", "ö", "dzs", "a"}, Markers: [WordLength]marker{Gray, Green, Orange, Green, Gray}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGuess_grays(t *testing.T) {
	type fields struct {
		Word    word
		Markers [WordLength]marker
	}
	tests := []struct {
		name   string
		fields fields
		want   letterSet
	}{
		{
			name:   "Empty",
			fields: fields{Word: word{}, Markers: [WordLength]marker{}},
			want:   letterSet{},
		},
		{
			name:   "No match",
			fields: fields{Word: word{"p", "o", "r", "t", "a"}, Markers: [WordLength]marker{Orange, Orange, Orange, Orange, Orange}},
			want:   letterSet{},
		},
		{
			name:   "One match",
			fields: fields{Word: word{"p", "o", "r", "t", "a"}, Markers: [WordLength]marker{Gray, Orange, Orange, Orange, Orange}},
			want:   letterSet{"p": {}},
		},
		{
			name:   "Multiple match",
			fields: fields{Word: word{"p", "o", "r", "t", "a"}, Markers: [WordLength]marker{Green, Gray, Gray, Orange, Gray}},
			want:   letterSet{"o": {}, "r": {}, "a": {}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := Guess{
				Word:    tt.fields.Word,
				Markers: tt.fields.Markers,
			}
			if got := g.grays(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Guess.filterLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hasEmptySlot(t *testing.T) {
	type args struct {
		t template
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "All empty",
			args: args{t: template{word: word{emptyLetter, emptyLetter, emptyLetter, emptyLetter, emptyLetter}}},
			want: true,
		},
		{
			name: "One empty",
			args: args{t: template{word: word{"d", emptyLetter, "dzs", "ly", "ö"}}},
			want: true,
		},
		{
			name: "Not has empty",
			args: args{t: template{word: word{"d", "á", "dzs", "ly", "ö"}}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasEmptySlot(tt.args.t); got != tt.want {
				t.Errorf("hasEmptySlot() = %v, want %v", got, tt.want)
			}
		})
	}
}
