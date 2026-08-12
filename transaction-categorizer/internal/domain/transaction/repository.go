package transaction

import "github.com/google/uuid"

type Repository interface {
	GetByID(*uuid.UUID) (*Transaction, error)
	GetByUserID(*uuid.UUID) (*Transaction, error)
	Store(*Transaction) error
}
