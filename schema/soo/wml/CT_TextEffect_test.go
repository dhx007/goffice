package wml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/wml"
)

func TestCT_TextEffectConstructor(t *testing.T) {
	v := wml.NewCT_TextEffect()
	if v == nil {
		t.Errorf("wml.NewCT_TextEffect must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_TextEffect should validate: %s", err)
	}
}

func TestCT_TextEffectMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_TextEffect()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_TextEffect()
	xml.Unmarshal(buf, v2)
}
