package payments

import (
	"log"
	"net/http"

	"goji.io/pat"
)

type PaymentResource struct {
}

func NewPaymentResource() *PaymentResource {
	return &PaymentResource{}
}

func (pr *PaymentResource) GetPayment(w http.ResponseWriter, r *http.Request) {
	paymentId := pat.Param(r, "id")
	log.Printf("GET payment %s", paymentId)
}

func (pr *PaymentResource) DeletePayment(w http.ResponseWriter, r *http.Request) {
	paymentId := pat.Param(r, "id")
	log.Printf("DELETE payment %s", paymentId)
}

func (pr *PaymentResource) UpdatePayment(w http.ResponseWriter, r *http.Request) {
	paymentId := pat.Param(r, "id")
	log.Printf("UPDATE payment %s", paymentId)
}
