package models

type Episode struct {
	Title           string   `bson:"title"`
	Release         string   `bson:"release"`
	RemasterRelease string   `bson:"remaster_release,omitempty"`
	Characters      []string `bson:"characters"`
	Season          int      `bson:"season"`
	Episode         int      `bson:"episode"`
	Locations       []string `bson:"locations"`
	Opening         string   `bson:"opening"`
}
