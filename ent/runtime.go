// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/tauki/invoiceexchange/ent/balance"
	"github.com/tauki/invoiceexchange/ent/bid"
	"github.com/tauki/invoiceexchange/ent/investor"
	"github.com/tauki/invoiceexchange/ent/invoice"
	"github.com/tauki/invoiceexchange/ent/invoiceitem"
	"github.com/tauki/invoiceexchange/ent/issuer"
	"github.com/tauki/invoiceexchange/ent/ledger"
	"github.com/tauki/invoiceexchange/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	balanceFields := schema.Balance{}.Fields()
	_ = balanceFields
	// balanceDescCreatedAt is the schema descriptor for created_at field.
	balanceDescCreatedAt := balanceFields[4].Descriptor()
	// balance.DefaultCreatedAt holds the default value on creation for the created_at field.
	balance.DefaultCreatedAt = balanceDescCreatedAt.Default.(func() time.Time)
	// balanceDescUpdatedAt is the schema descriptor for updated_at field.
	balanceDescUpdatedAt := balanceFields[5].Descriptor()
	// balance.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	balance.DefaultUpdatedAt = balanceDescUpdatedAt.Default.(func() time.Time)
	// balance.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	balance.UpdateDefaultUpdatedAt = balanceDescUpdatedAt.UpdateDefault.(func() time.Time)
	// balanceDescID is the schema descriptor for id field.
	balanceDescID := balanceFields[0].Descriptor()
	// balance.DefaultID holds the default value on creation for the id field.
	balance.DefaultID = balanceDescID.Default.(func() uuid.UUID)
	bidFields := schema.Bid{}.Fields()
	_ = bidFields
	// bidDescAcceptedAmount is the schema descriptor for accepted_amount field.
	bidDescAcceptedAmount := bidFields[3].Descriptor()
	// bid.DefaultAcceptedAmount holds the default value on creation for the accepted_amount field.
	bid.DefaultAcceptedAmount = bidDescAcceptedAmount.Default.(float64)
	// bidDescCreatedAt is the schema descriptor for created_at field.
	bidDescCreatedAt := bidFields[4].Descriptor()
	// bid.DefaultCreatedAt holds the default value on creation for the created_at field.
	bid.DefaultCreatedAt = bidDescCreatedAt.Default.(func() time.Time)
	// bidDescUpdatedAt is the schema descriptor for updated_at field.
	bidDescUpdatedAt := bidFields[5].Descriptor()
	// bid.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	bid.DefaultUpdatedAt = bidDescUpdatedAt.Default.(func() time.Time)
	// bid.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	bid.UpdateDefaultUpdatedAt = bidDescUpdatedAt.UpdateDefault.(func() time.Time)
	// bidDescID is the schema descriptor for id field.
	bidDescID := bidFields[0].Descriptor()
	// bid.DefaultID holds the default value on creation for the id field.
	bid.DefaultID = bidDescID.Default.(func() uuid.UUID)
	investorFields := schema.Investor{}.Fields()
	_ = investorFields
	// investorDescJoinedAt is the schema descriptor for joined_at field.
	investorDescJoinedAt := investorFields[2].Descriptor()
	// investor.DefaultJoinedAt holds the default value on creation for the joined_at field.
	investor.DefaultJoinedAt = investorDescJoinedAt.Default.(func() time.Time)
	// investorDescID is the schema descriptor for id field.
	investorDescID := investorFields[0].Descriptor()
	// investor.DefaultID holds the default value on creation for the id field.
	investor.DefaultID = investorDescID.Default.(func() uuid.UUID)
	invoiceFields := schema.Invoice{}.Fields()
	_ = invoiceFields
	// invoiceDescAskingPrice is the schema descriptor for asking_price field.
	invoiceDescAskingPrice := invoiceFields[2].Descriptor()
	// invoice.DefaultAskingPrice holds the default value on creation for the asking_price field.
	invoice.DefaultAskingPrice = invoiceDescAskingPrice.Default.(float64)
	// invoiceDescIsLocked is the schema descriptor for is_locked field.
	invoiceDescIsLocked := invoiceFields[3].Descriptor()
	// invoice.DefaultIsLocked holds the default value on creation for the is_locked field.
	invoice.DefaultIsLocked = invoiceDescIsLocked.Default.(bool)
	// invoiceDescIsApproved is the schema descriptor for is_approved field.
	invoiceDescIsApproved := invoiceFields[4].Descriptor()
	// invoice.DefaultIsApproved holds the default value on creation for the is_approved field.
	invoice.DefaultIsApproved = invoiceDescIsApproved.Default.(bool)
	// invoiceDescCurrency is the schema descriptor for currency field.
	invoiceDescCurrency := invoiceFields[12].Descriptor()
	// invoice.DefaultCurrency holds the default value on creation for the currency field.
	invoice.DefaultCurrency = invoiceDescCurrency.Default.(string)
	// invoiceDescTotalAmount is the schema descriptor for total_amount field.
	invoiceDescTotalAmount := invoiceFields[13].Descriptor()
	// invoice.DefaultTotalAmount holds the default value on creation for the total_amount field.
	invoice.DefaultTotalAmount = invoiceDescTotalAmount.Default.(float64)
	// invoiceDescTotalVat is the schema descriptor for total_vat field.
	invoiceDescTotalVat := invoiceFields[14].Descriptor()
	// invoice.DefaultTotalVat holds the default value on creation for the total_vat field.
	invoice.DefaultTotalVat = invoiceDescTotalVat.Default.(float64)
	// invoiceDescCreatedAt is the schema descriptor for created_at field.
	invoiceDescCreatedAt := invoiceFields[15].Descriptor()
	// invoice.DefaultCreatedAt holds the default value on creation for the created_at field.
	invoice.DefaultCreatedAt = invoiceDescCreatedAt.Default.(func() time.Time)
	// invoiceDescID is the schema descriptor for id field.
	invoiceDescID := invoiceFields[0].Descriptor()
	// invoice.DefaultID holds the default value on creation for the id field.
	invoice.DefaultID = invoiceDescID.Default.(func() uuid.UUID)
	invoiceitemFields := schema.InvoiceItem{}.Fields()
	_ = invoiceitemFields
	// invoiceitemDescVatRate is the schema descriptor for vat_rate field.
	invoiceitemDescVatRate := invoiceitemFields[5].Descriptor()
	// invoiceitem.DefaultVatRate holds the default value on creation for the vat_rate field.
	invoiceitem.DefaultVatRate = invoiceitemDescVatRate.Default.(float64)
	// invoiceitemDescVatAmount is the schema descriptor for vat_amount field.
	invoiceitemDescVatAmount := invoiceitemFields[6].Descriptor()
	// invoiceitem.DefaultVatAmount holds the default value on creation for the vat_amount field.
	invoiceitem.DefaultVatAmount = invoiceitemDescVatAmount.Default.(float64)
	// invoiceitemDescID is the schema descriptor for id field.
	invoiceitemDescID := invoiceitemFields[0].Descriptor()
	// invoiceitem.DefaultID holds the default value on creation for the id field.
	invoiceitem.DefaultID = invoiceitemDescID.Default.(func() uuid.UUID)
	issuerFields := schema.Issuer{}.Fields()
	_ = issuerFields
	// issuerDescJoinedAt is the schema descriptor for joined_at field.
	issuerDescJoinedAt := issuerFields[2].Descriptor()
	// issuer.DefaultJoinedAt holds the default value on creation for the joined_at field.
	issuer.DefaultJoinedAt = issuerDescJoinedAt.Default.(func() time.Time)
	// issuerDescID is the schema descriptor for id field.
	issuerDescID := issuerFields[0].Descriptor()
	// issuer.DefaultID holds the default value on creation for the id field.
	issuer.DefaultID = issuerDescID.Default.(func() uuid.UUID)
	ledgerFields := schema.Ledger{}.Fields()
	_ = ledgerFields
	// ledgerDescCreatedAt is the schema descriptor for created_at field.
	ledgerDescCreatedAt := ledgerFields[6].Descriptor()
	// ledger.DefaultCreatedAt holds the default value on creation for the created_at field.
	ledger.DefaultCreatedAt = ledgerDescCreatedAt.Default.(func() time.Time)
	// ledgerDescUpdatedAt is the schema descriptor for updated_at field.
	ledgerDescUpdatedAt := ledgerFields[7].Descriptor()
	// ledger.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	ledger.DefaultUpdatedAt = ledgerDescUpdatedAt.Default.(func() time.Time)
	// ledger.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	ledger.UpdateDefaultUpdatedAt = ledgerDescUpdatedAt.UpdateDefault.(func() time.Time)
	// ledgerDescID is the schema descriptor for id field.
	ledgerDescID := ledgerFields[0].Descriptor()
	// ledger.DefaultID holds the default value on creation for the id field.
	ledger.DefaultID = ledgerDescID.Default.(func() uuid.UUID)
}