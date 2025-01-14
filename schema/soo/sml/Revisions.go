package sml

import (
	"encoding/xml"

	"github.com/dhx007/goffice"
)

type Revisions struct {
	CT_Revisions
}

func NewRevisions() *Revisions {
	ret := &Revisions{}
	ret.CT_Revisions = *NewCT_Revisions()
	return ret
}

func (m *Revisions) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/spreadsheetml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:ma"}, Value: "http://schemas.openxmlformats.org/spreadsheetml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:s"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xdr"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "ma:revisions"
	return m.CT_Revisions.MarshalXML(e, start)
}

func (m *Revisions) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_Revisions = *NewCT_Revisions()
lRevisions:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "rrc"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "rrc"}:
				tmp := NewCT_RevisionRowColumn()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Rrc = append(m.Rrc, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "rm"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "rm"}:
				tmp := NewCT_RevisionMove()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Rm = append(m.Rm, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "rcv"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "rcv"}:
				tmp := NewCT_RevisionCustomView()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Rcv = append(m.Rcv, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "rsnm"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "rsnm"}:
				tmp := NewCT_RevisionSheetRename()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Rsnm = append(m.Rsnm, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "ris"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "ris"}:
				tmp := NewCT_RevisionInsertSheet()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Ris = append(m.Ris, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "rcc"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "rcc"}:
				tmp := NewCT_RevisionCellChange()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Rcc = append(m.Rcc, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "rfmt"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "rfmt"}:
				tmp := NewCT_RevisionFormatting()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Rfmt = append(m.Rfmt, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "raf"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "raf"}:
				tmp := NewCT_RevisionAutoFormatting()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Raf = append(m.Raf, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "rdn"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "rdn"}:
				tmp := NewCT_RevisionDefinedName()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Rdn = append(m.Rdn, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "rcmt"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "rcmt"}:
				tmp := NewCT_RevisionComment()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Rcmt = append(m.Rcmt, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "rqt"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "rqt"}:
				tmp := NewCT_RevisionQueryTableField()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Rqt = append(m.Rqt, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "rcft"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "rcft"}:
				tmp := NewCT_RevisionConflict()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Rcft = append(m.Rcft, tmp)
			default:
				goffice.Log("skipping unsupported element on Revisions %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lRevisions
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the Revisions and its children
func (m *Revisions) Validate() error {
	return m.ValidateWithPath("Revisions")
}

// ValidateWithPath validates the Revisions and its children, prefixing error messages with path
func (m *Revisions) ValidateWithPath(path string) error {
	if err := m.CT_Revisions.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
