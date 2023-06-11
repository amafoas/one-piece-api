package models

type Crew struct {
	ID              string   `bson:"_id"`
	Name            string   `bson:"name"`
	RomanizedName   string   `bson:"romanized_name"`
	FirstAppearance []string `bson:"first_appearance"`
	Captain         string   `bson:"captain"`
	TotalBounty     string   `bson:"total_bounty"`
	MainShip        string   `bson:"main_ship"`
	Members         []string `bson:"members"`
	Allies          []string `bson:"allies"`
}
