package filter

import (
	"reflect"
	"testing"
)

func TestParseGrayLetterFlag(t *testing.T) {
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
			name:    "Zero length",
			args:    args{s: ""},
			want:    nil,
			wantErr: false,
		},
		{
			name:    "Only spaces",
			args:    args{s: "  "},
			want:    nil,
			wantErr: false,
		},
		{
			name:    "Invalid letter",
			args:    args{s: "ky"},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Valid letter",
			args: args{s: "ly"},
			want: []Marker{
				{Letter: "ly", M: Gray},
			},
			wantErr: false,
		},
		{
			name: "Valid letter",
			args: args{s: "ly dzs k"},
			want: []Marker{
				{Letter: "ly", M: Gray},
				{Letter: "dzs", M: Gray},
				{Letter: "k", M: Gray},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseGrayLetters(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseGrayLetterFlag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseGrayLetterFlag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseOrangeLetters(t *testing.T) {
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
			name:    "Zero length",
			args:    args{s: ""},
			want:    nil,
			wantErr: false,
		},
		{
			name:    "Only spaces",
			args:    args{s: "  "},
			want:    nil,
			wantErr: false,
		},
		{
			name:    "Invalid letter",
			args:    args{s: "ky"},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Valid letter",
			args: args{s: "ly"},
			want: []Marker{
				{Letter: "ly", M: Orange},
			},
			wantErr: false,
		},
		{
			name: "Valid letter",
			args: args{s: "ly dzs k"},
			want: []Marker{
				{Letter: "ly", M: Orange},
				{Letter: "dzs", M: Orange},
				{Letter: "k", M: Orange},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseOrangeLetters(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseOrangeLetters() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseOrangeLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseGreenLetters(t *testing.T) {
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
			name:    "Zero length",
			args:    args{s: ""},
			want:    nil,
			wantErr: false,
		},
		{
			name:    "Only spaces",
			args:    args{s: "  "},
			want:    nil,
			wantErr: false,
		},
		{
			name:    "Invalid letter:invalid index",
			args:    args{s: "ky:12"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Invalid letter valid index",
			args:    args{s: "az:1"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Valid letter invlid index I",
			args:    args{s: "ly:0"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Valid letter invlid index II",
			args:    args{s: "ly:6"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Valid letter invlid index III",
			args:    args{s: "ly:az"},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Valid letter valid index I",
			args: args{s: "ly:1"},
			want: []Marker{
				{Letter: "ly", Position: 0, M: Green},
			},
			wantErr: false,
		},
		{
			name: "Valid letter valid index II",
			args: args{s: "ly:1 g:2 dzs:5"},
			want: []Marker{
				{Letter: "ly", Position: 0, M: Green},
				{Letter: "g", Position: 1, M: Green},
				{Letter: "dzs", Position: 4, M: Green},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseGreenLetters(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseGreenLetters() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseGreenLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}
