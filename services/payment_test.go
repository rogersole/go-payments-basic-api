package services

/*
import (
	"errors"
	"testing"

	"github.com/rogersole/payments-basic-api/app"
	"github.com/rogersole/payments-basic-api/dtos"
	"github.com/stretchr/testify/assert"
)

func TestNewPaymentService(t *testing.T) {
	dao := newMockPaymentDAO()
	s := NewPaymentService(dao)
	assert.Equal(t, dao, s.dao)
}

func TestPaymentService_Get(t *testing.T) {
	s := NewPaymentService(newMockPaymentDAO())
	payment, err := s.Get(nil, 1)
	if assert.Nil(t, err) && assert.NotNil(t, payment) {
		assert.Equal(t, "aaa", payment.Name)
	}

	payment, err = s.Get(nil, 100)
	assert.NotNil(t, err)
}

func TestPaymentService_Create(t *testing.T) {
	s := NewPaymentService(newMockPaymentDAO())
	payment, err := s.Create(nil, &dtos.Payment{
		//Name: "ddd",
		// TODO
	})
	if assert.Nil(t, err) && assert.NotNil(t, payment) {
		assert.Equal(t, 4, payment.Id)
		assert.Equal(t, "ddd", payment.Name)
	}

	// dao error
	_, err = s.Create(nil, &dtos.Payment{
		//Id:   100,
		//Name: "ddd",
		// TODO
	})
	assert.NotNil(t, err)

	// validation error
	_, err = s.Create(nil, &dtos.Payment{
		//Name: "",
		// TODO
	})
	assert.NotNil(t, err)
}

func TestPaymentService_Update(t *testing.T) {
	s := NewPaymentService(newMockPaymentDAO())
	payment, err := s.Update(nil, 2, &dtos.Payment{
		//Name: "ddd",
		// TODO
	})
	if assert.Nil(t, err) && assert.NotNil(t, payment) {
		assert.Equal(t, 2, payment.Id)
		assert.Equal(t, "ddd", payment.Name)
	}

	// dao error
	_, err = s.Update(nil, 100, &dtos.Payment{
		//Name: "ddd",
		// TODO
	})
	assert.NotNil(t, err)

	// validation error
	_, err = s.Update(nil, 2, &dtos.Payment{
		Name: "",
	})
	assert.NotNil(t, err)
}

func TestPaymentService_Delete(t *testing.T) {
	s := NewPaymentService(newMockPaymentDAO())
	payment, err := s.Delete(nil, 2)
	if assert.Nil(t, err) && assert.NotNil(t, payment) {
		assert.Equal(t, 2, payment.Id)
		assert.Equal(t, "bbb", payment.Name)
	}

	_, err = s.Delete(nil, 2)
	assert.NotNil(t, err)
}

func TestPaymentService_Query(t *testing.T) {
	s := NewPaymentService(newMockPaymentDAO())
	result, err := s.Query(nil, 1, 2)
	if assert.Nil(t, err) {
		assert.Equal(t, 2, len(result))
	}
}

func newMockPaymentDAO() paymentDAO {
	return &mockPaymentDAO{
		records: []dtos.Payment{
			//{Id: 1, Name: "aaa"},
			//{Id: 2, Name: "bbb"},
			//{Id: 3, Name: "ccc"},
			// TODO
		},
	}
}

type mockPaymentDAO struct {
	records []dtos.Payment
}

func (m *mockPaymentDAO) Get(rs app.RequestScope, id int) (*dtos.Payment, error) {
	for _, record := range m.records {
		if record.Id == id {
			return &record, nil
		}
	}
	return nil, errors.New("not found")
}

func (m *mockPaymentDAO) Query(rs app.RequestScope, offset, limit int) ([]dtos.Payment, error) {
	return m.records[offset : offset+limit], nil
}

func (m *mockPaymentDAO) Count(rs app.RequestScope) (int, error) {
	return len(m.records), nil
}

func (m *mockPaymentDAO) Create(rs app.RequestScope, payment *dtos.Payment) error {
	if payment.Id != 0 {
		return errors.New("ID cannot be set")
	}
	payment.Id = len(m.records) + 1
	m.records = append(m.records, *payment)
	return nil
}

func (m *mockPaymentDAO) Update(rs app.RequestScope, id int, payment *dtos.Payment) error {
	payment.Id = id
	for i, record := range m.records {
		if record.Id == id {
			m.records[i] = *payment
			return nil
		}
	}
	return errors.New("not found")
}

func (m *mockPaymentDAO) Delete(rs app.RequestScope, id int) error {
	for i, record := range m.records {
		if record.Id == id {
			m.records = append(m.records[:i], m.records[i+1:]...)
			return nil
		}
	}
	return errors.New("not found")
}
*/
