package internal

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/ayt-szk/tweets-manager-go/pkg/domain/models"
)

func ExportTweets() error {
	fmt.Println("export tweets")
	if err := createTweetJsonFile(); err != nil {
		log.Printf("Create json file error: %#v", err)
		return err
	}

	if err := exportCsv(); err != nil {
		log.Printf("export json file error: %#v", err)
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

func exportCsv() error {
	// scan tweet.json
	file, err := ioutil.ReadFile("tmp/inputs/tweet.json")
	if err != nil {
		log.Printf("Open file error: %#v", err)
		return err
	}

	var tweets []models.TweetStructrure
	err = json.Unmarshal([]byte(file), &tweets)
	if err != nil {
		return err
	}

	csvFile, err := os.Create("tmp/outputs/tweet.csv")
	if err != nil {
		log.Printf("Create file error: %#v", err)
		return err
	}
	defer csvFile.Close()

	w := csv.NewWriter(csvFile)
	w.Comma = ','

	for _, t := range tweets {
		var row []string
		row = append(row, t.Tweet.ID)
		row = append(row, t.Tweet.IDStr)
		row = append(row, t.Tweet.Source)
		row = append(row, convNewline(t.Tweet.FullText, " "))
		// row = append(row, t.Tweet.Retweeted)
		row = append(row, t.Tweet.RetweetCount)
		// row = append(row, t.Tweet.Favorited)
		row = append(row, t.Tweet.FavoriteCount)
		// row = append(row, t.Tweet.Truncated)
		// row = append(row, t.Tweet.DisplayTextRange)
		row = append(row, t.Tweet.CreateAt)
		row = append(row, t.Tweet.Lang)
		// row = append(row, t.Tweet.Entities.Hashtags)
		// row = append(row, t.Tweet.Entities.UserMentions)

		w.Write(row)
		// w.Write()
	}

	w.Flush()

	return nil
}

func convNewline(str, repStr string) string {
	return strings.NewReplacer(
		"\r\n", repStr,
		"\r", repStr,
		"\n", repStr,
	).Replace(str)
}
