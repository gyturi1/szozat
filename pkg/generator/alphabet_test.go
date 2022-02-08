package generator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_letterSet_asSlice(t *testing.T) {
	tests := []struct {
		name string
		ls   letterSet
		want []letter
	}{
		{
			name: "Empty",
			ls:   mkLetterSet(),
			want: nil,
		},
		{
			name: "T1",
			ls:   mkLetterSet("a", "dzs", "ny"),
			want: []letter{"a", "dzs", "ny"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.ls.asSlice()
			assert.ElementsMatch(t, got, tt.want)
		})
	}
}

func Test_letterSet_addAll(t *testing.T) {
	type args struct {
		s letterSet
	}
	tests := []struct {
		name string
		ls   letterSet
		args args
		want letterSet
	}{
		{
			name: "Empty both",
			ls:   mkLetterSet(),
			args: args{s: mkLetterSet()},
			want: mkLetterSet(),
		},
		{
			name: "Empty left",
			ls:   mkLetterSet(),
			args: args{s: mkLetterSet("a")},
			want: mkLetterSet("a"),
		},
		{
			name: "Nil left",
			ls:   nil,
			args: args{s: mkLetterSet("a")},
			want: mkLetterSet("a"),
		},
		{
			name: "Empty arg",
			ls:   mkLetterSet("b"),
			args: args{s: nil},
			want: mkLetterSet("b"),
		},
		{
			name: "Nil arg",
			ls:   mkLetterSet("b"),
			args: args{s: nil},
			want: mkLetterSet("b"),
		},
		{
			name: "both non nil non empty",
			ls:   mkLetterSet("dzs", "k"),
			args: args{s: mkLetterSet("k", "m")},
			want: mkLetterSet("m", "dzs", "k"),
		},
		{
			name: "full alphabet",
			ls:   hunAlphabet.asSet(),
			args: args{s: hunAlphabet.asSet()},
			want: hunAlphabet.asSet(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ls.addAll(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("letterSet.addAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_letterSet_removeAll(t *testing.T) {
	type args struct {
		ls2 letterSet
	}
	tests := []struct {
		name string
		ls1  letterSet
		args args
		want letterSet
	}{
		{
			name: "Both empty",
			ls1:  mkLetterSet(),
			args: args{ls2: mkLetterSet()},
			want: mkLetterSet(),
		},
		{
			name: "Both nil",
			ls1:  nil,
			args: args{ls2: nil},
			want: mkLetterSet(),
		},
		{
			name: "Both non nil non empty",
			ls1:  mkLetterSet("a", "b", "dzs"),
			args: args{ls2: mkLetterSet("dzs")},
			want: mkLetterSet("a", "b"),
		},
		{
			name: "full alphabet",
			ls1:  hunAlphabet.asSet(),
			args: args{ls2: hunAlphabet.asSet()},
			want: mkLetterSet(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ls1.removeAll(tt.args.ls2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("letterSet.removeAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
