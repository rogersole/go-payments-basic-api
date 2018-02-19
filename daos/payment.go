package daos

import (
	"github.com/go-ozzo/ozzo-dbx"
	"github.com/rogersole/payments-basic-api/app"
	"github.com/rogersole/payments-basic-api/dtos"
	uuid "github.com/satori/go.uuid"
)

// PaymentDAO persists payment data in database
type PaymentDAO struct{}

// NewPaymentDAO creates a new PaymentDAO
func NewPaymentDAO() *PaymentDAO {
	return &PaymentDAO{}
}

// Get reads the payment with the specified ID from the database
func (dao *PaymentDAO) Get(rs app.RequestScope, id uuid.UUID) (*dtos.Payment, error) {
	return dao.getSimple(rs, id)
}

func (dao *PaymentDAO) getSimple(rs app.RequestScope, id uuid.UUID) (*dtos.Payment, error) {
	var paymentDB PaymentDB
	if err := rs.Tx().Select().Model(id, &paymentDB); err != nil {
		return nil, err
	}

	var paymentAttributesDB PaymentAttributeDB
	if err := rs.Tx().Select().Model(paymentDB.PaymentAttributeId, &paymentAttributesDB); err != nil {
		return nil, err
	}

	var fxDB FxDB
	if err := rs.Tx().Select().Model(paymentAttributesDB.FxId, &fxDB); err != nil {
		return nil, err
	}

	var chargesInformationDB ChargesInformationDB
	if err := rs.Tx().Select().Model(paymentAttributesDB.ChargesInformationId, &chargesInformationDB); err != nil {
		return nil, err
	}

	var senderChargesDB []SenderChargeDB
	if err := rs.Tx().
		Select().
		From("sender_charge").
		Where(dbx.NewExp("charges_information_id={:charges_info_id}",
			dbx.Params{"charges_info_id": chargesInformationDB.Id})).
		All(&senderChargesDB); err != nil {
		return nil, err
	}

	var beneficiaryPartyDB PartyDB
	if err := rs.Tx().Select().Model(paymentAttributesDB.BeneficiaryPartyId, &beneficiaryPartyDB); err != nil {
		return nil, err
	}

	var debtorPartyDB PartyDB
	if err := rs.Tx().Select().Model(paymentAttributesDB.DebtorPartyId, &debtorPartyDB); err != nil {
		return nil, err
	}

	var sponsorPartyDB PartyDB
	if err := rs.Tx().Select().Model(paymentAttributesDB.SponsorPartyId, &sponsorPartyDB); err != nil {
		return nil, err
	}

	payment := NewPayment(&paymentDB, &paymentAttributesDB, &chargesInformationDB,
		senderChargesDB, &beneficiaryPartyDB, &debtorPartyDB, &sponsorPartyDB, &fxDB)

	return payment, nil
}

// Create saves a new payment record in the database.
func (dao *PaymentDAO) Create(rs app.RequestScope, payment *dtos.Payment) error {
	payment.Id = uuid.NewV4()

	beneficiaryPartyDB := NewPartyDB(payment.Attributes.BeneficiaryParty)
	if err := rs.Tx().Model(&beneficiaryPartyDB).Insert(); err != nil {
		return err
	}

	debtorPartyDB := NewPartyDB(payment.Attributes.DebtorParty)
	if err := rs.Tx().Model(&debtorPartyDB).Insert(); err != nil {
		return err
	}

	sponsorPartyDB := NewPartyDB(payment.Attributes.SponsorParty)
	if err := rs.Tx().Model(&sponsorPartyDB).Insert(); err != nil {
		return err
	}

	fxDB := NewFxDB(payment.Attributes.FX)
	if err := rs.Tx().Model(&fxDB).Insert(); err != nil {
		return err
	}

	chargeInformationDB := NewChargesInformationDB(payment.Attributes.ChargesInformation)
	if err := rs.Tx().Model(&chargeInformationDB).Insert(); err != nil {
		return err
	}

	for _, senderCharge := range payment.Attributes.ChargesInformation.SenderCharges {
		senderChargeDB := NewSenderChargeDB(senderCharge)
		senderChargeDB.ChargesInformationId = chargeInformationDB.Id
		if err := rs.Tx().Model(&senderChargeDB).Insert(); err != nil {
			return err
		}
	}

	paymentAttributeDB := NewPaymentAttributeDB(payment.Attributes)
	paymentAttributeDB.ChargesInformationId = chargeInformationDB.Id
	paymentAttributeDB.BeneficiaryPartyId = beneficiaryPartyDB.Id
	paymentAttributeDB.DebtorPartyId = debtorPartyDB.Id
	paymentAttributeDB.SponsorPartyId = sponsorPartyDB.Id
	paymentAttributeDB.FxId = fxDB.Id
	if err := rs.Tx().Model(&paymentAttributeDB).Insert(); err != nil {
		return err
	}

	paymentDB := NewPaymentDB(payment)
	paymentDB.Id = payment.Id
	paymentDB.PaymentAttributeId = paymentAttributeDB.Id
	return rs.Tx().Model(&paymentDB).Insert()
}

// Update saves the changes to a payment in the database.
func (dao *PaymentDAO) Update(rs app.RequestScope, id uuid.UUID, payment *dtos.Payment) error {

	// TODO

	if _, err := dao.Get(rs, id); err != nil {
		return err
	}
	payment.Id = id
	return rs.Tx().Model(payment).Exclude("Id").Update()
}

// Delete deletes a payment with the specified ID from the database.
func (dao *PaymentDAO) Delete(rs app.RequestScope, id uuid.UUID) error {
	var paymentDB PaymentDB
	if err := rs.Tx().Select().Model(id, &paymentDB); err != nil {
		return err
	}

	var paymentAttributesDB PaymentAttributeDB
	if err := rs.Tx().Select().Model(paymentDB.PaymentAttributeId, &paymentAttributesDB); err != nil {
		return err
	}

	var chargesInformationDB ChargesInformationDB
	if err := rs.Tx().Select().Model(paymentAttributesDB.ChargesInformationId, &chargesInformationDB); err != nil {
		return err
	}

	// Delete all models in the Database
	if _, err := rs.Tx().Delete("payment", dbx.HashExp{"id": id}).Execute(); err != nil {
		return err
	}

	if _, err := rs.Tx().Delete("payment_attribute", dbx.HashExp{"id": paymentAttributesDB.Id}).Execute(); err != nil {
		return err
	}

	if _, err := rs.Tx().Delete("party", dbx.In("id", paymentAttributesDB.BeneficiaryPartyId,
		paymentAttributesDB.SponsorPartyId, paymentAttributesDB.DebtorPartyId)).Execute(); err != nil {
		return err
	}

	if _, err := rs.Tx().Delete("fx", dbx.HashExp{"id": paymentAttributesDB.FxId}).Execute(); err != nil {
		return err
	}

	if _, err := rs.Tx().Delete("charges_information", dbx.HashExp{"id": paymentAttributesDB.ChargesInformationId}).Execute(); err != nil {
		return err
	}

	if _, err := rs.Tx().Delete("sender_charge", dbx.HashExp{"id": paymentAttributesDB.ChargesInformationId}).Execute(); err != nil {
		return err
	}

	return nil
}

// Count returns the number of the payment records in the database.
func (dao *PaymentDAO) Count(rs app.RequestScope) (int, error) {
	var count int
	err := rs.Tx().Select("COUNT(*)").From("payment").Row(&count)
	return count, err
}

// Query retrieves the payment records with the specified offset and limit from the database.
func (dao *PaymentDAO) Query(rs app.RequestScope, offset, limit int) ([]dtos.Payment, error) {
	var paymentsDB []PaymentDB
	err := rs.Tx().Select("id").OrderBy("id").Offset(int64(offset)).Limit(int64(limit)).All(&paymentsDB)

	var payments []dtos.Payment
	for _, paymentDB := range paymentsDB {
		payment, err := dao.getSimple(rs, paymentDB.Id)
		if err != nil {
			return nil, err
		}
		payments = append(payments, *payment)
	}
	return payments, err
}
