package docPropsVTypes_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/ofc/docPropsVTypes"
)

func TestVariantConstructor(t *testing.T) {
	v := docPropsVTypes.NewVariant()
	if v == nil {
		t.Errorf("docPropsVTypes.NewVariant must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed docPropsVTypes.Variant should validate: %s", err)
	}
}

func TestVariantMarshalUnmarshal(t *testing.T) {
	v := docPropsVTypes.NewVariant()
	buf, _ := xml.Marshal(v)
	v2 := docPropsVTypes.NewVariant()
	xml.Unmarshal(buf, v2)
}
