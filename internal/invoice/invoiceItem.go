package invoice

import (
	"github.com/google/uuid"
	"github.com/tauki/invoiceexchange/internal/errors"
)

type Item struct {
	ID          uuid.UUID `json:"id"`
	Description string    `json:"description"`
	Quantity    int       `json:"quantity"`
	UnitPrice   float64   `json:"unit_price"`
	Amount      float64   `json:"amount"`
	VatRate     float64   `json:"vat_rate"`
	VatAmount   float64   `json:"vat_amount"`
}

func NewItem(
	description string,
	quantity int,
	unitPrice, vatAmount, total float64) (*Item, error) {
	if quantity <= 0 {
		return nil, errors.NewValidationError("quantity cannot be zero or negative", "quantity")
	}

	if unitPrice <= 0 {
		return nil, errors.NewValidationError("unit price cannot be zero or negative", "unitPrice")
	}

	if vatAmount < 0 {
		return nil, errors.NewValidationError("vat amount cannot be negative", "vatAmount")
	}

	if total < 0 {
		return nil, errors.NewValidationError("total amount cannot be negative", "total")
	}

	item := &Item{
		ID:          uuid.New(),
		Description: description,
		Quantity:    quantity,
		UnitPrice:   unitPrice,
		Amount:      total,
		VatAmount:   vatAmount,
	}

	// ignoring calculation, expecting to be provided
	//item.Amount = float64(quantity) * unitPrice
	//item.VatAmount = item.Amount * vatRate / 100

	return item, nil
}
