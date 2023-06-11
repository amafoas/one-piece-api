package models

type Chapter struct {
	Title   string `bson:"title"`
	Volume  int    `bson:"volume"`
	Chapter int    `bson:"chapter"`
	Pages   int    `bson:"pages"`
	Release string `bson:"release"`
}
