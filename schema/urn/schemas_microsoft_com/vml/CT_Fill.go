package vml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/dhx007/goffice"
	"github.com/dhx007/goffice/schema/soo/ofc/sharedTypes"
)

type CT_Fill struct {
	TypeAttr             ST_FillType
	OnAttr               sharedTypes.ST_TrueFalse
	ColorAttr            *string
	OpacityAttr          *string
	Color2Attr           *string
	SrcAttr              *string
	HrefAttr             *string
	AlthrefAttr          *string
	SizeAttr             *string
	OriginAttr           *string
	PositionAttr         *string
	AspectAttr           ST_ImageAspect
	ColorsAttr           *string
	AngleAttr            *float64
	AlignshapeAttr       sharedTypes.ST_TrueFalse
	FocusAttr            *string
	FocussizeAttr        *string
	FocuspositionAttr    *string
	MethodAttr           ST_FillMethod
	DetectmouseclickAttr sharedTypes.ST_TrueFalse
	TitleAttr            *string
	Opacity2Attr         *string
	RecolorAttr          sharedTypes.ST_TrueFalse
	RotateAttr           sharedTypes.ST_TrueFalse
	IdAttr               *string
	RelidAttr            *string
	Fill                 *OfcFill
	SIdAttr              *string
}

func NewCT_Fill() *CT_Fill {
	ret := &CT_Fill{}
	return ret
}

func (m *CT_Fill) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.TypeAttr != ST_FillTypeUnset {
		attr, err := m.TypeAttr.MarshalXMLAttr(xml.Name{Local: "type"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.OnAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.OnAttr.MarshalXMLAttr(xml.Name{Local: "on"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.ColorAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "color"},
			Value: fmt.Sprintf("%v", *m.ColorAttr)})
	}
	if m.OpacityAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "opacity"},
			Value: fmt.Sprintf("%v", *m.OpacityAttr)})
	}
	if m.Color2Attr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "color2"},
			Value: fmt.Sprintf("%v", *m.Color2Attr)})
	}
	if m.SrcAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "src"},
			Value: fmt.Sprintf("%v", *m.SrcAttr)})
	}
	if m.HrefAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "o:href"},
			Value: fmt.Sprintf("%v", *m.HrefAttr)})
	}
	if m.AlthrefAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "o:althref"},
			Value: fmt.Sprintf("%v", *m.AlthrefAttr)})
	}
	if m.SizeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "size"},
			Value: fmt.Sprintf("%v", *m.SizeAttr)})
	}
	if m.OriginAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "origin"},
			Value: fmt.Sprintf("%v", *m.OriginAttr)})
	}
	if m.PositionAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "position"},
			Value: fmt.Sprintf("%v", *m.PositionAttr)})
	}
	if m.AspectAttr != ST_ImageAspectUnset {
		attr, err := m.AspectAttr.MarshalXMLAttr(xml.Name{Local: "aspect"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.ColorsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "colors"},
			Value: fmt.Sprintf("%v", *m.ColorsAttr)})
	}
	if m.AngleAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "angle"},
			Value: fmt.Sprintf("%v", *m.AngleAttr)})
	}
	if m.AlignshapeAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.AlignshapeAttr.MarshalXMLAttr(xml.Name{Local: "alignshape"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.FocusAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "focus"},
			Value: fmt.Sprintf("%v", *m.FocusAttr)})
	}
	if m.FocussizeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "focussize"},
			Value: fmt.Sprintf("%v", *m.FocussizeAttr)})
	}
	if m.FocuspositionAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "focusposition"},
			Value: fmt.Sprintf("%v", *m.FocuspositionAttr)})
	}
	if m.MethodAttr != ST_FillMethodUnset {
		attr, err := m.MethodAttr.MarshalXMLAttr(xml.Name{Local: "method"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.DetectmouseclickAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.DetectmouseclickAttr.MarshalXMLAttr(xml.Name{Local: "detectmouseclick"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.TitleAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "o:title"},
			Value: fmt.Sprintf("%v", *m.TitleAttr)})
	}
	if m.Opacity2Attr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "o:opacity2"},
			Value: fmt.Sprintf("%v", *m.Opacity2Attr)})
	}
	if m.RecolorAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.RecolorAttr.MarshalXMLAttr(xml.Name{Local: "recolor"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.RotateAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.RotateAttr.MarshalXMLAttr(xml.Name{Local: "rotate"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.IdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r:id"},
			Value: fmt.Sprintf("%v", *m.IdAttr)})
	}
	if m.RelidAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "o:relid"},
			Value: fmt.Sprintf("%v", *m.RelidAttr)})
	}
	if m.SIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "id"},
			Value: fmt.Sprintf("%v", *m.SIdAttr)})
	}
	e.EncodeToken(start)
	if m.Fill != nil {
		sefill := xml.StartElement{Name: xml.Name{Local: "o:fill"}}
		e.EncodeElement(m.Fill, sefill)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Fill) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
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
lCT_Fill:
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
				goffice.Log("skipping unsupported element on CT_Fill %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Fill
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Fill and its children
func (m *CT_Fill) Validate() error {
	return m.ValidateWithPath("CT_Fill")
}

// ValidateWithPath validates the CT_Fill and its children, prefixing error messages with path
func (m *CT_Fill) ValidateWithPath(path string) error {
	if err := m.TypeAttr.ValidateWithPath(path + "/TypeAttr"); err != nil {
		return err
	}
	if err := m.OnAttr.ValidateWithPath(path + "/OnAttr"); err != nil {
		return err
	}
	if err := m.AspectAttr.ValidateWithPath(path + "/AspectAttr"); err != nil {
		return err
	}
	if err := m.AlignshapeAttr.ValidateWithPath(path + "/AlignshapeAttr"); err != nil {
		return err
	}
	if err := m.MethodAttr.ValidateWithPath(path + "/MethodAttr"); err != nil {
		return err
	}
	if err := m.DetectmouseclickAttr.ValidateWithPath(path + "/DetectmouseclickAttr"); err != nil {
		return err
	}
	if err := m.RecolorAttr.ValidateWithPath(path + "/RecolorAttr"); err != nil {
		return err
	}
	if err := m.RotateAttr.ValidateWithPath(path + "/RotateAttr"); err != nil {
		return err
	}
	if m.Fill != nil {
		if err := m.Fill.ValidateWithPath(path + "/Fill"); err != nil {
			return err
		}
	}
	return nil
}
