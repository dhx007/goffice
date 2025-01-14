package picture

import (
	"encoding/xml"

	"github.com/dhx007/goffice"
)

type Pic struct {
	CT_Picture
}

func NewPic() *Pic {
	ret := &Pic{}
	ret.CT_Picture = *NewCT_Picture()
	return ret
}

func (m *Pic) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/picture"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:pic"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/picture"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "pic:pic"
	return m.CT_Picture.MarshalXML(e, start)
}

func (m *Pic) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_Picture = *NewCT_Picture()
lPic:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/picture", Local: "nvPicPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/picture", Local: "nvPicPr"}:
				if err := d.DecodeElement(m.NvPicPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/picture", Local: "blipFill"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/picture", Local: "blipFill"}:
				if err := d.DecodeElement(m.BlipFill, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/picture", Local: "spPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/picture", Local: "spPr"}:
				if err := d.DecodeElement(m.SpPr, &el); err != nil {
					return err
				}
			default:
				goffice.Log("skipping unsupported element on Pic %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lPic
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the Pic and its children
func (m *Pic) Validate() error {
	return m.ValidateWithPath("Pic")
}

// ValidateWithPath validates the Pic and its children, prefixing error messages with path
func (m *Pic) ValidateWithPath(path string) error {
	if err := m.CT_Picture.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
