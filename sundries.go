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
	Ventil        int     `xml:"Ventil"`                   // Ventilatie
	Account       int     `xml:"Account,omitempty"`        // Grootboekrekening
	Amount        Decimal `xml:"Amount"`                   // Bedrag document
	AmountRef     Decimal `xml:"AmountRef"`                // Bedrag in referentiemunt
	CurrencyCode  string  `xml:"CurrencyCode,omitempty"`   // Munt
	CustomerPrime int     `xml:"Customer_Prime,omitempty"` // Klant
	DebCre        Side    `xml:"DebCre"`                   // D/C
	DocNumber     string  `xml:"DocNumber,omitempty"`      // Documentnr.
	DocType       string  `xml:"DocType,omitempty"`        // Documenttype
	Ref           string  `xml:"Ref,omitempty"`            // Referte
	SupplierPrime int     `xml:"Supplier_Prime,omitempty"` // Leverancier
	TypeDiv       string  `xml:"TypeDiv,omitempty"`        // Leverancier
	Unit1         string  `xml:"Unit1,omitempty"`          // Eenheid1
	Unit2         string  `xml:"Unit2,omitempty"`          // Eenheid2
}
