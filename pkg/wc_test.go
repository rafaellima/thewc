package wc

import "testing"

func TestCountChars(t *testing.T) {
	tests := []struct {
		name string
		args *Wc
		want int
	}{
		{
			name: "count chars",
			args: &Wc{
				Data: "hello world",
			},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.CountChars(); got != tt.want {
				t.Errorf("CountChars() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountBytes(t *testing.T) {
	tests := []struct {
		name string
		args *Wc
		want int
	}{
		{
			name: "count bytes",
			args: &Wc{
				Data: "hello world",
			},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.CountBytes(); got != tt.want {
				t.Errorf("CountBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountLines(t *testing.T) {
	tests := []struct {
		name string
		args *Wc
		want int
	}{
		{
			name: "count lines",
			args: &Wc{
				Data: "hello\nworld\n",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.CountLines(); got != tt.want {
				t.Errorf("CountLines() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountWords(t *testing.T) {
	tests := []struct {
		name string
		args *Wc
		want int
	}{
		{
			name: "count words",
			args: &Wc{
				Data: "hello world counting words",
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.CountWords(); got != tt.want {
				t.Errorf("CountWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
