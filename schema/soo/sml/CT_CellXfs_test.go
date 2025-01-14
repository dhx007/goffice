package sml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/sml"
)

func TestCT_CellXfsConstructor(t *testing.T) {
	v := sml.NewCT_CellXfs()
	if v == nil {
		t.Errorf("sml.NewCT_CellXfs must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_CellXfs should validate: %s", err)
	}
}

func TestCT_CellXfsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_CellXfs()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_CellXfs()
	xml.Unmarshal(buf, v2)
}
