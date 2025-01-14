package word

import (
	"encoding/xml"
	"fmt"

	"github.com/dhx007/goffice"
)

type ST_BorderType byte

const (
	ST_BorderTypeUnset                 ST_BorderType = 0
	ST_BorderTypeNone                  ST_BorderType = 1
	ST_BorderTypeSingle                ST_BorderType = 2
	ST_BorderTypeThick                 ST_BorderType = 3
	ST_BorderTypeDouble                ST_BorderType = 4
	ST_BorderTypeHairline              ST_BorderType = 5
	ST_BorderTypeDot                   ST_BorderType = 6
	ST_BorderTypeDash                  ST_BorderType = 7
	ST_BorderTypeDotDash               ST_BorderType = 8
	ST_BorderTypeDashDotDot            ST_BorderType = 9
	ST_BorderTypeTriple                ST_BorderType = 10
	ST_BorderTypeThinThickSmall        ST_BorderType = 11
	ST_BorderTypeThickThinSmall        ST_BorderType = 12
	ST_BorderTypeThickBetweenThinSmall ST_BorderType = 13
	ST_BorderTypeThinThick             ST_BorderType = 14
	ST_BorderTypeThickThin             ST_BorderType = 15
	ST_BorderTypeThickBetweenThin      ST_BorderType = 16
	ST_BorderTypeThinThickLarge        ST_BorderType = 17
	ST_BorderTypeThickThinLarge        ST_BorderType = 18
	ST_BorderTypeThickBetweenThinLarge ST_BorderType = 19
	ST_BorderTypeWave                  ST_BorderType = 20
	ST_BorderTypeDoubleWave            ST_BorderType = 21
	ST_BorderTypeDashedSmall           ST_BorderType = 22
	ST_BorderTypeDashDotStroked        ST_BorderType = 23
	ST_BorderTypeThreeDEmboss          ST_BorderType = 24
	ST_BorderTypeThreeDEngrave         ST_BorderType = 25
	ST_BorderTypeHTMLOutset            ST_BorderType = 26
	ST_BorderTypeHTMLInset             ST_BorderType = 27
)

func (e ST_BorderType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_BorderTypeUnset:
		attr.Value = ""
	case ST_BorderTypeNone:
		attr.Value = "none"
	case ST_BorderTypeSingle:
		attr.Value = "single"
	case ST_BorderTypeThick:
		attr.Value = "thick"
	case ST_BorderTypeDouble:
		attr.Value = "double"
	case ST_BorderTypeHairline:
		attr.Value = "hairline"
	case ST_BorderTypeDot:
		attr.Value = "dot"
	case ST_BorderTypeDash:
		attr.Value = "dash"
	case ST_BorderTypeDotDash:
		attr.Value = "dotDash"
	case ST_BorderTypeDashDotDot:
		attr.Value = "dashDotDot"
	case ST_BorderTypeTriple:
		attr.Value = "triple"
	case ST_BorderTypeThinThickSmall:
		attr.Value = "thinThickSmall"
	case ST_BorderTypeThickThinSmall:
		attr.Value = "thickThinSmall"
	case ST_BorderTypeThickBetweenThinSmall:
		attr.Value = "thickBetweenThinSmall"
	case ST_BorderTypeThinThick:
		attr.Value = "thinThick"
	case ST_BorderTypeThickThin:
		attr.Value = "thickThin"
	case ST_BorderTypeThickBetweenThin:
		attr.Value = "thickBetweenThin"
	case ST_BorderTypeThinThickLarge:
		attr.Value = "thinThickLarge"
	case ST_BorderTypeThickThinLarge:
		attr.Value = "thickThinLarge"
	case ST_BorderTypeThickBetweenThinLarge:
		attr.Value = "thickBetweenThinLarge"
	case ST_BorderTypeWave:
		attr.Value = "wave"
	case ST_BorderTypeDoubleWave:
		attr.Value = "doubleWave"
	case ST_BorderTypeDashedSmall:
		attr.Value = "dashedSmall"
	case ST_BorderTypeDashDotStroked:
		attr.Value = "dashDotStroked"
	case ST_BorderTypeThreeDEmboss:
		attr.Value = "threeDEmboss"
	case ST_BorderTypeThreeDEngrave:
		attr.Value = "threeDEngrave"
	case ST_BorderTypeHTMLOutset:
		attr.Value = "HTMLOutset"
	case ST_BorderTypeHTMLInset:
		attr.Value = "HTMLInset"
	}
	return attr, nil
}

func (e *ST_BorderType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "none":
		*e = 1
	case "single":
		*e = 2
	case "thick":
		*e = 3
	case "double":
		*e = 4
	case "hairline":
		*e = 5
	case "dot":
		*e = 6
	case "dash":
		*e = 7
	case "dotDash":
		*e = 8
	case "dashDotDot":
		*e = 9
	case "triple":
		*e = 10
	case "thinThickSmall":
		*e = 11
	case "thickThinSmall":
		*e = 12
	case "thickBetweenThinSmall":
		*e = 13
	case "thinThick":
		*e = 14
	case "thickThin":
		*e = 15
	case "thickBetweenThin":
		*e = 16
	case "thinThickLarge":
		*e = 17
	case "thickThinLarge":
		*e = 18
	case "thickBetweenThinLarge":
		*e = 19
	case "wave":
		*e = 20
	case "doubleWave":
		*e = 21
	case "dashedSmall":
		*e = 22
	case "dashDotStroked":
		*e = 23
	case "threeDEmboss":
		*e = 24
	case "threeDEngrave":
		*e = 25
	case "HTMLOutset":
		*e = 26
	case "HTMLInset":
		*e = 27
	}
	return nil
}

func (m ST_BorderType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_BorderType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "none":
			*m = 1
		case "single":
			*m = 2
		case "thick":
			*m = 3
		case "double":
			*m = 4
		case "hairline":
			*m = 5
		case "dot":
			*m = 6
		case "dash":
			*m = 7
		case "dotDash":
			*m = 8
		case "dashDotDot":
			*m = 9
		case "triple":
			*m = 10
		case "thinThickSmall":
			*m = 11
		case "thickThinSmall":
			*m = 12
		case "thickBetweenThinSmall":
			*m = 13
		case "thinThick":
			*m = 14
		case "thickThin":
			*m = 15
		case "thickBetweenThin":
			*m = 16
		case "thinThickLarge":
			*m = 17
		case "thickThinLarge":
			*m = 18
		case "thickBetweenThinLarge":
			*m = 19
		case "wave":
			*m = 20
		case "doubleWave":
			*m = 21
		case "dashedSmall":
			*m = 22
		case "dashDotStroked":
			*m = 23
		case "threeDEmboss":
			*m = 24
		case "threeDEngrave":
			*m = 25
		case "HTMLOutset":
			*m = 26
		case "HTMLInset":
			*m = 27
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

func (m ST_BorderType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "none"
	case 2:
		return "single"
	case 3:
		return "thick"
	case 4:
		return "double"
	case 5:
		return "hairline"
	case 6:
		return "dot"
	case 7:
		return "dash"
	case 8:
		return "dotDash"
	case 9:
		return "dashDotDot"
	case 10:
		return "triple"
	case 11:
		return "thinThickSmall"
	case 12:
		return "thickThinSmall"
	case 13:
		return "thickBetweenThinSmall"
	case 14:
		return "thinThick"
	case 15:
		return "thickThin"
	case 16:
		return "thickBetweenThin"
	case 17:
		return "thinThickLarge"
	case 18:
		return "thickThinLarge"
	case 19:
		return "thickBetweenThinLarge"
	case 20:
		return "wave"
	case 21:
		return "doubleWave"
	case 22:
		return "dashedSmall"
	case 23:
		return "dashDotStroked"
	case 24:
		return "threeDEmboss"
	case 25:
		return "threeDEngrave"
	case 26:
		return "HTMLOutset"
	case 27:
		return "HTMLInset"
	}
	return ""
}

func (m ST_BorderType) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_BorderType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_BorderShadow byte

const (
	ST_BorderShadowUnset ST_BorderShadow = 0
	ST_BorderShadowT     ST_BorderShadow = 1
	ST_BorderShadowTrue  ST_BorderShadow = 2
	ST_BorderShadowF     ST_BorderShadow = 3
	ST_BorderShadowFalse ST_BorderShadow = 4
)

func (e ST_BorderShadow) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_BorderShadowUnset:
		attr.Value = ""
	case ST_BorderShadowT:
		attr.Value = "t"
	case ST_BorderShadowTrue:
		attr.Value = "true"
	case ST_BorderShadowF:
		attr.Value = "f"
	case ST_BorderShadowFalse:
		attr.Value = "false"
	}
	return attr, nil
}

func (e *ST_BorderShadow) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "t":
		*e = 1
	case "true":
		*e = 2
	case "f":
		*e = 3
	case "false":
		*e = 4
	}
	return nil
}

func (m ST_BorderShadow) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_BorderShadow) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "t":
			*m = 1
		case "true":
			*m = 2
		case "f":
			*m = 3
		case "false":
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

func (m ST_BorderShadow) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "t"
	case 2:
		return "true"
	case 3:
		return "f"
	case 4:
		return "false"
	}
	return ""
}

func (m ST_BorderShadow) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_BorderShadow) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_WrapType byte

const (
	ST_WrapTypeUnset        ST_WrapType = 0
	ST_WrapTypeTopAndBottom ST_WrapType = 1
	ST_WrapTypeSquare       ST_WrapType = 2
	ST_WrapTypeNone         ST_WrapType = 3
	ST_WrapTypeTight        ST_WrapType = 4
	ST_WrapTypeThrough      ST_WrapType = 5
)

func (e ST_WrapType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_WrapTypeUnset:
		attr.Value = ""
	case ST_WrapTypeTopAndBottom:
		attr.Value = "topAndBottom"
	case ST_WrapTypeSquare:
		attr.Value = "square"
	case ST_WrapTypeNone:
		attr.Value = "none"
	case ST_WrapTypeTight:
		attr.Value = "tight"
	case ST_WrapTypeThrough:
		attr.Value = "through"
	}
	return attr, nil
}

func (e *ST_WrapType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "topAndBottom":
		*e = 1
	case "square":
		*e = 2
	case "none":
		*e = 3
	case "tight":
		*e = 4
	case "through":
		*e = 5
	}
	return nil
}

func (m ST_WrapType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_WrapType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "topAndBottom":
			*m = 1
		case "square":
			*m = 2
		case "none":
			*m = 3
		case "tight":
			*m = 4
		case "through":
			*m = 5
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

func (m ST_WrapType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "topAndBottom"
	case 2:
		return "square"
	case 3:
		return "none"
	case 4:
		return "tight"
	case 5:
		return "through"
	}
	return ""
}

func (m ST_WrapType) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_WrapType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_WrapSide byte

const (
	ST_WrapSideUnset   ST_WrapSide = 0
	ST_WrapSideBoth    ST_WrapSide = 1
	ST_WrapSideLeft    ST_WrapSide = 2
	ST_WrapSideRight   ST_WrapSide = 3
	ST_WrapSideLargest ST_WrapSide = 4
)

func (e ST_WrapSide) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_WrapSideUnset:
		attr.Value = ""
	case ST_WrapSideBoth:
		attr.Value = "both"
	case ST_WrapSideLeft:
		attr.Value = "left"
	case ST_WrapSideRight:
		attr.Value = "right"
	case ST_WrapSideLargest:
		attr.Value = "largest"
	}
	return attr, nil
}

func (e *ST_WrapSide) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "both":
		*e = 1
	case "left":
		*e = 2
	case "right":
		*e = 3
	case "largest":
		*e = 4
	}
	return nil
}

func (m ST_WrapSide) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_WrapSide) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "both":
			*m = 1
		case "left":
			*m = 2
		case "right":
			*m = 3
		case "largest":
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

func (m ST_WrapSide) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "both"
	case 2:
		return "left"
	case 3:
		return "right"
	case 4:
		return "largest"
	}
	return ""
}

func (m ST_WrapSide) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_WrapSide) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_HorizontalAnchor byte

const (
	ST_HorizontalAnchorUnset  ST_HorizontalAnchor = 0
	ST_HorizontalAnchorMargin ST_HorizontalAnchor = 1
	ST_HorizontalAnchorPage   ST_HorizontalAnchor = 2
	ST_HorizontalAnchorText   ST_HorizontalAnchor = 3
	ST_HorizontalAnchorChar   ST_HorizontalAnchor = 4
)

func (e ST_HorizontalAnchor) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_HorizontalAnchorUnset:
		attr.Value = ""
	case ST_HorizontalAnchorMargin:
		attr.Value = "margin"
	case ST_HorizontalAnchorPage:
		attr.Value = "page"
	case ST_HorizontalAnchorText:
		attr.Value = "text"
	case ST_HorizontalAnchorChar:
		attr.Value = "char"
	}
	return attr, nil
}

func (e *ST_HorizontalAnchor) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "margin":
		*e = 1
	case "page":
		*e = 2
	case "text":
		*e = 3
	case "char":
		*e = 4
	}
	return nil
}

func (m ST_HorizontalAnchor) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_HorizontalAnchor) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "margin":
			*m = 1
		case "page":
			*m = 2
		case "text":
			*m = 3
		case "char":
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

func (m ST_HorizontalAnchor) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "margin"
	case 2:
		return "page"
	case 3:
		return "text"
	case 4:
		return "char"
	}
	return ""
}

func (m ST_HorizontalAnchor) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_HorizontalAnchor) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_VerticalAnchor byte

const (
	ST_VerticalAnchorUnset  ST_VerticalAnchor = 0
	ST_VerticalAnchorMargin ST_VerticalAnchor = 1
	ST_VerticalAnchorPage   ST_VerticalAnchor = 2
	ST_VerticalAnchorText   ST_VerticalAnchor = 3
	ST_VerticalAnchorLine   ST_VerticalAnchor = 4
)

func (e ST_VerticalAnchor) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_VerticalAnchorUnset:
		attr.Value = ""
	case ST_VerticalAnchorMargin:
		attr.Value = "margin"
	case ST_VerticalAnchorPage:
		attr.Value = "page"
	case ST_VerticalAnchorText:
		attr.Value = "text"
	case ST_VerticalAnchorLine:
		attr.Value = "line"
	}
	return attr, nil
}

func (e *ST_VerticalAnchor) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "margin":
		*e = 1
	case "page":
		*e = 2
	case "text":
		*e = 3
	case "line":
		*e = 4
	}
	return nil
}

func (m ST_VerticalAnchor) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_VerticalAnchor) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "margin":
			*m = 1
		case "page":
			*m = 2
		case "text":
			*m = 3
		case "line":
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

func (m ST_VerticalAnchor) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "margin"
	case 2:
		return "page"
	case 3:
		return "text"
	case 4:
		return "line"
	}
	return ""
}

func (m ST_VerticalAnchor) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_VerticalAnchor) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

// init registers constructor functions for dynamically creating elements based off the XML namespace and name
func init() {
	goffice.RegisterConstructor("urn:schemas-microsoft-com:office:word", "CT_Border", NewCT_Border)
	goffice.RegisterConstructor("urn:schemas-microsoft-com:office:word", "CT_Wrap", NewCT_Wrap)
	goffice.RegisterConstructor("urn:schemas-microsoft-com:office:word", "CT_AnchorLock", NewCT_AnchorLock)
	goffice.RegisterConstructor("urn:schemas-microsoft-com:office:word", "bordertop", NewBordertop)
	goffice.RegisterConstructor("urn:schemas-microsoft-com:office:word", "borderleft", NewBorderleft)
	goffice.RegisterConstructor("urn:schemas-microsoft-com:office:word", "borderright", NewBorderright)
	goffice.RegisterConstructor("urn:schemas-microsoft-com:office:word", "borderbottom", NewBorderbottom)
	goffice.RegisterConstructor("urn:schemas-microsoft-com:office:word", "wrap", NewWrap)
	goffice.RegisterConstructor("urn:schemas-microsoft-com:office:word", "anchorlock", NewAnchorlock)
}
