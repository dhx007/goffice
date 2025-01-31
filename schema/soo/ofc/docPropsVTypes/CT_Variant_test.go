package docPropsVTypes_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/ofc/docPropsVTypes"
)

func TestCT_VariantConstructor(t *testing.T) {
	v := docPropsVTypes.NewCT_Variant()
	if v == nil {
		t.Errorf("docPropsVTypes.NewCT_Variant must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed docPropsVTypes.CT_Variant should validate: %s", err)
	}
}

func TestCT_VariantMarshalUnmarshal(t *testing.T) {
	v := docPropsVTypes.NewCT_Variant()
	buf, _ := xml.Marshal(v)
	v2 := docPropsVTypes.NewCT_Variant()
	xml.Unmarshal(buf, v2)
}
