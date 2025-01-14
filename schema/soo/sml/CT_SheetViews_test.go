package sml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/sml"
)

func TestCT_SheetViewsConstructor(t *testing.T) {
	v := sml.NewCT_SheetViews()
	if v == nil {
		t.Errorf("sml.NewCT_SheetViews must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_SheetViews should validate: %s", err)
	}
}

func TestCT_SheetViewsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_SheetViews()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_SheetViews()
	xml.Unmarshal(buf, v2)
}
