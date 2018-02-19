package daos

import (
	"testing"

	"github.com/rogersole/payments-basic-api/app"
	"github.com/rogersole/payments-basic-api/dtos"
	"github.com/rogersole/payments-basic-api/testdata"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestPaymentDAO(t *testing.T) {
	db := testdata.ResetDB()
	dao := NewPaymentDAO()

	{
		// Get
		testDBCall(db, func(rs app.RequestScope) {
			uuid, _ := uuid.NewV4()
			payment, err := dao.Get(rs, uuid)
			assert.Nil(t, err)
			if assert.NotNil(t, payment) {
				assert.Equal(t, uuid, payment.Id)
			}
		})
	}

	{
		// Create
		testDBCall(db, func(rs app.RequestScope) {
			payment := &dtos.Payment{
				Id:   1000,
				Name: "tester",
			}
			err := dao.Create(rs, payment)
			assert.Nil(t, err)
			assert.NotEqual(t, 1000, payment.Id)
			assert.NotZero(t, payment.Id)
		})
	}

	{
		// Update
		testDBCall(db, func(rs app.RequestScope) {
			payment := &dtos.Payment{
				//Id:   2,
				//Name: "tester",
				// TODO
			}
			err := dao.Update(rs, payment.Id, payment)
			assert.Nil(t, err)
		})
	}

	{
		// Update with error
		testDBCall(db, func(rs app.RequestScope) {
			payment := &dtos.Payment{
				Id:   2,
				Name: "tester",
			}
			err := dao.Update(rs, 99999, payment)
			assert.NotNil(t, err)
		})
	}

	{
		// Delete
		testDBCall(db, func(rs app.RequestScope) {
			err := dao.Delete(rs, 2)
			assert.Nil(t, err)
		})
	}

	{
		// Delete with error
		testDBCall(db, func(rs app.RequestScope) {
			err := dao.Delete(rs, 99999)
			assert.NotNil(t, err)
		})
	}

	{
		// Query
		testDBCall(db, func(rs app.RequestScope) {
			payments, err := dao.Query(rs, 1, 3)
			assert.Nil(t, err)
			assert.Equal(t, 3, len(payments))
		})
	}

	{
		// Count
		testDBCall(db, func(rs app.RequestScope) {
			count, err := dao.Count(rs)
			assert.Nil(t, err)
			assert.NotZero(t, count)
		})
	}
}
