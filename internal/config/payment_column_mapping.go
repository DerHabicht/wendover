package config

type PaymentCardColumnMapping struct {
	Last4      SpreadsheetColumnMap `yaml:"last_4"`
	NameOnCard SpreadsheetColumnMap `yaml:"name_on_card"`
	Expiration SpreadsheetColumnMap `yaml:"expiration"`
}

type PaymentInfoColumnMapping struct {
	TransactionID         SpreadsheetColumnMap `yaml:"transaction_id"`
	TransactionDate       SpreadsheetColumnMap `yaml:"transaction_date"`
	SuccessfulTransaction SpreadsheetColumnMap `yaml:"successful_transaction"`
	SettledStatus         SpreadsheetColumnMap `yaml:"settled_status"`
	SettledAmount         SpreadsheetColumnMap `yaml:"settled_amount"`
	CreditVoid            SpreadsheetColumnMap `yaml:"credit_void"`
	CreditAmount          SpreadsheetColumnMap `yaml:"credit_amount"`
}

type AttendeeInfoColumnMapping struct {
	FullName          SpreadsheetColumnMap `yaml:"full_name"`
	CAPID             SpreadsheetColumnMap `yaml:"capid"`
	MultipleAttendees SpreadsheetColumnMap `yaml:"multiple_attendees"`
}

type PaymentColumnMapping struct {
	CardInfo     PaymentCardColumnMapping  `yaml:"card_info"`
	PaymentInfo  PaymentInfoColumnMapping  `yaml:"payment_info"`
	AttendeeInfo AttendeeInfoColumnMapping `yaml:"attendee_info"`
}
