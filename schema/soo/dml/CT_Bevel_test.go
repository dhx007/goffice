package dml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/dml"
)

func TestCT_BevelConstructor(t *testing.T) {
	v := dml.NewCT_Bevel()
	if v == nil {
		t.Errorf("dml.NewCT_Bevel must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_Bevel should validate: %s", err)
	}
}

func TestCT_BevelMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_Bevel()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_Bevel()
	xml.Unmarshal(buf, v2)
}
