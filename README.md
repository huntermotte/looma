# Looma Audition

## Getting started

Requirements:
* The `Go` programming language: <a>https://go.dev</a>

To start the server with default environment variables, run: `go mod tidy && go run main.go`

Server should be running at `http://locahost:8080`

#### Supported environment variables
| Variable  | Sample Value |
|:----------|:------------:|
| NUM_USERS |    **10**    |
| NUM_NOTES |  **500000**  |
| NUM_TASKS | **1000000**  |

These variables can be adjusted when starting the server to simulate various data source sizes for load testing.
Ex. `NUM_USERS=20 NUM_NOTES=1000 NUM_TASKS=1500 go run main.go`

### Interacting with Server

* To return the N most recent tasks and notes for a user, you can send a request like this:
`http://localhost:8080/user/1/recent`
* Optionally, you can specify a query parameter `limit` that overrides the default limit of 10 using this request: `http://localhost:8080/user/1/recent?limit=5
