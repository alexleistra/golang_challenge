package main

import (
	"testing"
)

func TestFindMostCommonWord(t *testing.T) {
	findsMostCommonWord(t)
	findsTie(t)
	findsNone(t)
}

func findsMostCommonWord(t *testing.T) {
	words := []string{"something", "something", "dark", "side", "some", "more", "words", "about", "nothing"}
	result := findMostCommonWord(words)
	if result != "something" {
		t.Error()
	}
}

func findsTie(t *testing.T) {
	words := []string{"something", "words", "are", "random", "something", "dark", "dark", "side", "dark", "more", "words", "something"}
	result := findMostCommonWord(words)
	if result != "tie" {
		t.Error()
	}
}

func findsNone(t *testing.T) {
	words := []string{"all", "these", "words", "are", "unique"}
	result := findMostCommonWord(words)
	if result != "none" {
		t.Error()
	}
}
