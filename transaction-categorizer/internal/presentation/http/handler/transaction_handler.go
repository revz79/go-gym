package handler

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/revz79/go-gym/transaction-catgorizer/internal/application/transaction"
	"github.com/revz79/go-gym/transaction-catgorizer/internal/application/dto"
	"github.com/shopspring/decimal"
)

type TransactionResponse struct {
	ID     string          `json:"id"`
	UserID string          `json:"userId"`
	Amount decimal.Decimal `json:"amount"`
}

func toTransactionResponse(t *dto.Transaction) *TransactionResponse {
	return &TransactionResponse{
		ID:     t.ID.String(),
		UserID: t.UserID.String(),
		Amount: t.Amount,
	}
}

type TransactionService interface {
	GetByID(uuid.UUID) (*dto.Transaction, error)
	// GetByUserID(uuid.UUID) ([]*dto.Transaction, error)
	// Store(*TransactionResponse) error
}

type TransactionHandler struct {
	service TransactionService
}

type transactionHandlerConfig struct {
	service TransactionService
}

type TransactionHandlerOption func(*transactionHandlerConfig)

func NewTransactionHandler(opts ...TransactionHandlerOption) *TransactionHandler {
	config := transactionHandlerConfig{}
	for _, opt := range opts {
		opt(&config)
	}

	return &TransactionHandler{
		service: func(c *transactionHandlerConfig) TransactionService { 
			if c.service == nil { return  transaction.NewTransaction()} 
			return c.service
			}(&config),
	}
}

func WithTransactionService(s TransactionService) TransactionHandlerOption {
	return func(config *transactionHandlerConfig) {
		config.service = s
	}
}

func (h *TransactionHandler) GetByID(w http.ResponseWriter, req *http.Request) {
	
	
	w.WriteHeader(200)
	w.Header().Add("Content", "application/json")

	r, _ := json.Marshal(TransactionResponse{ID: "id", UserID: "userid", Amount: decimal.New(0, 1)})
	w.Write(r)
}

func (h *TransactionHandler) GetByUserID(w http.ResponseWriter, req *http.Request) {
	
	
	w.WriteHeader(400)
	// w.Header().Add("Content", "application/json")

	// r, _ := json.Marshal(TransactionResponse{ID: "id", UserID: "userid", Amount: decimal.New(0, 1)})
	// w.Write(r)

}
