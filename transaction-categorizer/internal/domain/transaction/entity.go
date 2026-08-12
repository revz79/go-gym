package transaction

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)


type Transaction struct {
	ID uuid.UUID
	UserID uuid.UUID
	Amount decimal.Decimal
}

