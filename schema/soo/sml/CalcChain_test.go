package sml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/sml"
)

func TestCalcChainConstructor(t *testing.T) {
	v := sml.NewCalcChain()
	if v == nil {
		t.Errorf("sml.NewCalcChain must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CalcChain should validate: %s", err)
	}
}

func TestCalcChainMarshalUnmarshal(t *testing.T) {
	v := sml.NewCalcChain()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCalcChain()
	xml.Unmarshal(buf, v2)
}
