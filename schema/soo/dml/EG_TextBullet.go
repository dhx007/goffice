package dml

import (
	"encoding/xml"

	"github.com/dhx007/goffice"
)

type EG_TextBullet struct {
	BuNone    *CT_TextNoBullet
	BuAutoNum *CT_TextAutonumberBullet
	BuChar    *CT_TextCharBullet
	BuBlip    *CT_TextBlipBullet
}

func NewEG_TextBullet() *EG_TextBullet {
	ret := &EG_TextBullet{}
	return ret
}

func (m *EG_TextBullet) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.BuNone != nil {
		sebuNone := xml.StartElement{Name: xml.Name{Local: "a:buNone"}}
		e.EncodeElement(m.BuNone, sebuNone)
	}
	if m.BuAutoNum != nil {
		sebuAutoNum := xml.StartElement{Name: xml.Name{Local: "a:buAutoNum"}}
		e.EncodeElement(m.BuAutoNum, sebuAutoNum)
	}
	if m.BuChar != nil {
		sebuChar := xml.StartElement{Name: xml.Name{Local: "a:buChar"}}
		e.EncodeElement(m.BuChar, sebuChar)
	}
	if m.BuBlip != nil {
		sebuBlip := xml.StartElement{Name: xml.Name{Local: "a:buBlip"}}
		e.EncodeElement(m.BuBlip, sebuBlip)
	}
	return nil
}

func (m *EG_TextBullet) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lEG_TextBullet:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "buNone"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "buNone"}:
				m.BuNone = NewCT_TextNoBullet()
				if err := d.DecodeElement(m.BuNone, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "buAutoNum"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "buAutoNum"}:
				m.BuAutoNum = NewCT_TextAutonumberBullet()
				if err := d.DecodeElement(m.BuAutoNum, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "buChar"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "buChar"}:
				m.BuChar = NewCT_TextCharBullet()
				if err := d.DecodeElement(m.BuChar, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "buBlip"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "buBlip"}:
				m.BuBlip = NewCT_TextBlipBullet()
				if err := d.DecodeElement(m.BuBlip, &el); err != nil {
					return err
				}
			default:
				goffice.Log("skipping unsupported element on EG_TextBullet %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lEG_TextBullet
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the EG_TextBullet and its children
func (m *EG_TextBullet) Validate() error {
	return m.ValidateWithPath("EG_TextBullet")
}

// ValidateWithPath validates the EG_TextBullet and its children, prefixing error messages with path
func (m *EG_TextBullet) ValidateWithPath(path string) error {
	if m.BuNone != nil {
		if err := m.BuNone.ValidateWithPath(path + "/BuNone"); err != nil {
			return err
		}
	}
	if m.BuAutoNum != nil {
		if err := m.BuAutoNum.ValidateWithPath(path + "/BuAutoNum"); err != nil {
			return err
		}
	}
	if m.BuChar != nil {
		if err := m.BuChar.ValidateWithPath(path + "/BuChar"); err != nil {
			return err
		}
	}
	if m.BuBlip != nil {
		if err := m.BuBlip.ValidateWithPath(path + "/BuBlip"); err != nil {
			return err
		}
	}
	return nil
}
