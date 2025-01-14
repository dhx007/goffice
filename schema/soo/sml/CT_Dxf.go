package sml

import (
	"encoding/xml"

	"github.com/dhx007/goffice"
)

type CT_Dxf struct {
	// Font Properties
	Font *CT_Font
	// Number Format
	NumFmt *CT_NumFmt
	// Fill
	Fill *CT_Fill
	// Alignment
	Alignment *CT_CellAlignment
	// Border Properties
	Border *CT_Border
	// Protection Properties
	Protection *CT_CellProtection
	// Future Feature Data Storage Area
	ExtLst *CT_ExtensionList
}

func NewCT_Dxf() *CT_Dxf {
	ret := &CT_Dxf{}
	return ret
}

func (m *CT_Dxf) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Font != nil {
		sefont := xml.StartElement{Name: xml.Name{Local: "ma:font"}}
		e.EncodeElement(m.Font, sefont)
	}
	if m.NumFmt != nil {
		senumFmt := xml.StartElement{Name: xml.Name{Local: "ma:numFmt"}}
		e.EncodeElement(m.NumFmt, senumFmt)
	}
	if m.Fill != nil {
		sefill := xml.StartElement{Name: xml.Name{Local: "ma:fill"}}
		e.EncodeElement(m.Fill, sefill)
	}
	if m.Alignment != nil {
		sealignment := xml.StartElement{Name: xml.Name{Local: "ma:alignment"}}
		e.EncodeElement(m.Alignment, sealignment)
	}
	if m.Border != nil {
		seborder := xml.StartElement{Name: xml.Name{Local: "ma:border"}}
		e.EncodeElement(m.Border, seborder)
	}
	if m.Protection != nil {
		seprotection := xml.StartElement{Name: xml.Name{Local: "ma:protection"}}
		e.EncodeElement(m.Protection, seprotection)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "ma:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Dxf) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Dxf:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "font"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "font"}:
				m.Font = NewCT_Font()
				if err := d.DecodeElement(m.Font, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "numFmt"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "numFmt"}:
				m.NumFmt = NewCT_NumFmt()
				if err := d.DecodeElement(m.NumFmt, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "fill"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "fill"}:
				m.Fill = NewCT_Fill()
				if err := d.DecodeElement(m.Fill, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "alignment"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "alignment"}:
				m.Alignment = NewCT_CellAlignment()
				if err := d.DecodeElement(m.Alignment, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "border"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "border"}:
				m.Border = NewCT_Border()
				if err := d.DecodeElement(m.Border, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "protection"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "protection"}:
				m.Protection = NewCT_CellProtection()
				if err := d.DecodeElement(m.Protection, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				goffice.Log("skipping unsupported element on CT_Dxf %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Dxf
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Dxf and its children
func (m *CT_Dxf) Validate() error {
	return m.ValidateWithPath("CT_Dxf")
}

// ValidateWithPath validates the CT_Dxf and its children, prefixing error messages with path
func (m *CT_Dxf) ValidateWithPath(path string) error {
	if m.Font != nil {
		if err := m.Font.ValidateWithPath(path + "/Font"); err != nil {
			return err
		}
	}
	if m.NumFmt != nil {
		if err := m.NumFmt.ValidateWithPath(path + "/NumFmt"); err != nil {
			return err
		}
	}
	if m.Fill != nil {
		if err := m.Fill.ValidateWithPath(path + "/Fill"); err != nil {
			return err
		}
	}
	if m.Alignment != nil {
		if err := m.Alignment.ValidateWithPath(path + "/Alignment"); err != nil {
			return err
		}
	}
	if m.Border != nil {
		if err := m.Border.ValidateWithPath(path + "/Border"); err != nil {
			return err
		}
	}
	if m.Protection != nil {
		if err := m.Protection.ValidateWithPath(path + "/Protection"); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
