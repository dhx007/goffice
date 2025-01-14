package chart

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_Double struct {
	ValAttr float64
}

func NewCT_Double() *CT_Double {
	ret := &CT_Double{}
	return ret
}

func (m *CT_Double) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "val"},
		Value: fmt.Sprintf("%v", m.ValAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Double) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			m.ValAttr = parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_Double: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_Double and its children
func (m *CT_Double) Validate() error {
	return m.ValidateWithPath("CT_Double")
}

// ValidateWithPath validates the CT_Double and its children, prefixing error messages with path
func (m *CT_Double) ValidateWithPath(path string) error {
	return nil
}
