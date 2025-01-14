package math_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/ofc/math"
)

func TestCT_FTypeConstructor(t *testing.T) {
	v := math.NewCT_FType()
	if v == nil {
		t.Errorf("math.NewCT_FType must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_FType should validate: %s", err)
	}
}

func TestCT_FTypeMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_FType()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_FType()
	xml.Unmarshal(buf, v2)
}
