package firstuniquechar

import "testing"

func TestFirstUniqCharLeetcode(t *testing.T) {
	if firstUniqChar("leetcode") != 0 {
		t.Errorf("got %d, want 0", firstUniqChar("leetcode"))
	}
}

func TestFirstUniqCharLove(t *testing.T) {
	if firstUniqChar("loveleetcode") != 2 {
		t.Errorf("got %d, want 2", firstUniqChar("loveleetcode"))
	}
}

func TestFirstUniqCharAllRepeat(t *testing.T) {
	if firstUniqChar("aabb") != -1 {
		t.Errorf("got %d, want -1", firstUniqChar("aabb"))
	}
}

func TestFirstUniqCharSingle(t *testing.T) {
	if firstUniqChar("a") != 0 {
		t.Errorf("got %d, want 0", firstUniqChar("a"))
	}
}

func TestFirstUniqCharEmpty(t *testing.T) {
	if firstUniqChar("") != -1 {
		t.Errorf("got %d, want -1", firstUniqChar(""))
	}
}

func TestFirstUniqCharLastUnique(t *testing.T) {
	if firstUniqChar("aabbc") != 4 {
		t.Errorf("got %d, want 4", firstUniqChar("aabbc"))
	}
}
