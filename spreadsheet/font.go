package spreadsheet

import (
	"github.com/dhx007/goffice/color"
	"github.com/dhx007/goffice/schema/soo/sml"
)

// Font allows editing fonts within a spreadsheet stylesheet.
type Font struct {
	font   *sml.CT_Font
	styles *sml.StyleSheet
}

// X returns the inner wrapped XML type.
func (f Font) X() *sml.CT_Font {
	return f.font
}

func (f Font) Index() uint32 {
	for i, sf := range f.styles.Fonts.Font {
		if f.font == sf {
			return uint32(i)
		}
	}
	return 0
}

func (f Font) SetBold(b bool) {
	if b {
		f.font.B = []*sml.CT_BooleanProperty{{}}
	} else {
		f.font.B = nil
	}
}
func (f Font) SetItalic(b bool) {
	if b {
		f.font.I = []*sml.CT_BooleanProperty{{}}
	} else {
		f.font.I = nil
	}
}

func (f Font) SetName(name string) {
	f.font.Name = []*sml.CT_FontName{{ValAttr: name}}
}

func (f Font) SetSize(size float64) {
	f.font.Sz = []*sml.CT_FontSize{{ValAttr: size}}
}

func (f Font) SetColor(c color.Color) {
	clr := sml.NewCT_Color()
	rgbAttr := "ff" + *c.AsRGBString()
	clr.RgbAttr = &rgbAttr
	f.font.Color = []*sml.CT_Color{clr}
}
