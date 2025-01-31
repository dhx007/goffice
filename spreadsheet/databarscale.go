package spreadsheet

import (
	"github.com/dhx007/goffice"
	"github.com/dhx007/goffice/color"
	"github.com/dhx007/goffice/schema/soo/sml"
)

// DataBarScale is a colored scale that fills the cell with a background
// gradeint depending on the value.
type DataBarScale struct {
	x *sml.CT_DataBar
}

// X returns the inner wrapped XML type.
func (d DataBarScale) X() *sml.CT_DataBar {
	return d.x
}

// AddFormatValue adds a format value (databars require two).
func (d DataBarScale) AddFormatValue(t sml.ST_CfvoType, val string) {
	v := sml.NewCT_Cfvo()
	v.TypeAttr = t
	v.ValAttr = goffice.String(val)
	d.x.Cfvo = append(d.x.Cfvo, v)
}

// SetColor sets teh color of the databar.
func (d DataBarScale) SetColor(c color.Color) {
	d.x.Color = sml.NewCT_Color()
	d.x.Color.RgbAttr = c.AsRGBAString()
}

// SetShowValue controls if the cell value is displayed.
func (d DataBarScale) SetShowValue(b bool) {
	d.x.ShowValueAttr = goffice.Bool(b)
}

// SetMinLength sets the minimum bar length in percent.
func (d DataBarScale) SetMinLength(l uint32) {
	d.x.MinLengthAttr = goffice.Uint32(l)
}

// SetMaxLength sets the maximum bar length in percent.
func (d DataBarScale) SetMaxLength(l uint32) {
	d.x.MaxLengthAttr = goffice.Uint32(l)
}
