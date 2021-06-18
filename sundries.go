package expertm

type Sundries []Sundry

type Sundry struct {
	JournalPrime     int           `xml:"Journal_Prime"`              // Dagboek
	DocNumber        string        `xml:"DocNumber"`                  // Documentnr.
	DocDate          Date          `xml:"DocDate"`                    // Datum
	YearAlfa         int           `xml:"Year_Alfa,omitempty"`        // Boekjaar
	AccountingPeriod string        `xml:"AccountingPeriod,omitempty"` // Boekhoudperiode
	VATMonth         Month         `xml:"VATMonth,omitempty"`         // Btw-periode
	Status           int           `xml:"Status"`                     // Status
	Annex            string        `xml:"Annex,omitempty"`            // Pad bijlage
	Details          SundryDetails `xml:"Details>Detail"`
}

type SundryDetails []SundryDetail

type SundryDetail struct {
	Account int     `xml:"Account"`
	Amount  Decimal `xml:"Amount"`
	DebCre  Side    `xml:"DebCre"`
	Ventil  int     `xml:"Ventil"`
	Ref     string  `xml:"Ref"`
}
