package expertm

type Sundries []Sundry

type Sundry struct {
	JournalPrime string        `xml:"Journal_Prime"`
	DocNumber    string        `xml:"DocNumber"`
	DocDate      Date          `xml:"DocDate"`
	Status       int           `xml:"Status"`
	Details      SundryDetails `xml:"Details>Detail"`
}

type SundryDetails []SundryDetail

type SundryDetail struct {
	Account string  `xml:"Account"`
	Amount  Decimal `xml:"Amount"`
	DebCre  Side    `xml:"DebCre"`
	Ventil  int     `xml:"Ventil"`
	Ref     string  `xml:"Ref"`
}
