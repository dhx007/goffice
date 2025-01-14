package wml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/wml"
)

func TestWdCT_TextboxInfoConstructor(t *testing.T) {
	v := wml.NewWdCT_TextboxInfo()
	if v == nil {
		t.Errorf("wml.NewWdCT_TextboxInfo must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.WdCT_TextboxInfo should validate: %s", err)
	}
}

func TestWdCT_TextboxInfoMarshalUnmarshal(t *testing.T) {
	v := wml.NewWdCT_TextboxInfo()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewWdCT_TextboxInfo()
	xml.Unmarshal(buf, v2)
}
