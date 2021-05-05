package internal

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func DeleteTweets() error {
	tweetIds, err := getTweetID()
	if err != nil {
		log.Printf("Failed to get tweet ids. %s", err)
		return err
	}
	fmt.Println(tweetIds)

	api := GetTwitterApi()

	excludeTweets := []int64{
		// ex. 999999999999999999
	}

	//Delete tweets
	for _, tweetId := range tweetIds {
		if contains(excludeTweets, tweetId) {
			log.Printf("Skip tweet(%d)", tweetId)
			continue
		}

		_, err := api.DeleteTweet(tweetId, true)
		if err != nil {
			log.Printf("Failed to delte tweet(%d)", tweetId)
			log.Printf("[ERROR]: %v", err)
			continue
		}

		log.Printf("Deleted tweet(%d)", tweetId)
	}

	log.Printf("Finished deleteing the tweets.")
	return nil
}

func getTweetID() ([]int64, error) {
	var tweetIds []int64

	f, err := os.Open("tmp/outputs/tweet.csv")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.LazyQuotes = true

	// Get header
	header, _ := reader.Read()
	fmt.Println(header)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		iTweetId, err := strconv.ParseInt(record[0], 10, 64)
		if err != nil {
			log.Printf("ParseInt error: %s", err)
			return nil, err
		}

		tweetIds = append(tweetIds, iTweetId)
	}

	return tweetIds, nil
}

func contains(s []int64, e int64) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}
