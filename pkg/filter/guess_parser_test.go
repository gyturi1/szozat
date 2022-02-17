package filter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseMarker(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    M
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
			args:    args{s: "#"},
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
			args:    args{s: ":"},
			want:    Green,
			wantErr: false,
		},
		{
			name:    "Orange",
			args:    args{s: "+"},
			want:    Orange,
			wantErr: false,
		},
		{
			name:    "Gray",
			args:    args{s: "-"},
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
		want    M
		want1   Letter
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
			args:    args{m: "#iy"},
			want:    "",
			want1:   "",
			wantErr: true,
		},
		{
			name:    "Invalid letter",
			args:    args{m: ":iy"},
			want:    "",
			want1:   "",
			wantErr: true,
		},
		{
			name:    "Invalid marker",
			args:    args{m: "$i"},
			want:    "",
			want1:   "",
			wantErr: true,
		},
		{
			name:    "Green M",
			args:    args{m: ":m"},
			want:    Green,
			want1:   "m",
			wantErr: false,
		},
		{
			name:    "Orange dzs",
			args:    args{m: "+dzs"},
			want:    Orange,
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
		want    []Marker
		wantErr bool
	}{
		{
			name:    "Empty",
			args:    args{s: ""},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Not WordLength",
			args:    args{s: "+g"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Invalid marker",
			args:    args{s: "*l*á*n*dzs*a"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Invalid letter",
			args:    args{s: "+l:my-n-dzs+a"},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Valid I",
			args: args{s: "-l-á:n:dzs+a"},
			want: []Marker{
				{Letter: "l", Position: 0, M: Gray},
				{Letter: "á", Position: 1, M: Gray},
				{Letter: "n", Position: 2, M: Green},
				{Letter: "dzs", Position: 3, M: Green},
				{Letter: "a", Position: 4, M: Orange},
			},
			wantErr: false,
		},
		{
			name: "Valid II",
			args: args{s: "+ny+gy+ö+dzs+a"},
			want: []Marker{
				{Letter: "ny", Position: 0, M: Orange},
				{Letter: "gy", Position: 1, M: Orange},
				{Letter: "ö", Position: 2, M: Orange},
				{Letter: "dzs", Position: 3, M: Orange},
				{Letter: "a", Position: 4, M: Orange},
			},
			wantErr: false,
		},
		{
			name: "Valid III",
			args: args{s: "-ny+gy:ö+dzs-a"},
			want: []Marker{
				{Letter: "ny", Position: 0, M: Gray},
				{Letter: "gy", Position: 1, M: Orange},
				{Letter: "ö", Position: 2, M: Green},
				{Letter: "dzs", Position: 3, M: Orange},
				{Letter: "a", Position: 4, M: Gray},
			},
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
			assert.ElementsMatch(t, got, tt.want)
		})
	}
}
