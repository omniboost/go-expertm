package expertm

type Sales []Sale

type Sale struct {
	JournalPrime     int          `xml:"Journal_Prime,omitempty"`    // Dagboek
	CustomerPrime    int          `xml:"Customer_Prime"`             // Klant
	CurrencyCode     string       `xml:"CurrencyCode"`               // Munt
	DocType          int          `xml:"DocType`                     // Document
	DocNumber        string       `xml:"DocNumber"`                  // Documentnummer
	DocDate          Date         `xml:"DocDate,omitempty"`          // Datum
	DueDate          Date         `xml:"DueDate,omitempty"`          // Vervaldatum
	OurRef           string       `xml:"OurRef,omitempty"`           // Onze referte
	YourRef          string       `xml:"YourRef,omitempty"`          // Uw referte
	Amount           Decimal      `xml:"Amount"`                     // Bedrag
	Discount         Decimal      `xml:"Decimal,omitempty"`          // Korting
	AccountingPeriod string       `xml:"AccountingPeriod,omitempty"` // Boekhoudperiode
	VATMonth         Month        `xml:"VATMonth,omitempty"`         // Btw-periode
	YearAlfa         int          `xml:"Year_Alfa,omitempty"`        // Jaar
	Ventil           int          `xml:"Ventil,omitempty"`           // Ventilatie
	VATAmount        Decimal      `xml:"VATAmount,omitempty"`        // Btw-bedrag
	Status           int          `xml:"Status"`                     // Status
	Rate             string       `xml:"Rate,omitempty"`             // Koers
	OGM              string       `xml:"OGM,omitempty"`              // Gestructureerde mededeling
	Annex            string       `xml:"Annex,omitempty"`            // Pad bijlage
	DocState         string       `xml:"DocState,omitempty"`         // Toestand document
	DelivDate        string       `xml:"DelivDate,omitempty"`        // Leverdatum
	DatePayed        string       `xml:"DatePayed"`                  // Datum betaald
	PayDate          Date         `xml:"PayDate,omitempty"`          // Kassa betaaldatum
	Details          SalesDetails `xml:"Details>Detail"`
}

type SalesDetails []SalesDetail

type SalesDetail struct {
	Account     int     `xml:"Account"`               // Grootboekrekening
	Amount      Decimal `xml:"Amount"`                // Bedrag
	DebCre      Side    `xml:"DebCre"`                // D/C
	Ventil      int     `xml:"Ventil"`                // Ventilatie
	Ref         string  `xml:"Ref,omitempty"`         // Referte
	Unit1       string  `xml:"Unit1,omitempty"`       // Eenheid 1
	Unit2       string  `xml:"Unit2,omitempty"`       // Eenheid 2
	VAT1        string  `xml:"VAT1,omitempty"`        // Btw-rooster 1
	VatProc     Decimal `xml:"VatProc,omitempty"`     // Btw-percentage
	MOSSDefault string  `xml:"MOSSDefault,omitempty"` // Normaal of verlaagd tarief
}
