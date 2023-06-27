# One Piece API

This is a GoLang API project that provides information about the popular manga and anime series, One Piece. The API allows you to retrieve data related to chapters, characters, crews, episodes, and devil fruits. The project is built using the Gin web framework, MongoDB for data storage, and Testify for testing.

## Installation

To use this API locally, make sure you have Go installed on your system. Then, follow these steps:

1. Clone the repository:

   ```bash
   git clone https://github.com/amafoas/one-piece-api
   ```

2. Navigate to the project directory:

   ```bash
   cd one-piece-api
   ```

3. Install the project dependencies:

   ```bash
   go mod download
   ```

4. Start the API server:

   ```bash
   go run main.go
   ```

The API server should now be running locally on `http://localhost:8080`.

## Usage

The following routes are available in the API:

- `GET /episode/:id`: Retrieve information about anime episode by its number.
  <details>
   <summary>Return object</summary>
     
     ```json
     {
       "_id": "1",
       "title": "I'm Luffy! The Man Who Will Become the Pirate King!",
       "release": "October 20, 1999",
       "remaster_release": "April 7, 2012",
       "characters": [
         "Nami",
         "Alvida",
         "Koby",
         "Heppoko",
         "Peppoko",
         "Poppoko",
         "Monkey D Luffy",
         "Roronoa Zoro"
       ],
       "season": 1,
       "episode": 1,
       "locations": [
         "Shells Town"
       ],
       "opening": "We Are!"
     }
     ```
  </details>
- `POST /episode`: Create a new episode.
- `PUT /episode/:id`: Update an existing episode.
- `DELETE /episode/:id`: Delete an episode.

---

- `GET /chapter/:id`: Retrieve information about manga chapter by its number.
  <details>
   <summary>Return object</summary>
     
     ```json
     {
       "_id": "1",
       "title": "Romance Dawn - The Dawn of the Adventure",
       "volume": 1,
       "chapter": 1,
       "pages": 53,
       "release": "July 19, 1997"
     }
     ```
  </details>
- `POST /chapter`: Create a new chapter.
- `PUT /chapter/:id`: Update an existing chapter.
- `DELETE /chapter/:id`: Delete a chapter.

---

- `GET /devil-fruit/:id`: Get details about a devil fruit by its ID.
  <details>
   <summary>Return object</summary>
     
     ```json
     {
       "_id": "gomu_gomu_no_mi",
       "name": "Gomu Gomu no Mi",
       "type": "Paramecia",
       "meaning": "Rubber Human",
       "first_apparition": [
         "Chapter 1",
         "Episode 4"
       ],
       "first_usage": [
         "Chapter 1",
         "Episode 1"
       ],
       "current_user": "Monkey D Luffy",
       "previous_user": ""
     }
     ```
  </details>
- `POST /devil-fruit`: Create a new devil fruit.
- `PUT /devil-fruit/:id`: Update an existing devil fruit.
- `DELETE /devil-fruit/:id`: Delete a devil fruit.

---

- `GET /crew/:id`: Retrieve information about a crew by its ID.
  <details>
   <summary>Return object</summary>
     
     ```json
     {
       "_id": "straw_hat_pirates",
       "name": "Straw Hat Pirates",
       "romanized_name": "Mugiwara no Ichimi",
       "first_appearance": [
         "Chapter 5",
         "Episode 3"
       ],
       "captain": "Monkey D Luffy",
       "total_bounty": "8,816,001,000",
       "main_ship": "Thousand Sunny",
       "members": [
         "Monkey D Luffy",
         "Roronoa Zoro",
         "Nami",
         "Usopp",
         "Sanji",
         "Tony Tony Chopper",
         "Nico Robin",
         "Franky",
         "Brook",
         "Jinbe",
         "Nefertari Vivi"
       ],
       "allies": [
         "Ninja-Pirate-Mink-Samurai Alliance",
         "Galley-La Company",
         "Franky Family",
         "Fire Tank Pirates"
       ]
     }
     ```
  </details>
- `POST /crew`: Create a new crew.
- `PUT /crew/:id`: Update an existing crew.
- `DELETE /crew/:id`: Delete a crew.

---

- `GET /character/:id`: Get details about a character by their ID.
  <details>
   <summary>Return object</summary>
     
     ```json
     {
       "_id": "monkey_d_luffy",
       "name": "Monkey D Luffy",
       "age": 19,
       "status": "Alive",
       "devil_fruit": "Gomu Gomu no Mi",
       "devil_fruit_id": "gomu_gomu_no_mi",
       "debut": [
         "Chapter 1",
         "Episode 1"
       ],
       "main_affiliation": "Straw Hat Pirates",
       "other_affiliations": [
         "Dadan Family",
         "Ninja-Pirate-Mink-Samurai Alliance"
       ],
       "occupations": "Pirate Captain",
       "origin": "East Blue",
       "race": "Human",
       "bounty": "3,000,000,000",
       "birthday": "May 5th",
       "height": "174 cm"
     }
     ```
  </details>
- `POST /character`: Create a new character.
- `PUT /character/:id`: Update an existing character.
- `DELETE /character/:id`: Delete a character.

Make HTTP requests to the corresponding routes to perform the desired operations.

### Example
These are some examples of use with the `curl` command

**Create a new chapter:**
```bash
curl -X POST -H "Content-Type: application/json" -d '{
  "_id": "1",
  "title": "Romance Dawn - The Dawn of the Adventure",
  "volume": 1,
  "chapter": 1,
  "pages": 53,
  "release": "July 19, 1997"
}' "http://localhost:8080/chapter"
```

**Read a chapter:**
```bash
curl "http://localhost:8080/chapter/1"
```

**Update a chapter:**
```bash
curl -X PUT -H "Content-Type: application/json" -d '{
   "title": "updated title"
}' "http://localhost:8080/chapter/1"
```

**Delete a chapter:**
```bash
curl -X DELETE "http://localhost:8080/chapter/1"
```


## Testing

This project includes tests to ensure the correctness of the API routes. To run the tests, execute the following command:

```bash
go test ./test
```
