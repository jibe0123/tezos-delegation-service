# 🚀 Tezos Delegation Service

This project implements a Golang service that gathers new delegations made on the Tezos protocol and exposes them through a public API.

## Features

- Polls new delegations from the Tzkt API endpoint
- Stores delegation data in a persistent store (MariaDB)
- Exposes the collected data through a public API endpoint `/xtz/delegations`

## API

### GET /xtz/delegations

- Retrieves all delegations, optionally filtered by year.
- Query Parameters:
   - `year` (optional): Filter delegations by year (format: YYYY)

Example Response:
```json
{
   "data": [
      {
         "timestamp": "2024-06-25T18:28:51Z",
         "amount": "99889244",
         "delegator": "tz1QBRYrGuieidHTkjeCtgJNUeCc34C2Bd6L",
         "level": "6800040"
      },
      {
         "timestamp": "2024-06-25T18:28:04Z",
         "amount": "99889748",
         "delegator": "tz1QBRYrGuieidHTkjeCtgJNUeCc34C2Bd6L",
         "level": "6800032"
      }
   ]
}
```

## Environment Variables

- `TZKT_API_BASE_URL`: Base URL for the Tezos API (default: `https://api.tzkt.io/v1/`)
- `DB_HOST`: Database host (default: `localhost`)
- `DB_PORT`: Database port (default: `3306`)
- `DB_USER`: Database user (default: `root`)
- `DB_PASSWORD`: Database password (default: `password`)
- `DB_NAME`: Database name (default: `tezos_delegations`)

## Running Locally

1. Clone the repository:
    ```sh
    git clone https://github.com/jibe0123/tezos-delegation-service.git
    cd tezos-delegation-service
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

3. Run the service:
    ```sh
    go run cmd/main.go
    ```

4. The API will be available at `http://localhost:8080`.

## Running with Docker

1. Build and run the containers:
    ```sh
    docker-compose up --build
    ```

2. The API will be available at `http://localhost:8080`.

## Running with Makefile

1. Build the Docker containers:
    ```sh
    make build
    ```

2. Run the Docker containers:
    ```sh
    make run
    ```

3. Clean up the Docker environment:
    ```sh
    make clean
    ```

## Running Tests

Run the following command to execute the tests:
```sh
go test ./...
```

## Makefile Features

- **build**: Builds the Docker containers without using the cache.
- **run**: Cleans up any previous instances and runs the Docker containers, rebuilding them if necessary.
- **clean**: Stops and removes the Docker containers, volumes, and images.
- **full-clean**: Performs a clean operation and also removes all Docker images and prunes volumes and the system.
- **rebuild**: Rebuilds and recreates the Docker containers.
- **clean-images**: Removes all Docker images and prunes volumes and the system.

## Docker Compose Features

- **app**: Builds and runs the Go application, exposing port 8080.
- **db**: Runs a MariaDB database with the necessary environment variables and volume configurations.