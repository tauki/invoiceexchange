package repos_test

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/tauki/invoiceexchange/internal/bid"
	"github.com/tauki/invoiceexchange/repos"
)

func TestBidRepo_CreateBid(t *testing.T) {
	sqlMock, entClient, dbClose := setupSQLMock(t)
	defer dbClose()

	repo := repos.NewBidRepo(entClient)

	ctx := context.Background()
	b := &bid.Bid{
		ID:         uuid.New(),
		InvoiceID:  uuid.New(),
		InvestorID: uuid.New(),
		Amount:     100,
		Status:     bid.Pending}

	sqlMock.ExpectExec(
		`INSERT INTO "bids" ("status", "amount", "accepted_amount", "created_at", "updated_at", "invoice_bids", "investor_bids", "id") VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`).
		WithArgs(b.Status, b.Amount, b.AcceptedAmount, b.CreatedAt, b.UpdatedAt, b.InvoiceID, b.InvestorID, b.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.CreateBid(ctx, b)
	assert.NoError(t, err)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
}

func TestBidRepo_GetBidByID(t *testing.T) {
	sqlMock, entClient, dbClose := setupSQLMock(t)
	defer dbClose()

	repo := repos.NewBidRepo(entClient)

	bidID := uuid.New()
	ctx := context.Background()

	sqlMock.ExpectQuery(`SELECT "bids"."id", "bids"."status", "bids"."amount", "bids"."accepted_amount", "bids"."created_at", "bids"."updated_at", "bids"."investor_bids", "bids"."invoice_bids" FROM "bids" WHERE "bids"."id" = $1 LIMIT 1`).
		WithArgs(bidID).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).
			AddRow(bidID))

	b, err := repo.GetBidByID(ctx, bidID)
	assert.NoError(t, err)
	assert.Equal(t, bidID, b.ID)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
}

func TestBidRepo_UpdateBid(t *testing.T) {
	sqlMock, entClient, dbClose := setupSQLMock(t)
	defer dbClose()

	repo := repos.NewBidRepo(entClient)
	ctx := context.Background()

	b := &bid.Bid{
		ID:         uuid.New(),
		InvoiceID:  uuid.New(),
		InvestorID: uuid.New(),
		Amount:     100,
		Status:     bid.Pending,
	}

	sqlMock.ExpectBegin()
	sqlMock.ExpectExec(`UPDATE "bids" SET "status" = $1, "amount" = $2, "accepted_amount" = $3, "updated_at" = $4 WHERE "id" = $5`).
		WithArgs(b.Status, b.Amount, b.AcceptedAmount, b.UpdatedAt, b.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectQuery(`SELECT "id", "status", "amount", "accepted_amount", "created_at", "updated_at" FROM "bids" WHERE "id" = $1`).
		WithArgs(b.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).
			AddRow(b.ID))
	sqlMock.ExpectCommit()

	err := repo.UpdateBid(ctx, b)
	assert.NoError(t, err)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
}

func TestBidRepo_DeleteBid(t *testing.T) {
	sqlMock, entClient, dbClose := setupSQLMock(t)
	defer dbClose()

	repo := repos.NewBidRepo(entClient)

	bidID := uuid.New()
	ctx := context.Background()

	sqlMock.ExpectExec(`DELETE FROM "bids" WHERE "bids"."id" = $1`).
		WithArgs(bidID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.DeleteBid(ctx, bidID)
	assert.NoError(t, err)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
}

// TODO Similarly, implement for ListBidsByInvoiceID and ListBidsByInvestorID
