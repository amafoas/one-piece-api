package models

type DevilFruit struct {
	ID              string   `bson:"_id"`
	Name            string   `bson:"name"`
	Type            string   `bson:"type"`
	Meaning         string   `bson:"meaning"`
	FirstApparition []string `bson:"first_apparition"`
	FirstUsage      []string `bson:"first_usage"`
	CurrentUser     string   `bson:"current_user"`
	PreviousUser    string   `bson:"previous_user"`
}
