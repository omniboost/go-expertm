package expertm

import "encoding/xml"

type Import struct {
	XMLName xml.Name `xml:"ImportExpMPlus"`
	// XmlNS   string   `xml:"xmlns:prefix,attr"`

	Sales     Sales     `xml:"Sales"`
	Customers Customers `xml:"Customers"`
}
