package math

import (
	"encoding/xml"

	"github.com/dhx007/goffice"
)

type CT_SSubSupPr struct {
	AlnScr *CT_OnOff
	CtrlPr *CT_CtrlPr
}

func NewCT_SSubSupPr() *CT_SSubSupPr {
	ret := &CT_SSubSupPr{}
	return ret
}

func (m *CT_SSubSupPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.AlnScr != nil {
		sealnScr := xml.StartElement{Name: xml.Name{Local: "m:alnScr"}}
		e.EncodeElement(m.AlnScr, sealnScr)
	}
	if m.CtrlPr != nil {
		sectrlPr := xml.StartElement{Name: xml.Name{Local: "m:ctrlPr"}}
		e.EncodeElement(m.CtrlPr, sectrlPr)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_SSubSupPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_SSubSupPr:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "alnScr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "alnScr"}:
				m.AlnScr = NewCT_OnOff()
				if err := d.DecodeElement(m.AlnScr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "ctrlPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "ctrlPr"}:
				m.CtrlPr = NewCT_CtrlPr()
				if err := d.DecodeElement(m.CtrlPr, &el); err != nil {
					return err
				}
			default:
				goffice.Log("skipping unsupported element on CT_SSubSupPr %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_SSubSupPr
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_SSubSupPr and its children
func (m *CT_SSubSupPr) Validate() error {
	return m.ValidateWithPath("CT_SSubSupPr")
}

// ValidateWithPath validates the CT_SSubSupPr and its children, prefixing error messages with path
func (m *CT_SSubSupPr) ValidateWithPath(path string) error {
	if m.AlnScr != nil {
		if err := m.AlnScr.ValidateWithPath(path + "/AlnScr"); err != nil {
			return err
		}
	}
	if m.CtrlPr != nil {
		if err := m.CtrlPr.ValidateWithPath(path + "/CtrlPr"); err != nil {
			return err
		}
	}
	return nil
}
