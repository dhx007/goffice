package wml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/wml"
)

func TestCT_DecimalNumberConstructor(t *testing.T) {
	v := wml.NewCT_DecimalNumber()
	if v == nil {
		t.Errorf("wml.NewCT_DecimalNumber must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_DecimalNumber should validate: %s", err)
	}
}

func TestCT_DecimalNumberMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_DecimalNumber()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_DecimalNumber()
	xml.Unmarshal(buf, v2)
}
