package apis

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/rogersole/payments-basic-api/daos"
	"github.com/rogersole/payments-basic-api/dtos"
	"github.com/rogersole/payments-basic-api/services"
	"github.com/rogersole/payments-basic-api/testdata"
	uuid "github.com/satori/go.uuid"
)

func TestPayment(t *testing.T) {
	testdata.ResetDB()
	router := newRouter()
	ServePaymentResource(&router.RouteGroup, services.NewPaymentService(daos.NewPaymentDAO()))

	notFoundError := `{"error_code":"NOT_FOUND", "message":"NOT_FOUND"}`
	paymentT3 := newMockPaymentDTO(uuid.Nil, uuid.Nil, 3)
	paymentT3JSON, _ := json.Marshal(paymentT3)
	paymentT4 := newInvalidMockPaymentDTO(uuid.Nil, 4)
	paymentT4JSON, _ := json.Marshal(paymentT4)
	attributesRequiredError := `{"error_code":"INVALID_DATA", "message":"INVALID_DATA", "details":[{"error":"cannot be blank", "field":"attributes"}]}`
	paymentT5 := newMockPaymentDTO(uuid.FromStringOrNil("0cfb60f4-953a-4af7-afd4-d94f867d3d1d"), uuid.FromStringOrNil("743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb"), 5)
	paymentT5JSON, _ := json.Marshal(paymentT5)
	paymentT6 := newInvalidMockPaymentDTO(uuid.Nil, 6)
	paymentT6JSON, _ := json.Marshal(paymentT6)

	runAPITests(t, router, []apiTestCase{
		{"t1 - get a payment", "GET", "/payments/f4b83d65-2cbc-45e2-8827-2e4b25792a7d", "", http.StatusOK,
			`{
    "id": "f4b83d65-2cbc-45e2-8827-2e4b25792a7d",
    "version": 0,
    "organisation_id": "743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb",
    "type": "Payment 1",
    "attributes": {
        "amount": "1.21",
        "currency": "GBP",
        "end_to_end_reference": "Wil piano 1",
        "numeric_reference": "1002001",
        "payment_id": "123456789012345671",
        "payment_purpose": "Paying for goods/services",
        "payment_scheme": "FPS",
        "payment_type": "Credit",
        "processing_date": "2017-01-18",
        "reference": "Payment for Em's piano lessons",
        "scheme_payment_sub_type": "InternetBanking",
        "scheme_payment_type": "ImmediatePayment",
        "charges_information": {
            "bearer_code": "SHAR",
            "sender_charges": [
                {
                    "amount": "1.00",
                    "currency": "GBP"
                },
                {
                    "amount": "2.00",
                    "currency": "USD"
                }
            ],
            "receiver_charges_amount": "1.00",
            "receiver_charges_currency": "USD"
        },
        "beneficiary_party": {
            "account_name": "Party 1",
            "account_number": "1",
            "account_number_code": "BBAN",
            "account_type": 0,
            "address": "1 The Beneficiary Localtown SE2",
            "bank_id": "403000",
            "bank_id_code": "GBDSC",
            "name": "Wilfred Jeremiah Owens"
        },
        "debtor_party": {
            "account_name": "Party 2",
            "account_number": "2",
            "account_number_code": "IBAN",
            "account_type": 0,
            "address": "10 Debtor Crescent Sourcetown NE1",
            "bank_id": "203301",
            "bank_id_code": "GBDSC",
            "name": "Emelia Jane Brown"
        },
        "sponsor_party": {
			"account_name": "Party 3",
            "account_number": "3",
            "bank_id": "123123",
			"account_type": 0,
            "bank_id_code": "GBDSC"
        },
        "fx": {
            "contract_reference": "FX123",
            "exchange_rate": "1.00000",
            "original_amount": "201.42",
            "original_currency": "USD"
        }
    }
}`},
		{"t2 - get a nonexisting payment", "GET", "/payments/aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa", "", http.StatusNotFound, notFoundError},
		{"t4 - create a payment with validation error", "POST", "/payments", string(paymentT4JSON), http.StatusBadRequest, attributesRequiredError},
		{"t5 - update a payment", "PUT", "/payments/0cfb60f4-953a-4af7-afd4-d94f867d3d1d", string(paymentT5JSON), http.StatusOK, string(paymentT5JSON)},
		{"t6 - update a payment with validation error", "PUT", "/payments/0cfb60f4-953a-4af7-afd4-d94f867d3d1d", string(paymentT6JSON), http.StatusBadRequest, attributesRequiredError},
		{"t7 - update a nonexisting payment", "PUT", "/payments/00000000-0000-0000-0000-000000000000", "{}", http.StatusNotFound, notFoundError},
		{"t8 - delete a payment", "DELETE", "/payments/0cfb60f4-953a-4af7-afd4-d94f867d3d1d", ``, http.StatusOK, string(paymentT5JSON)},
		{"t9 - delete a nonexisting payment", "DELETE", "/payments/00000000-0000-0000-0000-000000000000", "", http.StatusNotFound, notFoundError},
		{"t10 - get a list of payments", "GET", "/payments?page=1&per_page=2", "", http.StatusOK, `{
    "page": 1,
    "per_page": 2,
    "page_count": 4,
    "total_count": 7,
    "items": [
        {
            "id": "242fe0db-7a83-4194-85a3-07c8ddb822c2",
            "version": 0,
            "organisation_id": "743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb",
            "type": "Payment 7",
            "attributes": {
                "amount": "7.21",
                "currency": "GBP",
                "end_to_end_reference": "Wil piano 7",
                "numeric_reference": "1002007",
                "payment_id": "123456789012345677",
                "payment_purpose": "Paying for goods/services",
                "payment_scheme": "FPS",
                "payment_type": "Credit",
                "processing_date": "2017-01-18",
                "reference": "Payment for Em's piano lessons",
                "scheme_payment_sub_type": "InternetBanking",
                "scheme_payment_type": "ImmediatePayment",
                "charges_information": {
                    "bearer_code": "SHAR",
                    "sender_charges": [
                        {
                            "amount": "13.00",
                            "currency": "GBP"
                        },
                        {
                            "amount": "14.00",
                            "currency": "USD"
                        }
                    ],
                    "receiver_charges_amount": "7.00",
                    "receiver_charges_currency": "USD"
                },
                "beneficiary_party": {
                    "account_name": "Party 19",
                    "account_number": "19",
                    "account_number_code": "BBAN",
                    "account_type": 0,
                    "address": "1 The Beneficiary Localtown SE2",
                    "bank_id": "403000",
                    "bank_id_code": "GBDSC",
                    "name": "Wilfred Jeremiah Owens"
                },
                "debtor_party": {
                    "account_name": "Party 20",
                    "account_number": "20",
                    "account_number_code": "IBAN",
                    "account_type": 0,
                    "address": "10 Debtor Crescent Sourcetown NE1",
                    "bank_id": "203301",
                    "bank_id_code": "GBDSC",
                    "name": "Emelia Jane Brown"
                },
                "sponsor_party": {
                    "account_name": "Party 21",
                    "account_number": "21",
                    "account_type": 0,
                    "bank_id": "123123",
                    "bank_id_code": "GBDSC"
                },
                "fx": {
                    "contract_reference": "FX123",
                    "exchange_rate": "7.00000",
                    "original_amount": "207.42",
                    "original_currency": "USD"
                }
            }
        },
        {
            "id": "2b389028-3429-4e74-8ea7-9ba6482385f9",
            "version": 0,
            "organisation_id": "743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb",
            "type": "Payment 2",
            "attributes": {
                "amount": "2.21",
                "currency": "GBP",
                "end_to_end_reference": "Wil piano 2",
                "numeric_reference": "1002002",
                "payment_id": "123456789012345672",
                "payment_purpose": "Paying for goods/services",
                "payment_scheme": "FPS",
                "payment_type": "Credit",
                "processing_date": "2017-01-18",
                "reference": "Payment for Em's piano lessons",
                "scheme_payment_sub_type": "InternetBanking",
                "scheme_payment_type": "ImmediatePayment",
                "charges_information": {
                    "bearer_code": "SHAR",
                    "sender_charges": [
                        {
                            "amount": "3.00",
                            "currency": "GBP"
                        },
                        {
                            "amount": "4.00",
                            "currency": "USD"
                        }
                    ],
                    "receiver_charges_amount": "2.00",
                    "receiver_charges_currency": "USD"
                },
                "beneficiary_party": {
                    "account_name": "Party 4",
                    "account_number": "4",
                    "account_number_code": "BBAN",
                    "account_type": 0,
                    "address": "1 The Beneficiary Localtown SE2",
                    "bank_id": "403000",
                    "bank_id_code": "GBDSC",
                    "name": "Wilfred Jeremiah Owens"
                },
                "debtor_party": {
                    "account_name": "Party 5",
                    "account_number": "5",
                    "account_number_code": "IBAN",
                    "account_type": 0,
                    "address": "10 Debtor Crescent Sourcetown NE1",
                    "bank_id": "203301",
                    "bank_id_code": "GBDSC",
                    "name": "Emelia Jane Brown"
                },
                "sponsor_party": {
                    "account_name": "Party 6",
                    "account_number": "6",
                    "account_type": 0,
                    "bank_id": "123123",
                    "bank_id_code": "GBDSC"
                },
                "fx": {
                    "contract_reference": "FX123",
                    "exchange_rate": "2.00000",
                    "original_amount": "202.42",
                    "original_currency": "USD"
                }
            }
        }
    ]
}`},
	})

	runAPITestsIgnoringResponseId(t, router, []apiTestCase{
		{"t3 - create an artist", "POST", "/payments", string(paymentT3JSON), http.StatusOK, string(paymentT3JSON)},
	})
}

func newInvalidMockPaymentDTO(paymentId uuid.UUID, idx int) *dtos.Payment {
	return &dtos.Payment{
		Id:             paymentId,
		Type:           fmt.Sprintf("Payment %d", idx),
		OrganisationId: uuid.NewV4(),
		Version:        0,
		Attributes:     nil,
	}
}

func newMockPaymentDTO(paymentId uuid.UUID, organisationId uuid.UUID, idx int) *dtos.Payment {

	orgId := uuid.NewV4()
	if organisationId != uuid.Nil {
		orgId = organisationId
	}

	return &dtos.Payment{
		Id:             paymentId,
		Type:           fmt.Sprintf("Payment %d", idx),
		OrganisationId: orgId,
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
				AccountType:       0,
				Name:              fmt.Sprintf("Debtor Name %d", idx),
			},
			SponsorParty: &dtos.Party{
				AccountName:       fmt.Sprintf("Sponsor name %d", idx),
				AccountNumber:     fmt.Sprintf("30000%d", idx),
				AccountNumberCode: "BBAN",
				Address:           fmt.Sprintf("Sponsor Party Address %d", idx),
				BankId:            fmt.Sprintf("4033%d", idx),
				BankIdCode:        "GBDSC",
				AccountType:       0,
			},
		},
	}
}
