package pii

type PiiType interface {
	GetName() string
}

type Phone struct {
	// number string
}

func (p *Phone) GetName() string {
	return "phone"
}

type Email struct {
	// number string
}

func (e *Email) GetName() string {
	return "phone"
}

type CreditCard struct {
	// number string
}

func (cc *CreditCard) GetName() string {
	return "phone"
}

type Address struct {
	// number string
}

func (add *Address) GetName() string {
	return "phone"
}

type Person struct {
	// number string
}

func (pers *Person) GetName() string {
	return "phone"
}

type BirthDate struct{}

func (bd *BirthDate) GetName() string {
	return "phone"
}

type Gender struct{}

func (gender *Gender) GetName() string {
	return "phone"
}

type Nationality struct{}

func (nat *Nationality) GetName() string {
	return "phone"
}

type SSN struct{}

func (ssn *SSN) GetName() string {
	return "phone"
}

type ZipCode struct{}

func (zip *ZipCode) GetName() string {
	return "phone"
}

type PoBox struct{}

func (po *PoBox) GetName() string {
	return "phone"
}

type UserName struct{}

func (user *UserName) GetName() string {
	return "phone"
}

type Password struct{}

func (pass *Password) GetName() string {
	return "phone"
}
