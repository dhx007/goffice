package wml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/wml"
)

func TestCT_TrPrChangeConstructor(t *testing.T) {
	v := wml.NewCT_TrPrChange()
	if v == nil {
		t.Errorf("wml.NewCT_TrPrChange must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_TrPrChange should validate: %s", err)
	}
}

func TestCT_TrPrChangeMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_TrPrChange()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_TrPrChange()
	xml.Unmarshal(buf, v2)
}
