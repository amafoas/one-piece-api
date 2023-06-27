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

- `GET /episode/:num`: Retrieve information about an episode by its number.
- `GET /devil-fruit/:id`: Get details about a devil fruit by its ID.
- `GET /crew/:id`: Retrieve information about a crew by its ID.
- `GET /character/:id`: Get details about a character by their ID.
- `GET /character/fruit/:fruit`: Get details about characters who possess a specific devil fruit.
- `GET /chapter/:num`: Retrieve information about a chapter by its number.

Make HTTP requests to the corresponding routes to retrieve the desired data.

## Testing

This project includes tests to ensure the correctness of the API routes. To run the tests, execute the following command:

```bash
go test ./test
```
