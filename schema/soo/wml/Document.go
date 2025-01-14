package wml

import (
	"encoding/xml"

	"github.com/dhx007/goffice"
)

type Document struct {
	CT_Document
}

func NewDocument() *Document {
	ret := &Document{}
	ret.CT_Document = *NewCT_Document()
	return ret
}

func (m *Document) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	start.Name.Local = "w:document"
	return m.CT_Document.MarshalXML(e, start)
}

func (m *Document) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_Document = *NewCT_Document()
	for _, attr := range start.Attr {
		if attr.Name.Local == "conformance" {
			m.ConformanceAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
lDocument:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "background"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "background"}:
				m.Background = NewCT_Background()
				if err := d.DecodeElement(m.Background, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "body"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "body"}:
				m.Body = NewCT_Body()
				if err := d.DecodeElement(m.Body, &el); err != nil {
					return err
				}
			default:
				goffice.Log("skipping unsupported element on Document %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lDocument
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the Document and its children
func (m *Document) Validate() error {
	return m.ValidateWithPath("Document")
}

// ValidateWithPath validates the Document and its children, prefixing error messages with path
func (m *Document) ValidateWithPath(path string) error {
	if err := m.CT_Document.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
