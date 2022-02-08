package generator

import (
	"testing"
)

func TestMkLetter(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		want    letter
		wantErr bool
	}{
		{"Empty", "", "", true},
		{"Invalid", "df", "", true},
		{"Valid one letter", "f", "f", false},
		{"Valid double letter", "dz", "dz", false},
		{"Valid tripple letter", "dzs", "dzs", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := mkLetter(tt.s)
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
