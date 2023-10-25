package fast_and_slow

import "testing"

func TestIsHappy(t *testing.T) {

	tests := []struct {
		name string
		num  int
		want bool
	}{
		{
			name: "Case 1",
			num:  2147483646,
			want: false,
		},
		{
			name: "Case 2",
			num:  1,
			want: true,
		},
		{
			name: "Case 3",
			num:  19,
			want: true,
		},
		{
			name: "Case 4",
			num:  8,
			want: false,
		},
		{
			name: "Case 5",
			num:  7,
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsHappy(tt.num); got != tt.want {
				t.Errorf("IsHappy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sumOfDigitsSquared(t *testing.T) {

	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "Case 1",
			n:    28,
			want: 68,
		},
		{
			name: "Case 2",
			n:    4,
			want: 16,
		},
		{
			name: "Case 3",
			n:    100,
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfDigitsSquared(tt.n); got != tt.want {
				t.Errorf("sumOfDigitsSquared() = %v, want %v", got, tt.want)
			}
		})
	}
}
