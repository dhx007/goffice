package sml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/sml"
)

func TestCT_SetConstructor(t *testing.T) {
	v := sml.NewCT_Set()
	if v == nil {
		t.Errorf("sml.NewCT_Set must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Set should validate: %s", err)
	}
}

func TestCT_SetMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Set()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Set()
	xml.Unmarshal(buf, v2)
}
