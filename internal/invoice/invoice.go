package invoice

import (
	"github.com/tauki/invoiceexchange/internal/errors"
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	StatusPending   Status = "pending"
	StatusProcessed Status = "processed"
)

type Invoice struct {
	ID            uuid.UUID   `json:"id"`
	Status        Status      `json:"status"`
	AskingPrice   float64     `json:"asking_price"`
	IsLocked      bool        `json:"is_locked"`
	IsApproved    bool        `json:"is_approved"`
	InvoiceNumber string      `json:"invoice_number"`
	InvoiceDate   time.Time   `json:"invoice_date"`
	DueDate       time.Time   `json:"due_date"`
	AmountDue     float64     `json:"amount_due"`
	CustomerName  string      `json:"customer_name"`
	Reference     string      `json:"reference"`
	CompanyName   string      `json:"company_name"`
	Currency      string      `json:"currency"`
	TotalAmount   float64     `json:"total_amount"`
	TotalVAT      float64     `json:"total_vat"`
	IssuerID      uuid.UUID   `json:"issuer_id"`
	InvestorIDs   []uuid.UUID `json:"investor_ids"`
	Items         []*Item     `json:"items"`
}

func New(
	invoiceDate,
	dueDate time.Time,
	invoiceNumber,
	reference,
	companyName,
	customerName,
	currency string,
	totalAmount,
	totalVAT,
	askingPrice float64,
	issuerID uuid.UUID,
) (*Invoice, error) {
	if invoiceNumber == "" {
		return nil, errors.NewValidationError("invoice number cannot be empty", "invoiceNumber")
	}

	if companyName == "" {
		return nil, errors.NewValidationError("company name cannot be empty", "companyName")
	}

	if currency == "" {
		return nil, errors.NewValidationError("currency cannot be empty", "currency")
	}

	if issuerID == uuid.Nil {
		return nil, errors.NewValidationError("issuerID cannot be nil", "issuerID")
	}

	if askingPrice < 0 {
		return nil, errors.NewValidationError("asking price cannot be negative", "askingPrice")
	}

	if totalVAT < 0 {
		return nil, errors.NewValidationError("total VAT cannot be negative", "totalVAT")
	}

	return &Invoice{
		ID:            uuid.New(),
		Status:        StatusPending,
		InvoiceDate:   invoiceDate,
		InvoiceNumber: invoiceNumber,
		Reference:     reference,
		CompanyName:   companyName,
		CustomerName:  customerName,
		Currency:      currency,
		TotalAmount:   totalAmount,
		TotalVAT:      totalVAT,
		AskingPrice:   askingPrice,
		AmountDue:     askingPrice,
		DueDate:       dueDate,
		IssuerID:      issuerID,
	}, nil
}

func (i *Invoice) AddInvestor(investorID uuid.UUID) error {
	if investorID == uuid.Nil {
		return errors.NewValidationError("investorID cannot be nil", "investorID")
	}
	i.InvestorIDs = append(i.InvestorIDs, investorID)
	return nil
}
