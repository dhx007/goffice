package dml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/dml"
)

func TestCT_TableBackgroundStyleConstructor(t *testing.T) {
	v := dml.NewCT_TableBackgroundStyle()
	if v == nil {
		t.Errorf("dml.NewCT_TableBackgroundStyle must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_TableBackgroundStyle should validate: %s", err)
	}
}

func TestCT_TableBackgroundStyleMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_TableBackgroundStyle()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_TableBackgroundStyle()
	xml.Unmarshal(buf, v2)
}
