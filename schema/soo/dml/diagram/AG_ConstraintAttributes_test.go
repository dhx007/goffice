package diagram_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/dml/diagram"
)

func TestAG_ConstraintAttributesConstructor(t *testing.T) {
	v := diagram.NewAG_ConstraintAttributes()
	if v == nil {
		t.Errorf("diagram.NewAG_ConstraintAttributes must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.AG_ConstraintAttributes should validate: %s", err)
	}
}

func TestAG_ConstraintAttributesMarshalUnmarshal(t *testing.T) {
	v := diagram.NewAG_ConstraintAttributes()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewAG_ConstraintAttributes()
	xml.Unmarshal(buf, v2)
}
