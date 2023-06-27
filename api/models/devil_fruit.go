package models

type DevilFruit struct {
	ID              string   `bson:"_id" json:"_id"`
	Name            string   `bson:"name" json:"name"`
	Type            string   `bson:"type" json:"type"`
	Meaning         string   `bson:"meaning" json:"meaning"`
	FirstApparition []string `bson:"first_apparition" json:"first_apparition"`
	FirstUsage      []string `bson:"first_usage" json:"first_usage"`
	CurrentUser     string   `bson:"current_user" json:"current_user"`
	PreviousUser    string   `bson:"previous_user" json:"previous_user"`
}
