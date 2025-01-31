package math_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/ofc/math"
)

func TestCT_MRConstructor(t *testing.T) {
	v := math.NewCT_MR()
	if v == nil {
		t.Errorf("math.NewCT_MR must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_MR should validate: %s", err)
	}
}

func TestCT_MRMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_MR()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_MR()
	xml.Unmarshal(buf, v2)
}
