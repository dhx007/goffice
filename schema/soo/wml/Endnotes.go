package wml

import (
	"encoding/xml"

	"github.com/dhx007/goffice"
)

type Endnotes struct {
	CT_Endnotes
}

func NewEndnotes() *Endnotes {
	ret := &Endnotes{}
	ret.CT_Endnotes = *NewCT_Endnotes()
	return ret
}

func (m *Endnotes) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/wordprocessingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:a"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:m"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/math"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:ma"}, Value: "http://schemas.openxmlformats.org/schemaLibrary/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:pic"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/picture"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:s"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:w"}, Value: "http://schemas.openxmlformats.org/wordprocessingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:wp"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "w:endnotes"
	return m.CT_Endnotes.MarshalXML(e, start)
}

func (m *Endnotes) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_Endnotes = *NewCT_Endnotes()
lEndnotes:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "endnote"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "endnote"}:
				tmp := NewCT_FtnEdn()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Endnote = append(m.Endnote, tmp)
			default:
				goffice.Log("skipping unsupported element on Endnotes %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lEndnotes
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the Endnotes and its children
func (m *Endnotes) Validate() error {
	return m.ValidateWithPath("Endnotes")
}

// ValidateWithPath validates the Endnotes and its children, prefixing error messages with path
func (m *Endnotes) ValidateWithPath(path string) error {
	if err := m.CT_Endnotes.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
