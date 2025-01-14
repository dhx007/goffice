package dml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/dml"
)

func TestCT_DashStopListConstructor(t *testing.T) {
	v := dml.NewCT_DashStopList()
	if v == nil {
		t.Errorf("dml.NewCT_DashStopList must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_DashStopList should validate: %s", err)
	}
}

func TestCT_DashStopListMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_DashStopList()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_DashStopList()
	xml.Unmarshal(buf, v2)
}
