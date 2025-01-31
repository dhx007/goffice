package math_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/ofc/math"
)

func TestCT_CtrlPrConstructor(t *testing.T) {
	v := math.NewCT_CtrlPr()
	if v == nil {
		t.Errorf("math.NewCT_CtrlPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_CtrlPr should validate: %s", err)
	}
}

func TestCT_CtrlPrMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_CtrlPr()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_CtrlPr()
	xml.Unmarshal(buf, v2)
}
