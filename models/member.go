package models

import (
	"time"
)

type MemberNames struct {
	Name string
	On   time.Time
}

type MemberProfile struct {
	Name   string        `json:"name"`
	Names  []MemberNames `json:"names"`
	Email  string        `json:"email"`
	Joined time.Time     `json:"joined"`
}

// I have not verified this structure just yet, it may vary in the future
type MemberOAuth struct {
	ServiceName string
	Username    string
	Token       string
}

type Member struct {
	Id       string        `bson:"_id" json:"id"`
	Name     string        `json:"name"`
	Password string        `json:"-"`
	Profile  MemberProfile `json:"memberProfile"`
	Private  MemberProfile `json:"-"`
	Settings map[string]string
	Oauth    MemberOAuth `json:"-"`
	Groups   []string    `json:"groups"`
}
