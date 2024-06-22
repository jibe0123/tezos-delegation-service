package repository

import (
	"technical-test/internal/delegation/domain"
	database "technical-test/pkg/sqlite"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaveAndFindAll(t *testing.T) {
	db, err := database.InitDB(":memory:")
	assert.NoError(t, err)
	defer db.Close()

	repo := NewRepository(db)

	delegation := domain.Delegation{
		Timestamp: "2022-05-05T06:29:14Z",
		Amount:    125896,
		Sender: struct {
			Address string `json:"address"`
		}{
			Address: "tz1a1SAaXRt9yoGMx29rh9FsBF4UzmvojdTL",
		},
		Level: 2338084,
	}

	err = repo.Save(delegation)
	assert.NoError(t, err)

	delegations, err := repo.FindAll("")
	assert.NoError(t, err)
	assert.Equal(t, 1, len(delegations))
	assert.Equal(t, delegation, delegations[0])
}
