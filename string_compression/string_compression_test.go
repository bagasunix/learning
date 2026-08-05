package stringcompression

import "testing"

func TestCompressBasic(t *testing.T) {
	if compress("aabcccccaaa") != "a2b1c5a3" {
		t.Errorf("got %s, want a2b1c5a3", compress("aabcccccaaa"))
	}
}

func TestCompressSingleChars(t *testing.T) {
	if compress("abc") != "a1b1c1" {
		t.Errorf("got %s, want a1b1c1", compress("abc"))
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
	if compress("a") != "a1" {
		t.Errorf("got %s, want a1", compress("a"))
	}
}

func TestCompressLongRun(t *testing.T) {
	if compress("aaaaaaaaaaaa") != "a12" {
		t.Errorf("got %s, want a12", compress("aaaaaaaaaaaa"))
	}
}
