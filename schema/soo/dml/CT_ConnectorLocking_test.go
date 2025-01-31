package dml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/dml"
)

func TestCT_ConnectorLockingConstructor(t *testing.T) {
	v := dml.NewCT_ConnectorLocking()
	if v == nil {
		t.Errorf("dml.NewCT_ConnectorLocking must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_ConnectorLocking should validate: %s", err)
	}
}

func TestCT_ConnectorLockingMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_ConnectorLocking()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_ConnectorLocking()
	xml.Unmarshal(buf, v2)
}
