package transaction

import (
	"github.com/google/uuid"

	dt "github.com/revz79/go-gym/transaction-catgorizer/internal/domain/transaction"
	"github.com/revz79/go-gym/transaction-catgorizer/internal/application/dto"
)


type transaction struct {
	repo dt.Repository
}

func NewTransaction() *transaction {
	return &transaction{}
}

func (t *transaction) GetByID(id uuid.UUID) (*dto.Transaction, error) {
	panic("not implemented")
}