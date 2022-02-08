package wordmap

import "testing"

func Test_Contains(t *testing.T) {
	type args struct {
		w string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Empty",
			args: args{w: ""},
			want: false,
		},
		{
			name: "Contains",
			args: args{w: "lándzsa"},
			want: true,
		},
		{
			name: "Not Contains",
			args: args{w: "kottyo"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.w); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
