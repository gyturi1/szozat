package lib

import (
	"reflect"
	"testing"
)

func Test_putLetterToFirstEmptySlot(t *testing.T) {
	type args struct {
		w Word
		l Letter
	}
	tests := []struct {
		name string
		args args
		want Word
	}{
		{"T1", args{[5]Letter{"a", "a", "a", "a", "a"}, Letter("b")}, [5]Letter{"a", "a", "a", "a", "a"}},
		{"T2", args{[5]Letter{"_", "a", "a", "a", "a"}, Letter("b")}, [5]Letter{"b", "a", "a", "a", "a"}},
		{"T3", args{[5]Letter{"a", "_", "a", "a", "a"}, Letter("b")}, [5]Letter{"a", "b", "a", "a", "a"}},
		{"T4", args{[5]Letter{"a", "a", "_", "a", "a"}, Letter("b")}, [5]Letter{"a", "a", "b", "a", "a"}},
		{"T5", args{[5]Letter{"a", "a", "a", "_", "a"}, Letter("b")}, [5]Letter{"a", "a", "a", "b", "a"}},
		{"T6", args{[5]Letter{"a", "a", "a", "a", "_"}, Letter("b")}, [5]Letter{"a", "a", "a", "a", "b"}},

		{"T7", args{[5]Letter{"_", "_", "_", "_", "_"}, Letter("b")}, [5]Letter{"b", "_", "_", "_", "_"}},
		{"T8", args{[5]Letter{"a", "_", "_", "_", "_"}, Letter("b")}, [5]Letter{"a", "b", "_", "_", "_"}},
		{"T8", args{[5]Letter{"a", "a", "_", "_", "_"}, Letter("b")}, [5]Letter{"a", "a", "b", "_", "_"}},
		{"T8", args{[5]Letter{"a", "a", "a", "_", "_"}, Letter("b")}, [5]Letter{"a", "a", "a", "b", "_"}},
		{"T8", args{[5]Letter{"a", "_", "a", "_", "a"}, Letter("b")}, [5]Letter{"a", "b", "a", "_", "a"}},
		{"T8", args{[5]Letter{"_", "_", "a", "_", "a"}, Letter("b")}, [5]Letter{"b", "_", "a", "_", "a"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := putLetterToFirstEmptySlot(tt.args.w, tt.args.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("putLetterToFirstEmptySlot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_putAllLettersToFirstEmptySlot(t *testing.T) {
	type args struct {
		ls []Letter
		w  Word
	}
	tests := []struct {
		name string
		args args
		want []Word
	}{
		{"T1", args{[]Letter{}, Word{"a", "a", "a", "a", "a"}}, []Word{{"a", "a", "a", "a", "a"}}},
		{"T2", args{[]Letter{"b"}, Word{"a", "a", "a", "a", "a"}}, []Word{{"a", "a", "a", "a", "a"}}},
		{"T3", args{[]Letter{"b"}, Word{"_", "a", "a", "a", "a"}}, []Word{{"b", "a", "a", "a", "a"}}},
		{"T4", args{[]Letter{"b", "c"}, Word{"_", "a", "a", "a", "a"}}, []Word{{"b", "a", "a", "a", "a"}, {"c", "a", "a", "a", "a"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := putAllLettersToFirstEmptySlot(tt.args.ls, tt.args.w); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("putAllLettersToFirstEmptySlot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_split(t *testing.T) {
	tests := []struct {
		name          string
		ws            []Word
		wantFinshed   []Word
		wantUnfinshed []Word
	}{
		{
			name:          "T0",
			ws:            []Word{{}},
			wantFinshed:   []Word{{}},
			wantUnfinshed: []Word{},
		},
		{
			name:          "T1",
			ws:            []Word{{"a", "a", "a", "a", "a"}},
			wantFinshed:   []Word{{"a", "a", "a", "a", "a"}},
			wantUnfinshed: []Word{},
		},
		{
			name:          "T2",
			ws:            []Word{{"a", "a", "a", "a", "a"}, {"_", "a", "a", "a", "a"}},
			wantFinshed:   []Word{{"a", "a", "a", "a", "a"}},
			wantUnfinshed: []Word{{"_", "a", "a", "a", "a"}},
		},
		{
			name:          "T3",
			ws:            []Word{{"a", "a", "a", "a", "a"}, {"_", "_", "a", "a", "a"}, {"_", "_", "_", "a", "a"}},
			wantFinshed:   []Word{{"a", "a", "a", "a", "a"}},
			wantUnfinshed: []Word{{"_", "_", "a", "a", "a"}, {"_", "_", "_", "a", "a"}},
		},
		{
			name:          "T4",
			ws:            []Word{{"a", "a", "a", "a", "a"}, {"_", "_", "a", "a", "a"}, {"_", "_", "_", "a", "a"}, {"a", "a", "a", "a", "b"}},
			wantFinshed:   []Word{{"a", "a", "a", "a", "a"}, {"a", "a", "a", "a", "b"}},
			wantUnfinshed: []Word{{"_", "_", "a", "a", "a"}, {"_", "_", "_", "a", "a"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFinshed, gotUnfinshed := split(tt.ws)
			if !reflect.DeepEqual(gotFinshed, tt.wantFinshed) {
				t.Errorf("split() gotFinshed = %v, want %v", gotFinshed, tt.wantFinshed)
			}
			if !reflect.DeepEqual(gotUnfinshed, tt.wantUnfinshed) {
				t.Errorf("split() gotUnfinshed = %v, want %v", gotUnfinshed, tt.wantUnfinshed)
			}
		})
	}
}

func Test_filterByPattern(t *testing.T) {
	type args struct {
		ws      []Word
		pattern Word
	}
	tests := []struct {
		name string
		args args
		want []Word
	}{
		{
			name: "T1",
			args: args{
				ws:      []Word{},
				pattern: Word{},
			},
			want: []Word{},
		},
		{
			name: "T2",
			args: args{
				ws:      []Word{{"a", "ly", "c", "v", "ny"}},
				pattern: EmptyWord,
			},
			want: []Word{{"a", "ly", "c", "v", "ny"}},
		},
		{
			name: "T3",
			args: args{
				ws:      []Word{{"a", "ly", "c", "v", "ny"}},
				pattern: Word{EmptyLetterMarker, "ly", EmptyLetterMarker, EmptyLetterMarker, EmptyLetterMarker},
			},
			want: []Word{{"a", "ly", "c", "v", "ny"}},
		},
		{
			name: "T4",
			args: args{
				ws:      []Word{{"a", "ly", "c", "v", "ny"}, {"a", "v", "c", "v", "ny"}},
				pattern: Word{EmptyLetterMarker, "ly", EmptyLetterMarker, EmptyLetterMarker, EmptyLetterMarker},
			},
			want: []Word{{"a", "ly", "c", "v", "ny"}},
		},
		{
			name: "T5",
			args: args{
				ws:      []Word{{"a", "ly", "c", "v", "ny"}, {"a", "ly", "k", "v", "ny"}},
				pattern: Word{EmptyLetterMarker, "ly", "c", EmptyLetterMarker, EmptyLetterMarker},
			},
			want: []Word{{"a", "ly", "c", "v", "ny"}},
		},
		{
			name: "T6",
			args: args{
				ws:      []Word{{"a", "ly", "c", "v", "ny"}, {"a", "ly", "c", "a", "ny"}, {"a", "ly", "b", "v", "ny"}},
				pattern: Word{EmptyLetterMarker, "ly", "c", EmptyLetterMarker, EmptyLetterMarker},
			},
			want: []Word{{"a", "ly", "c", "v", "ny"}, {"a", "ly", "c", "a", "ny"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := filterByPattern(tt.args.ws, tt.args.pattern); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filterByPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateAll(t *testing.T) {
	type args struct {
		initialWord      []Word
		availableLetters []Letter
	}
	tests := []struct {
		name string
		args args
		want []Word
	}{
		{
			name: "T0",
			args: args{
				initialWord:      []Word{},
				availableLetters: []Letter{},
			},
			want: []Word{},
		},
		{
			name: "T1",
			args: args{
				initialWord:      []Word{{"a", "ű", "c", "cs", "dzs"}},
				availableLetters: []Letter{},
			},
			want: []Word{{"a", "ű", "c", "cs", "dzs"}},
		},
		{
			name: "T2",
			args: args{
				initialWord:      []Word{},
				availableLetters: []Letter{"a"},
			},
			want: []Word{},
		},
		{
			name: "T3",
			args: args{
				initialWord:      []Word{EmptyWord},
				availableLetters: []Letter{"a"},
			},
			want: []Word{{"a", "a", "a", "a", "a"}},
		},
		{
			name: "T4",
			args: args{
				initialWord:      []Word{{"b", EmptyLetterMarker, EmptyLetterMarker, EmptyLetterMarker, EmptyLetterMarker}},
				availableLetters: []Letter{"a"},
			},
			want: []Word{{"b", "a", "a", "a", "a"}},
		},
		{
			name: "T5",
			args: args{
				initialWord:      []Word{{"b", "d", "ly", EmptyLetterMarker, EmptyLetterMarker}},
				availableLetters: []Letter{"a", "cs"},
			},
			want: []Word{{"b", "d", "ly", "a", "a"}, {"b", "d", "ly", "a", "cs"}, {"b", "d", "ly", "cs", "a"}, {"b", "d", "ly", "cs", "cs"}},
		},
		{
			name: "T6",
			args: args{
				initialWord:      []Word{{"b", "d", "ly", EmptyLetterMarker, EmptyLetterMarker}},
				availableLetters: []Letter{"a", "ly"},
			},
			want: []Word{{"b", "d", "ly", "a", "a"}, {"b", "d", "ly", "a", "ly"}, {"b", "d", "ly", "ly", "a"}, {"b", "d", "ly", "ly", "ly"}},
		},
		{
			name: "T7",
			args: args{
				initialWord:      []Word{{EmptyLetterMarker, "d", "ly", EmptyLetterMarker, "dzs"}},
				availableLetters: []Letter{"a", "dzs"},
			},
			want: []Word{{"a", "d", "ly", "a", "dzs"}, {"a", "d", "ly", "dzs", "dzs"}, {"dzs", "d", "ly", "a", "dzs"}, {"dzs", "d", "ly", "dzs", "dzs"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateAll(tt.args.initialWord, tt.args.availableLetters); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
