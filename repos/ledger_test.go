package repos_test

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/tauki/invoiceexchange/internal/ledger"
	"github.com/tauki/invoiceexchange/repos"
)

func TestLedgerRepo(t *testing.T) {
	sqlMock, entClient, dbClose := setupSQLMock(t)
	defer dbClose()

	repo := repos.NewLedgerRepo(entClient)

	ctx := context.Background()

	t.Run("Add Ledger Entry", func(t *testing.T) {
		l := &ledger.Ledger{
			Status:    ledger.StatusPending,
			ID:        uuid.New(),
			InvoiceID: uuid.New(),
			Entity:    ledger.EntityTypeIssuer,
			EntityID:  uuid.New(),
			Amount:    1000,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		sqlMock.ExpectExec(
			`INSERT INTO "ledgers" ("status", "invoice_id", "entity", "entity_id", "amount", "created_at", "updated_at", "id") VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`).
			WithArgs(l.Status, l.InvoiceID, l.Entity, l.EntityID, l.Amount, l.CreatedAt, l.UpdatedAt, l.ID).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := repo.AddLedgerEntry(ctx, l)
		assert.NoError(t, err)
		assert.NoError(t, sqlMock.ExpectationsWereMet())
	})

	t.Run("Get Ledger Entries By Invoice ID", func(t *testing.T) {
		invoiceID := uuid.New()

		sqlMock.ExpectQuery(`SELECT "ledgers"."id", "ledgers"."status", "ledgers"."invoice_id", "ledgers"."entity", "ledgers"."entity_id", "ledgers"."amount", "ledgers"."created_at", "ledgers"."updated_at" FROM "ledgers" WHERE "ledgers"."invoice_id" = $1`).
			WithArgs(invoiceID).
			WillReturnRows(sqlmock.NewRows([]string{"id", "entity", "entity_id", "amount", "created_at", "updated_at", "invoice_id"}).
				AddRow(uuid.New(), ledger.EntityTypeIssuer, uuid.New(), 1000, time.Now(), time.Now(), invoiceID))

		entries, err := repo.GetLedgerEntriesByInvoiceID(ctx, invoiceID)
		assert.NoError(t, err)
		assert.Len(t, entries, 1)
		assert.NoError(t, sqlMock.ExpectationsWereMet())
	})
}
