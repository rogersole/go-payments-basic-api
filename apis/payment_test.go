package apis

import (
	"net/http"
	"testing"

	"github.com/rogersole/payments-basic-api/daos"
	"github.com/rogersole/payments-basic-api/services"
	"github.com/rogersole/payments-basic-api/testdata"
)

func TestPayment(t *testing.T) {
	testdata.ResetDB()
	router := newRouter()
	ServePaymentResource(&router.RouteGroup, services.NewPaymentService(daos.NewPaymentDAO()))

	notFoundError := `{"error_code":"NOT_FOUND", "message":"NOT_FOUND"}`
	nameRequiredError := `{"error_code":"INVALID_DATA","message":"INVALID_DATA","details":[{"field":"name","error":"cannot be blank"}]}`

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
            "account_type": 0,
            "bank_id": "123123",
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
		{"t3 - create an artist", "POST", "/artists", `{"name":"Qiang"}`, http.StatusOK, `{"id": 276, "name":"Qiang"}`},
		{"t4 - create an artist with validation error", "POST", "/artists", `{"name":""}`, http.StatusBadRequest, nameRequiredError},
		{"t5 - update an artist", "PUT", "/artists/2", `{"name":"Qiang"}`, http.StatusOK, `{"id": 2, "name":"Qiang"}`},
		{"t6 - update an artist with validation error", "PUT", "/artists/2", `{"name":""}`, http.StatusBadRequest, nameRequiredError},
		{"t7 - update a nonexisting artist", "PUT", "/artists/99999", "{}", http.StatusNotFound, notFoundError},
		{"t8 - delete an artist", "DELETE", "/artists/2", ``, http.StatusOK, `{"id": 2, "name":"Qiang"}`},
		{"t9 - delete a nonexisting artist", "DELETE", "/artists/99999", "", http.StatusNotFound, notFoundError},
		{"t10 - get a list of artists", "GET", "/artists?page=3&per_page=2", "", http.StatusOK, `{"page":3,"per_page":2,"page_count":138,"total_count":275,"items":[{"id":6,"name":"Antônio Carlos Jobim"},{"id":7,"name":"Apocalyptica"}]}`},
	})
}
