package daos

import (
	"github.com/rogersole/payments-basic-api/app"
	"github.com/rogersole/payments-basic-api/models"
)

// PaymentDAO persists payment data in database
type PaymentDAO struct{}

// NewPaymentDAO creates a new PaymentDAO
func NewPaymentDAO() *PaymentDAO {
	return &PaymentDAO{}
}

// Get reads the payment with the specified ID from the database
func (dao *PaymentDAO) Get(rs app.RequestScope, id models.CustomUUID) (*models.Payment, error) {
	var payment models.Payment
	err := rs.Tx().Select().Model(id, &payment)
	return &payment, err
}

// Create saves a new payment record in the database.
func (dao *PaymentDAO) Create(rs app.RequestScope, payment *models.Payment) error {
	payment.Id = models.NewCustomUUID()

	subPaymentDB := models.SubPaymentDB{
		Simple: payment.SubPayment.Simple,
	}

	var err error
	err = rs.Tx().Model(&subPaymentDB).Insert()
	if err != nil {
		return err
	}

	paymentDB := models.PaymentDB{
		Id:           payment.Id,
		Name:         payment.Name,
		SubPaymentId: subPaymentDB.Id,
	}

	return rs.Tx().Model(&paymentDB).Insert()
}

// Update saves the changes to a payment in the database.
func (dao *PaymentDAO) Update(rs app.RequestScope, id models.CustomUUID, payment *models.Payment) error {
	if _, err := dao.Get(rs, id); err != nil {
		return err
	}
	payment.Id = id
	return rs.Tx().Model(payment).Exclude("Id").Update()
}

// Delete deletes a payment with the specified ID from the database.
func (dao *PaymentDAO) Delete(rs app.RequestScope, id models.CustomUUID) error {
	payment, err := dao.Get(rs, id)
	if err != nil {
		return err
	}
	return rs.Tx().Model(payment).Delete()
}

// Count returns the number of the payment records in the database.
func (dao *PaymentDAO) Count(rs app.RequestScope) (int, error) {
	var count int
	err := rs.Tx().Select("COUNT(*)").From("payment").Row(&count)
	return count, err
}

// Query retrieves the payment records with the specified offset and limit from the database.
func (dao *PaymentDAO) Query(rs app.RequestScope, offset, limit int) ([]models.Payment, error) {
	var payments []models.Payment
	err := rs.Tx().Select().OrderBy("id").Offset(int64(offset)).Limit(int64(limit)).All(&payments)
	return payments, err
}
