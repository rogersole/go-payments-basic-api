DROP TABLE IF EXISTS payment;
DROP TABLE IF EXISTS payment_attribute;
DROP TABLE IF EXISTS fx;
DROP TABLE IF EXISTS sender_charge;
DROP TABLE IF EXISTS charges_information;
DROP TABLE IF EXISTS party;

CREATE TABLE party
(
  id SERIAL PRIMARY KEY,
  account_name VARCHAR(255),
  account_number VARCHAR(255),
  account_number_code VARCHAR(255),
  account_type INTEGER,
  address VARCHAR(255),
  bank_id VARCHAR(255),
  bank_id_code VARCHAR(255),
  name VARCHAR(255)
);

CREATE TABLE charges_information
(
  id SERIAL PRIMARY KEY,
  bearer_code VARCHAR(255),
  receiver_charges_amount VARCHAR(255),
  receiver_charges_currency VARCHAR(255)
);

CREATE TABLE sender_charge
(
  id SERIAL PRIMARY KEY,
  amount VARCHAR(255),
  currency VARCHAR(255),
  charges_information_id INTEGER NOT NULL,
  FOREIGN KEY (charges_information_id) REFERENCES charges_information (id) ON DELETE CASCADE
);

CREATE TABLE fx
(
  id SERIAL PRIMARY KEY,
  contract_reference VARCHAR(255),
  exchange_rate VARCHAR(255),
  original_amount VARCHAR(255),
  original_currency VARCHAR(255)
);

CREATE TABLE payment_attribute
(
  id SERIAL PRIMARY KEY,
  amount VARCHAR(255),
  currency VARCHAR(255),
  end_to_end_reference VARCHAR(255),
  numeric_reference VARCHAR(255),
  payment_id VARCHAR(255),
  payment_purpose VARCHAR(255),
  payment_scheme VARCHAR(255),
  payment_type VARCHAR(255),
  processing_date TIME without time zone NOT NULL DEFAULT (current_timestamp AT TIME ZONE 'UTC'),
  reference VARCHAR(255),
  scheme_payment_sub_type VARCHAR(255),
  scheme_payment_type VARCHAR(255),
  charges_information_id INTEGER NOT NULL,
  beneficiary_party_id INTEGER NOT NULL,
  debtor_party_id INTEGER NOT NULL,
  sponsor_party_id INTEGER NOT NULL,
  fx_id INTEGER NOT NULL,

  FOREIGN KEY (charges_information_id) REFERENCES charges_information (id) ON DELETE CASCADE,
  FOREIGN KEY (beneficiary_party_id) REFERENCES party (id) ON DELETE CASCADE,
  FOREIGN KEY (debtor_party_id) REFERENCES party (id) ON DELETE CASCADE,
  FOREIGN KEY (sponsor_party_id) REFERENCES party (id) ON DELETE CASCADE,
  FOREIGN KEY (fx_id) REFERENCES fx (id) ON DELETE CASCADE
);

CREATE TABLE payment
(
  id UUID PRIMARY KEY,
  type VARCHAR(120) NOT NULL,
  version INTEGER NOT NULL,
  organisation_id UUID NOT NULL,
  payment_attribute_id INTEGER,
  FOREIGN KEY (payment_attribute_id) REFERENCES payment_attribute (id) ON DELETE CASCADE
);