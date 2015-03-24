package models

import (
	"time"
)

type Thread struct {
	Id         string    `bson:"_id"`
	Title      string    `json:"title"`
	Author     string    `json:"author"`
	AuthorName string    `json:"authorName"`
	Created    time.Time `json:"created"`
	Updated    time.Time `json:"updated"`
	Count      int64     `json:"count"`
	Deleted    bool      `json:"deleted"`
	Locked     bool      `json:"locked"`
	Tags       []string  `json:"tags"`
}
