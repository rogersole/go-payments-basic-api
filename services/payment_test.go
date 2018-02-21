package services

import (
	"errors"
	"fmt"
	"testing"

	"github.com/rogersole/payments-basic-api/app"
	"github.com/rogersole/payments-basic-api/dtos"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

var (
	paymentIds = []uuid.UUID{
		uuid.NewV4(), // position never used by the tests
		uuid.NewV4(),
		uuid.NewV4(),
		uuid.NewV4(),
		uuid.NewV4(),
		uuid.NewV4(),
	}
	organisationId = uuid.NewV4()
	globalPayments []*dtos.Payment
)

func init() {
	for i := 0; i < len(paymentIds); i++ {
		globalPayments = append(globalPayments, newMockPaymentDTO(paymentIds[i], organisationId, i))
	}
}

func TestNewPaymentService(t *testing.T) {
	dao := newMockPaymentDAO()
	s := NewPaymentService(dao)
	assert.Equal(t, dao, s.dao)
}

func TestPaymentService_Get(t *testing.T) {
	s := NewPaymentService(newMockPaymentDAO())
	id := paymentIds[1]
	payment, err := s.Get(nil, id)
	if assert.Nil(t, err) && assert.NotNil(t, payment) {
		assert.Equal(t, globalPayments[1].Type, payment.Type)
	}

	payment, err = s.Get(nil, uuid.NewV4())
	assert.NotNil(t, err)
}

func TestPaymentService_Create(t *testing.T) {
	s := NewPaymentService(newMockPaymentDAO())
	paymentId := uuid.NewV4()
	newPayment := newMockPaymentDTO(paymentId, organisationId, 1)
	newPayment.Id = uuid.Nil
	payment, err := s.Create(nil, newPayment)
	if assert.Nil(t, err) && assert.NotNil(t, payment) {
		assert.NotEqual(t, paymentId, payment.Id)
		assert.Equal(t, organisationId, payment.OrganisationId)
		assert.Equal(t, globalPayments[1].Type, payment.Type)
	}

	// basic validation error
	_, err = s.Create(nil, &dtos.Payment{
		OrganisationId: organisationId,
		Version:        0,
		Type:           "type",
	})
	assert.NotNil(t, err)
}

func TestPaymentService_Update(t *testing.T) {
	s := NewPaymentService(newMockPaymentDAO())

	targetUpdatePayment := globalPayments[len(paymentIds)-1]

	payment, err := s.Update(nil, paymentIds[0], targetUpdatePayment)
	if assert.Nil(t, err) && assert.NotNil(t, payment) {
		assert.Equal(t, paymentIds[0], payment.Id)
		assert.Equal(t, targetUpdatePayment.Type, payment.Type)
	}

	// validation error
	// basic validation error
	_, err = s.Create(nil, &dtos.Payment{
		OrganisationId: organisationId,
		Version:        0,
		Type:           "type",
	})
	assert.NotNil(t, err)
}

func TestPaymentService_Delete(t *testing.T) {
	s := NewPaymentService(newMockPaymentDAO())
	paymentIdx := 2
	toRemoveId := globalPayments[paymentIdx].Id
	toRemoveType := globalPayments[paymentIdx].Type
	payment, err := s.Delete(nil, toRemoveId)
	if assert.Nil(t, err) && assert.NotNil(t, payment) {
		assert.Equal(t, toRemoveId, payment.Id)
		assert.Equal(t, toRemoveType, payment.Type)
	}

	_, err = s.Delete(nil, uuid.NewV4())
	assert.NotNil(t, err)
}

func TestPaymentService_Query(t *testing.T) {
	s := NewPaymentService(newMockPaymentDAO())
	result, err := s.Query(nil, 1, 5)
	if assert.Nil(t, err) {
		assert.Equal(t, 5, len(result))
	}
}

func newMockPaymentDTO(paymentId uuid.UUID, organisationId uuid.UUID, idx int) *dtos.Payment {
	return &dtos.Payment{
		Id:             paymentId,
		Type:           fmt.Sprintf("Payment %d", idx),
		OrganisationId: organisationId,
		Version:        0,
		Attributes: &dtos.PaymentAttributes{
			Amount:               fmt.Sprintf("%d.2%d", idx, idx),
			Currency:             "GBP",
			EndToEndReference:    fmt.Sprintf("Reference %d", idx),
			NumericReference:     fmt.Sprintf("1000%d", idx),
			PaymentId:            fmt.Sprintf("%d", idx),
			PaymentPurpose:       fmt.Sprintf("Purpose %d", idx),
			PaymentScheme:        "FPS",
			PaymentType:          "Credit",
			ProcessingDate:       "2017-01-18",
			Reference:            fmt.Sprintf("Reference %d", idx),
			SchemePaymentSubType: fmt.Sprintf("Subtype %d", idx),
			SchemePaymentType:    fmt.Sprintf("Type %d", idx),
			ChargesInformation: &dtos.ChargesInformation{
				BearerCode:              "SHAR",
				ReceiverChargesAmount:   fmt.Sprintf("%d", idx),
				ReceiverChargesCurrency: "GBP",
				SenderCharges: []*dtos.Charge{
					{Currency: "GBP", Amount: fmt.Sprintf("%d1", idx)},
					{Currency: "GBP", Amount: fmt.Sprintf("%d2", idx)},
				},
			},
			FX: &dtos.FX{
				ContractReference: fmt.Sprintf("REF%d", idx),
				ExchangeRate:      fmt.Sprintf("%d.00000", idx),
				OriginalAmount:    fmt.Sprintf("%d00.%d", idx, idx),
				OriginalCurrency:  "USD",
			},
			BeneficiaryParty: &dtos.Party{
				AccountName:       fmt.Sprintf("Beneficiary Account name %d", idx),
				AccountNumber:     fmt.Sprintf("10000%d", idx),
				AccountNumberCode: "BBAN",
				AccountType:       0,
				Address:           fmt.Sprintf("Beneficiary Party Address %d", idx),
				BankId:            fmt.Sprintf("4033%d", idx),
				BankIdCode:        "GBDSC",
				Name:              fmt.Sprintf("Beneficiary Name %d", idx),
			},
			DebtorParty: &dtos.Party{
				AccountName:       fmt.Sprintf("Debtor name %d", idx),
				AccountNumber:     fmt.Sprintf("20000%d", idx),
				AccountNumberCode: "BBAN",
				Address:           fmt.Sprintf("Debtor Party Address %d", idx),
				BankId:            fmt.Sprintf("2033%d", idx),
				BankIdCode:        "GBDSC",
				Name:              fmt.Sprintf("Debtor Name %d", idx),
			},
			SponsorParty: &dtos.Party{
				AccountNumber: fmt.Sprintf("30000%d", idx),
				BankId:        fmt.Sprintf("4033%d", idx),
				BankIdCode:    "GBDSC",
			},
		},
	}
}

type mockPaymentDAO struct {
	payments []*dtos.Payment
}

func newMockPaymentDAO() paymentDAO {
	var payments []*dtos.Payment
	for i := 1; i < 4; i++ {
		payments = append(payments, newMockPaymentDTO(paymentIds[i], organisationId, i))
	}

	return &mockPaymentDAO{
		payments: globalPayments,
	}
}

func (m *mockPaymentDAO) Get(rs app.RequestScope, id uuid.UUID) (*dtos.Payment, error) {
	for _, record := range m.payments {
		if record.Id == id {
			return record, nil
		}
	}
	return nil, errors.New("not found")
}

func (m *mockPaymentDAO) Query(rs app.RequestScope, offset, limit int) ([]*dtos.Payment, error) {
	return m.payments[offset : offset+limit], nil
}

func (m *mockPaymentDAO) Count(rs app.RequestScope) (int, error) {
	return len(m.payments), nil
}

func (m *mockPaymentDAO) Create(rs app.RequestScope, payment *dtos.Payment) error {
	if payment.Id != uuid.Nil {
		return errors.New("ID cannot be set")
	}
	payment.Id = uuid.NewV4()
	m.payments = append(m.payments, payment)
	return nil
}

func (m *mockPaymentDAO) Update(rs app.RequestScope, id uuid.UUID, payment *dtos.Payment) error {
	payment.Id = id
	for i, record := range m.payments {
		if record.Id == id {
			m.payments[i] = payment
			return nil
		}
	}
	return errors.New("not found")
}

func (m *mockPaymentDAO) Delete(rs app.RequestScope, id uuid.UUID) error {
	for i, record := range m.payments {
		if record.Id == id {
			m.payments = append(m.payments[:i], m.payments[i+1:]...)
			return nil
		}
	}
	return errors.New("not found")
}
