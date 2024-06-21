package domain

import "time"

// Delegation represents a Tezos delegation operation
type Delegation struct {
	Delegator string    `json:"delegator"` // Sender's address
	Timestamp time.Time `json:"timestamp"` // Time of delegation
	Amount    int64     `json:"amount"`    // Amount delegated
	Level     int64     `json:"level"`     // Block level
}
