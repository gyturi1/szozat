package lib

import (
	"reflect"
	"testing"
)

func TestParseWords(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		want    Word
		want1   Word
		wantErr bool
	}{
		{"T1", []string{"", "", "", "", ""}, Word{}, Word{}, true},
		{"T2", []string{"", "", "dfadf", "", ""}, Word{}, Word{}, true},
		{"T3", []string{"a", "b", "CS", "d", "k"}, Word{"a", "b", "cs", "d", "k"}, EmptyWord, false},
		{"T4", []string{"a", "_", "CS", "d", "k"}, Word{"a", "_", "cs", "d", "k"}, EmptyWord, false},
		{"T5", []string{"*a", "_", "CS", "d", "k"}, Word{"a", "_", "cs", "d", "k"}, Word{"a", EmptyLetterMarker, EmptyLetterMarker, EmptyLetterMarker, EmptyLetterMarker}, false},
		{"T5", []string{"*a", "_", "*CS", "d", "k"}, Word{"a", "_", "cs", "d", "k"}, Word{"a", EmptyLetterMarker, "cs", EmptyLetterMarker, EmptyLetterMarker}, false},
		{"T5", []string{"*a", "_", "?CS", "d", "k"}, Word{}, Word{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := ParseWords(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseWords() error = %+v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseWords() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ParseWords() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestParseRemainingLetters(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		want    []Letter
		wantErr bool
	}{
		{"T1", []string{}, []Letter{}, false},
		{"T1", []string{"ffdagd"}, []Letter{}, true},
		{"T1", []string{"*l"}, []Letter{}, true},
		{"T1", []string{"a"}, []Letter{"a"}, false},
		{"T1", []string{"LY"}, []Letter{"ly"}, false},
		{"T1", []string{"LY", string(EmptyLetterMarker)}, []Letter{"ly"}, false},
		{"T1", []string{"LY", string(EmptyLetterMarker), "a"}, []Letter{"ly", "a"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseRemainingLetters(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseRemainingLetters() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseRemainingLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}
