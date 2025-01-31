package sml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/sml"
)

func TestCT_DrawingConstructor(t *testing.T) {
	v := sml.NewCT_Drawing()
	if v == nil {
		t.Errorf("sml.NewCT_Drawing must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Drawing should validate: %s", err)
	}
}

func TestCT_DrawingMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Drawing()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Drawing()
	xml.Unmarshal(buf, v2)
}
