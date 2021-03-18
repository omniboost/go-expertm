package expertm

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strings"
	"time"
)

var (
	SideDebet  Side = 1
	SideCredit Side = -1
)

type Date struct {
	time.Time
	Layout string
}

func (d Date) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if d.Time.IsZero() {
		return e.Encode(nil)
	}

	return e.EncodeElement(d.Time.Format(d.Layout), start)
}

func (d Date) MarshalJSON() ([]byte, error) {
	if d.Time.IsZero() {
		return json.Marshal(nil)
	}

	return json.Marshal(d.Time.Format(d.Layout))
}

// func (d Date) MarshalSchema() string {
// 	layout := "2006-01-02"
// 	return d.Time.Format(layout)
// }

func (d Date) IsEmpty() bool {
	return d.Time.IsZero()
}

func (d *Date) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	var value string
	err := dec.DecodeElement(&value, &start)
	if err != nil {
		return err
	}

	if value == "" {
		return nil
	}

	// first try standard date
	layout := time.RFC3339
	d.Time, err = time.Parse(layout, value)
	if err == nil {
		d.Layout = layout
		return nil
	}

	layout = "02/01/2006"
	d.Time, err = time.Parse(layout, value)
	if err == nil {
		d.Layout = layout
		return err
	}

	return err
}

func (d *Date) UnmarshalJSON(data []byte) (err error) {
	var value string
	err = json.Unmarshal(data, &value)
	if err != nil {
		return err
	}

	if value == "" {
		return nil
	}

	// first try standard date
	layout := time.RFC3339
	d.Time, err = time.Parse(layout, value)
	if err == nil {
		return nil
	}

	// try iso8601 date format
	layout = "2006-01-02"
	d.Time, err = time.Parse(layout, value)
	if err == nil {
		d.Layout = layout
		return nil
	}

	layout = "02/01/2006"
	d.Time, err = time.Parse(layout, value)
	d.Layout = layout
	return err
}

type Month struct {
	time.Time
	Layout string
}

func (m Month) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.Time.IsZero() {
		return e.Encode(nil)
	}

	return e.EncodeElement(m.Time.Format(m.Layout), start)
}

func (m Month) MarshalJSON() ([]byte, error) {
	if m.Time.IsZero() {
		return json.Marshal(nil)
	}

	return json.Marshal(m.Time.Format(m.Layout))
}

// func (d Month) MarshalSchema() string {
// 	layout := "2006-01-02"
// 	return d.Time.Format(layout)
// }

func (m Month) IsEmpty() bool {
	return m.Time.IsZero()
}

func (m *Month) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	var value string
	err := dec.DecodeElement(&value, &start)
	if err != nil {
		return err
	}

	if value == "" {
		return nil
	}

	// first try standard date
	layout := time.RFC3339
	m.Time, err = time.Parse(layout, value)
	if err == nil {
		m.Layout = layout
		return nil
	}

	layout = "200601"
	m.Time, err = time.Parse(layout, value)
	if err == nil {
		m.Layout = layout
		return err
	}

	return nil
}

func (m *Month) UnmarshalJSON(data []byte) (err error) {
	var value string
	err = json.Unmarshal(data, &value)
	if err != nil {
		return err
	}

	if value == "" {
		return nil
	}

	// first try standard date
	layout := time.RFC3339
	m.Time, err = time.Parse(layout, value)
	if err == nil {
		m.Layout = layout
		return nil
	}

	// try iso8601 date format
	layout = "2006-01-02"
	m.Time, err = time.Parse(layout, value)
	if err == nil {
		m.Layout = layout
		return nil
	}

	layout = "200601"
	m.Time, err = time.Parse(layout, value)
	m.Layout = layout
	return err
}

type Decimal float64

func (d Decimal) String() string {
	if d == 0.0 {
		return ""
	}

	return d.Round2()
}

func (d Decimal) Round2() string {
	s := fmt.Sprintf("%.2f", float64(d))
	return strings.Replace(s, ".", ",", -1)
}

func (d Decimal) MarshalText() ([]byte, error) {
	s := d.String()
	s = strings.Replace(d.String(), ".", ",", -1)
	return []byte(s), nil
}

type Side int
