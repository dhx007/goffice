package pml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/pml"
)

func TestCT_TLTimeConditionConstructor(t *testing.T) {
	v := pml.NewCT_TLTimeCondition()
	if v == nil {
		t.Errorf("pml.NewCT_TLTimeCondition must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_TLTimeCondition should validate: %s", err)
	}
}

func TestCT_TLTimeConditionMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_TLTimeCondition()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_TLTimeCondition()
	xml.Unmarshal(buf, v2)
}
