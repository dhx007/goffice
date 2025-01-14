package diagram

import (
	"encoding/xml"
	"fmt"
)

type CT_RelIds struct {
	DmAttr string
	LoAttr string
	QsAttr string
	CsAttr string
}

func NewCT_RelIds() *CT_RelIds {
	ret := &CT_RelIds{}
	return ret
}

func (m *CT_RelIds) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r:dm"},
		Value: fmt.Sprintf("%v", m.DmAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r:lo"},
		Value: fmt.Sprintf("%v", m.LoAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r:qs"},
		Value: fmt.Sprintf("%v", m.QsAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r:cs"},
		Value: fmt.Sprintf("%v", m.CsAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_RelIds) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Space == "http://schemas.openxmlformats.org/officeDocument/2006/relationships" && attr.Name.Local == "dm" ||
			attr.Name.Space == "http://purl.oclc.org/ooxml/officeDocument/relationships" && attr.Name.Local == "dm" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.DmAttr = parsed
			continue
		}
		if attr.Name.Space == "http://schemas.openxmlformats.org/officeDocument/2006/relationships" && attr.Name.Local == "lo" ||
			attr.Name.Space == "http://purl.oclc.org/ooxml/officeDocument/relationships" && attr.Name.Local == "lo" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.LoAttr = parsed
			continue
		}
		if attr.Name.Space == "http://schemas.openxmlformats.org/officeDocument/2006/relationships" && attr.Name.Local == "qs" ||
			attr.Name.Space == "http://purl.oclc.org/ooxml/officeDocument/relationships" && attr.Name.Local == "qs" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.QsAttr = parsed
			continue
		}
		if attr.Name.Space == "http://schemas.openxmlformats.org/officeDocument/2006/relationships" && attr.Name.Local == "cs" ||
			attr.Name.Space == "http://purl.oclc.org/ooxml/officeDocument/relationships" && attr.Name.Local == "cs" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CsAttr = parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_RelIds: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_RelIds and its children
func (m *CT_RelIds) Validate() error {
	return m.ValidateWithPath("CT_RelIds")
}

// ValidateWithPath validates the CT_RelIds and its children, prefixing error messages with path
func (m *CT_RelIds) ValidateWithPath(path string) error {
	return nil
}
