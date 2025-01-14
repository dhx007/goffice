package wml

import (
	"encoding/xml"

	"github.com/dhx007/goffice"
)

type CT_Ruby struct {
	// Phonetic Guide Properties
	RubyPr *CT_RubyPr
	// Phonetic Guide Text
	Rt *CT_RubyContent
	// Phonetic Guide Base Text
	RubyBase *CT_RubyContent
}

func NewCT_Ruby() *CT_Ruby {
	ret := &CT_Ruby{}
	ret.RubyPr = NewCT_RubyPr()
	ret.Rt = NewCT_RubyContent()
	ret.RubyBase = NewCT_RubyContent()
	return ret
}

func (m *CT_Ruby) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	serubyPr := xml.StartElement{Name: xml.Name{Local: "w:rubyPr"}}
	e.EncodeElement(m.RubyPr, serubyPr)
	sert := xml.StartElement{Name: xml.Name{Local: "w:rt"}}
	e.EncodeElement(m.Rt, sert)
	serubyBase := xml.StartElement{Name: xml.Name{Local: "w:rubyBase"}}
	e.EncodeElement(m.RubyBase, serubyBase)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Ruby) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.RubyPr = NewCT_RubyPr()
	m.Rt = NewCT_RubyContent()
	m.RubyBase = NewCT_RubyContent()
lCT_Ruby:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "rubyPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "rubyPr"}:
				if err := d.DecodeElement(m.RubyPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "rt"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "rt"}:
				if err := d.DecodeElement(m.Rt, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "rubyBase"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "rubyBase"}:
				if err := d.DecodeElement(m.RubyBase, &el); err != nil {
					return err
				}
			default:
				goffice.Log("skipping unsupported element on CT_Ruby %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Ruby
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Ruby and its children
func (m *CT_Ruby) Validate() error {
	return m.ValidateWithPath("CT_Ruby")
}

// ValidateWithPath validates the CT_Ruby and its children, prefixing error messages with path
func (m *CT_Ruby) ValidateWithPath(path string) error {
	if err := m.RubyPr.ValidateWithPath(path + "/RubyPr"); err != nil {
		return err
	}
	if err := m.Rt.ValidateWithPath(path + "/Rt"); err != nil {
		return err
	}
	if err := m.RubyBase.ValidateWithPath(path + "/RubyBase"); err != nil {
		return err
	}
	return nil
}
