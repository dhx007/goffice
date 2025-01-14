package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/dhx007/goffice"
)

type CT_RevisionCellChange struct {
	// Sheet Id
	SIdAttr uint32
	// Old Formatting Information
	OdxfAttr *bool
	// Row Column Formatting Change
	XfDxfAttr *bool
	// Style Revision
	SAttr *bool
	// Formatting
	DxfAttr *bool
	// Number Format Id
	NumFmtIdAttr *uint32
	// Quote Prefix
	QuotePrefixAttr *bool
	// Old Quote Prefix
	OldQuotePrefixAttr *bool
	// Phonetic Text
	PhAttr *bool
	// Old Phonetic Text
	OldPhAttr *bool
	// End of List Formula Update
	EndOfListFormulaUpdateAttr *bool
	// Old Cell Data
	Oc *CT_Cell
	// New Cell Data
	Nc *CT_Cell
	// Old Formatting Information
	Odxf *CT_Dxf
	// New Formatting Information
	Ndxf    *CT_Dxf
	ExtLst  *CT_ExtensionList
	RIdAttr *uint32
	UaAttr  *bool
	RaAttr  *bool
}

func NewCT_RevisionCellChange() *CT_RevisionCellChange {
	ret := &CT_RevisionCellChange{}
	ret.Nc = NewCT_Cell()
	return ret
}

func (m *CT_RevisionCellChange) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "sId"},
		Value: fmt.Sprintf("%v", m.SIdAttr)})
	if m.OdxfAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "odxf"},
			Value: fmt.Sprintf("%d", b2i(*m.OdxfAttr))})
	}
	if m.XfDxfAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xfDxf"},
			Value: fmt.Sprintf("%d", b2i(*m.XfDxfAttr))})
	}
	if m.SAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "s"},
			Value: fmt.Sprintf("%d", b2i(*m.SAttr))})
	}
	if m.DxfAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "dxf"},
			Value: fmt.Sprintf("%d", b2i(*m.DxfAttr))})
	}
	if m.NumFmtIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "numFmtId"},
			Value: fmt.Sprintf("%v", *m.NumFmtIdAttr)})
	}
	if m.QuotePrefixAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "quotePrefix"},
			Value: fmt.Sprintf("%d", b2i(*m.QuotePrefixAttr))})
	}
	if m.OldQuotePrefixAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "oldQuotePrefix"},
			Value: fmt.Sprintf("%d", b2i(*m.OldQuotePrefixAttr))})
	}
	if m.PhAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ph"},
			Value: fmt.Sprintf("%d", b2i(*m.PhAttr))})
	}
	if m.OldPhAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "oldPh"},
			Value: fmt.Sprintf("%d", b2i(*m.OldPhAttr))})
	}
	if m.EndOfListFormulaUpdateAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "endOfListFormulaUpdate"},
			Value: fmt.Sprintf("%d", b2i(*m.EndOfListFormulaUpdateAttr))})
	}
	if m.RIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "rId"},
			Value: fmt.Sprintf("%v", *m.RIdAttr)})
	}
	if m.UaAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ua"},
			Value: fmt.Sprintf("%d", b2i(*m.UaAttr))})
	}
	if m.RaAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ra"},
			Value: fmt.Sprintf("%d", b2i(*m.RaAttr))})
	}
	e.EncodeToken(start)
	if m.Oc != nil {
		seoc := xml.StartElement{Name: xml.Name{Local: "ma:oc"}}
		e.EncodeElement(m.Oc, seoc)
	}
	senc := xml.StartElement{Name: xml.Name{Local: "ma:nc"}}
	e.EncodeElement(m.Nc, senc)
	if m.Odxf != nil {
		seodxf := xml.StartElement{Name: xml.Name{Local: "ma:odxf"}}
		e.EncodeElement(m.Odxf, seodxf)
	}
	if m.Ndxf != nil {
		sendxf := xml.StartElement{Name: xml.Name{Local: "ma:ndxf"}}
		e.EncodeElement(m.Ndxf, sendxf)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "ma:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_RevisionCellChange) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Nc = NewCT_Cell()
	for _, attr := range start.Attr {
		if attr.Name.Local == "oldPh" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.OldPhAttr = &parsed
			continue
		}
		if attr.Name.Local == "endOfListFormulaUpdate" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.EndOfListFormulaUpdateAttr = &parsed
			continue
		}
		if attr.Name.Local == "odxf" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.OdxfAttr = &parsed
			continue
		}
		if attr.Name.Local == "s" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.SAttr = &parsed
			continue
		}
		if attr.Name.Local == "numFmtId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.NumFmtIdAttr = &pt
			continue
		}
		if attr.Name.Local == "oldQuotePrefix" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.OldQuotePrefixAttr = &parsed
			continue
		}
		if attr.Name.Local == "ph" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.PhAttr = &parsed
			continue
		}
		if attr.Name.Local == "sId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.SIdAttr = uint32(parsed)
			continue
		}
		if attr.Name.Local == "xfDxf" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.XfDxfAttr = &parsed
			continue
		}
		if attr.Name.Local == "dxf" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DxfAttr = &parsed
			continue
		}
		if attr.Name.Local == "quotePrefix" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.QuotePrefixAttr = &parsed
			continue
		}
		if attr.Name.Local == "rId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.RIdAttr = &pt
			continue
		}
		if attr.Name.Local == "ua" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.UaAttr = &parsed
			continue
		}
		if attr.Name.Local == "ra" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.RaAttr = &parsed
			continue
		}
	}
lCT_RevisionCellChange:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "oc"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "oc"}:
				m.Oc = NewCT_Cell()
				if err := d.DecodeElement(m.Oc, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "nc"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "nc"}:
				if err := d.DecodeElement(m.Nc, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "odxf"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "odxf"}:
				m.Odxf = NewCT_Dxf()
				if err := d.DecodeElement(m.Odxf, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "ndxf"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "ndxf"}:
				m.Ndxf = NewCT_Dxf()
				if err := d.DecodeElement(m.Ndxf, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				goffice.Log("skipping unsupported element on CT_RevisionCellChange %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_RevisionCellChange
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_RevisionCellChange and its children
func (m *CT_RevisionCellChange) Validate() error {
	return m.ValidateWithPath("CT_RevisionCellChange")
}

// ValidateWithPath validates the CT_RevisionCellChange and its children, prefixing error messages with path
func (m *CT_RevisionCellChange) ValidateWithPath(path string) error {
	if m.Oc != nil {
		if err := m.Oc.ValidateWithPath(path + "/Oc"); err != nil {
			return err
		}
	}
	if err := m.Nc.ValidateWithPath(path + "/Nc"); err != nil {
		return err
	}
	if m.Odxf != nil {
		if err := m.Odxf.ValidateWithPath(path + "/Odxf"); err != nil {
			return err
		}
	}
	if m.Ndxf != nil {
		if err := m.Ndxf.ValidateWithPath(path + "/Ndxf"); err != nil {
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
