package config

type MemberDataMapping struct {
	CAPID                SpreadsheetColumnMap `yaml:"capid"`
	Grade                SpreadsheetColumnMap `yaml:"grade"`
	MembershipExpiration SpreadsheetColumnMap `yaml:"membership_expiration"`
	AgeAtStart           SpreadsheetColumnMap `yaml:"age_at_start"`
	AgeAtEnd             SpreadsheetColumnMap `yaml:"age_at_end"`
}

type UnitDataMapping struct {
	Region           SpreadsheetColumnMap `yaml:"region"`
	Wing             SpreadsheetColumnMap `yaml:"wing"`
	UnitNumber       SpreadsheetColumnMap `yaml:"unit_number"`
	UnitCC           SpreadsheetColumnMap `yaml:"unit_cc_cc"`
	UnitCCEmail      SpreadsheetColumnMap `yaml:"unit_cc_email"`
	WingCC           SpreadsheetColumnMap `yaml:"wing_cc"`
	WingCCEmail      SpreadsheetColumnMap `yaml:"wing_cc_email"`
	UnitApproved     SpreadsheetColumnMap `yaml:"unit_approved"`
	UnitApprovedDate SpreadsheetColumnMap `yaml:"unit_approved_date"`
	WingApproved     SpreadsheetColumnMap `yaml:"wing_approved"`
	WingApprovedDate SpreadsheetColumnMap `yaml:"wing_approved_date"`
}

type MailingAddressMapping struct {
	Addr1 SpreadsheetColumnMap `yaml:"addr1"`
	Addr2 SpreadsheetColumnMap `yaml:"addr2"`
	City  SpreadsheetColumnMap `yaml:"city"`
	State SpreadsheetColumnMap `yaml:"state"`
	Zip   SpreadsheetColumnMap `yaml:"zip"`
}

type CadetParentMapping struct {
	PrimaryPhone   SpreadsheetColumnMap `yaml:"primary_phone"`
	PrimaryEmail   SpreadsheetColumnMap `yaml:"primary_email"`
	EmergencyPhone SpreadsheetColumnMap `yaml:"emergency_phone"`
	EmergencyEmail SpreadsheetColumnMap `yaml:"emergency_email"`
}

type EServicesContactInfoMapping struct {
	Email          SpreadsheetColumnMap  `yaml:"email"`
	MailingAddress MailingAddressMapping `yaml:"mailing_address"`
	CadetParent    CadetParentMapping    `yaml:"cadet_parent"`
}

type TrainingMapping struct {
	LastEncampment SpreadsheetColumnMap `yaml:"last_encampment"`
	HighestORide   SpreadsheetColumnMap `yaml:"highest_oride"`
	AGH            SpreadsheetColumnMap `yaml:"agh"`
	WingRunner     SpreadsheetColumnMap `yaml:"wing_runner"`
	ORMB           SpreadsheetColumnMap `yaml:"ormb"`
	ORMI           SpreadsheetColumnMap `yaml:"ormi"`
	CpptExpiration SpreadsheetColumnMap `yaml:"cppt_expiration"`
	ICUT           SpreadsheetColumnMap `yaml:"icut"`
	IS100          SpreadsheetColumnMap `yaml:"is_100"`
	IS700          SpreadsheetColumnMap `yaml:"is_700"`
	CAPT116        SpreadsheetColumnMap `yaml:"capt_116"`
	CAPT1171       SpreadsheetColumnMap `yaml:"capt_117_1"`
	CAPT1172       SpreadsheetColumnMap `yaml:"capt_117_2"`
	CAPT1173       SpreadsheetColumnMap `yaml:"capt_117_3"`
	FirstAid       SpreadsheetColumnMap `yaml:"first_aid"`
	MeetsPrereqs   SpreadsheetColumnMap `yaml:"meets_prereqs"`
}

type RegistrationZoneColumnMapping struct {
	MemberData           MemberDataMapping           `yaml:"member_data"`
	Unit                 UnitDataMapping             `yaml:"unit"`
	EServicesContactInfo EServicesContactInfoMapping `yaml:"eservices_contact_info"`
	Training             TrainingMapping             `yaml:"training"`
}
