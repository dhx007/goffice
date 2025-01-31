package sml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/sml"
)

func TestCT_SmartTagTypesConstructor(t *testing.T) {
	v := sml.NewCT_SmartTagTypes()
	if v == nil {
		t.Errorf("sml.NewCT_SmartTagTypes must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_SmartTagTypes should validate: %s", err)
	}
}

func TestCT_SmartTagTypesMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_SmartTagTypes()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_SmartTagTypes()
	xml.Unmarshal(buf, v2)
}
