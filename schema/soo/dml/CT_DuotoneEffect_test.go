package dml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/dml"
)

func TestCT_DuotoneEffectConstructor(t *testing.T) {
	v := dml.NewCT_DuotoneEffect()
	if v == nil {
		t.Errorf("dml.NewCT_DuotoneEffect must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_DuotoneEffect should validate: %s", err)
	}
}

func TestCT_DuotoneEffectMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_DuotoneEffect()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_DuotoneEffect()
	xml.Unmarshal(buf, v2)
}
