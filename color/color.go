package color

import (
	"github.com/dhx007/goffice"
)

// Color is a 24 bit color that can be converted to
// internal ECMA-376 formats as needed.
type Color struct {
	r, g, b, a uint8
	isAuto     bool
}

// RGB constructs a new RGB color with a given red, green and blue value.
func RGB(r, g, b uint8) Color {
	return Color{r, g, b, 255, false}
}

// RGBA constructs a new RGBA color with a given red, green, blue and alpha
// value.
func RGBA(r, g, b, a uint8) Color {
	return Color{r, g, b, a, false}
}

// IsAuto returns true if the color is the 'Auto' type.  If the
// field doesn't support an Auto color, then black is used.
func (c Color) IsAuto() bool {
	return c.isAuto
}

// AsRGBString is used by the various wrappers to return a pointer
// to a string containing a six digit hex RGB value.
func (c Color) AsRGBString() *string {
	return goffice.Stringf("%02x%02x%02x", c.r, c.g, c.b)
}

// AsRGBAString is used by the various wrappers to return a pointer
// to a string containing a six digit hex RGB value.
func (c Color) AsRGBAString() *string {
	return goffice.Stringf("%02x%02x%02x%02x", c.a, c.r, c.g, c.b)
}
