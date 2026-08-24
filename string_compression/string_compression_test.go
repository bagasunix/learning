package stringcompression

import "testing"

func TestCompressBasic(t *testing.T) {
	if compress("aabcccccaaa") != "a2b1c5a3" {
		t.Errorf("got %s, want a2b1c5a3", compress("aabcccccaaa"))
	}
}

func TestCompressSingleChars(t *testing.T) {
	// "a1b1c1" (6 chars) lebih panjang dari "abc" (3 chars) -> kembalikan asli (bonus)
	if compress("abc") != "abc" {
		t.Errorf("got %s, want abc", compress("abc"))
	}
}

func TestCompressEmpty(t *testing.T) {
	if compress("") != "" {
		t.Errorf("got %s, want empty", compress(""))
	}
}

func TestCompressAllSame(t *testing.T) {
	if compress("aaaa") != "a4" {
		t.Errorf("got %s, want a4", compress("aaaa"))
	}
}

func TestCompressSingleChar(t *testing.T) {
	// "a1" (2 chars) lebih panjang dari "a" (1 char) -> kembalikan asli (bonus)
	if compress("a") != "a" {
		t.Errorf("got %s, want a", compress("a"))
	}
}

func TestCompressLongRun(t *testing.T) {
	if compress("aaaaaaaaaaaa") != "a12" {
		t.Errorf("got %s, want a12", compress("aaaaaaaaaaaa"))
	}
}
