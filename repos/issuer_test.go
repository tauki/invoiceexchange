package repos_test

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/tauki/invoiceexchange/internal/balance"
	domain "github.com/tauki/invoiceexchange/internal/issuer"
	"github.com/tauki/invoiceexchange/repos"
	"testing"
	"time"
)

func TestInvestorRepo_CreateIssuer(t *testing.T) {
	sqlMock, entClient, dbClose := setupSQLMock(t)
	defer dbClose()

	repo := repos.NewIssuerRepo(entClient)

	ctx := context.Background()
	issID := uuid.New()
	iss := &domain.Issuer{
		ID:       issID,
		Name:     "Issuer 1",
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
		WithArgs(iss.Balance.TotalAmount, iss.Balance.AvailableAmount, iss.ID, iss.Balance.CreatedAt, iss.Balance.UpdatedAt, iss.Balance.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectExec(`INSERT INTO "issuers" ("name", "joined_at", "issuer_balance", "id") VALUES ($1, $2, $3, $4)`).
		WithArgs(iss.Name, iss.JoinedAt, iss.Balance.ID, iss.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.CreateIssuer(ctx, iss)
	assert.NoError(t, err)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
}

func TestInvestorRepo_GetIssuerByID(t *testing.T) {
	sqlMock, entClient, dbClose := setupSQLMock(t)
	defer dbClose()

	repo := repos.NewIssuerRepo(entClient)

	ctx := context.Background()
	issID := uuid.New()

	sqlMock.ExpectQuery(`SELECT "issuers"."id", "issuers"."name", "issuers"."joined_at", "issuers"."issuer_balance" FROM "issuers" WHERE "issuers"."id" = $1 LIMIT 1`).
		WithArgs(issID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "balance_id", "joined_at", "balances_id", "balances_entity_id", "balances_total_amount", "balances_available_amount", "balances_created_at", "balances_updated_at"}).
			AddRow(issID, "Issuer 1", uuid.New(), time.Now(), uuid.New(), uuid.New(), 100, 100, time.Now(), time.Now()))

	inv, err := repo.GetIssuerByID(ctx, issID)
	assert.NoError(t, err)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.NotNil(t, inv)
}

func TestInvestorRepo_UpdateIssuer(t *testing.T) {
	sqlMock, entClient, dbClose := setupSQLMock(t)
	defer dbClose()

	repo := repos.NewIssuerRepo(entClient)

	ctx := context.Background()
	iss := &domain.Issuer{
		ID:       uuid.New(),
		Name:     "Issuer 1",
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
	sqlMock.ExpectExec(`UPDATE "issuers" SET "name" = $1 WHERE "id" = $2`).
		WithArgs(iss.Name, iss.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectQuery(`SELECT "id", "name", "joined_at" FROM "issuers" WHERE "id" = $1`).
		WithArgs(iss.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "joined_at"}).
			AddRow(iss.ID, "Issuer 1", time.Now()))
	sqlMock.ExpectCommit()
	sqlMock.ExpectBegin()
	sqlMock.ExpectExec(`UPDATE "balances" SET "total_amount" = $1, "available_amount" = $2, "updated_at" = $3 WHERE "id" = $4`).
		WithArgs(iss.Balance.TotalAmount, iss.Balance.AvailableAmount, iss.Balance.UpdatedAt, iss.Balance.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectQuery(`SELECT "id", "total_amount", "available_amount", "entity_id", "created_at", "updated_at" FROM "balances" WHERE "id" = $1`).
		WithArgs(iss.Balance.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "total_amount", "available_amount", "entity_id", "created_at", "updated_at"}).
			AddRow(iss.Balance.ID, iss.Balance.TotalAmount, iss.Balance.AvailableAmount, iss.Balance.EntityID, iss.Balance.CreatedAt, iss.Balance.UpdatedAt))
	sqlMock.ExpectCommit()
	err := repo.UpdateIssuer(ctx, iss)
	assert.NoError(t, err)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
}

func TestInvestorRepo_DeleteIssuer(t *testing.T) {
	sqlMock, entClient, dbClose := setupSQLMock(t)
	defer dbClose()

	repo := repos.NewIssuerRepo(entClient)

	ctx := context.Background()
	issID := uuid.New()

	sqlMock.ExpectExec(`DELETE FROM "issuers" WHERE "issuers"."id" = $1`).
		WithArgs(issID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.DeleteIssuer(ctx, issID)
	assert.NoError(t, err)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
}
