package pii

import (
	"regexp"

	cregex "github.com/mingrammer/commonregex"
)

const (
	personColumnPattern = `(?i)(^.*(firstname|fname|lastname|lname|"
	"fullname|maidenname|_name|"
	"nickname|name_suffix|name|person).*)`
	emailColumnPattern     = `(?i)(^.*(email|e-mail|mail).*)`
	birthDateColumnPattern = `(?i)(^.*(date_of_birth|dateofbirth|dob|"
	"birthday|date_of_death|dateofdeath|birthdate).*)`
	genderColumnPattern      = `(?i)(^.*(gender).*)`
	nationalityColumnPattern = `(?i)(^.*(nationality).*)`
	addressColumnPattern     = `(?i)(^.*(address|city|state|county|country|zone|borough).*)`
	zipCodeColumnPattern     = `(?i)(^.*(zipcode|zip_code|postal|postal_code|zip).*)`
	userNameColumnPattern    = `(?i)(^.*user(id|name|).*)`
	passwordColumnPattern    = `(?i)(^.*pass.*)`
	ssnColumnPattern         = `(?i)(^.*(ssn|social_number|social_security|"
	"social_security_number|social_security_no).*)`
	poBoxColumnPattern      = `(?i)(^.*(po_box|pobox).*)`
	creditCardColumnPattern = `(?i)(^.*(credit_card|cc_number|cc_num|creditcard|"
	"credit_card_num|creditcardnumber).*)`
	phoneColumnPattern = `(?i)(^.*(phone|phone_number|phone_no|phone_num|"
	"telephone|telephone_num|telephone_no).*)`
)

var (
	personRegexp      = regexp.MustCompile(personColumnPattern)
	emailRegexp       = regexp.MustCompile(emailColumnPattern)
	birthDateRegexp   = regexp.MustCompile(birthDateColumnPattern)
	genderRegexp      = regexp.MustCompile(genderColumnPattern)
	nationalityRegexp = regexp.MustCompile(nationalityColumnPattern)
	addressRegexp     = regexp.MustCompile(addressColumnPattern)
	zipCodeRegexp     = regexp.MustCompile(zipCodeColumnPattern)
	userNameRegexp    = regexp.MustCompile(userNameColumnPattern)
	passwordRegexp    = regexp.MustCompile(passwordColumnPattern)
	ssnRegexp         = regexp.MustCompile(ssnColumnPattern)
	poBoxRegexp       = regexp.MustCompile(poBoxColumnPattern)
	creditCardRegexp  = regexp.MustCompile(creditCardColumnPattern)
	phoneRegexp       = regexp.MustCompile(phoneColumnPattern)
)

type Detector interface {
	Detect(columnName string) PiiType
}

type DatumRegexDetector struct{}

type ColumnNameRegexDetector struct{}

func (c *ColumnNameRegexDetector) Detect(columnName string) string {
	// TODO:
	var columnRegexp map[string]*regexp.Regexp = map[string]*regexp.Regexp{
		"person":      personRegexp,
		"email":       emailRegexp,
		"birthDate":   birthDateRegexp,
		"gender":      genderRegexp,
		"nationality": nationalityRegexp,
		"address":     addressRegexp,
		"zipcode":     zipCodeRegexp,
		"username":    userNameRegexp,
		"password":    passwordRegexp,
		"ssn":         ssnRegexp,
		"pobox":       poBoxRegexp,
		"creditcard":  creditCardRegexp,
		"phone":       phoneRegexp,
	}

	for pii_type, ex := range columnRegexp {
		if ex.MatchString(columnName) {
			return pii_type
		}
	}

	return ""
}

func (d *DatumRegexDetector) Detect(columnData string) PiiType {
	// TODO:
	_ = cregex.BtcAddressRegex
	return &Phone{}
}
