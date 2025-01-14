package sml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/sml"
)

func TestCT_ColConstructor(t *testing.T) {
	v := sml.NewCT_Col()
	if v == nil {
		t.Errorf("sml.NewCT_Col must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Col should validate: %s", err)
	}
}

func TestCT_ColMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Col()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Col()
	xml.Unmarshal(buf, v2)
}
