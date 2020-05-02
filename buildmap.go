package anagrambf

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

// ReadWords takes a file name pointing to a file that contains single words
// separated by newlines.  It constructs an array of strings that is sorted
// alphabetically and has duplicates removed.  Further all words are converted
// to lowercase.
func ReadWords(fileName string) ([]string, error) {
	wordsFile, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("Could not open given words file %s with error: %v", fileName, err)
	}
	defer wordsFile.Close()

	// Read the words from the file, lower case them, and sort them all.
	var words []string
	scanner := bufio.NewScanner(wordsFile)
	for scanner.Scan() {
		words = append(words, strings.ToLower(scanner.Text()))
	}
	sort.Strings(words)

	// Remove duplicates.
	var dedupedWords []string
	lastWord := ""
	for _, word := range words {
		if word != lastWord {
			dedupedWords = append(dedupedWords, word)
			lastWord = word
		}
	}

	return dedupedWords, nil
}

// BuildMap constructs a map keyed by words that map to a bool which is
// true if the word is a prefix for another word and false otherwise.
// It is assumed that words is sorted.  As a result, it is wise to create
// words using the ReadWords function above.
func BuildMap(words []string) map[string]bool {
	var dictionary map[string]bool
	for i, word := range words {
		isPrefix := false
		if i < len(words) {
			nextWord := words[i+1]
			if strings.Contains(nextWord, word) {
				isPrefix = true
			}
		}
		dictionary[word] = isPrefix
	}

	return dictionary
}
