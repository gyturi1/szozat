package main

import (
	"reflect"
	"testing"

	"github.com/gyturi1/szozat/pkg/lib"
)

func Test_parseFlags(t *testing.T) {
	type args struct {
		guess string
		al    string
	}
	tests := []struct {
		name  string
		args  args
		want  lib.Word
		want1 lib.Word
		want2 []lib.Letter
	}{
		{
			name: "T1",
			args: args{
				guess: "_ _ cs ö k",
				al:    "ü f",
			},
			want:  lib.Word{lib.EmptyLetterMarker, lib.EmptyLetterMarker, "cs", "ö", "k"},
			want1: lib.EmptyWord,
			want2: []lib.Letter{"ü", "f"},
		},
		{
			name: "T2",
			args: args{
				guess: "_ _ *cs ö k",
				al:    "ü f",
			},
			want:  lib.Word{lib.EmptyLetterMarker, lib.EmptyLetterMarker, "cs", "ö", "k"},
			want1: lib.Word{lib.EmptyLetterMarker, lib.EmptyLetterMarker, "cs", lib.EmptyLetterMarker, lib.EmptyLetterMarker},
			want2: []lib.Letter{"ü", "f"},
		},
		{
			name: "T3",
			args: args{
				guess: "_ _ *cs *ö *k",
				al:    "ü f g",
			},
			want:  lib.Word{lib.EmptyLetterMarker, lib.EmptyLetterMarker, "cs", "ö", "k"},
			want1: lib.Word{lib.EmptyLetterMarker, lib.EmptyLetterMarker, "cs", "ö", "k"},
			want2: []lib.Letter{"ü", "f", "g"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := parseFlags(tt.args.guess, tt.args.al)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseArgs() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("parseArgs() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("parseArgs() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
