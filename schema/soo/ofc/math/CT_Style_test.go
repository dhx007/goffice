package math_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/ofc/math"
)

func TestCT_StyleConstructor(t *testing.T) {
	v := math.NewCT_Style()
	if v == nil {
		t.Errorf("math.NewCT_Style must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_Style should validate: %s", err)
	}
}

func TestCT_StyleMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_Style()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_Style()
	xml.Unmarshal(buf, v2)
}
