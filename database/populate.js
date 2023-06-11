// Connect to the database
var db = db.getSiblingDB('one-piece');

// Drop the database
db.dropDatabase();

// adding chapters
if (!db.chapters.count() === 0) {
  db.createCollection('chapters');
}
var chapters = JSON.parse(cat('./data/chapters.json'));
db.chapters.insertMany(chapters);

// adding characters
if (!db.characters.count() === 0) {
  db.createCollection('characters');
}
var characters = JSON.parse(cat('./data/characters.json'));
db.characters.insertMany(characters);

// adding crews
if (!db.crews.count() === 0) {
  db.createCollection('crews');
}
var crews = JSON.parse(cat('./data/crews.json'));
db.crews.insertMany(crews);

// adding devil fruits
if (!db.devil_fruits.count() === 0) {
  db.createCollection('devil_fruits');
}
var devil_fruits = JSON.parse(cat('./data/devil_fruits.json'));
db.devil_fruits.insertMany(devil_fruits);

// adding episodes
if (!db.episodes.count() === 0) {
  db.createCollection('episodes');
}
var episodes = JSON.parse(cat('./data/episodes.json'));
db.episodes.insertMany(episodes);
