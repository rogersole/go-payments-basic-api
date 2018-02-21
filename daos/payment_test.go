package daos

import (
	"fmt"
	"testing"

	"github.com/rogersole/payments-basic-api/app"
	"github.com/rogersole/payments-basic-api/dtos"
	"github.com/rogersole/payments-basic-api/testdata"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestPaymentDAO(t *testing.T) {
	db := testdata.ResetDB()
	dao := NewPaymentDAO()

	{
		// Get
		testDBCall(db, func(rs app.RequestScope) {
			id, _ := uuid.FromString("f4b83d65-2cbc-45e2-8827-2e4b25792a7d")
			payment, err := dao.Get(rs, id)
			assert.Nil(t, err)
			if assert.NotNil(t, payment) {
				assert.Equal(t, id, payment.Id)
			}
		})
	}

	{
		// Create
		testDBCall(db, func(rs app.RequestScope) {
			paymentId := uuid.NewV4()
			organisationId := uuid.NewV4()
			payment := newMockPaymentDTO(paymentId, organisationId, 0)
			err := dao.Create(rs, payment)
			assert.Nil(t, err)
			assert.NotEqual(t, paymentId, payment.Id)
			assert.NotZero(t, payment.Id)
		})
	}

	{
		// Update
		testDBCall(db, func(rs app.RequestScope) {
			id, _ := uuid.FromString("f4b83d65-2cbc-45e2-8827-2e4b25792a7d")
			organisationId := uuid.NewV4()
			payment := newMockPaymentDTO(id, organisationId, 0)
			err := dao.Update(rs, payment.Id, payment)
			assert.Nil(t, err)
			assert.Equal(t, payment.OrganisationId, organisationId)
		})
	}

	{
		// Update with error
		testDBCall(db, func(rs app.RequestScope) {
			paymentId := uuid.NewV4()
			organisationId := uuid.NewV4()
			payment := newMockPaymentDTO(paymentId, organisationId, 0)
			err := dao.Update(rs, uuid.NewV4(), payment)
			assert.NotNil(t, err)
		})
	}

	{
		// Delete
		testDBCall(db, func(rs app.RequestScope) {
			id, _ := uuid.FromString("f4b83d65-2cbc-45e2-8827-2e4b25792a7d")
			err := dao.Delete(rs, id)
			assert.Nil(t, err)
		})
	}

	{
		// Delete with error
		testDBCall(db, func(rs app.RequestScope) {
			err := dao.Delete(rs, uuid.NewV4())
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
