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
	"time"

	"github.com/ayt-szk/tweets-manager-go/pkg/domain/models"
)

func ExportTweets() error {
	if err := createTweetJsonFile(); err != nil {
		log.Printf("Create json file error: %#v", err)
		return err
	}

	if err := exportCsv(); err != nil {
		log.Printf("export csv file error: %#v", err)
		return err
	}

	return nil
}

func createTweetJsonFile() error {
	// write tweet.json
	jsonFile, err := os.OpenFile("tmp/inputs/tweet.json", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	// read tweet.js
	jsFile, err := os.Open("tmp/inputs/tweet.js")
	if err != nil {
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
			return err
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
		return err
	}

	var tweets []models.TweetStructrure
	err = json.Unmarshal([]byte(file), &tweets)
	if err != nil {
		return err
	}

	csvFile, err := os.Create("tmp/outputs/tweet.csv")
	if err != nil {
		return err
	}
	defer csvFile.Close()

	w := csv.NewWriter(csvFile)
	w.Comma = ','

	header := []string{
		"ID_STR",
		"FULL_TEXT",
		"HASHTAGS",
		"USER_MENTIONS",
		"RETWEET_COUNT",
		"FAVORITE_COUNT",
		"CREATE_AT",
	}
	w.Write(header)

	for _, t := range tweets {

		idStr := t.Tweet.IDStr
		fullText := convNewline(t.Tweet.FullText, " ")
		hashtags := joinHashtags(t.Tweet.Entities.Hashtags)
		userMentions := joinUserMentions(t.Tweet.Entities.UserMentions)
		retweetCount := t.Tweet.RetweetCount
		favoriteCount := t.Tweet.FavoriteCount
		createAt, err := convTimeToJst(t.Tweet.CreateAt)
		if err != nil {
			return err
		}

		row := []string{
			idStr,
			fullText,
			hashtags,
			userMentions,
			retweetCount,
			favoriteCount,
			createAt,
		}

		w.Write(row)
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

func convTimeToJst(createAt string) (string, error) {
	layout := "Mon Jan 2 15:04:05 +0000 2006"

	t, err := time.Parse(layout, createAt)
	if err != nil {
		return "", err
	}

	// UTC+9時間
	t = t.Add(9 * time.Hour)
	jstTime := t.Format("2006-01-02 15:04:05")

	return jstTime, nil
}

func joinHashtags(hashtags []models.HashtagEntity) string {
	var arrHashtags []string

	for _, hashtag := range hashtags {
		h := fmt.Sprintf("#%s", hashtag.Text)
		arrHashtags = append(arrHashtags, h)
	}
	joinedHashtags := strings.Join(arrHashtags, ",")

	return joinedHashtags
}

func joinUserMentions(userMentions []models.MentionEntity) string {
	var arrUserMentions []string

	for _, userMention := range userMentions {
		u := fmt.Sprintf("@%s", userMention.ScreenName)
		arrUserMentions = append(arrUserMentions, u)
	}
	joinedUserMentions := strings.Join(arrUserMentions, ",")

	return joinedUserMentions
}
