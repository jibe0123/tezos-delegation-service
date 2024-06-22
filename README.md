# 🚀 Tezos Delegation Service

This project implements a Golang service that gathers new delegations made on the Tezos protocol and exposes them through a public API.

## Features

- Polls new delegations from the Tzkt API endpoint
- Stores delegation data in a persistent store (SQLite)
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
      "timestamp": "2022-05-05T06:29:14Z",
      "amount": "125896",
      "delegator": "tz1a1SAaXRt9yoGMx29rh9FsBF4UzmvojdTL",
      "level": "2338084"
    },
    {
      "timestamp": "2021-05-07T14:48:07Z",
      "amount": "9856354",
      "delegator": "KT1JejNYjmQYh8yw95u5kfQDRuxJcaUPjUnf",
      "level": "1461334"
    }
  ]
}
```

## Environment Variables

- `TZKT_API_BASE_URL`: Base URL for the Tezos API (default: `https://api.tzkt.io/v1/`)
- `DATABASE_PATH`: Path to the SQLite database file (default: `delegations.db`)

## Running Locally

1. Clone the repository:
    ```sh
    git clone https://github.com/your_username/tezos-delegation-service.git
    cd tezos-delegation-service
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

3. Set up environment variables:
   - Create a `.env` file in the root directory of your project and add the following lines:
    ```env
    TZKT_API_BASE_URL=https://api.tzkt.io/v1/
    DATABASE_PATH=delegations.db
    ```

4. Run the service:
    ```sh
    go run cmd/main.go
    ```

5. The API will be available at `http://localhost:8080`.

## Running Tests

Run the following command to execute the tests:
```sh
go test ./...
```