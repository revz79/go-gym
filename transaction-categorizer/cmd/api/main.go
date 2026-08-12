package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/revz79/go-gym/transaction-catgorizer/internal/application/transaction"
	. "github.com/revz79/go-gym/transaction-catgorizer/internal/presentation/http"
	"github.com/revz79/go-gym/transaction-catgorizer/internal/presentation/http/handler"
)

func main() {

	transactionService := transaction.NewTransaction()

	transactionHandler := handler.NewTransactionHandler(
		handler.WithTransactionService(transactionService),
	)

	server := &http.Server{
		Addr: "localhost:8080",
		Handler: SetupRoutes(transactionHandler),
	}

	go func ()  {
		fmt.Println("Server starting...")
		server.ListenAndServe()
	}()
	
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	server.Shutdown(ctx)

	fmt.Println("Server stopped.")
	
}