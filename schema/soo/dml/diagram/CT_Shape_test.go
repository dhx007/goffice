package diagram_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/dml/diagram"
)

func TestCT_ShapeConstructor(t *testing.T) {
	v := diagram.NewCT_Shape()
	if v == nil {
		t.Errorf("diagram.NewCT_Shape must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_Shape should validate: %s", err)
	}
}

func TestCT_ShapeMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_Shape()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_Shape()
	xml.Unmarshal(buf, v2)
}
