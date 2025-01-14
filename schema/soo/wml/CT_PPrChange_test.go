package wml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/wml"
)

func TestCT_PPrChangeConstructor(t *testing.T) {
	v := wml.NewCT_PPrChange()
	if v == nil {
		t.Errorf("wml.NewCT_PPrChange must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_PPrChange should validate: %s", err)
	}
}

func TestCT_PPrChangeMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_PPrChange()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_PPrChange()
	xml.Unmarshal(buf, v2)
}
