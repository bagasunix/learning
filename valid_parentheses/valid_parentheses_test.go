package validparentheses

import "testing"

func TestIsValidSimple(t *testing.T) {
	if !isValid("()") {
		t.Error("() should be valid")
	}
}

func TestIsValidMixed(t *testing.T) {
	if !isValid("()[]{}") {
		t.Error("()[]{} should be valid")
	}
}

func TestIsValidWrongClose(t *testing.T) {
	if isValid("(]") {
		t.Error("(] should be invalid")
	}
}

func TestIsValidWrongOrder(t *testing.T) {
	if isValid("([)]") {
		t.Error("([)] should be invalid")
	}
}

func TestIsValidNested(t *testing.T) {
	if !isValid("{[]}") {
		t.Error("{[]} should be valid")
	}
}

func TestIsValidEmpty(t *testing.T) {
	if !isValid("") {
		t.Error("empty should be valid")
	}
}

func TestIsValidOnlyOpen(t *testing.T) {
	if isValid("(") {
		t.Error("( should be invalid")
	}
}

func TestIsValidOnlyClose(t *testing.T) {
	if isValid(")") {
		t.Error(") should be invalid")
	}
}
