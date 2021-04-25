package main

import (
	"testing"
)

// runs all tests for findMostCommonWord
func TestFindMostCommonWord(t *testing.T) {
	findsMostCommonWord(t)
	findsTie(t)
	findsNone(t)
}

// checks that findsMostCommonWord finds the most duplicated word in a list of words
func findsMostCommonWord(t *testing.T) {
	words := []string{"something", "something", "dark", "side", "some", "more", "words", "about", "nothing"}
	result := findMostCommonWord(words)
	if result != "something" {
		t.Error()
	}
}

// checks that findsMostCommonWord correctly finds a tie
func findsTie(t *testing.T) {
	words := []string{"something", "words", "are", "random", "something", "dark", "dark", "side", "dark", "more", "words", "something"}
	result := findMostCommonWord(words)
	if result != "tie" {
		t.Error()
	}
}

// checks that findsMostCommonWord correctly determines when there are no duplicated words
func findsNone(t *testing.T) {
	words := []string{"all", "these", "words", "are", "unique"}
	result := findMostCommonWord(words)
	if result != "none" {
		t.Error()
	}
}
