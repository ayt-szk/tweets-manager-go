package models

type (
	TweetStructrure struct {
		Tweet tweet `json:"tweet"`
	}

	tweet struct {
		ID               string   `json:"id"`
		IDStr            string   `json:"id_str"`
		Source           string   `json:"source"`
		FullText         string   `json:"full_text"`
		Retweeted        bool     `json:"retweeted"`
		RetweetCount     string   `json:"retweet_count"`
		Favorited        bool     `json:"favorited"`
		FavoriteCount    string   `json:"favorite_count"`
		Truncated        bool     `json:"truncated"`
		DisplayTextRange []string `json:"display_text_range"`
		CreateAt         string   `json:"created_at"`
		Lang             string   `json:"lang"`
		Entities         entities
	}

	entities struct {
		HashTags     []hashtag      `json:"hashtags"`
		UserMentions []userMentions `json:"user_mentions"`
	}

	hashtag struct {
		Text string `json:"text"`
	}

	userMentions struct {
		ID         string `json:"id"`
		IDStr      string `json:"id_str"`
		Name       string `json:"name"`
		ScreenName string `json:"screen_name"`
	}
)
