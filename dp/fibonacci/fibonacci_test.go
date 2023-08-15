package fibonacci

import "testing"

func TestFibonacci(t *testing.T) {

	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "0th Fibonacci number",
			n:    0,
			want: 0,
		},
		{
			name: "1st Fibonacci number",
			n:    1,
			want: 1,
		},
		{
			name: "2nd Fibonacci number",
			n:    2,
			want: 1,
		},
		{
			name: "Fibonacci number of positive index",
			n:    6,
			want: 8,
		},
		{
			name: "Fibonacci number of a large positive index",
			n:    10,
			want: 55,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fibonacci(tt.n); got != tt.want {
				t.Errorf("Fibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMemoFibonacci(t *testing.T) {

	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "0th Fibonacci number",
			n:    0,
			want: 0,
		},
		{
			name: "1st Fibonacci number",
			n:    1,
			want: 1,
		},
		{
			name: "2nd Fibonacci number",
			n:    2,
			want: 1,
		},
		{
			name: "Fibonacci number of positive index",
			n:    6,
			want: 8,
		},
		{
			name: "Fibonacci number of a large positive index",
			n:    10,
			want: 55,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MemoFibonacci(tt.n); got != tt.want {
				t.Errorf("Fibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBottomUpDPFibonacci(t *testing.T) {

	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "0th Fibonacci number",
			n:    0,
			want: 0,
		},
		{
			name: "1st Fibonacci number",
			n:    1,
			want: 1,
		},
		{
			name: "2nd Fibonacci number",
			n:    2,
			want: 1,
		},
		{
			name: "Fibonacci number of positive index",
			n:    6,
			want: 8,
		},
		{
			name: "Fibonacci number of a large positive index",
			n:    10,
			want: 55,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BottomUpDPFibonacci(tt.n); got != tt.want {
				t.Errorf("Fibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}
