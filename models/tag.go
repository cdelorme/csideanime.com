package models

// import (
// "gopkg.in/mgo.v2"
// "gopkg.in/mgo.v2/bson"
// )

// can I use a string?  Or do I **need** a bson.ObjectId?
// @note: just some of the thoughts for when I start testing with mongo/mgo

type Tag struct {
	Id          string   `bson:"_id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Parents     []string `json:"parents"`
}
