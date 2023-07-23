package main

import (
	"reflect"
	"testing"
)

func TestSumNumbers(t *testing.T) {
	var (
		expected = 10
		x        = 5
		y        = 5
	)

	have := sumNumbers(x, y)

	if have != expected {
		t.Errorf("Expected %d, but got %d", expected, have)
	}
}

func TestPlayersAreEqual(t *testing.T) {
	expected := Player{
		name: "One",
		hp:   100,
	}

	have := Player{
		name: "Two",
		hp:   100,
	}

	if !reflect.DeepEqual(expected, have) {
		t.Errorf("Expected %+v, but got %+v\n", expected, have)
	}
}
