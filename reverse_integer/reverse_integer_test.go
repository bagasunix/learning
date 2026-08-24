package reverseinteger

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int
	}{
		{"positive", 123, 321},
		{"negative", -123, -321},
		{"trailing zero", 120, 21},
		{"single digit", 9, 9},
		{"zero", 0, 0},
		{"overflow positive", 1534236469, 0},
		{"overflow negative", -1534236469, 0},
		{"palindrome", 1221, 1221},
		{"negative trailing zero", -120, -21},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverse(tt.input)
			if got != tt.want {
				t.Errorf("reverse(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}
