package wml

import (
	"encoding/xml"

	"github.com/dhx007/goffice"
)

type CT_DocPart struct {
	// Glossary Document Entry Properties
	DocPartPr *CT_DocPartPr
	// Contents of Glossary Document Entry
	DocPartBody *CT_Body
}

func NewCT_DocPart() *CT_DocPart {
	ret := &CT_DocPart{}
	return ret
}

func (m *CT_DocPart) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.DocPartPr != nil {
		sedocPartPr := xml.StartElement{Name: xml.Name{Local: "w:docPartPr"}}
		e.EncodeElement(m.DocPartPr, sedocPartPr)
	}
	if m.DocPartBody != nil {
		sedocPartBody := xml.StartElement{Name: xml.Name{Local: "w:docPartBody"}}
		e.EncodeElement(m.DocPartBody, sedocPartBody)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_DocPart) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_DocPart:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "docPartPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "docPartPr"}:
				m.DocPartPr = NewCT_DocPartPr()
				if err := d.DecodeElement(m.DocPartPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "docPartBody"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "docPartBody"}:
				m.DocPartBody = NewCT_Body()
				if err := d.DecodeElement(m.DocPartBody, &el); err != nil {
					return err
				}
			default:
				goffice.Log("skipping unsupported element on CT_DocPart %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_DocPart
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_DocPart and its children
func (m *CT_DocPart) Validate() error {
	return m.ValidateWithPath("CT_DocPart")
}

// ValidateWithPath validates the CT_DocPart and its children, prefixing error messages with path
func (m *CT_DocPart) ValidateWithPath(path string) error {
	if m.DocPartPr != nil {
		if err := m.DocPartPr.ValidateWithPath(path + "/DocPartPr"); err != nil {
			return err
		}
	}
	if m.DocPartBody != nil {
		if err := m.DocPartBody.ValidateWithPath(path + "/DocPartBody"); err != nil {
			return err
		}
	}
	return nil
}
