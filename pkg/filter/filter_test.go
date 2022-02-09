package filter

import (
	"reflect"
	"testing"
)

var wl Wordlist = [][]string{
	{"b", "a", "n", "dzs", "a"},
	{"b", "é", "k", "ly", "ó"},
	{"k", "o", "r", "o", "m"},
	{"m", "a", "r", "a", "d"},
	{"p", "á", "r", "n", "a"},
	{"k", "é", "sz", "e", "m"},
}

func TestPattern_Filter(t *testing.T) {
	type args struct {
		wl Wordlist
	}
	tests := []struct {
		name string
		p    Pattern
		args args
		want Wordlist
	}{
		{
			name: "Empty",
			p:    nil,
			args: args{wl: nil},
			want: nil,
		},
		{
			name: "Empty pattern",
			p:    nil,
			args: args{wl: wl},
			want: wl,
		},
		{
			name: "Pattern no filter",
			p:    Pattern{Mark{Letter: "cs", Position: 0, Marker: Gray}},
			args: args{wl: wl},
			want: wl,
		},
		{
			name: "Pattern filter orange",
			p:    Pattern{Mark{Letter: "a", Position: 0, Marker: Orange}},
			args: args{wl: wl},
			want: [][]string{
				{"b", "a", "n", "dzs", "a"},
				{"m", "a", "r", "a", "d"},
				{"p", "á", "r", "n", "a"},
			},
		},
		{
			name: "Pattern filter green",
			p:    Pattern{Mark{Letter: "n", Position: 2, Marker: Green}},
			args: args{wl: wl},
			want: [][]string{
				{"b", "a", "n", "dzs", "a"},
			},
		},
		{
			name: "Pattern filter green, orange",
			p: Pattern{
				Mark{Letter: "r", Position: 2, Marker: Green},
				Mark{Letter: "a", Position: 0, Marker: Orange},
			},
			args: args{wl: wl},
			want: [][]string{
				{"m", "a", "r", "a", "d"},
				{"p", "á", "r", "n", "a"},
			},
		},
		{
			name: "Pattern filter green, orange, gray I",
			p: Pattern{
				Mark{Letter: "r", Position: 2, Marker: Green},
				Mark{Letter: "a", Position: 0, Marker: Orange},
				Mark{Letter: "b", Position: 0, Marker: Gray},
			},
			args: args{wl: wl},
			want: [][]string{
				{"m", "a", "r", "a", "d"},
				{"p", "á", "r", "n", "a"},
			},
		},
		{
			name: "Pattern filter green, orange, gray II",
			p: Pattern{
				Mark{Letter: "r", Position: 2, Marker: Green},
				Mark{Letter: "a", Position: 0, Marker: Orange},
				Mark{Letter: "p", Position: 0, Marker: Gray},
			},
			args: args{wl: wl},
			want: [][]string{
				{"m", "a", "r", "a", "d"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Filter(tt.args.wl); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pattern.Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}
