package repos_test

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/tauki/invoiceexchange/internal/balance"
	domain "github.com/tauki/invoiceexchange/internal/investor"
	"github.com/tauki/invoiceexchange/repos"
	"testing"
	"time"
)

func TestInvestorRepo_CreateInvestor(t *testing.T) {
	sqlMock, entClient, dbClose := setupSQLMock(t)
	defer dbClose()

	repo := repos.NewInvestorRepo(entClient)

	ctx := context.Background()
	invID := uuid.New()
	inv := &domain.Investor{
		ID:       invID,
		Name:     "Investor 1",
		JoinedAt: time.Now(),
		Balance: &balance.Balance{
			ID:              uuid.New(),
			TotalAmount:     100,
			AvailableAmount: 100,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		},
	}

	sqlMock.ExpectExec(`INSERT INTO "balances" ("total_amount", "available_amount", "entity_id", "created_at", "updated_at", "id") VALUES ($1, $2, $3, $4, $5, $6)`).
		WithArgs(inv.Balance.TotalAmount, inv.Balance.AvailableAmount, inv.ID, inv.Balance.CreatedAt, inv.Balance.UpdatedAt, inv.Balance.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectExec(`INSERT INTO "investors" ("name", "joined_at", "investor_balance", "id") VALUES ($1, $2, $3, $4)`).
		WithArgs(inv.Name, inv.JoinedAt, inv.Balance.ID, inv.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.CreateInvestor(ctx, inv)
	assert.NoError(t, err)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
}

func TestInvestorRepo_GetInvestorByID(t *testing.T) {
	sqlMock, entClient, dbClose := setupSQLMock(t)
	defer dbClose()

	repo := repos.NewInvestorRepo(entClient)

	ctx := context.Background()
	invID := uuid.New()

	sqlMock.ExpectQuery(`SELECT "investors"."id", "investors"."name", "investors"."joined_at", "investors"."investor_balance" FROM "investors" WHERE "investors"."id" = $1 LIMIT 1`).
		WithArgs(invID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "balance_id", "joined_at", "balances_id", "balances_entity_id", "balances_total_amount", "balances_available_amount", "balances_created_at", "balances_updated_at"}).
			AddRow(invID, "Investor 1", uuid.New(), time.Now(), uuid.New(), uuid.New(), 100, 100, time.Now(), time.Now()))

	inv, err := repo.GetInvestorByID(ctx, invID)
	assert.NoError(t, err)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.NotNil(t, inv)
}

func TestInvestorRepo_UpdateInvestor(t *testing.T) {
	sqlMock, entClient, dbClose := setupSQLMock(t)
	defer dbClose()

	repo := repos.NewInvestorRepo(entClient)

	ctx := context.Background()
	inv := &domain.Investor{
		ID:       uuid.New(),
		Name:     "Investor 1",
		JoinedAt: time.Now(),
		Balance: &balance.Balance{
			ID:              uuid.New(),
			EntityID:        uuid.New(),
			TotalAmount:     100,
			AvailableAmount: 100,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		},
	}

	sqlMock.ExpectBegin()
	sqlMock.ExpectExec(`UPDATE "investors" SET "name" = $1 WHERE "id" = $2`).
		WithArgs(inv.Name, inv.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectQuery(`SELECT "id", "name", "joined_at" FROM "investors" WHERE "id" = $1`).
		WithArgs(inv.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "joined_at"}).
			AddRow(inv.ID, "Investor 1", time.Now()))
	sqlMock.ExpectCommit()
	sqlMock.ExpectBegin()
	sqlMock.ExpectExec(`UPDATE "balances" SET "total_amount" = $1, "available_amount" = $2, "updated_at" = $3 WHERE "id" = $4`).
		WithArgs(inv.Balance.TotalAmount, inv.Balance.AvailableAmount, inv.Balance.UpdatedAt, inv.Balance.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectQuery(`SELECT "id", "total_amount", "available_amount", "entity_id", "created_at", "updated_at" FROM "balances" WHERE "id" = $1`).
		WithArgs(inv.Balance.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "total_amount", "available_amount", "entity_id", "created_at", "updated_at"}).
			AddRow(inv.Balance.ID, inv.Balance.TotalAmount, inv.Balance.AvailableAmount, inv.Balance.EntityID, inv.Balance.CreatedAt, inv.Balance.UpdatedAt))
	sqlMock.ExpectCommit()
	err := repo.UpdateInvestor(ctx, inv)
	assert.NoError(t, err)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
}

func TestInvestorRepo_DeleteInvestor(t *testing.T) {
	sqlMock, entClient, dbClose := setupSQLMock(t)
	defer dbClose()

	repo := repos.NewInvestorRepo(entClient)

	ctx := context.Background()
	invID := uuid.New()

	sqlMock.ExpectExec(`DELETE FROM "investors" WHERE "investors"."id" = $1`).
		WithArgs(invID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.DeleteInvestor(ctx, invID)
	assert.NoError(t, err)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
}
