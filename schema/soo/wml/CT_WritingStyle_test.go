package wml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/wml"
)

func TestCT_WritingStyleConstructor(t *testing.T) {
	v := wml.NewCT_WritingStyle()
	if v == nil {
		t.Errorf("wml.NewCT_WritingStyle must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_WritingStyle should validate: %s", err)
	}
}

func TestCT_WritingStyleMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_WritingStyle()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_WritingStyle()
	xml.Unmarshal(buf, v2)
}
