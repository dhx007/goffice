package math_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/ofc/math"
)

func TestOMathConstructor(t *testing.T) {
	v := math.NewOMath()
	if v == nil {
		t.Errorf("math.NewOMath must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.OMath should validate: %s", err)
	}
}

func TestOMathMarshalUnmarshal(t *testing.T) {
	v := math.NewOMath()
	buf, _ := xml.Marshal(v)
	v2 := math.NewOMath()
	xml.Unmarshal(buf, v2)
}
