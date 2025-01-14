package wml

import (
	"encoding/xml"

	"github.com/dhx007/goffice"
)

type CT_DocPartCategory struct {
	// Category Associated With Entry
	Name *CT_String
	// Gallery Associated With Entry
	Gallery *CT_DocPartGallery
}

func NewCT_DocPartCategory() *CT_DocPartCategory {
	ret := &CT_DocPartCategory{}
	ret.Name = NewCT_String()
	ret.Gallery = NewCT_DocPartGallery()
	return ret
}

func (m *CT_DocPartCategory) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	sename := xml.StartElement{Name: xml.Name{Local: "w:name"}}
	e.EncodeElement(m.Name, sename)
	segallery := xml.StartElement{Name: xml.Name{Local: "w:gallery"}}
	e.EncodeElement(m.Gallery, segallery)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_DocPartCategory) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Name = NewCT_String()
	m.Gallery = NewCT_DocPartGallery()
lCT_DocPartCategory:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "name"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "name"}:
				if err := d.DecodeElement(m.Name, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "gallery"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "gallery"}:
				if err := d.DecodeElement(m.Gallery, &el); err != nil {
					return err
				}
			default:
				goffice.Log("skipping unsupported element on CT_DocPartCategory %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_DocPartCategory
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_DocPartCategory and its children
func (m *CT_DocPartCategory) Validate() error {
	return m.ValidateWithPath("CT_DocPartCategory")
}

// ValidateWithPath validates the CT_DocPartCategory and its children, prefixing error messages with path
func (m *CT_DocPartCategory) ValidateWithPath(path string) error {
	if err := m.Name.ValidateWithPath(path + "/Name"); err != nil {
		return err
	}
	if err := m.Gallery.ValidateWithPath(path + "/Gallery"); err != nil {
		return err
	}
	return nil
}
