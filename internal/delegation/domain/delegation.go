package domain

import "time"

// Delegation represents a Tezos delegation operation
type Delegation struct {
	TxId      int64     `json:"id"`              // Unique ID of the operation, stored in the TzKT indexer database
	Sender    Address   `json:"sender"`          // Information about the delegated account
	Timestamp time.Time `json:"timestamp"`       // Datetime of the block, in which the operation was included (ISO 8601)
	Amount    int64     `json:"amount"`          // Sender's balance at the time of delegation operation (aka delegation amount)
	Level     int32     `json:"level"`           // The height of the block from the genesis block, in which the operation was included
	Block     *string   `json:"block,omitempty"` // Hash of the block, in which the operation was included (nullable)
	Hash      *string   `json:"hash,omitempty"`  // Hash of the operation (nullable)
}

// Address represents an address in the Tezos network
type Address struct {
	Address string `json:"address"`
}
