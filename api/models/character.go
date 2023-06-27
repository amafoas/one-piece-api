package models

type Character struct {
	ID               string   `bson:"_id" json:"_id"`
	Name             string   `bson:"name" json:"name"`
	Age              int      `bson:"age" json:"age"`
	Status           string   `bson:"status" json:"status"`
	DevilFruit       string   `bson:"devil_fruit" json:"devil_fruit"`
	DevilFruitID     string   `bson:"devil_fruit_id" json:"devil_fruit_id"`
	Debut            []string `bson:"debut" json:"debut"`
	MainAffiliation  string   `bson:"main_affiliation" json:"main_affiliation"`
	OtherAffiliation []string `bson:"other_affiliations" json:"other_affiliations"`
	Occupations      string   `bson:"occupations" json:"occupations"`
	Origin           string   `bson:"origin" json:"origin"`
	Race             string   `bson:"race" json:"race"`
	Bounty           string   `bson:"bounty" json:"bounty"`
	Birthday         string   `bson:"birthday" json:"birthday"`
	Height           string   `bson:"height" json:"height"`
}
