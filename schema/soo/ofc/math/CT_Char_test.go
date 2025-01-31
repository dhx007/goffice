package math_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/ofc/math"
)

func TestCT_CharConstructor(t *testing.T) {
	v := math.NewCT_Char()
	if v == nil {
		t.Errorf("math.NewCT_Char must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_Char should validate: %s", err)
	}
}

func TestCT_CharMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_Char()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_Char()
	xml.Unmarshal(buf, v2)
}
