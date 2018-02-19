package dtos

import (
	"github.com/satori/go.uuid"
)

// Payment represents a payment record
type Payment struct {
	Id             uuid.UUID          `json:"id"`
	Version        int                `json:"version"`
	OrganisationId uuid.UUID          `json:"organisation_id"`
	Type           string             `json:"type"`
	Attributes     *PaymentAttributes `json:"attributes"`
}

type PaymentAttributes struct {
	Amount               string              `json:"amount"`
	Currency             string              `json:"currency"`
	EndToEndReference    string              `json:"end_to_end_reference"`
	NumericReference     string              `json:"numeric_reference"`
	PaymentId            string              `json:"payment_id"`
	PaymentPurpose       string              `json:"payment_purpose"`
	PaymentScheme        string              `json:"payment_scheme"`
	PaymentType          string              `json:"payment_type"`
	ProcessingDate       string              `json:"processing_date"`
	Reference            string              `json:"reference"`
	SchemePaymentSubType string              `json:"scheme_payment_sub_type"`
	SchemePaymentType    string              `json:"scheme_payment_type"`
	ChargesInformation   *ChargesInformation `json:"charges_information"`
	BeneficiaryParty     *Party              `json:"beneficiary_party"`
	DebtorParty          *Party              `json:"debtor_party"`
	SponsorParty         *Party              `json:"sponsor_party"`
	FX                   *FX                 `json:"fx"`
}

type Party struct {
	AccountName       string `json:"account_name,omitempty"`
	AccountNumber     string `json:"account_number,omitempty"`
	AccountNumberCode string `json:"account_number_code,omitempty"`
	AccountType       int    `json:"account_type"`
	Address           string `json:"address,omitempty"`
	BankId            string `json:"bank_id,omitempty"`
	BankIdCode        string `json:"bank_id_code,omitempty"`
	Name              string `json:"name,omitempty"`
}

type ChargesInformation struct {
	BearerCode              string    `json:"bearer_code"`
	SenderCharges           []*Charge `json:"sender_charges"`
	ReceiverChargesAmount   string    `json:"receiver_charges_amount"`
	ReceiverChargesCurrency string    `json:"receiver_charges_currency"`
}

type Charge struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type FX struct {
	ContractReference string `json:"contract_reference"`
	ExchangeRate      string `json:"exchange_rate"`
	OriginalAmount    string `json:"original_amount"`
	OriginalCurrency  string `json:"original_currency"`
}

// Validate validates the Payment fields.
func (p Payment) Validate() error {
	//return validation.ValidateStruct(&p,
	//	validation.Field(&p.Version, validation.Required, validation.Length(0, 120)),
	//)
	return nil
}
