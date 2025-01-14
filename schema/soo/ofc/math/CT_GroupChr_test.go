package math_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/ofc/math"
)

func TestCT_GroupChrConstructor(t *testing.T) {
	v := math.NewCT_GroupChr()
	if v == nil {
		t.Errorf("math.NewCT_GroupChr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_GroupChr should validate: %s", err)
	}
}

func TestCT_GroupChrMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_GroupChr()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_GroupChr()
	xml.Unmarshal(buf, v2)
}
