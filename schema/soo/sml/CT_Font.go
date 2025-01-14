package sml

import (
	"encoding/xml"
	"fmt"

	"github.com/dhx007/goffice"
)

type CT_Font struct {
	// Font Name
	Name []*CT_FontName
	// Character Set
	Charset []*CT_IntProperty
	// Font Family
	Family []*CT_FontFamily
	// Bold
	B []*CT_BooleanProperty
	// Italic
	I []*CT_BooleanProperty
	// Strike Through
	Strike []*CT_BooleanProperty
	// Outline
	Outline []*CT_BooleanProperty
	// Shadow
	Shadow []*CT_BooleanProperty
	// Condense
	Condense []*CT_BooleanProperty
	// Extend
	Extend []*CT_BooleanProperty
	// Text Color
	Color []*CT_Color
	// Font Size
	Sz []*CT_FontSize
	// Underline
	U []*CT_UnderlineProperty
	// Text Vertical Alignment
	VertAlign []*CT_VerticalAlignFontProperty
	// Scheme
	Scheme []*CT_FontScheme
}

func NewCT_Font() *CT_Font {
	ret := &CT_Font{}
	return ret
}

func (m *CT_Font) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Name != nil {
		sename := xml.StartElement{Name: xml.Name{Local: "ma:name"}}
		for _, c := range m.Name {
			e.EncodeElement(c, sename)
		}
	}
	if m.Charset != nil {
		secharset := xml.StartElement{Name: xml.Name{Local: "ma:charset"}}
		for _, c := range m.Charset {
			e.EncodeElement(c, secharset)
		}
	}
	if m.Family != nil {
		sefamily := xml.StartElement{Name: xml.Name{Local: "ma:family"}}
		for _, c := range m.Family {
			e.EncodeElement(c, sefamily)
		}
	}
	if m.B != nil {
		seb := xml.StartElement{Name: xml.Name{Local: "ma:b"}}
		for _, c := range m.B {
			e.EncodeElement(c, seb)
		}
	}
	if m.I != nil {
		sei := xml.StartElement{Name: xml.Name{Local: "ma:i"}}
		for _, c := range m.I {
			e.EncodeElement(c, sei)
		}
	}
	if m.Strike != nil {
		sestrike := xml.StartElement{Name: xml.Name{Local: "ma:strike"}}
		for _, c := range m.Strike {
			e.EncodeElement(c, sestrike)
		}
	}
	if m.Outline != nil {
		seoutline := xml.StartElement{Name: xml.Name{Local: "ma:outline"}}
		for _, c := range m.Outline {
			e.EncodeElement(c, seoutline)
		}
	}
	if m.Shadow != nil {
		seshadow := xml.StartElement{Name: xml.Name{Local: "ma:shadow"}}
		for _, c := range m.Shadow {
			e.EncodeElement(c, seshadow)
		}
	}
	if m.Condense != nil {
		secondense := xml.StartElement{Name: xml.Name{Local: "ma:condense"}}
		for _, c := range m.Condense {
			e.EncodeElement(c, secondense)
		}
	}
	if m.Extend != nil {
		seextend := xml.StartElement{Name: xml.Name{Local: "ma:extend"}}
		for _, c := range m.Extend {
			e.EncodeElement(c, seextend)
		}
	}
	if m.Color != nil {
		secolor := xml.StartElement{Name: xml.Name{Local: "ma:color"}}
		for _, c := range m.Color {
			e.EncodeElement(c, secolor)
		}
	}
	if m.Sz != nil {
		sesz := xml.StartElement{Name: xml.Name{Local: "ma:sz"}}
		for _, c := range m.Sz {
			e.EncodeElement(c, sesz)
		}
	}
	if m.U != nil {
		seu := xml.StartElement{Name: xml.Name{Local: "ma:u"}}
		for _, c := range m.U {
			e.EncodeElement(c, seu)
		}
	}
	if m.VertAlign != nil {
		severtAlign := xml.StartElement{Name: xml.Name{Local: "ma:vertAlign"}}
		for _, c := range m.VertAlign {
			e.EncodeElement(c, severtAlign)
		}
	}
	if m.Scheme != nil {
		sescheme := xml.StartElement{Name: xml.Name{Local: "ma:scheme"}}
		for _, c := range m.Scheme {
			e.EncodeElement(c, sescheme)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Font) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Font:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "name"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "name"}:
				tmp := NewCT_FontName()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Name = append(m.Name, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "charset"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "charset"}:
				tmp := NewCT_IntProperty()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Charset = append(m.Charset, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "family"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "family"}:
				tmp := NewCT_FontFamily()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Family = append(m.Family, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "b"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "b"}:
				tmp := NewCT_BooleanProperty()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.B = append(m.B, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "i"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "i"}:
				tmp := NewCT_BooleanProperty()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.I = append(m.I, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "strike"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "strike"}:
				tmp := NewCT_BooleanProperty()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Strike = append(m.Strike, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "outline"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "outline"}:
				tmp := NewCT_BooleanProperty()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Outline = append(m.Outline, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "shadow"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "shadow"}:
				tmp := NewCT_BooleanProperty()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Shadow = append(m.Shadow, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "condense"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "condense"}:
				tmp := NewCT_BooleanProperty()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Condense = append(m.Condense, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "extend"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "extend"}:
				tmp := NewCT_BooleanProperty()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Extend = append(m.Extend, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "color"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "color"}:
				tmp := NewCT_Color()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Color = append(m.Color, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "sz"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "sz"}:
				tmp := NewCT_FontSize()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Sz = append(m.Sz, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "u"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "u"}:
				tmp := NewCT_UnderlineProperty()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.U = append(m.U, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "vertAlign"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "vertAlign"}:
				tmp := NewCT_VerticalAlignFontProperty()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.VertAlign = append(m.VertAlign, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "scheme"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "scheme"}:
				tmp := NewCT_FontScheme()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Scheme = append(m.Scheme, tmp)
			default:
				goffice.Log("skipping unsupported element on CT_Font %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Font
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Font and its children
func (m *CT_Font) Validate() error {
	return m.ValidateWithPath("CT_Font")
}

// ValidateWithPath validates the CT_Font and its children, prefixing error messages with path
func (m *CT_Font) ValidateWithPath(path string) error {
	for i, v := range m.Name {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Name[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Charset {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Charset[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Family {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Family[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.B {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/B[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.I {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/I[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Strike {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Strike[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Outline {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Outline[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Shadow {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Shadow[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Condense {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Condense[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Extend {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Extend[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Color {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Color[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Sz {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Sz[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.U {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/U[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.VertAlign {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/VertAlign[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Scheme {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Scheme[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
