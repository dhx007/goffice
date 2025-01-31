package math_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/ofc/math"
)

func TestCT_RadConstructor(t *testing.T) {
	v := math.NewCT_Rad()
	if v == nil {
		t.Errorf("math.NewCT_Rad must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_Rad should validate: %s", err)
	}
}

func TestCT_RadMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_Rad()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_Rad()
	xml.Unmarshal(buf, v2)
}
