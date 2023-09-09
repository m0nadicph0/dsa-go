package wordladder

import "testing"

func TestWordLadderLength(t *testing.T) {

	tests := []struct {
		name  string
		begin string
		end   string
		words []string
		want  int
	}{
		{
			name:  "Case 1",
			begin: "hit",
			end:   "cog",
			words: []string{"hot", "dot", "dog", "lot", "log", "cog"},
			want:  5,
		},
		{
			name:  "Case 2",
			begin: "hit",
			end:   "cog",
			words: []string{"hot", "dot", "dog", "lot", "log"},
			want:  0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WordLadderLength(tt.begin, tt.end, tt.words); got != tt.want {
				t.Errorf("WordLadderLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
