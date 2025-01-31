package sml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/sml"
)

func TestCT_ColorScaleConstructor(t *testing.T) {
	v := sml.NewCT_ColorScale()
	if v == nil {
		t.Errorf("sml.NewCT_ColorScale must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_ColorScale should validate: %s", err)
	}
}

func TestCT_ColorScaleMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_ColorScale()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_ColorScale()
	xml.Unmarshal(buf, v2)
}
