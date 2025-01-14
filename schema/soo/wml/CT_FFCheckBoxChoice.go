package wml

import (
	"encoding/xml"

	"github.com/dhx007/goffice"
)

type CT_FFCheckBoxChoice struct {
	Size     *CT_HpsMeasure
	SizeAuto *CT_OnOff
}

func NewCT_FFCheckBoxChoice() *CT_FFCheckBoxChoice {
	ret := &CT_FFCheckBoxChoice{}
	return ret
}

func (m *CT_FFCheckBoxChoice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.Size != nil {
		sesize := xml.StartElement{Name: xml.Name{Local: "w:size"}}
		e.EncodeElement(m.Size, sesize)
	}
	if m.SizeAuto != nil {
		sesizeAuto := xml.StartElement{Name: xml.Name{Local: "w:sizeAuto"}}
		e.EncodeElement(m.SizeAuto, sesizeAuto)
	}
	return nil
}

func (m *CT_FFCheckBoxChoice) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_FFCheckBoxChoice:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "size"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "size"}:
				m.Size = NewCT_HpsMeasure()
				if err := d.DecodeElement(m.Size, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "sizeAuto"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "sizeAuto"}:
				m.SizeAuto = NewCT_OnOff()
				if err := d.DecodeElement(m.SizeAuto, &el); err != nil {
					return err
				}
			default:
				goffice.Log("skipping unsupported element on CT_FFCheckBoxChoice %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_FFCheckBoxChoice
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_FFCheckBoxChoice and its children
func (m *CT_FFCheckBoxChoice) Validate() error {
	return m.ValidateWithPath("CT_FFCheckBoxChoice")
}

// ValidateWithPath validates the CT_FFCheckBoxChoice and its children, prefixing error messages with path
func (m *CT_FFCheckBoxChoice) ValidateWithPath(path string) error {
	if m.Size != nil {
		if err := m.Size.ValidateWithPath(path + "/Size"); err != nil {
			return err
		}
	}
	if m.SizeAuto != nil {
		if err := m.SizeAuto.ValidateWithPath(path + "/SizeAuto"); err != nil {
			return err
		}
	}
	return nil
}
