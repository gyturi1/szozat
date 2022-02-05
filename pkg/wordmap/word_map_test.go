package wordmap

import (
	"reflect"
	"testing"

	"github.com/gyturi1/szozat/pkg/lib"
)

func TestFilter(t *testing.T) {
	tests := []struct {
		name string
		ws   []lib.Word
		want []lib.Word
	}{
		{
			name: "T1",
			ws:   []lib.Word{{"a", "l", "a", "t", "t"}},
			want: []lib.Word{{"a", "l", "a", "t", "t"}},
		},
		{
			name: "T2",
			ws:   []lib.Word{{"t", "ü", "cs", "ö", "k"}},
			want: []lib.Word{{"t", "ü", "cs", "ö", "k"}},
		},
		{
			name: "T3",
			ws:   []lib.Word{{"t", "ü", "cs", "ö", "k"}, {"a", "l", "a", "t", "t"}},
			want: []lib.Word{{"t", "ü", "cs", "ö", "k"}, {"a", "l", "a", "t", "t"}},
		},
		{
			name: "T4",
			ws:   []lib.Word{{"t", "ü", "cs", "ö", "k"}, {"a", "l", "a", "t", "t"}, {"l", "á", "z"}},
			want: []lib.Word{{"t", "ü", "cs", "ö", "k"}, {"a", "l", "a", "t", "t"}},
		},
		{
			name: "T5 nyári is definitly a 4 letter word see: github.com/gyturi1/cleanwords/README.md for details",
			ws:   []lib.Word{{"t", "ü", "cs", "ö", "k"}, {"a", "l", "a", "t", "t"}, {"l", "á", "z"}, {"n", "y", "á", "r", "i"}},
			want: []lib.Word{{"t", "ü", "cs", "ö", "k"}, {"a", "l", "a", "t", "t"}, {"n", "y", "á", "r", "i"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.ws); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}
