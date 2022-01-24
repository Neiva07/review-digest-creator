package persistence

import "time"

type Application struct {
	ApplicationId string    `bson:"_id"`
	SleepTime     uint64    `bson:"sleepTime"`
	LastDigest    time.Time `bson:"lastDigest,omitempty"`
	StartDigestAt time.Time `bson:"startDigestAt"`
	IsActive      bool      `bson:"isActive"`
	PageSize      uint64    `bson:"pageSize"`
}

type Review struct {
	ApplciationId string    `bson:"applicationId"`
	ReviewId      string    `bson:"_id"`
	Title         string    `bson:"title"`
	PostedAt      time.Time `bson:"postedAt"`
	Rating        uint8     `bson:"rating"`
	Author        string    `bson:"author"`
	Content       string    `bson:"content"`
	WasDigested   bool      `bson:"wasDigested"`
}
