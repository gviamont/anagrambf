package main

import (
	"fmt"
	"go/build"
	"log"
	"os"

	"github.com/gviamont/anagrambf"
)

func main() {
	if (len(os.Args) < 2) || (len(os.Args) > 3) {
		fmt.Println("usage: anagram-brute-forcer anagram [dictionary_file]")
		os.Exit(1)
	}

	// Open the dictionary file and organize the words.
	wordsFileName := build.Default.GOPATH + "/src/github.com/gviamont/anagrambf/words.txt"
	if len(os.Args) == 3 {
		wordsFileName = os.Args[2]
	}
	words, err := anagrambf.ReadWords(wordsFileName)
	if err != nil {
		log.Fatal(err)
	}

	// Construct the dictionary prefix map.
	dictionary := anagrambf.BuildMap(words)
	for key, value := range dictionary {
		fmt.Printf("%s: %t\n", key, value)
	}
}
