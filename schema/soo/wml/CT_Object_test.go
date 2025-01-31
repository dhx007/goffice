package wml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/wml"
)

func TestCT_ObjectConstructor(t *testing.T) {
	v := wml.NewCT_Object()
	if v == nil {
		t.Errorf("wml.NewCT_Object must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Object should validate: %s", err)
	}
}

func TestCT_ObjectMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Object()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Object()
	xml.Unmarshal(buf, v2)
}
