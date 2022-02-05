package lib

import (
	"reflect"
	"testing"
)

func TestParseGuess(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		want    Word
		want1   Word
		wantErr bool
	}{
		{"T0", []string{}, Word{}, Word{}, false},
		{"T1", []string{"", "", "", "", ""}, Word{}, Word{}, true},
		{"T2", []string{" ", " ", " ", " ", " "}, Word{}, Word{}, true},
		{"T3", []string{"", "", "dfadf", "", ""}, Word{}, Word{}, true},
		{"T4", []string{"a", "b", "CS", "d", "k"}, Word{"a", "b", "cs", "d", "k"}, EmptyWord, false},
		{"T5", []string{"a", "_", "CS", "d", "k"}, Word{"a", "_", "cs", "d", "k"}, EmptyWord, false},
		{"T6", []string{"*a", "_", "CS", "d", "k"}, Word{"a", "_", "cs", "d", "k"}, Word{"a", EmptyLetterMarker, EmptyLetterMarker, EmptyLetterMarker, EmptyLetterMarker}, false},
		{"T7", []string{"*a", "_", "*CS", "d", "k"}, Word{"a", "_", "cs", "d", "k"}, Word{"a", EmptyLetterMarker, "cs", EmptyLetterMarker, EmptyLetterMarker}, false},
		{"T8", []string{"*a", "_", "?CS", "d", "k"}, Word{}, Word{}, true},
		{"T9", []string{"f", "g"}, Word{}, Word{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := ParseGuess(tt.args)
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
		{"T0", []string{}, []Letter{}, false},
		{"T1", []string{""}, []Letter{}, false},
		{"T2", []string{"", ""}, []Letter{}, false},
		{"T3", []string{" ", " "}, []Letter{}, false},
		{"T4", []string{"ffdagd"}, []Letter{}, true},
		{"T5", []string{"*l"}, []Letter{}, true},
		{"T6", []string{"a"}, []Letter{"a"}, false},
		{"T7", []string{"LY"}, []Letter{"ly"}, false},
		{"T8", []string{"LY", string(EmptyLetterMarker)}, []Letter{"ly"}, false},
		{"T9", []string{"LY", string(EmptyLetterMarker), "a"}, []Letter{"ly", "a"}, false},
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
