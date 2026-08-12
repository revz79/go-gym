package presentation

import (
	"net/http"

	"github.com/revz79/go-gym/transaction-catgorizer/internal/presentation/http/handler"
)

func SetupRoutes(th *handler.TransactionHandler) http.Handler {
	
	mux := http.NewServeMux()
	mux.HandleFunc("GET /transactions/{id}", th.GetByID)
	mux.HandleFunc("GET /users/{id}", th.GetByUserID)


	return mux
} 