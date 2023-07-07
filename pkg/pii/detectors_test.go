package pii

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestColumnNameRegexDetector_Detect(t *testing.T) {
	detector := &ColumnNameRegexDetector{}

	Convey("Column name regex", t, func() {
		Convey("PersonColumnRegex", func() {
			Convey("pii column name", func() {
				columnNames := []string{
					"fname", "full_name", "name", "FNAME", "FULL_NAME", "NAME",
					"firstname", "lastname", "lname", "maidenname", "_name",
					"nickname", "name_suffix", "name|person",
				}

				assertTrue(columnNames, detector, t)
			})

			Convey("non pii column name", func() {
				columnNames := []string{
					"test", "n_ame",
				}

				assertFalse(columnNames, detector, t)
			})
		})

		Convey("EmailColumnRegex", func() {
			Convey("pii column name", func() {
				columnNames := []string{
					"email", "e-mail", "mail", "mailid", "mail-id", "mail_id", "identity",
					"e_mail",
				}

				assertTrue(columnNames, detector, t)
			})

			Convey("non pii column name", func() {
				columnNames := []string{
					"test", "em_ail",
				}

				assertFalse(columnNames, detector, t)
			})
		})

		Convey("BirthDateColumnRegex", func() {
			Convey("pii column name", func() {
				columnNames := []string{
					"date_of_birth", "dateofbirth", "dob",
					"birthday", "date_of_death", "dateofdeath", "birthdate",
				}

				assertTrue(columnNames, detector, t)
			})

			Convey("non pii column name", func() {
				columnNames := []string{
					"test", "d_ate",
				}

				assertFalse(columnNames, detector, t)
			})
		})

		Convey("GenderColumnRegex", func() {
			Convey("pii column name", func() {
				columnNames := []string{
					"gender", "GENDER", "gEnDeR",
				}

				assertTrue(columnNames, detector, t)
			})

			Convey("non pii column name", func() {
				columnNames := []string{
					"test", "g_ender",
				}

				assertFalse(columnNames, detector, t)
			})
		})

		Convey("NationalityColumnRegex", func() {
			Convey("pii column name", func() {
				columnNames := []string{
					"nationality", "NATIONALITY", "_nAtionAlitY",
				}

				assertTrue(columnNames, detector, t)
			})

			Convey("non pii column name", func() {
				columnNames := []string{
					"test", "g_ender",
				}

				assertFalse(columnNames, detector, t)
			})
		})

		Convey("AddressColumnRegex", func() {
			Convey("pii column name", func() {
				columnNames := []string{
					"address", "city", "state", "county", "country", "zone", "borough",
				}

				assertTrue(columnNames, detector, t)
			})

			Convey("non pii column name", func() {
				columnNames := []string{
					"test", "g_ender",
				}

				assertFalse(columnNames, detector, t)
			})
		})

		Convey("ZipCodeColumnRegex", func() {
			Convey("pii column name", func() {
				columnNames := []string{
					"zipcode", "zip_code", "postal", "postal_code", "zip",
				}

				assertTrue(columnNames, detector, t)
			})

			Convey("non pii column name", func() {
				columnNames := []string{
					"test", "g_ender",
				}

				assertFalse(columnNames, detector, t)
			})
		})

		Convey("UsernameColumnRegex", func() {
			Convey("pii column name", func() {
				columnNames := []string{
					"user", "userid", "name", "username",
				}

				assertTrue(columnNames, detector, t)
			})

			Convey("non pii column name", func() {
				columnNames := []string{
					"test", "g_ender",
				}

				assertFalse(columnNames, detector, t)
			})
		})

		Convey("PasswordColumnRegex", func() {
			Convey("pii column name", func() {
				columnNames := []string{
					"pass", "passThebeer", "password", "PASSWORD",
				}

				assertTrue(columnNames, detector, t)
			})

			Convey("non pii column name", func() {
				columnNames := []string{
					"test",
				}

				assertFalse(columnNames, detector, t)
			})
		})

		Convey("SSNColumnRegex", func() {
			Convey("pii column name", func() {
				columnNames := []string{
					"ssn", "social_number", "social_security", "social_security_number", "social_security_no",
				}

				assertTrue(columnNames, detector, t)
			})

			Convey("non pii column name", func() {
				columnNames := []string{
					"test",
				}

				assertFalse(columnNames, detector, t)
			})
		})

		Convey("POBoxColumnRegex", func() {
			Convey("pii column name", func() {
				columnNames := []string{
					"po_box", "pobox",
				}

				assertTrue(columnNames, detector, t)
			})

			Convey("non pii column name", func() {
				columnNames := []string{
					"test",
				}

				assertFalse(columnNames, detector, t)
			})
		})

		Convey("CreditCardColumnRegex", func() {
			Convey("pii column name", func() {
				columnNames := []string{
					"credit_card", "cc_number", "cc_num", "creditcard", "credit_card_num", "creditcardnumber",
				}

				assertTrue(columnNames, detector, t)
			})

			Convey("non pii column name", func() {
				columnNames := []string{
					"test",
				}

				assertFalse(columnNames, detector, t)
			})
		})

		Convey("PhoneColumnRegex", func() {
			Convey("pii column name", func() {
				columnNames := []string{
					"phone", "phone_number", "phone_no", "phone_num", "telephone", "telephone_num", "telephone_no",
				}

				assertTrue(columnNames, detector, t)
			})

			Convey("non pii column name", func() {
				columnNames := []string{
					"test",
				}

				assertFalse(columnNames, detector, t)
			})
		})
	})
}

func assertTrue(columnNames []string, detector Detector, t *testing.T) {
	for _, column := range columnNames {
		isPii := detector.Detect(column)
		So(isPii, func(actual any, expected ...any) string {
			result := ShouldBeTrue(actual, expected...)
			if result != "" {
				t.Logf("failed for input %s", column)
			}

			return result
		})
	}
}

func assertFalse(columnNames []string, detector Detector, t *testing.T) {
	for _, column := range columnNames {
		isPii := detector.Detect(column)
		So(isPii, func(actual any, expected ...any) string {
			result := ShouldBeFalse(actual, expected...)
			if result != "" {
				t.Logf("failed for input %s", column)
			}

			return result
		})
	}
}
