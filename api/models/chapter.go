package models

type Chapter struct {
	ID      string `bson:"_id" json:"_id"`
	Title   string `bson:"title" json:"title"`
	Volume  int    `bson:"volume" json:"volume"`
	Chapter int    `bson:"chapter" json:"chapter"`
	Pages   int    `bson:"pages" json:"pages"`
	Release string `bson:"release" json:"release"`
}
