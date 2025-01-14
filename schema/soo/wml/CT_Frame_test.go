package wml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/wml"
)

func TestCT_FrameConstructor(t *testing.T) {
	v := wml.NewCT_Frame()
	if v == nil {
		t.Errorf("wml.NewCT_Frame must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Frame should validate: %s", err)
	}
}

func TestCT_FrameMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Frame()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Frame()
	xml.Unmarshal(buf, v2)
}
