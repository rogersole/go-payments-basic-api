package daos

import (
	"time"

	"github.com/rogersole/payments-basic-api/dtos"
)

type PartyDB struct {
	Id                int64
	AccountName       string
	AccountNumber     string
	AccountNumberCode string
	AccountType       int
	Address           string
	BankId            string
	BankIdCode        string
	Name              string
}

func (c PartyDB) TableName() string {
	return "party"
}

func NewPartyDB(dto *dtos.Party) PartyDB {
	return PartyDB{
		AccountName:       dto.AccountName,
		AccountNumber:     dto.AccountNumber,
		AccountNumberCode: dto.AccountNumberCode,
		AccountType:       dto.AccountType,
		Address:           dto.Address,
		BankId:            dto.BankId,
		BankIdCode:        dto.BankIdCode,
		Name:              dto.Name,
	}
}

type ChargesInformationDB struct {
	Id                      int64
	BearerCode              string
	ReceiverChargesAmount   string
	ReceiverChargesCurrency string
}

func (c ChargesInformationDB) TableName() string {
	return "charges_information"
}

func NewChargesInformationDB(dto *dtos.ChargesInformation) ChargesInformationDB {
	return ChargesInformationDB{
		BearerCode:              dto.BearerCode,
		ReceiverChargesAmount:   dto.ReceiverChargesAmount,
		ReceiverChargesCurrency: dto.ReceiverChargesCurrency,
	}
}

type SenderChargeDB struct {
	Id                   int64
	Amount               string
	Currency             string
	ChargesInformationId int64
}

func (c SenderChargeDB) TableName() string {
	return "sender_charge"
}

func NewSenderChargeDB(dto *dtos.Charge) SenderChargeDB {
	return SenderChargeDB{
		Amount:   dto.Amount,
		Currency: dto.Currency,
	}
}

type FxDB struct {
	Id                int64
	ContractReference string
	ExchangeRate      string
	OriginalAmount    string
	OriginalCurrency  string
}

func (c FxDB) TableName() string {
	return "fx"
}

func NewFxDB(dto *dtos.FX) FxDB {
	return FxDB{
		ContractReference: dto.ContractReference,
		ExchangeRate:      dto.ExchangeRate,
		OriginalAmount:    dto.OriginalAmount,
		OriginalCurrency:  dto.OriginalCurrency,
	}
}

type PaymentAttributeDB struct {
	Id                   int64
	Amount               string
	Currency             string
	EndToEndReference    string
	NumericReference     string
	PaymentId            string
	PaymentPurpose       string
	PaymentScheme        string
	PaymentType          string
	ProcessingDate       time.Time
	Reference            string
	SchemePaymentSubType string
	SchemePaymentType    string
	ChargesInformationId int64
	BeneficiaryPartyId   int64
	DebtorPartyId        int64
	SponsorPartyId       int64
	FxId                 int64
}

func (c PaymentAttributeDB) TableName() string {
	return "payment_attribute"
}

func NewPaymentAttributeDB(dto *dtos.PaymentAttributes) PaymentAttributeDB {
	return PaymentAttributeDB{
		Amount:               dto.Amount,
		Currency:             dto.Currency,
		EndToEndReference:    dto.EndToEndReference,
		NumericReference:     dto.NumericReference,
		PaymentId:            dto.PaymentId,
		PaymentPurpose:       dto.PaymentPurpose,
		PaymentScheme:        dto.PaymentScheme,
		PaymentType:          dto.PaymentType,
		ProcessingDate:       dto.ProcessingDate.Time,
		Reference:            dto.Reference,
		SchemePaymentSubType: dto.SchemePaymentSubType,
		SchemePaymentType:    dto.SchemePaymentType,
	}
}

type PaymentDB struct {
	Id                 dtos.CustomUUID
	Type               string
	Version            int
	OrganisationId     dtos.CustomUUID
	PaymentAttributeId int64
}

func (c PaymentDB) TableName() string {
	return "payment"
}

func NewPaymentDB(dto *dtos.Payment) PaymentDB {
	return PaymentDB{
		Type:           dto.Type,
		Version:        dto.Version,
		OrganisationId: dto.OrganisationID,
	}
}
