package models

type Crew struct {
	ID              string   `bson:"_id" json:"_id"`
	Name            string   `bson:"name" json:"name"`
	RomanizedName   string   `bson:"romanized_name" json:"romanized_name"`
	FirstAppearance []string `bson:"first_appearance" json:"first_appearance"`
	Captain         string   `bson:"captain" json:"captain"`
	TotalBounty     string   `bson:"total_bounty" json:"total_bounty"`
	MainShip        string   `bson:"main_ship" json:"main_ship"`
	Members         []string `bson:"members" json:"members"`
	Allies          []string `bson:"allies" json:"allies"`
}
