package vml

import (
	"encoding/xml"
	"fmt"

	"github.com/dhx007/goffice/schema/soo/ofc/sharedTypes"
)

type AG_ImageAttributes struct {
	SrcAttr        *string
	CropleftAttr   *string
	CroptopAttr    *string
	CroprightAttr  *string
	CropbottomAttr *string
	GainAttr       *string
	BlacklevelAttr *string
	GammaAttr      *string
	GrayscaleAttr  sharedTypes.ST_TrueFalse
	BilevelAttr    sharedTypes.ST_TrueFalse
}

func NewAG_ImageAttributes() *AG_ImageAttributes {
	ret := &AG_ImageAttributes{}
	return ret
}

func (m *AG_ImageAttributes) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.SrcAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "src"},
			Value: fmt.Sprintf("%v", *m.SrcAttr)})
	}
	if m.CropleftAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "cropleft"},
			Value: fmt.Sprintf("%v", *m.CropleftAttr)})
	}
	if m.CroptopAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "croptop"},
			Value: fmt.Sprintf("%v", *m.CroptopAttr)})
	}
	if m.CroprightAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "cropright"},
			Value: fmt.Sprintf("%v", *m.CroprightAttr)})
	}
	if m.CropbottomAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "cropbottom"},
			Value: fmt.Sprintf("%v", *m.CropbottomAttr)})
	}
	if m.GainAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "gain"},
			Value: fmt.Sprintf("%v", *m.GainAttr)})
	}
	if m.BlacklevelAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "blacklevel"},
			Value: fmt.Sprintf("%v", *m.BlacklevelAttr)})
	}
	if m.GammaAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "gamma"},
			Value: fmt.Sprintf("%v", *m.GammaAttr)})
	}
	if m.GrayscaleAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.GrayscaleAttr.MarshalXMLAttr(xml.Name{Local: "grayscale"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.BilevelAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.BilevelAttr.MarshalXMLAttr(xml.Name{Local: "bilevel"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	return nil
}

func (m *AG_ImageAttributes) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "src" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.SrcAttr = &parsed
			continue
		}
		if attr.Name.Local == "cropleft" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CropleftAttr = &parsed
			continue
		}
		if attr.Name.Local == "croptop" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CroptopAttr = &parsed
			continue
		}
		if attr.Name.Local == "cropright" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CroprightAttr = &parsed
			continue
		}
		if attr.Name.Local == "cropbottom" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CropbottomAttr = &parsed
			continue
		}
		if attr.Name.Local == "gain" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.GainAttr = &parsed
			continue
		}
		if attr.Name.Local == "blacklevel" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.BlacklevelAttr = &parsed
			continue
		}
		if attr.Name.Local == "gamma" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.GammaAttr = &parsed
			continue
		}
		if attr.Name.Local == "grayscale" {
			m.GrayscaleAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "bilevel" {
			m.BilevelAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing AG_ImageAttributes: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the AG_ImageAttributes and its children
func (m *AG_ImageAttributes) Validate() error {
	return m.ValidateWithPath("AG_ImageAttributes")
}

// ValidateWithPath validates the AG_ImageAttributes and its children, prefixing error messages with path
func (m *AG_ImageAttributes) ValidateWithPath(path string) error {
	if err := m.GrayscaleAttr.ValidateWithPath(path + "/GrayscaleAttr"); err != nil {
		return err
	}
	if err := m.BilevelAttr.ValidateWithPath(path + "/BilevelAttr"); err != nil {
		return err
	}
	return nil
}
