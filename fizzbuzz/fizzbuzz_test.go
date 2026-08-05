package fizzbuzz

import (
	"reflect"
	"testing"
)

func TestFizzBuzzBasic15(t *testing.T) {
	expected := []string{
		"1", "2", "Fizz", "4", "Buzz",
		"Fizz", "7", "8", "Fizz", "Buzz",
		"11", "Fizz", "13", "14", "FizzBuzz",
	}
	result := fizzbuzz(15)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("got %v, want %v", result, expected)
	}
}

func TestFizzBuzzEmpty(t *testing.T) {
	result := fizzbuzz(0)
	if len(result) != 0 {
		t.Errorf("expected empty, got %v", result)
	}
}

func TestFizzBuzzSingle(t *testing.T) {
	result := fizzbuzz(1)
	if !reflect.DeepEqual(result, []string{"1"}) {
		t.Errorf("got %v, want [1]", result)
	}
}

func TestFizzBuzzFizzAt3(t *testing.T) {
	result := fizzbuzz(3)
	if len(result) < 3 {
		t.Fatalf("expected at least 3 elements, got %d", len(result))
	}
	if result[2] != "Fizz" {
		t.Errorf("expected Fizz at index 2, got %s", result[2])
	}
}

func TestFizzBuzzBuzzAt5(t *testing.T) {
	result := fizzbuzz(5)
	if result[4] != "Buzz" {
		t.Errorf("expected Buzz at index 4, got %s", result[4])
	}
}

func TestFizzBuzzFizzBuzzAt15(t *testing.T) {
	result := fizzbuzz(15)
	if result[14] != "FizzBuzz" {
		t.Errorf("expected FizzBuzz at index 14, got %s", result[14])
	}
}
