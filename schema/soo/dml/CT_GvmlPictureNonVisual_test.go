package dml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/dml"
)

func TestCT_GvmlPictureNonVisualConstructor(t *testing.T) {
	v := dml.NewCT_GvmlPictureNonVisual()
	if v == nil {
		t.Errorf("dml.NewCT_GvmlPictureNonVisual must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_GvmlPictureNonVisual should validate: %s", err)
	}
}

func TestCT_GvmlPictureNonVisualMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_GvmlPictureNonVisual()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_GvmlPictureNonVisual()
	xml.Unmarshal(buf, v2)
}
