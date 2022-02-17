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
		p    Markers
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
			p:    Markers{Marker{Letter: "cs", Position: 0, M: Gray}},
			args: args{wl: wl},
			want: wl,
		},
		{
			name: "Pattern filter orange",
			p:    Markers{Marker{Letter: "a", Position: 0, M: Orange}},
			args: args{wl: wl},
			want: [][]string{
				{"b", "a", "n", "dzs", "a"},
				{"m", "a", "r", "a", "d"},
				{"p", "á", "r", "n", "a"},
			},
		},
		{
			name: "Pattern filter green",
			p:    Markers{Marker{Letter: "n", Position: 2, M: Green}},
			args: args{wl: wl},
			want: [][]string{
				{"b", "a", "n", "dzs", "a"},
			},
		},
		{
			name: "Pattern filter green, orange",
			p: Markers{
				Marker{Letter: "r", Position: 2, M: Green},
				Marker{Letter: "a", Position: 0, M: Orange},
			},
			args: args{wl: wl},
			want: [][]string{
				{"m", "a", "r", "a", "d"},
				{"p", "á", "r", "n", "a"},
			},
		},
		{
			name: "Pattern filter green, orange, gray I",
			p: Markers{
				Marker{Letter: "r", Position: 2, M: Green},
				Marker{Letter: "a", Position: 0, M: Orange},
				Marker{Letter: "b", Position: 0, M: Gray},
			},
			args: args{wl: wl},
			want: [][]string{
				{"m", "a", "r", "a", "d"},
				{"p", "á", "r", "n", "a"},
			},
		},
		{
			name: "Pattern filter green, orange, gray II",
			p: Markers{
				Marker{Letter: "r", Position: 2, M: Green},
				Marker{Letter: "a", Position: 0, M: Orange},
				Marker{Letter: "p", Position: 0, M: Gray},
			},
			args: args{wl: wl},
			want: [][]string{
				{"m", "a", "r", "a", "d"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.args.wl, tt.p.ToPredicates()); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pattern.Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}
