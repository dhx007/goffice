package math_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/ofc/math"
)

func TestCT_MathPrConstructor(t *testing.T) {
	v := math.NewCT_MathPr()
	if v == nil {
		t.Errorf("math.NewCT_MathPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_MathPr should validate: %s", err)
	}
}

func TestCT_MathPrMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_MathPr()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_MathPr()
	xml.Unmarshal(buf, v2)
}
