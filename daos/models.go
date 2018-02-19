package daos

import (
	"time"

	"github.com/rogersole/payments-basic-api/dtos"
	uuid "github.com/satori/go.uuid"
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

func NewParty(partyDB *PartyDB) *dtos.Party {
	return &dtos.Party{
		AccountName:       partyDB.AccountName,
		AccountNumber:     partyDB.AccountNumber,
		AccountNumberCode: partyDB.AccountNumberCode,
		AccountType:       partyDB.AccountType,
		Address:           partyDB.Address,
		BankId:            partyDB.BankId,
		BankIdCode:        partyDB.BankIdCode,
		Name:              partyDB.Name,
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

func NewChargesInformation(chargesInformationDB *ChargesInformationDB, sc []SenderChargeDB) *dtos.ChargesInformation {

	var senderCharges []*dtos.Charge
	for _, senderChargeDB := range sc {
		senderCharge := NewSenderCharge(&senderChargeDB)
		senderCharges = append(senderCharges, senderCharge)
	}

	return &dtos.ChargesInformation{
		BearerCode:              chargesInformationDB.BearerCode,
		SenderCharges:           senderCharges,
		ReceiverChargesAmount:   chargesInformationDB.ReceiverChargesAmount,
		ReceiverChargesCurrency: chargesInformationDB.ReceiverChargesCurrency,
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

func NewSenderCharge(senderChargeDB *SenderChargeDB) *dtos.Charge {
	return &dtos.Charge{
		Amount:   senderChargeDB.Amount,
		Currency: senderChargeDB.Currency,
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

func NewFx(fxDB *FxDB) *dtos.FX {
	return &dtos.FX{
		ContractReference: fxDB.ContractReference,
		ExchangeRate:      fxDB.ExchangeRate,
		OriginalAmount:    fxDB.OriginalAmount,
		OriginalCurrency:  fxDB.OriginalCurrency,
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

func NewPaymentAttributes(paymentAttributesDB *PaymentAttributeDB,
	chargesInformationDB *ChargesInformationDB,
	senderChargesDB []SenderChargeDB,
	beneficiaryPartyDB *PartyDB, debtorPartyDB *PartyDB,
	sponsorPartyDB *PartyDB, fxDB *FxDB) *dtos.PaymentAttributes {

	chargesInformation := NewChargesInformation(chargesInformationDB, senderChargesDB)
	beneficiaryParty := NewParty(beneficiaryPartyDB)
	debtorParty := NewParty(debtorPartyDB)
	sponsorParty := NewParty(sponsorPartyDB)
	fx := NewFx(fxDB)

	return &dtos.PaymentAttributes{
		Amount:            paymentAttributesDB.Amount,
		Currency:          paymentAttributesDB.Currency,
		EndToEndReference: paymentAttributesDB.EndToEndReference,
		NumericReference:  paymentAttributesDB.NumericReference,
		PaymentId:         paymentAttributesDB.PaymentId,
		PaymentPurpose:    paymentAttributesDB.PaymentPurpose,
		PaymentScheme:     paymentAttributesDB.PaymentScheme,
		PaymentType:       paymentAttributesDB.PaymentType,
		ProcessingDate: dtos.CustomTime{
			Time: paymentAttributesDB.ProcessingDate,
		},
		Reference:            paymentAttributesDB.Reference,
		SchemePaymentSubType: paymentAttributesDB.SchemePaymentSubType,
		SchemePaymentType:    paymentAttributesDB.SchemePaymentType,
		ChargesInformation:   chargesInformation,
		BeneficiaryParty:     beneficiaryParty,
		DebtorParty:          debtorParty,
		SponsorParty:         sponsorParty,
		FX:                   fx,
	}
}

type PaymentDB struct {
	Id                 uuid.UUID
	Type               string
	Version            int
	OrganisationId     uuid.UUID
	PaymentAttributeId int64
}

func (c PaymentDB) TableName() string {
	return "payment"
}

func NewPayment(paymentDB *PaymentDB, paymentAttributesDB *PaymentAttributeDB,
	chargesInformationDB *ChargesInformationDB,
	senderChargesDB []SenderChargeDB,
	beneficiaryPartyDB *PartyDB, debtorPartyDB *PartyDB,
	sponsorPartyDB *PartyDB, fxDB *FxDB) *dtos.Payment {

	paymentAttributes := NewPaymentAttributes(paymentAttributesDB, chargesInformationDB, senderChargesDB,
		beneficiaryPartyDB, debtorPartyDB, sponsorPartyDB, fxDB)
	return &dtos.Payment{
		Id:             paymentDB.Id,
		Version:        paymentDB.Version,
		Type:           paymentDB.Type,
		OrganisationId: paymentDB.OrganisationId,
		Attributes:     paymentAttributes,
	}
}

func NewPaymentDB(dto *dtos.Payment) PaymentDB {
	return PaymentDB{
		Type:           dto.Type,
		Version:        dto.Version,
		OrganisationId: dto.OrganisationId,
	}
}
