package pml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/pml"
)

func TestCT_TLTimeNodeExclusiveConstructor(t *testing.T) {
	v := pml.NewCT_TLTimeNodeExclusive()
	if v == nil {
		t.Errorf("pml.NewCT_TLTimeNodeExclusive must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_TLTimeNodeExclusive should validate: %s", err)
	}
}

func TestCT_TLTimeNodeExclusiveMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_TLTimeNodeExclusive()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_TLTimeNodeExclusive()
	xml.Unmarshal(buf, v2)
}
