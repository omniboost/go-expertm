package expertm

import "encoding/xml"

type Import struct {
	XMLName xml.Name `xml:"ImportExpMPlus"`
	// XmlNS   string   `xml:"xmlns:prefix,attr"`

	Sales     Sales     `xml:"Sales>Sale"`
	Customers Customers `xml:"Customers>Customer"`
}
