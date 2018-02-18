package models

import (
	"fmt"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
)

const ctLayout = "2006-01-02"

type CustomTime struct {
	time.Time
}

func (t CustomTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", t.Time.Format(ctLayout))), nil
}

func (t CustomTime) UnmarshalJSON(b []byte) error {
	var err error
	fmt.Printf(">>>>  unmarshaling %s\n\n", string(b))
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		t.Time = time.Time{}
		return nil
	}
	t.Time, err = time.Parse(ctLayout, s)
	fmt.Printf(">>>>> TIME: %v\n\n", t.Time)
	return err
}

func (t CustomTime) String() string {
	return t.Time.String()
}

type CustomUUID struct {
	uuid.UUID
}

func NewCustomUUID() CustomUUID {
	uuid, _ := uuid.NewV4()
	return CustomUUID{uuid}
}

func NewCustomUUIDFromString(s string) CustomUUID {
	uuid, _ := uuid.FromString(s)
	return CustomUUID{uuid}
}

func (u CustomUUID) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", u.UUID.String())), nil
}

func (u CustomUUID) UnmarshalJSON(b []byte) error {
	fmt.Printf(">>>>> unmarshaling UUID %s\n\n", string(b))
	var err error
	s := strings.Trim(string(b), "\"")
	u.UUID, err = uuid.FromString(s)
	if err != nil {
		return err
	}
	fmt.Printf(">>>>> UUID: %v\n\n", u.UUID)
	return nil
}

func (u CustomUUID) String() string {
	return u.UUID.String()
}

/*
// Payment represents a payment record
type Payment struct {
	Id             CustomUUID        `json:"id"`
	Version        int               `json:"version,omitempty"`
	OrganisationID CustomUUID        `json:"organisation_id,omitempty"`
	Type           string            `json:"type,omitempty"`
	Attributes     PaymentAttributes `json:"attributes,omitempty" db:"attr"`
}

type PaymentAttributes struct {
	Amount               string             `json:"amount,omitempty"`
	Currency             string             `json:"currency,omitempty"`
	EndToEndReference    string             `json:"end_to_end_reference,omitempty"`
	NumericReference     string             `json:"numeric_reference,omitempty"`
	PaymentId            string             `json:"payment_id,omitempty"`
	PaymentPurpose       string             `json:"payment_purpose,omitempty"`
	PaymentScheme        string             `json:"payment_scheme,omitempty"`
	PaymentType          string             `json:"payment_type,omitempty"`
	ProcessingDate       CustomTime         `json:"processing_date,omitempty"`
	Reference            string             `json:"reference,omitempty"`
	SchemePaymentSubType string             `json:"scheme_payment_sub_type,omitempty"`
	SchemePaymentType    string             `json:"scheme_payment_type,omitempty"`
	ChargesInformation   ChargesInformation `json:"charges_information,omitempty"`
	BeneficiaryParty     Party              `json:"beneficiary_party,omitempty" db:"b_prt"`
	DebtorParty          Party              `json:"debtor_party,omitempty" db:"d_prt"`
	SponsorParty         Party              `json:"sponsor_party,omitempty" db:"s_prt"`
	FX                   FX                 `json:"fx,omitempty" db:"fx"`
}

type Party struct {
	AccountName       string `json:"account_name,omitempty"`
	AccountNumber     string `json:"account_number,omitempty"`
	AccountNumberCode string `json:"account_number_code,omitempty"`
	AccountType       int    `json:"account_type,omitempty"`
	Address           string `json:"address,omitempty"`
	BankId            string `json:"bank_id,omitempty"`
	BankIdCode        string `json:"bank_id_code,omitempty"`
	Name              string `json:"name,omitempty"`
}

type ChargesInformation struct {
	BearerCode              string   `json:"bearer_code,omitempty"`
	SenderCharges           []Charge `json:"sender_charges,omitempty"`
	ReceiverChargesAmount   string   `json:"receiver_charges_amount,omitempty"`
	ReceiverChargesCurrency string   `json:"receiver_charges_currency,omitempty"`
}

type Charge struct {
	Amount   string `json:"amount,omitempty"`
	Currency string `json:"currency,omitempty"`
}

type FX struct {
	ContractReference string `json:"contract_reference,omitempty"`
	ExchangeRate      string `json:"exchange_rate,omitempty"`
	OriginalAmount    string `json:"original_amount,omitempty"`
	OriginalCurrency  string `json:"original_currency,omitempty"`
}

*/

type Payment struct {
	Id         CustomUUID `json:"id"`
	Name       string     `json:"name"`
	SubPayment SubPayment `json:"sub_payment"`
}

type SubPayment struct {
	Simple string `json:"simple"`
}

type PaymentDB struct {
	Id           CustomUUID
	Name         string
	SubPaymentId int64
}

func (c PaymentDB) TableName() string {
	return "payment"
}

type SubPaymentDB struct {
	Id     int64
	Simple string
}

func (c SubPaymentDB) TableName() string {
	return "sub_payment"
}

// Validate validates the Payment fields.
func (m Payment) Validate() error {
	//return validation.ValidateStruct(&m,
	//	validation.Field(&m.Name, validation.Required, validation.Length(0, 120)),
	//)
	return nil
	// TODO: implement validation
}
