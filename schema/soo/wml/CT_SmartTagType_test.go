package wml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/wml"
)

func TestCT_SmartTagTypeConstructor(t *testing.T) {
	v := wml.NewCT_SmartTagType()
	if v == nil {
		t.Errorf("wml.NewCT_SmartTagType must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_SmartTagType should validate: %s", err)
	}
}

func TestCT_SmartTagTypeMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_SmartTagType()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_SmartTagType()
	xml.Unmarshal(buf, v2)
}
