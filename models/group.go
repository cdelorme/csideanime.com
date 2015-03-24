package models

type Group struct {
	Id         string   `bson:"_id" json:"id"`
	Name       string   `json:"name"`
	Actions    []string `json:"actions"`
	Icon       string   `json:"icon"`
	Management bool     `json:"management"`
	Dogtag     bool     `json:"dogtag"`
}
