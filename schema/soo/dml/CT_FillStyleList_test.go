package dml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/dml"
)

func TestCT_FillStyleListConstructor(t *testing.T) {
	v := dml.NewCT_FillStyleList()
	if v == nil {
		t.Errorf("dml.NewCT_FillStyleList must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_FillStyleList should validate: %s", err)
	}
}

func TestCT_FillStyleListMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_FillStyleList()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_FillStyleList()
	xml.Unmarshal(buf, v2)
}
