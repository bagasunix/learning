package countvowels

import (
	"reflect"
	"testing"
)

func TestCountVowelsHello(t *testing.T) {
	if countVowels("hello") != 2 {
		t.Errorf("got %d, want 2", countVowels("hello"))
	}
}

func TestCountVowelsHermes(t *testing.T) {
	if countVowels("Hermes") != 2 {
		t.Errorf("got %d, want 2", countVowels("Hermes"))
	}
}

func TestCountVowelsNoVowels(t *testing.T) {
	if countVowels("sky") != 0 {
		t.Errorf("got %d, want 0", countVowels("sky"))
	}
}

func TestCountVowelsAllVowels(t *testing.T) {
	if countVowels("aiueo") != 5 {
		t.Errorf("got %d, want 5", countVowels("aiueo"))
	}
}

func TestCountVowelsEmpty(t *testing.T) {
	if countVowels("") != 0 {
		t.Errorf("got %d, want 0", countVowels(""))
	}
}

func TestCountVowelsDetail(t *testing.T) {
	result := countVowelsDetail("hello")
	expected := map[string]int{"a": 0, "i": 0, "u": 0, "e": 1, "o": 1}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("got %v, want %v", result, expected)
	}
}

func TestCountVowelsDetailAll(t *testing.T) {
	result := countVowelsDetail("aiueo")
	for _, v := range "aiueo" {
		if result[string(v)] != 1 {
			t.Errorf("expected 1 for %c, got %d", v, result[string(v)])
		}
	}
}
