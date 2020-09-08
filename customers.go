package expertm

type Customers []Customer

type Customer struct {
	Prime                int    `xml:"Prime"`                          // Nummer
	Alfa                 string `xml:"Alfa,omitempty"`                 // Alfacode
	Name                 string `xml:"Name"`                           // Name
	Country              string `xml:"Country"`                        // Country
	Street               string `xml:"Street,omitempty"`               // Straat
	HouseNumber          string `xml:"HouseNumber,omitempty"`          // Huisnr.
	MailboxNumber        string `xml:"MailboxNumber,omitempty"`        // Bus
	ZipCode              string `xml:"ZipCode,omitempty"`              // Postnummer
	City                 string `xml:"City,omitempty"`                 // Woonplaats
	Language             int    `xml:"Language"`                       // Taal
	CurrencyCode         string `xml:"CurrencyCode"`                   // Munt
	VATCode              int    `xml:"VATCode"`                        // Btw-plichtig
	VATStatus            int    `xml:"VATStatus"`                      // Btw-status
	VATNumber            string `xml:"VATNumber,omitempty"`            // Ondernemingsnummer
	Rappel               string `xml:"Rappel,omitempty"`               // Betalingsherinnering
	Dom                  string `xml:"Dom,omitempty"`                  // Code domiciliëring
	Extra1               string `xml:"Extra1,omitempty"`               // Extra 1
	Extra2               string `xml:"Extra2,omitempty"`               // Extra 2
	Extra3               string `xml:"Extra3,omitempty"`               // Extra 3
	Phone                string `xml:"Phone,omitempty"`                // Telefoonnummer
	Fax                  string `xml:"Fax,omitempty"`                  // Faxnummer
	GSM                  string `xml:"GSM,omitempty"`                  // Gsm-nummer
	Email                string `xml:"Email,omitempty"`                // E-mail
	AccountSale          string `xml:"AccountSale,omitempty"`          // Verkooprekening
	GoodsCode            string `xml:"GoodsCode,omitempty"`            // Intrastat goederencode
	IntrastatTransport   string `xml:"IntrastatTransport,omitempty"`   // Intrastat vervoerswijze
	IntrastatTransaction string `xml:"IntrastatTransaction,omitempty"` // Intrastat aard transactie
	BankNumber           string `xml:"BankNumber,omitempty"`           // Banknummer
	BankNumberDom        string `xml:"BankNumberDom,omitempty"`        // Banknummer domiciliëring
	FreeField1           string `xml:"FreeField1,omitempty"`           // Vrij veld 1
	FreeField2           string `xml:"FreeField2,omitempty"`           // Vrij veld 2
	Title                string `xml:"Title,omitempty"`                // Aanspreektitel
	Website              string `xml:"Website,omitempty"`              // Website
	CountryVATNumber     string `xml:"CountryVATNumber,omitempty"`     // Land ondernemingsnummer
	CountryBankNumber    string `xml:"CountryBankNumber,omitempty"`    // Land banknummer
	CountryBankNumberDom string `xml:"CountryBankNumberDom,omitempty"` // Land banknummer domiciliëring
	IntrastatDelivery    string `xml:"IntrastatDelivery,omitempty"`    // Intrastat leveringsvoorwaarden
	DueDays              string `xml:"DueDays,omitempty"`              // Vervaldagen
	CustomerGroup        string `xml:"CustomerGroup,omitempty"`        // Klantengroep
	DISCOUNTPERC         string `xml:"DISCOUNTPERC,omitempty"`         // % Korting contant
	AccountCentral       string `xml:"AccountCentral,omitempty"`       // Centralisatierekening
	Due                  string `xml:"Due,omitempty"`                  // Vervaldag
	Ventil               string `xml:"Ventil,omitempty"`               // Ventilatiecode
	Status               int    `xml:"Status"`                         // Status
	BIC                  string `xml:"BIC,omitempty"`                  // BIC
	BICDOM               string `xml:"BICDOM,omitempty"`               // BIC domiciliëring
	MEMO                 string `xml:"MEMO,omitempty"`                 // Memotekst
	DelivDate            string `xml:"DelivDate,omitempty"`            // Voorstel leverdatum
	DatePayed            string `xml:"DatePayed,omitempty"`            // Voorstel datum betaald
	ElecTb               string `xml:"ElecTb,omitempty"`               // Terugbetaling klant
}
