package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	gg "go_reloaded/modifytext"
)

func modifications(words []string) []string {
	words = gg.Hexadecimal(words)
	words = gg.Binary(words)
	words = gg.Capitalize(words)
	words = gg.UpperCase(words)
	words = gg.LowerCase(words)
	words = gg.VowelArticle(words)
	words = gg.FixPunctuation(words)
	words = gg.FixApostrophe(words)

	return words
}

// read file in the command line, convert/modify data then write file
func main() {
	if len(os.Args) != 3 { // files sample.txt result.txt
		log.Fatal("Incorrect filename arguments")
		return
	}

	args := os.Args[1:]
	data, err := os.ReadFile(args[0])
	if err != nil {
		errorMsg := fmt.Sprintf("Error reading file '%s': %v", args[0], err)
		log.Fatal(errorMsg)
	}

	words := strings.Split(string(data), " ")
	words = modifications(words)
	final := strings.Join(words, " ") + "\n"

	error := os.WriteFile(args[1], []byte(final), 0o644)
	if error != nil {
		errorMsg := fmt.Sprintf("Error writing file '%s': %v", args[1], err)
		log.Fatal(errorMsg)
	}
}

// func main() {
// open sample.txt
// read contents of sample.txt
// scan contents of sample.txt
// modify the contents of sample.txt
// create result.txt
// write into result.txt
