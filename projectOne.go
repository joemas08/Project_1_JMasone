package main //Joe Masone's Project

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

func main() {
	var fileName string
	fmt.Println("Please enter the file's name: ")
	_, err := fmt.Scan(&fileName) //getting file name from user
	if err != nil {
		return
	}

	fileasBytes, err := ioutil.ReadFile(fileName) //converting file to bytes
	if err != nil {
		log.Fatalln("Error reading this file: ", err)
	}
	fileContents := string(fileasBytes)

	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(fileContents, " ") //filtering file passed
	words := strings.Split(processedString, " ")

	reportResults(countWords(words))
} //end main

func countWords(wordsPassed []string) map[string]int {
	wordMap := make(map[string]int)

	for _, word := range wordsPassed {
		word = strings.ToLower(word)
		_, check := wordMap[word]
		if word == "" { // - was adding a blank key so skipped iteration if it comes across the blank
			continue
		}
		if check {
			wordMap[word] += 1
		} else {
			wordMap[word] = 1
		}
	}
	return wordMap
} //end countWords

func reportResults(wordMapPassed map[string]int) {
	for key, numTimes := range wordMapPassed {
		if numTimes < 2 {
			fmt.Println("Word:", key, "| Usage:", numTimes, "time")
		} else {
			fmt.Println("Word:", key, "| Usage:", numTimes, "times")
		}
	}
} //end reportResults
