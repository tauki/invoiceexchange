package repos_test

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/tauki/invoiceexchange/internal/balance"
	"github.com/tauki/invoiceexchange/repos"
	"testing"
	"time"
)

func TestBalanceRepo_CreateBalance(t *testing.T) {
	sqlMock, entClient, dbClose := setupSQLMock(t)
	defer dbClose()

	repo := repos.NewBalanceRepo(entClient)

	ctx := context.Background()
	b := &balance.Balance{
		ID:              uuid.New(),
		EntityID:        uuid.New(),
		TotalAmount:     100,
		AvailableAmount: 100,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	sqlMock.ExpectExec(`INSERT INTO "balances" ("total_amount", "available_amount", "entity_id", "created_at", "updated_at", "id") VALUES ($1, $2, $3, $4, $5, $6)`).
		WithArgs(b.TotalAmount, b.AvailableAmount, b.EntityID, b.CreatedAt, b.UpdatedAt, b.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.CreateBalance(ctx, b)
	assert.NoError(t, err)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
}

func TestBalanceRepo_UpdateBalance(t *testing.T) {
	sqlMock, entClient, dbClose := setupSQLMock(t)
	defer dbClose()

	repo := repos.NewBalanceRepo(entClient)
	ctx := context.Background()

	b := &balance.Balance{
		ID:              uuid.New(),
		EntityID:        uuid.New(),
		TotalAmount:     100,
		AvailableAmount: 100,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	sqlMock.ExpectBegin()
	sqlMock.ExpectExec(`UPDATE "balances" SET "total_amount" = $1, "available_amount" = $2, "updated_at" = $3 WHERE "id" = $4`).
		WithArgs(b.TotalAmount, b.AvailableAmount, b.UpdatedAt, b.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectQuery(`SELECT "id", "total_amount", "available_amount", "entity_id", "created_at", "updated_at" FROM "balances" WHERE "id" = $1`).
		WithArgs(b.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "entity_id", "total_amount", "available_amount", "created_at", "updated_at"}).
			AddRow(b.ID, b.EntityID, b.TotalAmount, b.AvailableAmount, b.CreatedAt, b.UpdatedAt))
	sqlMock.ExpectCommit()

	err := repo.UpdateBalance(ctx, b)
	assert.NoError(t, err)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
}

func TestBalanceRepo_GetBalance(t *testing.T) {
	sqlMock, entClient, dbClose := setupSQLMock(t)
	defer dbClose()

	repo := repos.NewBalanceRepo(entClient)
	ctx := context.Background()

	b := &balance.Balance{
		ID:              uuid.New(),
		EntityID:        uuid.New(),
		TotalAmount:     100,
		AvailableAmount: 100,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	sqlMock.ExpectQuery(`SELECT "balances"."id", "balances"."total_amount", "balances"."available_amount", "balances"."entity_id", "balances"."created_at", "balances"."updated_at" FROM "balances" WHERE "balances"."entity_id" = $1 LIMIT 1`).
		WithArgs(b.EntityID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "entity_id", "total_amount", "available_amount", "created_at", "updated_at"}).
			AddRow(b.ID, b.EntityID, b.TotalAmount, b.AvailableAmount, b.CreatedAt, b.UpdatedAt))

	bal, err := repo.GetBalance(ctx, b.EntityID)
	assert.NoError(t, err)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Equal(t, b, bal)
}
