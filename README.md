# Looma Audition

## Getting started

`go mod tidy && go run main.go`

Server should be running at `http://locahost:8080`

### Interacting with Server

* To return the N most recent tasks and notes for a user, you can send a request like this:
`http://localhost:8080/user/1/recent`
* Optionally, you can specify a query parameter `limit` that overrides the default limit of 10 using this request: `http://localhost:8080/user/1/recent?limit=5

## Running Tests

To get a code coverage report:

## Considerations:
