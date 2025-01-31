package sml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/sml"
)

func TestCT_SmartTagTypeConstructor(t *testing.T) {
	v := sml.NewCT_SmartTagType()
	if v == nil {
		t.Errorf("sml.NewCT_SmartTagType must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_SmartTagType should validate: %s", err)
	}
}

func TestCT_SmartTagTypeMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_SmartTagType()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_SmartTagType()
	xml.Unmarshal(buf, v2)
}
