package models

import (
	"time"
)

type Post struct {
	Id         string    `bson:"_id" json:"id"`
	Message    string    `json:"message"`
	Author     string    `json:"author"`
	AuthorName string    `json:"authorName"`
	Created    time.Time `json:"created"`
	Updated    time.Time `json:"updated"`
	Deleted    bool      `json:"deleted"`
	Sticky     bool      `json:"sticky"`
	Tags       []string  `json:"tags"`
}
