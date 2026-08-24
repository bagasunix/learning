package reversewords

import "testing"

func TestReverseWords(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"normal sentence", "the sky is blue", "blue is sky the"},
		{"leading spaces", "  hello world  ", "world hello"},
		{"multiple spaces between", "a good   example", "example good a"},
		{"single word", "hello", "hello"},
		{"two words", "foo bar", "bar foo"},
		{"all spaces", "   ", ""},
		{"empty string", "", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseWords(tt.input)
			if got != tt.want {
				t.Errorf("reverseWords(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
