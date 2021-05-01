package internal

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ExportTweets() error {
	fmt.Println("export tweets")
	if err := createTweetJsonFile(); err != nil {
		log.Printf("Create json file error: %#v", err)
		return err
	}

	return nil
}

func createTweetJsonFile() error {
	// write tweet.json
	jsonFile, err := os.Create("tmp/inputs/tweet.json")
	if err != nil {
		log.Printf("Create file error: %#v", err)
		return err
	}
	defer jsonFile.Close()

	// read tweet.js
	jsFile, err := os.Open("tmp/inputs/tweet.js")
	if err != nil {
		log.Printf("Open file error: %#v", err)
		return err
	}
	defer jsFile.Close()

	scanner := bufio.NewScanner(jsFile)

	lineCount := 1
	for scanner.Scan() {

		// replace "window.YTD.tweet.part0" to "[ { \n"
		if lineCount == 1 {
			jsonFile.WriteString("[ { \n")
			lineCount++
			continue
		}

		line := fmt.Sprintf("%s\n", scanner.Text())
		_, err := jsonFile.WriteString(line)
		if err != nil {
			log.Printf("Write file error: %#v", err)
		}

		lineCount++
	}

	log.Printf("Finished creating the tweet.json")
	return nil
}
