package math

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/dhx007/goffice"
	"github.com/dhx007/goffice/schema/soo/ofc/sharedTypes"
)

func ParseUnionST_OnOff(s string) (sharedTypes.ST_OnOff, error) {
	return sharedTypes.ParseUnionST_OnOff(s)
}
func ParseUnionST_TwipsMeasure(s string) (sharedTypes.ST_TwipsMeasure, error) {
	ret := sharedTypes.ST_TwipsMeasure{}
	if sharedTypes.ST_PositiveUniversalMeasurePatternRe.MatchString(s) {
		ret.ST_PositiveUniversalMeasure = &s
	} else {
		v, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return ret, fmt.Errorf("parsing %s as uint: %s", s, err)
		}
		ret.ST_UnsignedDecimalNumber = goffice.Uint64(uint64(v))
	}
	return ret, nil
}

type ST_Shp byte

const (
	ST_ShpUnset    ST_Shp = 0
	ST_ShpCentered ST_Shp = 1
	ST_ShpMatch    ST_Shp = 2
)

func (e ST_Shp) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_ShpUnset:
		attr.Value = ""
	case ST_ShpCentered:
		attr.Value = "centered"
	case ST_ShpMatch:
		attr.Value = "match"
	}
	return attr, nil
}

func (e *ST_Shp) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "centered":
		*e = 1
	case "match":
		*e = 2
	}
	return nil
}

func (m ST_Shp) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_Shp) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		*m = 1
		return nil
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "centered":
			*m = 1
		case "match":
			*m = 2
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_Shp) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "centered"
	case 2:
		return "match"
	}
	return ""
}

func (m ST_Shp) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_Shp) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_FType byte

const (
	ST_FTypeUnset ST_FType = 0
	ST_FTypeBar   ST_FType = 1
	ST_FTypeSkw   ST_FType = 2
	ST_FTypeLin   ST_FType = 3
	ST_FTypeNoBar ST_FType = 4
)

func (e ST_FType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_FTypeUnset:
		attr.Value = ""
	case ST_FTypeBar:
		attr.Value = "bar"
	case ST_FTypeSkw:
		attr.Value = "skw"
	case ST_FTypeLin:
		attr.Value = "lin"
	case ST_FTypeNoBar:
		attr.Value = "noBar"
	}
	return attr, nil
}

func (e *ST_FType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "bar":
		*e = 1
	case "skw":
		*e = 2
	case "lin":
		*e = 3
	case "noBar":
		*e = 4
	}
	return nil
}

func (m ST_FType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_FType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		*m = 1
		return nil
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "bar":
			*m = 1
		case "skw":
			*m = 2
		case "lin":
			*m = 3
		case "noBar":
			*m = 4
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_FType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "bar"
	case 2:
		return "skw"
	case 3:
		return "lin"
	case 4:
		return "noBar"
	}
	return ""
}

func (m ST_FType) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_FType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_LimLoc byte

const (
	ST_LimLocUnset  ST_LimLoc = 0
	ST_LimLocUndOvr ST_LimLoc = 1
	ST_LimLocSubSup ST_LimLoc = 2
)

func (e ST_LimLoc) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_LimLocUnset:
		attr.Value = ""
	case ST_LimLocUndOvr:
		attr.Value = "undOvr"
	case ST_LimLocSubSup:
		attr.Value = "subSup"
	}
	return attr, nil
}

func (e *ST_LimLoc) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "undOvr":
		*e = 1
	case "subSup":
		*e = 2
	}
	return nil
}

func (m ST_LimLoc) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_LimLoc) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		*m = 1
		return nil
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "undOvr":
			*m = 1
		case "subSup":
			*m = 2
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_LimLoc) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "undOvr"
	case 2:
		return "subSup"
	}
	return ""
}

func (m ST_LimLoc) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_LimLoc) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TopBot byte

const (
	ST_TopBotUnset ST_TopBot = 0
	ST_TopBotTop   ST_TopBot = 1
	ST_TopBotBot   ST_TopBot = 2
)

func (e ST_TopBot) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TopBotUnset:
		attr.Value = ""
	case ST_TopBotTop:
		attr.Value = "top"
	case ST_TopBotBot:
		attr.Value = "bot"
	}
	return attr, nil
}

func (e *ST_TopBot) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "top":
		*e = 1
	case "bot":
		*e = 2
	}
	return nil
}

func (m ST_TopBot) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_TopBot) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		*m = 1
		return nil
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "top":
			*m = 1
		case "bot":
			*m = 2
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_TopBot) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "top"
	case 2:
		return "bot"
	}
	return ""
}

func (m ST_TopBot) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_TopBot) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_Script byte

const (
	ST_ScriptUnset         ST_Script = 0
	ST_ScriptRoman         ST_Script = 1
	ST_ScriptScript        ST_Script = 2
	ST_ScriptFraktur       ST_Script = 3
	ST_ScriptDouble_struck ST_Script = 4
	ST_ScriptSans_serif    ST_Script = 5
	ST_ScriptMonospace     ST_Script = 6
)

func (e ST_Script) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_ScriptUnset:
		attr.Value = ""
	case ST_ScriptRoman:
		attr.Value = "roman"
	case ST_ScriptScript:
		attr.Value = "script"
	case ST_ScriptFraktur:
		attr.Value = "fraktur"
	case ST_ScriptDouble_struck:
		attr.Value = "double-struck"
	case ST_ScriptSans_serif:
		attr.Value = "sans-serif"
	case ST_ScriptMonospace:
		attr.Value = "monospace"
	}
	return attr, nil
}

func (e *ST_Script) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "roman":
		*e = 1
	case "script":
		*e = 2
	case "fraktur":
		*e = 3
	case "double-struck":
		*e = 4
	case "sans-serif":
		*e = 5
	case "monospace":
		*e = 6
	}
	return nil
}

func (m ST_Script) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_Script) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		*m = 1
		return nil
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "roman":
			*m = 1
		case "script":
			*m = 2
		case "fraktur":
			*m = 3
		case "double-struck":
			*m = 4
		case "sans-serif":
			*m = 5
		case "monospace":
			*m = 6
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_Script) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "roman"
	case 2:
		return "script"
	case 3:
		return "fraktur"
	case 4:
		return "double-struck"
	case 5:
		return "sans-serif"
	case 6:
		return "monospace"
	}
	return ""
}

func (m ST_Script) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_Script) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_Style byte

const (
	ST_StyleUnset ST_Style = 0
	ST_StyleP     ST_Style = 1
	ST_StyleB     ST_Style = 2
	ST_StyleI     ST_Style = 3
	ST_StyleBi    ST_Style = 4
)

func (e ST_Style) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_StyleUnset:
		attr.Value = ""
	case ST_StyleP:
		attr.Value = "p"
	case ST_StyleB:
		attr.Value = "b"
	case ST_StyleI:
		attr.Value = "i"
	case ST_StyleBi:
		attr.Value = "bi"
	}
	return attr, nil
}

func (e *ST_Style) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "p":
		*e = 1
	case "b":
		*e = 2
	case "i":
		*e = 3
	case "bi":
		*e = 4
	}
	return nil
}

func (m ST_Style) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_Style) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		*m = 1
		return nil
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "p":
			*m = 1
		case "b":
			*m = 2
		case "i":
			*m = 3
		case "bi":
			*m = 4
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_Style) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "p"
	case 2:
		return "b"
	case 3:
		return "i"
	case 4:
		return "bi"
	}
	return ""
}

func (m ST_Style) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_Style) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_Jc byte

const (
	ST_JcUnset       ST_Jc = 0
	ST_JcLeft        ST_Jc = 1
	ST_JcRight       ST_Jc = 2
	ST_JcCenter      ST_Jc = 3
	ST_JcCenterGroup ST_Jc = 4
)

func (e ST_Jc) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_JcUnset:
		attr.Value = ""
	case ST_JcLeft:
		attr.Value = "left"
	case ST_JcRight:
		attr.Value = "right"
	case ST_JcCenter:
		attr.Value = "center"
	case ST_JcCenterGroup:
		attr.Value = "centerGroup"
	}
	return attr, nil
}

func (e *ST_Jc) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "left":
		*e = 1
	case "right":
		*e = 2
	case "center":
		*e = 3
	case "centerGroup":
		*e = 4
	}
	return nil
}

func (m ST_Jc) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_Jc) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		*m = 1
		return nil
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "left":
			*m = 1
		case "right":
			*m = 2
		case "center":
			*m = 3
		case "centerGroup":
			*m = 4
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_Jc) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "left"
	case 2:
		return "right"
	case 3:
		return "center"
	case 4:
		return "centerGroup"
	}
	return ""
}

func (m ST_Jc) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_Jc) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_BreakBin byte

const (
	ST_BreakBinUnset  ST_BreakBin = 0
	ST_BreakBinBefore ST_BreakBin = 1
	ST_BreakBinAfter  ST_BreakBin = 2
	ST_BreakBinRepeat ST_BreakBin = 3
)

func (e ST_BreakBin) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_BreakBinUnset:
		attr.Value = ""
	case ST_BreakBinBefore:
		attr.Value = "before"
	case ST_BreakBinAfter:
		attr.Value = "after"
	case ST_BreakBinRepeat:
		attr.Value = "repeat"
	}
	return attr, nil
}

func (e *ST_BreakBin) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "before":
		*e = 1
	case "after":
		*e = 2
	case "repeat":
		*e = 3
	}
	return nil
}

func (m ST_BreakBin) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_BreakBin) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		*m = 1
		return nil
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "before":
			*m = 1
		case "after":
			*m = 2
		case "repeat":
			*m = 3
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_BreakBin) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "before"
	case 2:
		return "after"
	case 3:
		return "repeat"
	}
	return ""
}

func (m ST_BreakBin) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_BreakBin) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_BreakBinSub byte

const (
	ST_BreakBinSubUnset ST_BreakBinSub = 0
	ST_BreakBinSub__    ST_BreakBinSub = 1
	ST_BreakBinSub___   ST_BreakBinSub = 2
	ST_BreakBinSub____  ST_BreakBinSub = 3
)

func (e ST_BreakBinSub) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_BreakBinSubUnset:
		attr.Value = ""
	case ST_BreakBinSub__:
		attr.Value = "--"
	case ST_BreakBinSub___:
		attr.Value = "-+"
	case ST_BreakBinSub____:
		attr.Value = "+-"
	}
	return attr, nil
}

func (e *ST_BreakBinSub) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "--":
		*e = 1
	case "-+":
		*e = 2
	case "+-":
		*e = 3
	}
	return nil
}

func (m ST_BreakBinSub) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_BreakBinSub) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		*m = 1
		return nil
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "--":
			*m = 1
		case "-+":
			*m = 2
		case "+-":
			*m = 3
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_BreakBinSub) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "--"
	case 2:
		return "-+"
	case 3:
		return "+-"
	}
	return ""
}

func (m ST_BreakBinSub) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_BreakBinSub) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

// init registers constructor functions for dynamically creating elements based off the XML namespace and name
func init() {
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_Integer255", NewCT_Integer255)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_Integer2", NewCT_Integer2)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_SpacingRule", NewCT_SpacingRule)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_UnSignedInteger", NewCT_UnSignedInteger)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_Char", NewCT_Char)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_OnOff", NewCT_OnOff)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_String", NewCT_String)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_XAlign", NewCT_XAlign)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_YAlign", NewCT_YAlign)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_Shp", NewCT_Shp)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_FType", NewCT_FType)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_LimLoc", NewCT_LimLoc)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_TopBot", NewCT_TopBot)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_Script", NewCT_Script)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_Style", NewCT_Style)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_ManualBreak", NewCT_ManualBreak)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_RPR", NewCT_RPR)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_Text", NewCT_Text)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_R", NewCT_R)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_CtrlPr", NewCT_CtrlPr)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_AccPr", NewCT_AccPr)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_Acc", NewCT_Acc)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_BarPr", NewCT_BarPr)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_Bar", NewCT_Bar)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_BoxPr", NewCT_BoxPr)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_Box", NewCT_Box)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_BorderBoxPr", NewCT_BorderBoxPr)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_BorderBox", NewCT_BorderBox)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_DPr", NewCT_DPr)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_D", NewCT_D)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_EqArrPr", NewCT_EqArrPr)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_EqArr", NewCT_EqArr)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_FPr", NewCT_FPr)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_F", NewCT_F)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_FuncPr", NewCT_FuncPr)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_Func", NewCT_Func)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_GroupChrPr", NewCT_GroupChrPr)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_GroupChr", NewCT_GroupChr)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_LimLowPr", NewCT_LimLowPr)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_LimLow", NewCT_LimLow)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_LimUppPr", NewCT_LimUppPr)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_LimUpp", NewCT_LimUpp)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_MCPr", NewCT_MCPr)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_MC", NewCT_MC)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_MCS", NewCT_MCS)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_MPr", NewCT_MPr)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_MR", NewCT_MR)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_M", NewCT_M)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_NaryPr", NewCT_NaryPr)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_Nary", NewCT_Nary)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_PhantPr", NewCT_PhantPr)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_Phant", NewCT_Phant)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_RadPr", NewCT_RadPr)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_Rad", NewCT_Rad)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_SPrePr", NewCT_SPrePr)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_SPre", NewCT_SPre)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_SSubPr", NewCT_SSubPr)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_SSub", NewCT_SSub)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_SSubSupPr", NewCT_SSubSupPr)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_SSubSup", NewCT_SSubSup)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_SSupPr", NewCT_SSupPr)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_SSup", NewCT_SSup)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_OMathArgPr", NewCT_OMathArgPr)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_OMathArg", NewCT_OMathArg)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_OMathJc", NewCT_OMathJc)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_OMathParaPr", NewCT_OMathParaPr)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_TwipsMeasure", NewCT_TwipsMeasure)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_BreakBin", NewCT_BreakBin)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_BreakBinSub", NewCT_BreakBinSub)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_MathPr", NewCT_MathPr)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_OMathPara", NewCT_OMathPara)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "CT_OMath", NewCT_OMath)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "mathPr", NewMathPr)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "oMathPara", NewOMathPara)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "oMath", NewOMath)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "EG_ScriptStyle", NewEG_ScriptStyle)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "EG_OMathMathElements", NewEG_OMathMathElements)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/math", "EG_OMathElements", NewEG_OMathElements)
}
