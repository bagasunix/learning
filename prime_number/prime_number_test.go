package primenumber

import (
	"reflect"
	"testing"
)

func TestIsPrime7(t *testing.T) {
	if !isPrime(7) {
		t.Error("7 should be prime")
	}
}

func TestIsPrimeNot10(t *testing.T) {
	if isPrime(10) {
		t.Error("10 should not be prime")
	}
}

func TestIsPrime2(t *testing.T) {
	if !isPrime(2) {
		t.Error("2 should be prime")
	}
}

func TestIsPrimeNot1(t *testing.T) {
	if isPrime(1) {
		t.Error("1 should not be prime")
	}
}

func TestIsPrimeNot0(t *testing.T) {
	if isPrime(0) {
		t.Error("0 should not be prime")
	}
}

func TestIsPrimeNegative(t *testing.T) {
	if isPrime(-5) {
		t.Error("-5 should not be prime")
	}
}

func TestIsPrimeLarge(t *testing.T) {
	if !isPrime(97) {
		t.Error("97 should be prime")
	}
}

func TestSievePrimes20(t *testing.T) {
	expected := []int{2, 3, 5, 7, 11, 13, 17, 19}
	result := sievePrimes(20)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("got %v, want %v", result, expected)
	}
}

func TestSievePrimes2(t *testing.T) {
	result := sievePrimes(2)
	if !reflect.DeepEqual(result, []int{2}) {
		t.Errorf("got %v, want [2]", result)
	}
}

func TestSievePrimesBelow2(t *testing.T) {
	result := sievePrimes(1)
	if len(result) != 0 {
		t.Errorf("expected empty, got %v", result)
	}
}
