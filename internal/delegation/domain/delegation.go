package domain

// Delegation represents a Tezos delegation operation
type Delegation struct {
	Sender struct {
		Address string `json:"address"`
	} `json:"sender"` // Delegator Sender's address
	Timestamp string `json:"timestamp"` // Timestamp Time of delegation
	Amount    int64  `json:"amount"`    // Amount delegated
	Level     int64  `json:"level"`     // Level Block level
}
