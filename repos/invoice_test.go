package repos_test

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/tauki/invoiceexchange/internal/invoice"
	"github.com/tauki/invoiceexchange/repos"
)

func TestInvoiceRepo_CreateInvoice(t *testing.T) {
	sqlMock, entClient, dbClose := setupSQLMock(t)
	defer dbClose()

	repo := repos.NewInvoiceRepo(entClient)

	ctx := context.Background()

	inv := &invoice.Invoice{
		ID:            uuid.New(),
		Status:        invoice.StatusPending,
		AskingPrice:   100,
		IsLocked:      false,
		IsApproved:    false,
		InvoiceNumber: "INV-123",
		InvoiceDate:   time.Now(),
		DueDate:       time.Now().AddDate(0, 1, 0),
		AmountDue:     100,
		CustomerName:  "John Doe",
		Reference:     "REF-123",
		CompanyName:   "ABC Company",
		Currency:      "USD",
		TotalAmount:   100,
		TotalVAT:      10,
		IssuerID:      uuid.New(),
		InvestorIDs:   []uuid.UUID{uuid.New()},
	}

	// Mock invoice creation
	sqlMock.ExpectExec(`INSERT INTO "invoices" ("status", "asking_price", "is_locked", "is_approved", "invoice_number", "invoice_date", "due_date", "amount_due", "customer_name", "reference", "company_name", "currency", "total_amount", "total_vat", "created_at", "issuer_invoices", "id") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17)`).
		WithArgs(inv.Status, inv.AskingPrice, inv.IsLocked, inv.IsApproved, inv.InvoiceNumber, inv.InvoiceDate, inv.DueDate, inv.AmountDue, inv.CustomerName, inv.Reference, inv.CompanyName, inv.Currency, inv.TotalAmount, inv.TotalVAT, sqlmock.AnyArg(),
			inv.IssuerID, inv.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.CreateInvoice(ctx, inv)
	assert.NoError(t, err)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
}

func TestInvoiceRepo_GetInvoiceByID(t *testing.T) {
	sqlMock, entClient, dbClose := setupSQLMock(t)
	defer dbClose()

	repo := repos.NewInvoiceRepo(entClient)

	ctx := context.Background()

	inv := &invoice.Invoice{
		ID:            uuid.New(),
		Status:        invoice.StatusPending,
		AskingPrice:   100,
		IsLocked:      false,
		IsApproved:    false,
		InvoiceNumber: "INV-123",
		InvoiceDate:   time.Now(),
		DueDate:       time.Now().AddDate(0, 1, 0),
		AmountDue:     100,
		CustomerName:  "John Doe",
		Reference:     "REF-123",
		CompanyName:   "ABC Company",
		Currency:      "USD",
		TotalAmount:   100,
		TotalVAT:      10,
		InvestorIDs:   []uuid.UUID{},
		Items:         []*invoice.Item{},
	}

	sqlMock.ExpectQuery(`SELECT "invoices"."id", "invoices"."status", "invoices"."asking_price", "invoices"."is_locked", "invoices"."is_approved", "invoices"."invoice_number", "invoices"."invoice_date", "invoices"."due_date", "invoices"."amount_due", "invoices"."customer_name", "invoices"."reference", "invoices"."company_name", "invoices"."currency", "invoices"."total_amount", "invoices"."total_vat", "invoices"."created_at", "invoices"."issuer_invoices" FROM "invoices" WHERE "invoices"."id" = $1 LIMIT 1`).
		WithArgs(inv.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "status", "asking_price", "is_locked", "is_approved", "invoice_number", "invoice_date", "due_date", "amount_due", "customer_name", "reference", "company_name", "currency", "total_amount", "total_vat", "issuer_id"}).
			AddRow(inv.ID, inv.Status, inv.AskingPrice, inv.IsLocked, inv.IsApproved, inv.InvoiceNumber, inv.InvoiceDate, inv.DueDate, inv.AmountDue, inv.CustomerName, inv.Reference, inv.CompanyName, inv.Currency, inv.TotalAmount, inv.TotalVAT, inv.IssuerID))
	sqlMock.ExpectQuery(`SELECT "invoice_items"."id", "invoice_items"."description", "invoice_items"."quantity", "invoice_items"."unit_price", "invoice_items"."amount", "invoice_items"."vat_rate", "invoice_items"."vat_amount", "invoice_items"."invoice_items" FROM "invoice_items" WHERE "invoice_items"."invoice_items" IN ($1)`).
		WithArgs(inv.ID).WillReturnRows(sqlmock.NewRows([]string{"id", "description", "quantity", "unit_price", "amount", "vat_rate", "vat_amount", "invoice_id"}))
	sqlMock.ExpectQuery(`SELECT "t1"."invoice_id", "investors"."id", "investors"."name", "investors"."joined_at" FROM "investors" JOIN "investor_invoices" AS "t1" ON "investors"."id" = "t1"."investor_id" WHERE "t1"."invoice_id" IN ($1)`).
		WithArgs(inv.ID).
		WillReturnRows(
			sqlmock.NewRows([]string{"invoice_id", "id", "name", "joined_at"}))

	resultInv, err := repo.GetInvoiceByID(ctx, inv.ID)
	assert.NoError(t, err)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Equal(t, inv, resultInv)
}

func TestInvoiceRepo_GetInvoiceByID_NotFound(t *testing.T) {
	sqlMock, entClient, dbClose := setupSQLMock(t)
	defer dbClose()

	repo := repos.NewInvoiceRepo(entClient)

	ctx := context.Background()

	inv := &invoice.Invoice{
		ID:            uuid.New(),
		Status:        invoice.StatusPending,
		AskingPrice:   100,
		IsLocked:      false,
		IsApproved:    false,
		InvoiceNumber: "INV-123",
		InvoiceDate:   time.Now(),
		DueDate:       time.Now().AddDate(0, 1, 0),
		AmountDue:     100,
		CustomerName:  "John Doe",
		Reference:     "REF-123",
		CompanyName:   "ABC Company",
		Currency:      "USD",
		TotalAmount:   100,
		TotalVAT:      10,
		InvestorIDs:   []uuid.UUID{},
		Items:         []*invoice.Item{},
	}

	sqlMock.ExpectQuery(`SELECT "invoices"."id", "invoices"."status", "invoices"."asking_price", "invoices"."is_locked", "invoices"."is_approved", "invoices"."invoice_number", "invoices"."invoice_date", "invoices"."due_date", "invoices"."amount_due", "invoices"."customer_name", "invoices"."reference", "invoices"."company_name", "invoices"."currency", "invoices"."total_amount", "invoices"."total_vat", "invoices"."created_at", "invoices"."issuer_invoices" FROM "invoices" WHERE "invoices"."id" = $1 LIMIT 1`).
		WithArgs(inv.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "status", "asking_price", "is_locked", "is_approved", "invoice_number", "invoice_date", "due_date", "amount_due", "customer_name", "reference", "company_name", "currency", "total_amount", "total_vat", "issuer_id"}))

	resultInv, err := repo.GetInvoiceByID(ctx, inv.ID)
	assert.NoError(t, err)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Nil(t, resultInv)
}
