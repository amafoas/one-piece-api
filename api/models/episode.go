package models

type Episode struct {
	ID              string   `bson:"_id" json:"_id"`
	Title           string   `bson:"title" json:"title"`
	Release         string   `bson:"release" json:"release"`
	RemasterRelease string   `bson:"remaster_release" json:"remaster_release"`
	Characters      []string `bson:"characters" json:"characters"`
	Season          int      `bson:"season" json:"season"`
	Episode         int      `bson:"episode" json:"episode"`
	Locations       []string `bson:"locations" json:"locations"`
	Opening         string   `bson:"opening" json:"opening"`
}
