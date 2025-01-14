package sml

import (
	"encoding/xml"

	"github.com/dhx007/goffice"
)

type CalcChain struct {
	CT_CalcChain
}

func NewCalcChain() *CalcChain {
	ret := &CalcChain{}
	ret.CT_CalcChain = *NewCT_CalcChain()
	return ret
}

func (m *CalcChain) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/spreadsheetml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:ma"}, Value: "http://schemas.openxmlformats.org/spreadsheetml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:s"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xdr"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "ma:calcChain"
	return m.CT_CalcChain.MarshalXML(e, start)
}

func (m *CalcChain) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_CalcChain = *NewCT_CalcChain()
lCalcChain:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "c"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "c"}:
				tmp := NewCT_CalcCell()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.C = append(m.C, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				goffice.Log("skipping unsupported element on CalcChain %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCalcChain
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CalcChain and its children
func (m *CalcChain) Validate() error {
	return m.ValidateWithPath("CalcChain")
}

// ValidateWithPath validates the CalcChain and its children, prefixing error messages with path
func (m *CalcChain) ValidateWithPath(path string) error {
	if err := m.CT_CalcChain.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
