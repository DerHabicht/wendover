package config

type PersonalInfoMapping struct {
	CAPID               SpreadsheetColumnMap `yaml:"capid"`
	LastName            SpreadsheetColumnMap `yaml:"last_name"`
	FirstName           SpreadsheetColumnMap `yaml:"first_name"`
	MemberType          SpreadsheetColumnMap `yaml:"member_type"`
	GradeInSchool       SpreadsheetColumnMap `yaml:"grade_in_school"`
	Gender              SpreadsheetColumnMap `yaml:"gender"`
	ReligiousPreference SpreadsheetColumnMap `yaml:"religious_preference"`
	DateOfBirth         SpreadsheetColumnMap `yaml:"date_of_birth"`
	ShirtSize           SpreadsheetColumnMap `yaml:"shirt_size"`
	Height              SpreadsheetColumnMap `yaml:"height"`
	Weight              SpreadsheetColumnMap `yaml:"weight"`
	HairColor           SpreadsheetColumnMap `yaml:"hair_color"`
	EyeColor            SpreadsheetColumnMap `yaml:"eye_color"`
}

type ApplicantContactInfoMapping struct {
	PrimaryPhoneNumber   SpreadsheetColumnMap `yaml:"primary_phone_number"`
	SecondaryPhoneNumber SpreadsheetColumnMap `yaml:"secondary_phone_number"`
	EmailAddress         SpreadsheetColumnMap `yaml:"email_address"`
}

type FlightInfoMapping struct {
	ArrivalDate  SpreadsheetColumnMap `yaml:"arrival_date"`
	ArrivalTime  SpreadsheetColumnMap `yaml:"arrival_time"`
	Airline      SpreadsheetColumnMap `yaml:"airline"`
	FlightNumber SpreadsheetColumnMap `yaml:"flight_number"`
}

type PreferredRolesMapping struct {
	First  SpreadsheetColumnMap `yaml:"first"`
	Second SpreadsheetColumnMap `yaml:"second"`
	Third  SpreadsheetColumnMap `yaml:"third"`
	Fourth SpreadsheetColumnMap `yaml:"fourth"`
	Fifth  SpreadsheetColumnMap `yaml:"fifth"`
}

type ResumeMapping struct {
	SquadronRoles       SpreadsheetColumnMap `yaml:"squadron_roles"`
	EncampmentsAttended SpreadsheetColumnMap `yaml:"encampments_attended"`
	CadreRoles          SpreadsheetColumnMap `yaml:"cadre_roles"`
}

type CadreApplicationMapping struct {
	ClassesToTeach       SpreadsheetColumnMap `yaml:"classes_to_teach"`
	PreferredRolesSenior SpreadsheetColumnMap `yaml:"preferred_roles_senior"`
	PreferredRolesCadet  SpreadsheetColumnMap `yaml:"preferred_roles_cadet"`
	Resume               ResumeMapping        `yaml:"resume"`
}

type VaccineMapping struct {
	Tetanus       SpreadsheetColumnMap `yaml:"tetanus"`
	Hepatitis     SpreadsheetColumnMap `yaml:"hepatitis"`
	Pneumonia     SpreadsheetColumnMap `yaml:"pneumonia"`
	VaricellaVax  SpreadsheetColumnMap `yaml:"varicella_vax"`
	VaricellaDate SpreadsheetColumnMap `yaml:"varicella_date"`
	Influenza     SpreadsheetColumnMap `yaml:"influenza"`
	Covid19       SpreadsheetColumnMap `yaml:"covid19"`
}

type CadetOtcMapping struct {
	OtcBlacklist SpreadsheetColumnMap `yaml:"otc_blacklist"`
	OtcReactions SpreadsheetColumnMap `yaml:"otc_reactions"`
}

type HealthHistoryMapping struct {
	FoodAllergies           SpreadsheetColumnMap `yaml:"food_allergies"`
	OtherAllergies          SpreadsheetColumnMap `yaml:"other_allergies"`
	MedicalConditions       SpreadsheetColumnMap `yaml:"medical_conditions"`
	MedicalConditionDetails SpreadsheetColumnMap `yaml:"medical_condition_details"`
	Vaccines                VaccineMapping       `yaml:"vaccines"`
	Medications             SpreadsheetColumnMap `yaml:"medications"`
	CadetOtcs               CadetOtcMapping      `yaml:"cadet_otcs"`
}

type EmergencyContactDetailsMapping struct {
	Name         SpreadsheetColumnMap `yaml:"name"`
	Relationship SpreadsheetColumnMap `yaml:"relationship"`
	PhoneNumber  SpreadsheetColumnMap `yaml:"phone_number"`
}

type EmergencyContactsMapping struct {
	PrimaryContact   EmergencyContactDetailsMapping `yaml:"primary_contact"`
	SecondaryContact EmergencyContactDetailsMapping `yaml:"secondary_contact"`
}

type InsuranceDetailsMapping struct {
	Company      SpreadsheetColumnMap `yaml:"company"`
	PolicyNumber SpreadsheetColumnMap `yaml:"policy_number"`
	GroupNumber  SpreadsheetColumnMap `yaml:"group_number"`
	Copay        SpreadsheetColumnMap `yaml:"copay"`
}

type PhysicianContactInfoMapping struct {
	Name           SpreadsheetColumnMap `yaml:"name"`
	PhoneNumber    SpreadsheetColumnMap `yaml:"phone_number"`
	MailingAddress SpreadsheetColumnMap `yaml:"mailing_address"`
}

type InsuranceMapping struct {
	MedicalInsurance      InsuranceDetailsMapping     `yaml:"medical_insurance"`
	PrescriptionInsurance InsuranceDetailsMapping     `yaml:"prescription_insurance"`
	FamilyPhysician       PhysicianContactInfoMapping `yaml:"family_physician"`
}

type ApplicationPackageColumnMapping struct {
	PersonalInfo      PersonalInfoMapping         `yaml:"personal_info"`
	ContactInfo       ApplicantContactInfoMapping `yaml:"contact_info"`
	InboundFlight     FlightInfoMapping           `yaml:"inbound_flight"`
	OutboundFlight    FlightInfoMapping           `yaml:"outbound_flight"`
	CadreApplication  CadreApplicationMapping     `yaml:"cadre_application"`
	HealthHistory     HealthHistoryMapping        `yaml:"health_history"`
	EmergencyContacts EmergencyContactsMapping    `yaml:"emergency_contacts"`
	Insurance         InsuranceMapping            `yaml:"insurance"`
	Signature         SpreadsheetColumnMap        `yaml:"signature"`
}
