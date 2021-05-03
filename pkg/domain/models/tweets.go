package models

type (
	TweetStructrure struct {
		Tweet *Tweet `json:"tweet"`
	}

	Tweet struct {
		IDStr         string `json:"id_str"`
		FullText      string `json:"full_text"`
		Retweeted     bool   `json:"retweeted"`
		RetweetCount  string `json:"retweet_count"`
		Favorited     bool   `json:"favorited"`
		FavoriteCount string `json:"favorite_count"`
		CreateAt      string `json:"created_at"`
		Source        string `json:"source"`
		Entities      *Entities
	}

	Entities struct {
		Hashtags     []HashtagEntity `json:"hashtags"`
		UserMentions []MentionEntity `json:"user_mentions"`
	}

	MentionEntity struct {
		IDStr      string `json:"id_str"`
		Name       string `json:"name"`
		ScreenName string `json:"screen_name"`
	}

	HashtagEntity struct {
		Text string `json:"text"`
	}
)
