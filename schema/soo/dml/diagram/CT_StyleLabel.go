package diagram

import (
	"encoding/xml"
	"fmt"

	"github.com/dhx007/goffice"
	"github.com/dhx007/goffice/schema/soo/dml"
)

type CT_StyleLabel struct {
	NameAttr string
	Scene3d  *dml.CT_Scene3D
	Sp3d     *dml.CT_Shape3D
	TxPr     *CT_TextProps
	Style    *dml.CT_ShapeStyle
	ExtLst   *dml.CT_OfficeArtExtensionList
}

func NewCT_StyleLabel() *CT_StyleLabel {
	ret := &CT_StyleLabel{}
	return ret
}

func (m *CT_StyleLabel) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "name"},
		Value: fmt.Sprintf("%v", m.NameAttr)})
	e.EncodeToken(start)
	if m.Scene3d != nil {
		sescene3d := xml.StartElement{Name: xml.Name{Local: "scene3d"}}
		e.EncodeElement(m.Scene3d, sescene3d)
	}
	if m.Sp3d != nil {
		sesp3d := xml.StartElement{Name: xml.Name{Local: "sp3d"}}
		e.EncodeElement(m.Sp3d, sesp3d)
	}
	if m.TxPr != nil {
		setxPr := xml.StartElement{Name: xml.Name{Local: "txPr"}}
		e.EncodeElement(m.TxPr, setxPr)
	}
	if m.Style != nil {
		sestyle := xml.StartElement{Name: xml.Name{Local: "style"}}
		e.EncodeElement(m.Style, sestyle)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_StyleLabel) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = parsed
			continue
		}
	}
lCT_StyleLabel:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "scene3d"}:
				m.Scene3d = dml.NewCT_Scene3D()
				if err := d.DecodeElement(m.Scene3d, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "sp3d"}:
				m.Sp3d = dml.NewCT_Shape3D()
				if err := d.DecodeElement(m.Sp3d, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "txPr"}:
				m.TxPr = NewCT_TextProps()
				if err := d.DecodeElement(m.TxPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "style"}:
				m.Style = dml.NewCT_ShapeStyle()
				if err := d.DecodeElement(m.Style, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "extLst"}:
				m.ExtLst = dml.NewCT_OfficeArtExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				goffice.Log("skipping unsupported element on CT_StyleLabel %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_StyleLabel
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_StyleLabel and its children
func (m *CT_StyleLabel) Validate() error {
	return m.ValidateWithPath("CT_StyleLabel")
}

// ValidateWithPath validates the CT_StyleLabel and its children, prefixing error messages with path
func (m *CT_StyleLabel) ValidateWithPath(path string) error {
	if m.Scene3d != nil {
		if err := m.Scene3d.ValidateWithPath(path + "/Scene3d"); err != nil {
			return err
		}
	}
	if m.Sp3d != nil {
		if err := m.Sp3d.ValidateWithPath(path + "/Sp3d"); err != nil {
			return err
		}
	}
	if m.TxPr != nil {
		if err := m.TxPr.ValidateWithPath(path + "/TxPr"); err != nil {
			return err
		}
	}
	if m.Style != nil {
		if err := m.Style.ValidateWithPath(path + "/Style"); err != nil {
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
