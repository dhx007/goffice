package dml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/dml"
)

func TestCT_CameraConstructor(t *testing.T) {
	v := dml.NewCT_Camera()
	if v == nil {
		t.Errorf("dml.NewCT_Camera must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_Camera should validate: %s", err)
	}
}

func TestCT_CameraMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_Camera()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_Camera()
	xml.Unmarshal(buf, v2)
}
