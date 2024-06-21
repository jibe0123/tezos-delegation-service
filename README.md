# Tezos Delegation Service 🚀

## Description 📋

This service collects new delegations made on the Tezos protocol and exposes them through a public API.

## Requirements 🛠️

- Go 1.16 or later
- `godotenv` package

## Installation 💻

1. Clone the repository:
   ```sh
   git clone https://github.com/jibe0123/tezos-delegation-service.git
   cd tezos-delegation-service
   ```

2. Install dependencies:
   ```sh
   go mod tidy
   ```

3. Create a `.env` file at the root of the project with the following content:
   ```plaintext
   TZKT_API_BASE_URL=https://api.tzkt.io
   ```

## Usage 🚀

1. Start the service:
   ```sh
   go run cmd/main.go
   ```

2. The service will start polling the Tezos Tzkt API and storing delegation data in memory.

3. Access the delegations through the public API at `http://localhost:8080/xtz/delegations`.

### Example API Request 🔍

- **Endpoint**: `/xtz/delegations`
- **Method**: GET
- **Query Parameters**:
    - `year` (optional): Filter delegations by the specified year (format: YYYY)

### Example API Response 📄

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

## Testing 🧪

To run tests:

```sh
go test ./...
```

## Project Structure 🗂️

- `cmd/`: Contains the entry point of the application.
- `internal/api/`: Contains the HTTP handler.
- `internal/domain/`: Contains domain models.
- `internal/repository/`: Contains the repository interface and in-memory implementation.
- `internal/service/`: Contains the business logic.
- `internal/sync/`: Contains the polling logic.
- `pkg/tzkt/`: Contains the client for the Tzkt API.

## Environment Variables 🌍

- `TZKT_API_BASE_URL`: The base URL for the Tzkt API.