package sml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/sml"
)

func TestCT_CalcCellConstructor(t *testing.T) {
	v := sml.NewCT_CalcCell()
	if v == nil {
		t.Errorf("sml.NewCT_CalcCell must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_CalcCell should validate: %s", err)
	}
}

func TestCT_CalcCellMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_CalcCell()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_CalcCell()
	xml.Unmarshal(buf, v2)
}
