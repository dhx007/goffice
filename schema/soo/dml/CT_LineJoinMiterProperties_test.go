package dml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/dml"
)

func TestCT_LineJoinMiterPropertiesConstructor(t *testing.T) {
	v := dml.NewCT_LineJoinMiterProperties()
	if v == nil {
		t.Errorf("dml.NewCT_LineJoinMiterProperties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_LineJoinMiterProperties should validate: %s", err)
	}
}

func TestCT_LineJoinMiterPropertiesMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_LineJoinMiterProperties()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_LineJoinMiterProperties()
	xml.Unmarshal(buf, v2)
}
