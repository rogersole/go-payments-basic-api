package services

import (
	"github.com/rogersole/payments-basic-api/app"
	"github.com/rogersole/payments-basic-api/dtos"
	"github.com/satori/go.uuid"
)

// paymentDAO specifies the interface of the payment DAO needed by PaymentService.
type paymentDAO interface {
	// Get returns the payment with the specified ID
	Get(rs app.RequestScope, id uuid.UUID) (*dtos.Payment, error)
	// Count returns the number of payments
	Count(rs app.RequestScope) (int, error)
	// Query returns the list of payments with the given offset and limit
	Query(rs app.RequestScope, offset, limit int) ([]*dtos.Payment, error)
	// Create saves a new payment in the storage
	Create(rs app.RequestScope, payment *dtos.Payment) error
	// Update updates the payment with given ID in the storage
	Update(rs app.RequestScope, id uuid.UUID, payment *dtos.Payment) error
	// Delete removes the payment with given ID from the storage
	Delete(rs app.RequestScope, id uuid.UUID) error
}

// PaymentService provides services related with payments
type PaymentService struct {
	dao paymentDAO
}

// NewPaymentService creates a new PaymentService with the given payment DAO
func NewPaymentService(dao paymentDAO) *PaymentService {
	return &PaymentService{dao}
}

// Get returns the payment with the specified ID
func (s *PaymentService) Get(rs app.RequestScope, id uuid.UUID) (*dtos.Payment, error) {
	return s.dao.Get(rs, id)
}

// Create creates a new payment
func (s *PaymentService) Create(rs app.RequestScope, model *dtos.Payment) (*dtos.Payment, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Create(rs, model); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, model.Id)
}

// Update updates the payment with the specified ID
func (s *PaymentService) Update(rs app.RequestScope, id uuid.UUID, model *dtos.Payment) (*dtos.Payment, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Update(rs, id, model); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, id)
}

// Delete deletes the payment with the specified ID
func (s *PaymentService) Delete(rs app.RequestScope, id uuid.UUID) (*dtos.Payment, error) {
	payment, err := s.dao.Get(rs, id)
	if err != nil {
		return nil, err
	}
	err = s.dao.Delete(rs, id)
	return payment, err
}

// Count returns the number of payments
func (s *PaymentService) Count(rs app.RequestScope) (int, error) {
	return s.dao.Count(rs)
}

// Query returns the payments with the specified offset and limit
func (s *PaymentService) Query(rs app.RequestScope, offset, limit int) ([]*dtos.Payment, error) {
	return s.dao.Query(rs, offset, limit)
}
