# Go REST API

This is a simple Go web server that serves as a REST API. It is designed to help me practice the Go programming language.

## Main takeaways

- Go is an interesting language
- I like the concept of pointers and how by default Go passes by value
- I understood how setting up a REST API with a low-level HTTP module is difficult and time-consuming
- I learned how to use the `net/http` module to set up a web server
- I learned how to use the `encoding/json` module to encode and decode JSON
- I learned how to use the `strconv` module to convert strings to integers

## Prerequisites

Before running this application, make sure you have the following installed:

- Go (1.21.x)

## Getting Started

1. Clone this repository:

   ```shell
   git clone https://github.com/flacial/go-rest-api.git
   ```

2. Navigate to the project directory:

   ```shell
   cd go-rest-api
   ```

3. Build the application:

   ```shell
   go build
   ```

4. Run the application:

   ```shell
   ./go-rest-api
   ```

5. Open your web browser and visit `http://localhost:8080` to access the API.

## API Endpoints

- `GET /`: Home page
- `POST /increment-number`: Increment a number stored in memory.

---

It has been a smooth experience overall. I am looking forward to learning more about Go and using it in the future.
