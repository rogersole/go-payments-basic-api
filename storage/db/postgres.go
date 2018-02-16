package db

import (
	"context"

	"github.com/rogersole/payments-basic-api/api/payments"
)

type postgresStorage struct {
}

func New() payments.Storage {
	return &postgresStorage{}
}

func (ps *postgresStorage) GetAllPayments(_ context.Context) error {

}

func (ps *postgresStorage) InsertPayment(_ context.Context) error {

}

func (ps *postgresStorage) GetPayment(_ context.Context, paymentId int) error {

}

func (ps *postgresStorage) UpdatePayment(_ context.Context, paymentId int) error {

}

func (ps *postgresStorage) DeletePayment(_ context.Context, paymentId int) error {

}
