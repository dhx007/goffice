package math

import (
	"encoding/xml"

	"github.com/dhx007/goffice"
)

type CT_PhantPr struct {
	Show     *CT_OnOff
	ZeroWid  *CT_OnOff
	ZeroAsc  *CT_OnOff
	ZeroDesc *CT_OnOff
	Transp   *CT_OnOff
	CtrlPr   *CT_CtrlPr
}

func NewCT_PhantPr() *CT_PhantPr {
	ret := &CT_PhantPr{}
	return ret
}

func (m *CT_PhantPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Show != nil {
		seshow := xml.StartElement{Name: xml.Name{Local: "m:show"}}
		e.EncodeElement(m.Show, seshow)
	}
	if m.ZeroWid != nil {
		sezeroWid := xml.StartElement{Name: xml.Name{Local: "m:zeroWid"}}
		e.EncodeElement(m.ZeroWid, sezeroWid)
	}
	if m.ZeroAsc != nil {
		sezeroAsc := xml.StartElement{Name: xml.Name{Local: "m:zeroAsc"}}
		e.EncodeElement(m.ZeroAsc, sezeroAsc)
	}
	if m.ZeroDesc != nil {
		sezeroDesc := xml.StartElement{Name: xml.Name{Local: "m:zeroDesc"}}
		e.EncodeElement(m.ZeroDesc, sezeroDesc)
	}
	if m.Transp != nil {
		setransp := xml.StartElement{Name: xml.Name{Local: "m:transp"}}
		e.EncodeElement(m.Transp, setransp)
	}
	if m.CtrlPr != nil {
		sectrlPr := xml.StartElement{Name: xml.Name{Local: "m:ctrlPr"}}
		e.EncodeElement(m.CtrlPr, sectrlPr)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_PhantPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_PhantPr:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "show"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "show"}:
				m.Show = NewCT_OnOff()
				if err := d.DecodeElement(m.Show, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "zeroWid"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "zeroWid"}:
				m.ZeroWid = NewCT_OnOff()
				if err := d.DecodeElement(m.ZeroWid, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "zeroAsc"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "zeroAsc"}:
				m.ZeroAsc = NewCT_OnOff()
				if err := d.DecodeElement(m.ZeroAsc, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "zeroDesc"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "zeroDesc"}:
				m.ZeroDesc = NewCT_OnOff()
				if err := d.DecodeElement(m.ZeroDesc, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "transp"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "transp"}:
				m.Transp = NewCT_OnOff()
				if err := d.DecodeElement(m.Transp, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "ctrlPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "ctrlPr"}:
				m.CtrlPr = NewCT_CtrlPr()
				if err := d.DecodeElement(m.CtrlPr, &el); err != nil {
					return err
				}
			default:
				goffice.Log("skipping unsupported element on CT_PhantPr %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_PhantPr
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_PhantPr and its children
func (m *CT_PhantPr) Validate() error {
	return m.ValidateWithPath("CT_PhantPr")
}

// ValidateWithPath validates the CT_PhantPr and its children, prefixing error messages with path
func (m *CT_PhantPr) ValidateWithPath(path string) error {
	if m.Show != nil {
		if err := m.Show.ValidateWithPath(path + "/Show"); err != nil {
			return err
		}
	}
	if m.ZeroWid != nil {
		if err := m.ZeroWid.ValidateWithPath(path + "/ZeroWid"); err != nil {
			return err
		}
	}
	if m.ZeroAsc != nil {
		if err := m.ZeroAsc.ValidateWithPath(path + "/ZeroAsc"); err != nil {
			return err
		}
	}
	if m.ZeroDesc != nil {
		if err := m.ZeroDesc.ValidateWithPath(path + "/ZeroDesc"); err != nil {
			return err
		}
	}
	if m.Transp != nil {
		if err := m.Transp.ValidateWithPath(path + "/Transp"); err != nil {
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
