package main

import (
	"flag"
	"fmt"
	"net/http"

	"log"

	"github.com/rogersole/payments-basic-api/api/payments"
	"goji.io"
	"goji.io/pat"
)

var (
	flagPort = flag.String("port", "8080", "--port <port> [default: 8080]")
)

func main() {
	flag.Parse()

	//postgresURL := os.Getenv("POSTGRES_URL")
	//if postgresURL == "" {
	//	log.Fatal("POSTGRES_URL environment variable not found. Aborting")
	//}

	paymentsResource := payments.NewPaymentsResource()
	paymentResource := payments.NewPaymentResource()

	mux := goji.NewMux()
	mux.HandleFunc(pat.Post("/payments"), paymentsResource.CreatePayment)
	mux.HandleFunc(pat.Get("/payments"), paymentsResource.ListPayments)
	mux.HandleFunc(pat.Get("/payments/:id"), paymentResource.GetPayment)
	mux.HandleFunc(pat.Delete("/payments/:id"), paymentResource.DeletePayment)
	mux.HandleFunc(pat.Put("/payments/:id"), paymentResource.UpdatePayment)

	addr := fmt.Sprintf(":%s", *flagPort)

	log.Printf("listening and serve on addr: %s", addr)
	log.Fatalf("listening and serve error: %s", http.ListenAndServe(addr, mux))
}
