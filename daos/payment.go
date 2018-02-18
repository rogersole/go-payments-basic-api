package daos

import (
	"github.com/rogersole/payments-basic-api/app"
	"github.com/rogersole/payments-basic-api/dtos"
)

// PaymentDAO persists payment data in database
type PaymentDAO struct{}

// NewPaymentDAO creates a new PaymentDAO
func NewPaymentDAO() *PaymentDAO {
	return &PaymentDAO{}
}

// Get reads the payment with the specified ID from the database
func (dao *PaymentDAO) Get(rs app.RequestScope, id dtos.CustomUUID) (*dtos.Payment, error) {
	var payment dtos.Payment
	err := rs.Tx().Select().Model(id, &payment)
	return &payment, err
}

// Create saves a new payment record in the database.
func (dao *PaymentDAO) Create(rs app.RequestScope, payment *dtos.Payment) error {
	beneficiaryPartyDB := NewPartyDB(&payment.Attributes.BeneficiaryParty)
	if err := rs.Tx().Model(&beneficiaryPartyDB).Insert(); err != nil {
		return err
	}

	debtorPartyDB := NewPartyDB(&payment.Attributes.DebtorParty)
	if err := rs.Tx().Model(&debtorPartyDB).Insert(); err != nil {
		return err
	}

	sponsorPartyDB := NewPartyDB(&payment.Attributes.SponsorParty)
	if err := rs.Tx().Model(&sponsorPartyDB).Insert(); err != nil {
		return err
	}

	fxDB := NewFxDB(&payment.Attributes.FX)
	if err := rs.Tx().Model(&fxDB).Insert(); err != nil {
		return err
	}

	chargeInformationDB := NewChargesInformationDB(&payment.Attributes.ChargesInformation)
	if err := rs.Tx().Model(&chargeInformationDB).Insert(); err != nil {
		return err
	}

	for _, senderCharge := range payment.Attributes.ChargesInformation.SenderCharges {
		senderChargeDB := NewSenderChargeDB(&senderCharge)
		senderChargeDB.ChargesInformationId = chargeInformationDB.Id
		if err := rs.Tx().Model(&senderChargeDB).Insert(); err != nil {
			return err
		}
	}

	paymentAttributeDB := NewPaymentAttributeDB(&payment.Attributes)
	paymentAttributeDB.ChargesInformationId = chargeInformationDB.Id
	paymentAttributeDB.BeneficiaryPartyId = beneficiaryPartyDB.Id
	paymentAttributeDB.DebtorPartyId = debtorPartyDB.Id
	paymentAttributeDB.SponsorPartyId = sponsorPartyDB.Id
	paymentAttributeDB.FxId = fxDB.Id
	if err := rs.Tx().Model(&paymentAttributeDB).Insert(); err != nil {
		return err
	}

	paymentDB := NewPaymentDB(payment)
	paymentDB.Id = dtos.NewCustomUUID()
	paymentDB.PaymentAttributeId = paymentAttributeDB.Id
	return rs.Tx().Model(&paymentDB).Insert()
}

// Update saves the changes to a payment in the database.
func (dao *PaymentDAO) Update(rs app.RequestScope, id dtos.CustomUUID, payment *dtos.Payment) error {

	// TODO

	if _, err := dao.Get(rs, id); err != nil {
		return err
	}
	payment.Id = id
	return rs.Tx().Model(payment).Exclude("Id").Update()
}

// Delete deletes a payment with the specified ID from the database.
func (dao *PaymentDAO) Delete(rs app.RequestScope, id dtos.CustomUUID) error {

	// TODO

	payment, err := dao.Get(rs, id)
	if err != nil {
		return err
	}
	return rs.Tx().Model(payment).Delete()
}

// Count returns the number of the payment records in the database.
func (dao *PaymentDAO) Count(rs app.RequestScope) (int, error) {

	// TODO

	var count int
	err := rs.Tx().Select("COUNT(*)").From("payment").Row(&count)
	return count, err
}

// Query retrieves the payment records with the specified offset and limit from the database.
func (dao *PaymentDAO) Query(rs app.RequestScope, offset, limit int) ([]dtos.Payment, error) {

	// TODO

	var payments []dtos.Payment
	err := rs.Tx().Select().OrderBy("id").Offset(int64(offset)).Limit(int64(limit)).All(&payments)
	return payments, err
}
