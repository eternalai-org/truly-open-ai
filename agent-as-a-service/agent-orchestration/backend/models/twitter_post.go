package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type TwitterPostStatus string
type TwitterPostType string

const (
	TwitterPostTypeText  TwitterPostType = "text"
	TwitterPostTypeImage TwitterPostType = "image"

	TwitterPostStatusNew            TwitterPostStatus = "new"
	TwitterPostStatusInvalid        TwitterPostStatus = "invalid"
	TwitterPostStatusValid          TwitterPostStatus = "valid"
	TwitterPostStatusInferNew       TwitterPostStatus = "infer_new"
	TwitterPostStatusInferSubmitted TwitterPostStatus = "infer_submitted"
	TwitterPostStatusInferFailed    TwitterPostStatus = "infer_failed"
	TwitterPostStatusInferResolved  TwitterPostStatus = "infer_resolved"
	TwitterPostStatusReplied        TwitterPostStatus = "replied"
)

type TwitterPost struct {
	gorm.Model
	TwitterID       string
	TwitterPostID   string `gorm:"unique_index"`
	Type            TwitterPostType
	PostAt          *time.Time
	Content         string `gorm:"type:longtext"`
	InferData       string `gorm:"type:longtext"`
	ReplyContent    string `gorm:"type:longtext"`
	ImageUrl        string
	InferId         uint `gorm:"default:0"`
	InferTxHash     string
	InferAt         *time.Time
	InferNum        uint `gorm:"default:0"`
	ResolveTxHash   string
	Status          TwitterPostStatus `gorm:"index"`
	IsGenerateImage bool              `gorm:"default:0"`
	Prompt          string            `gorm:"type:longtext"`
}

type UserAgentInferDataItem struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type TwitterTweet struct {
	gorm.Model
	TweetID          string `gorm:"unique_index"`
	TwitterID        string
	LikeCount        int
	RetweetCount     int
	ReplyCount       int
	QuoteCount       int
	ImpressionCount  int
	FullText         string `gorm:"type:longtext"`
	PostedAt         time.Time
	InReplyToUserID  string
	InReplyToTweetID string
	IsReply          bool
	OriginalText     string `gorm:"type:longtext"`

	// Retweet-related fields
	IsRetweet     bool
	RepostTweetID string
	RepostText    string `gorm:"type:longtext"`

	// Quote-related fields
	IsQuote      bool
	QuoteTweetID string
	QuoteText    string `gorm:"type:longtext"`
}

type TwitterUser struct {
	gorm.Model
	TwitterID       string `gorm:"index"`
	TwitterUsername string `gorm:"type:text collate utf8mb4_unicode_ci"`
	Name            string `gorm:"type:text collate utf8mb4_unicode_ci"`
	ProfileUrl      string `gorm:"type:text collate utf8mb4_unicode_ci"`
	FollowersCount  uint
	FollowingsCount uint
	IsBlueVerified  bool
	JoinedAt        time.Time
}

type TwitterFollowing struct {
	gorm.Model
	OwnerTwitterID  string `gorm:"index"`
	TwitterID       string `gorm:"index"`
	TwitterUsername string `gorm:"type:text collate utf8mb4_unicode_ci"`
	Name            string `gorm:"type:text collate utf8mb4_unicode_ci"`
	ProfileUrl      string `gorm:"type:text collate utf8mb4_unicode_ci"`
	FollowersCount  uint
	FollowingsCount uint
	IsBlueVerified  bool
	JoinedAt        *time.Time
}

type TwitterScan struct {
	gorm.Model
	Username  string `gorm:"type:varchar(256);"`
	TwitterID string `gorm:"type:varchar(256);"`
	Enabled   bool   `gorm:"default:true"`
	IsMention bool   `gorm:"default:false"`
	Scanned   bool   `gorm:"default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TwitterTweetLiked struct {
	gorm.Model
	LikedUserID      string `gorm:"unique_index:liked_main_unique_index;index"`
	TweetID          string `gorm:"unique_index:liked_main_unique_index;index"`
	TwitterID        string
	LikeCount        int
	RetweetCount     int
	ReplyCount       int
	QuoteCount       int
	ImpressionCount  int
	FullText         string `gorm:"type:longtext"`
	PostedAt         time.Time
	InReplyToUserID  string
	InReplyToTweetID string
	IsReply          bool
	OriginalText     string `gorm:"type:longtext"`

	// Retweet-related fields
	IsRetweet     bool
	RepostTweetID string
	RepostText    string `gorm:"type:longtext"`

	// Quote-related fields
	IsQuote      bool
	QuoteTweetID string
	QuoteText    string `gorm:"type:longtext"`
}
