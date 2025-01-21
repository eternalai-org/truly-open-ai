package mysql_models

import (
	"gorm.io/gorm"
	"time"
)

type TwitterTweet struct {
	gorm.Model

	TwitterID string `json:"twitter_id"`
	TweetId   string `json:"tweet_id"`
	FullText  string `json:"full_text"`

	RetweetCount  int       `json:"retweet_count"`
	FavoriteCount int       `json:"favorite_count"`
	PostedAt      time.Time `json:"posted_at"`
}

type TwitterTweetOpenAIFormat struct {
	TweetId        string    `json:"tweet_id"`
	Text           string    `json:"text"`
	CreatedAt      time.Time `json:"created_at"`
	UserId         string    `json:"user_id"`
	UserScreenName string    `json:"user_screen_name"`
	RetweetCount   int       `json:"retweet_count"`
	LikeCount      int       `json:"like_count"`
}

func (TwitterTweet) TableName() string {
	return "twitter_tweets"
}
