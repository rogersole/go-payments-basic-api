package payments

import "context"

type Storage interface {
	GetAllPayments(context.Context) error
	InsertPayment(context.Context) error
	GetPayment(context.Context, int) error
	UpdatePayment(context.Context, int) error
	DeletePayment(context.Context, int) error
}
