package main

import (
	"fmt"
	"os"
)

// launching point for this program
func main() {

	// get arguments skipping program name
	args := os.Args[1:]

	// process list of words and print result
	fmt.Println(findMostCommonWord(args))
}

// finds the most duplicated word in a list of words
func findMostCommonWord(words []string) string {
	var mostCommonWord1 string  // we only need to know the most duplicated word
	var mostCommonCount1 uint32 // keep track of the count of 1st and 2nd most common words so we know if there is a tie
	var mostCommonCount2 uint32

	wordCounts := make(map[string]uint32)

	for _, word := range words {
		// check if the map has the word and increment the count if it is already in the map
		if count, found := wordCounts[word]; found { // cool that this can be done in 1 line with Go
			newCount := count + 1
			wordCounts[word] = newCount
			if newCount >= mostCommonCount1 {
				// move #1 spot down to #2, this word is the new #1 most common
				mostCommonCount2 = mostCommonCount1
				mostCommonWord1 = word
				mostCommonCount1 = newCount
			}
		} else {
			// first time seeing this word
			wordCounts[word] = 1
		}
	}

	// no duplicates
	if mostCommonCount1 == 0 {
		return "none"
	}

	// there was a tie
	if mostCommonCount1 == mostCommonCount2 {
		return "tie"
	}

	// the most duplicated word
	return mostCommonWord1
}
