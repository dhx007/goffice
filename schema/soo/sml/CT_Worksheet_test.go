package sml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/sml"
)

func TestCT_WorksheetConstructor(t *testing.T) {
	v := sml.NewCT_Worksheet()
	if v == nil {
		t.Errorf("sml.NewCT_Worksheet must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Worksheet should validate: %s", err)
	}
}

func TestCT_WorksheetMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Worksheet()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Worksheet()
	xml.Unmarshal(buf, v2)
}
