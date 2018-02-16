package payments

import (
	"log"
	"net/http"
)

type PaymentsResource struct {
}

func NewPaymentsResource() *PaymentsResource {
	return &PaymentsResource{}
}

func (pr *PaymentsResource) CreatePayment(w http.ResponseWriter, r *http.Request) {
	log.Printf("POST payments")
}

func (pr *PaymentsResource) ListPayments(w http.ResponseWriter, r *http.Request) {
	log.Printf("GET payments")
}
