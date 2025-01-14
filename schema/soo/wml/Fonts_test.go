package wml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/wml"
)

func TestFontsConstructor(t *testing.T) {
	v := wml.NewFonts()
	if v == nil {
		t.Errorf("wml.NewFonts must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.Fonts should validate: %s", err)
	}
}

func TestFontsMarshalUnmarshal(t *testing.T) {
	v := wml.NewFonts()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewFonts()
	xml.Unmarshal(buf, v2)
}
