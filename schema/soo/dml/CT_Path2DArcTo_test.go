package dml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/dml"
)

func TestCT_Path2DArcToConstructor(t *testing.T) {
	v := dml.NewCT_Path2DArcTo()
	if v == nil {
		t.Errorf("dml.NewCT_Path2DArcTo must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_Path2DArcTo should validate: %s", err)
	}
}

func TestCT_Path2DArcToMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_Path2DArcTo()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_Path2DArcTo()
	xml.Unmarshal(buf, v2)
}
