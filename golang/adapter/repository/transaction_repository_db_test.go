package repository

import (
	"os"
	"testing"

	"github.com/jeisonborba/full-cycle-gateway/adapter/repository/fixture"
	"github.com/jeisonborba/full-cycle-gateway/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestTransactionRepositoryDbInsert(t *testing.T) {
	migrationsDir := os.DirFS("fixture/sql")
	db := fixture.Up(migrationsDir)
	defer fixture.Down(db, migrationsDir)

	repository := NewTransactionRepositoryDB(db)
	err := repository.Insert("1", "1", 50, entity.APPROVED, "")
	assert.Nil(t, err)
}
