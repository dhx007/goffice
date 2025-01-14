package pml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/dhx007/goffice"
	"github.com/dhx007/goffice/schema/soo/dml"
)

type CT_CommonViewProperties struct {
	// Variable Scale
	VarScaleAttr *bool
	// View Scale
	Scale *dml.CT_Scale2D
	// View Origin
	Origin *dml.CT_Point2D
}

func NewCT_CommonViewProperties() *CT_CommonViewProperties {
	ret := &CT_CommonViewProperties{}
	ret.Scale = dml.NewCT_Scale2D()
	ret.Origin = dml.NewCT_Point2D()
	return ret
}

func (m *CT_CommonViewProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.VarScaleAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "varScale"},
			Value: fmt.Sprintf("%d", b2i(*m.VarScaleAttr))})
	}
	e.EncodeToken(start)
	sescale := xml.StartElement{Name: xml.Name{Local: "p:scale"}}
	e.EncodeElement(m.Scale, sescale)
	seorigin := xml.StartElement{Name: xml.Name{Local: "p:origin"}}
	e.EncodeElement(m.Origin, seorigin)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_CommonViewProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Scale = dml.NewCT_Scale2D()
	m.Origin = dml.NewCT_Point2D()
	for _, attr := range start.Attr {
		if attr.Name.Local == "varScale" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.VarScaleAttr = &parsed
			continue
		}
	}
lCT_CommonViewProperties:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "scale"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "scale"}:
				if err := d.DecodeElement(m.Scale, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "origin"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "origin"}:
				if err := d.DecodeElement(m.Origin, &el); err != nil {
					return err
				}
			default:
				goffice.Log("skipping unsupported element on CT_CommonViewProperties %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_CommonViewProperties
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_CommonViewProperties and its children
func (m *CT_CommonViewProperties) Validate() error {
	return m.ValidateWithPath("CT_CommonViewProperties")
}

// ValidateWithPath validates the CT_CommonViewProperties and its children, prefixing error messages with path
func (m *CT_CommonViewProperties) ValidateWithPath(path string) error {
	if err := m.Scale.ValidateWithPath(path + "/Scale"); err != nil {
		return err
	}
	if err := m.Origin.ValidateWithPath(path + "/Origin"); err != nil {
		return err
	}
	return nil
}
