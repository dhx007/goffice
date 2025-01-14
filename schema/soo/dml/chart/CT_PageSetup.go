package chart

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/dhx007/goffice/schema/soo/ofc/sharedTypes"
)

type CT_PageSetup struct {
	PaperSizeAttr          *uint32
	PaperHeightAttr        *string
	PaperWidthAttr         *string
	FirstPageNumberAttr    *uint32
	OrientationAttr        ST_PageSetupOrientation
	BlackAndWhiteAttr      *bool
	DraftAttr              *bool
	UseFirstPageNumberAttr *bool
	HorizontalDpiAttr      *int32
	VerticalDpiAttr        *int32
	CopiesAttr             *uint32
}

func NewCT_PageSetup() *CT_PageSetup {
	ret := &CT_PageSetup{}
	return ret
}

func (m *CT_PageSetup) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.PaperSizeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "paperSize"},
			Value: fmt.Sprintf("%v", *m.PaperSizeAttr)})
	}
	if m.PaperHeightAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "paperHeight"},
			Value: fmt.Sprintf("%v", *m.PaperHeightAttr)})
	}
	if m.PaperWidthAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "paperWidth"},
			Value: fmt.Sprintf("%v", *m.PaperWidthAttr)})
	}
	if m.FirstPageNumberAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "firstPageNumber"},
			Value: fmt.Sprintf("%v", *m.FirstPageNumberAttr)})
	}
	if m.OrientationAttr != ST_PageSetupOrientationUnset {
		attr, err := m.OrientationAttr.MarshalXMLAttr(xml.Name{Local: "orientation"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.BlackAndWhiteAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "blackAndWhite"},
			Value: fmt.Sprintf("%d", b2i(*m.BlackAndWhiteAttr))})
	}
	if m.DraftAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "draft"},
			Value: fmt.Sprintf("%d", b2i(*m.DraftAttr))})
	}
	if m.UseFirstPageNumberAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "useFirstPageNumber"},
			Value: fmt.Sprintf("%d", b2i(*m.UseFirstPageNumberAttr))})
	}
	if m.HorizontalDpiAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "horizontalDpi"},
			Value: fmt.Sprintf("%v", *m.HorizontalDpiAttr)})
	}
	if m.VerticalDpiAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "verticalDpi"},
			Value: fmt.Sprintf("%v", *m.VerticalDpiAttr)})
	}
	if m.CopiesAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "copies"},
			Value: fmt.Sprintf("%v", *m.CopiesAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_PageSetup) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "paperSize" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.PaperSizeAttr = &pt
			continue
		}
		if attr.Name.Local == "paperHeight" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.PaperHeightAttr = &parsed
			continue
		}
		if attr.Name.Local == "paperWidth" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.PaperWidthAttr = &parsed
			continue
		}
		if attr.Name.Local == "firstPageNumber" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.FirstPageNumberAttr = &pt
			continue
		}
		if attr.Name.Local == "orientation" {
			m.OrientationAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "blackAndWhite" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.BlackAndWhiteAttr = &parsed
			continue
		}
		if attr.Name.Local == "draft" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DraftAttr = &parsed
			continue
		}
		if attr.Name.Local == "useFirstPageNumber" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.UseFirstPageNumberAttr = &parsed
			continue
		}
		if attr.Name.Local == "horizontalDpi" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.HorizontalDpiAttr = &pt
			continue
		}
		if attr.Name.Local == "verticalDpi" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.VerticalDpiAttr = &pt
			continue
		}
		if attr.Name.Local == "copies" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.CopiesAttr = &pt
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_PageSetup: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_PageSetup and its children
func (m *CT_PageSetup) Validate() error {
	return m.ValidateWithPath("CT_PageSetup")
}

// ValidateWithPath validates the CT_PageSetup and its children, prefixing error messages with path
func (m *CT_PageSetup) ValidateWithPath(path string) error {
	if m.PaperHeightAttr != nil {
		if !sharedTypes.ST_PositiveUniversalMeasurePatternRe.MatchString(*m.PaperHeightAttr) {
			return fmt.Errorf(`%s/m.PaperHeightAttr must match '%s' (have %v)`, path, sharedTypes.ST_PositiveUniversalMeasurePatternRe, *m.PaperHeightAttr)
		}
	}
	if m.PaperHeightAttr != nil {
		if !sharedTypes.ST_UniversalMeasurePatternRe.MatchString(*m.PaperHeightAttr) {
			return fmt.Errorf(`%s/m.PaperHeightAttr must match '%s' (have %v)`, path, sharedTypes.ST_UniversalMeasurePatternRe, *m.PaperHeightAttr)
		}
	}
	if m.PaperWidthAttr != nil {
		if !sharedTypes.ST_PositiveUniversalMeasurePatternRe.MatchString(*m.PaperWidthAttr) {
			return fmt.Errorf(`%s/m.PaperWidthAttr must match '%s' (have %v)`, path, sharedTypes.ST_PositiveUniversalMeasurePatternRe, *m.PaperWidthAttr)
		}
	}
	if m.PaperWidthAttr != nil {
		if !sharedTypes.ST_UniversalMeasurePatternRe.MatchString(*m.PaperWidthAttr) {
			return fmt.Errorf(`%s/m.PaperWidthAttr must match '%s' (have %v)`, path, sharedTypes.ST_UniversalMeasurePatternRe, *m.PaperWidthAttr)
		}
	}
	if err := m.OrientationAttr.ValidateWithPath(path + "/OrientationAttr"); err != nil {
		return err
	}
	return nil
}
