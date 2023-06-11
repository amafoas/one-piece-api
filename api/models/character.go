package models

type Character struct {
	ID               string   `bson:"_id"`
	Name             string   `bson:"name"`
	Age              int      `bson:"age"`
	Status           string   `bson:"status"`
	DevilFruit       string   `bson:"devil_fruit,omitempty"`
	DevilFruitID     string   `bson:"devil_fruit_id,omitempty"`
	Debut            []string `bson:"debut"`
	MainAffiliation  string   `bson:"main_affiliation"`
	OtherAffiliation []string `bson:"other_affiliations"`
	Occupations      string   `bson:"occupations"`
	Origin           string   `bson:"origin"`
	Race             string   `bson:"race"`
	Bounty           string   `bson:"bounty,omitempty"`
	Birthday         string   `bson:"birthday"`
	Height           string   `bson:"height"`
}
