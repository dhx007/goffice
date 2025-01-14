package wml

import (
	"encoding/xml"

	"github.com/dhx007/goffice"
)

type CT_SdtDocPart struct {
	// Document Part Gallery Filter
	DocPartGallery *CT_String
	// Document Part Category Filter
	DocPartCategory *CT_String
	// Built-In Document Part
	DocPartUnique *CT_OnOff
}

func NewCT_SdtDocPart() *CT_SdtDocPart {
	ret := &CT_SdtDocPart{}
	return ret
}

func (m *CT_SdtDocPart) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.DocPartGallery != nil {
		sedocPartGallery := xml.StartElement{Name: xml.Name{Local: "w:docPartGallery"}}
		e.EncodeElement(m.DocPartGallery, sedocPartGallery)
	}
	if m.DocPartCategory != nil {
		sedocPartCategory := xml.StartElement{Name: xml.Name{Local: "w:docPartCategory"}}
		e.EncodeElement(m.DocPartCategory, sedocPartCategory)
	}
	if m.DocPartUnique != nil {
		sedocPartUnique := xml.StartElement{Name: xml.Name{Local: "w:docPartUnique"}}
		e.EncodeElement(m.DocPartUnique, sedocPartUnique)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_SdtDocPart) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_SdtDocPart:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "docPartGallery"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "docPartGallery"}:
				m.DocPartGallery = NewCT_String()
				if err := d.DecodeElement(m.DocPartGallery, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "docPartCategory"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "docPartCategory"}:
				m.DocPartCategory = NewCT_String()
				if err := d.DecodeElement(m.DocPartCategory, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "docPartUnique"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "docPartUnique"}:
				m.DocPartUnique = NewCT_OnOff()
				if err := d.DecodeElement(m.DocPartUnique, &el); err != nil {
					return err
				}
			default:
				goffice.Log("skipping unsupported element on CT_SdtDocPart %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_SdtDocPart
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_SdtDocPart and its children
func (m *CT_SdtDocPart) Validate() error {
	return m.ValidateWithPath("CT_SdtDocPart")
}

// ValidateWithPath validates the CT_SdtDocPart and its children, prefixing error messages with path
func (m *CT_SdtDocPart) ValidateWithPath(path string) error {
	if m.DocPartGallery != nil {
		if err := m.DocPartGallery.ValidateWithPath(path + "/DocPartGallery"); err != nil {
			return err
		}
	}
	if m.DocPartCategory != nil {
		if err := m.DocPartCategory.ValidateWithPath(path + "/DocPartCategory"); err != nil {
			return err
		}
	}
	if m.DocPartUnique != nil {
		if err := m.DocPartUnique.ValidateWithPath(path + "/DocPartUnique"); err != nil {
			return err
		}
	}
	return nil
}
