package wml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/wml"
)

func TestEG_RPrMathConstructor(t *testing.T) {
	v := wml.NewEG_RPrMath()
	if v == nil {
		t.Errorf("wml.NewEG_RPrMath must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.EG_RPrMath should validate: %s", err)
	}
}

func TestEG_RPrMathMarshalUnmarshal(t *testing.T) {
	v := wml.NewEG_RPrMath()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewEG_RPrMath()
	xml.Unmarshal(buf, v2)
}
