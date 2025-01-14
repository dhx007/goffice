package vml

import (
	"encoding/xml"
	"strconv"

	"github.com/dhx007/goffice"
)

type Fill struct {
	CT_Fill
}

func NewFill() *Fill {
	ret := &Fill{}
	ret.CT_Fill = *NewCT_Fill()
	return ret
}

func (m *Fill) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return m.CT_Fill.MarshalXML(e, start)
}

func (m *Fill) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_Fill = *NewCT_Fill()
	for _, attr := range start.Attr {
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "detectmouseclick" {
			m.DetectmouseclickAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "relid" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RelidAttr = &parsed
			continue
		}
		if attr.Name.Space == "http://schemas.openxmlformats.org/officeDocument/2006/relationships" && attr.Name.Local == "id" ||
			attr.Name.Space == "http://purl.oclc.org/ooxml/officeDocument/relationships" && attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdAttr = &parsed
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "opacity2" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.Opacity2Attr = &parsed
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "title" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.TitleAttr = &parsed
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "href" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.HrefAttr = &parsed
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "althref" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.AlthrefAttr = &parsed
			continue
		}
		if attr.Name.Local == "alignshape" {
			m.AlignshapeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "method" {
			m.MethodAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "size" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.SizeAttr = &parsed
			continue
		}
		if attr.Name.Local == "position" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.PositionAttr = &parsed
			continue
		}
		if attr.Name.Local == "src" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.SrcAttr = &parsed
			continue
		}
		if attr.Name.Local == "colors" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ColorsAttr = &parsed
			continue
		}
		if attr.Name.Local == "color2" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.Color2Attr = &parsed
			continue
		}
		if attr.Name.Local == "type" {
			m.TypeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "focussize" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.FocussizeAttr = &parsed
			continue
		}
		if attr.Name.Local == "focusposition" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.FocuspositionAttr = &parsed
			continue
		}
		if attr.Name.Local == "origin" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.OriginAttr = &parsed
			continue
		}
		if attr.Name.Local == "aspect" {
			m.AspectAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "angle" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			m.AngleAttr = &parsed
			continue
		}
		if attr.Name.Local == "focus" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.FocusAttr = &parsed
			continue
		}
		if attr.Name.Local == "recolor" {
			m.RecolorAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "rotate" {
			m.RotateAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "opacity" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.OpacityAttr = &parsed
			continue
		}
		if attr.Name.Local == "color" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ColorAttr = &parsed
			continue
		}
		if attr.Name.Local == "on" {
			m.OnAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.SIdAttr = &parsed
			continue
		}
	}
lFill:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "urn:schemas-microsoft-com:office:office", Local: "fill"}:
				m.Fill = NewOfcFill()
				if err := d.DecodeElement(m.Fill, &el); err != nil {
					return err
				}
			default:
				goffice.Log("skipping unsupported element on Fill %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lFill
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the Fill and its children
func (m *Fill) Validate() error {
	return m.ValidateWithPath("Fill")
}

// ValidateWithPath validates the Fill and its children, prefixing error messages with path
func (m *Fill) ValidateWithPath(path string) error {
	if err := m.CT_Fill.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
