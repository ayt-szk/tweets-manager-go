package internal

import (
	"fmt"
	"os"

	"github.com/ChimeraCoder/anaconda"
)

func GetTwitterApi() *anaconda.TwitterApi {
	anaconda.SetConsumerKey(os.Getenv("TWITTER_CONSUMER_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("TWITTER_CONSUMER_SECRET"))
	api := anaconda.NewTwitterApi(os.Getenv("TWITTER_ACCESS_TOKEN"), os.Getenv("TWITTER_ACCESS_TOKEN_SECRET"))

	fmt.Printf("TWITTER_CONSUMER_KEY:%s\n", os.Getenv("TWITTER_CONSUMER_KEY"))
	fmt.Printf("TWITTER_CONSUMER_SECRET:%s\n", os.Getenv("TWITTER_CONSUMER_SECRET"))
	fmt.Printf("TWITTER_ACCESS_TOKEN:%s\n", os.Getenv("TWITTER_ACCESS_TOKEN"))
	fmt.Printf("TWITTER_ACCESS_TOKEN_SECRET:%s\n", os.Getenv("TWITTER_ACCESS_TOKEN_SECRET"))

	return api
}
