package sml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/sml"
)

func TestCT_PageBreakConstructor(t *testing.T) {
	v := sml.NewCT_PageBreak()
	if v == nil {
		t.Errorf("sml.NewCT_PageBreak must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_PageBreak should validate: %s", err)
	}
}

func TestCT_PageBreakMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_PageBreak()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_PageBreak()
	xml.Unmarshal(buf, v2)
}
