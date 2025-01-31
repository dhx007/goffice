package dml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/dml"
)

func TestCT_TableCellBorderStyleConstructor(t *testing.T) {
	v := dml.NewCT_TableCellBorderStyle()
	if v == nil {
		t.Errorf("dml.NewCT_TableCellBorderStyle must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_TableCellBorderStyle should validate: %s", err)
	}
}

func TestCT_TableCellBorderStyleMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_TableCellBorderStyle()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_TableCellBorderStyle()
	xml.Unmarshal(buf, v2)
}
