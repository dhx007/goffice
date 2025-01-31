package vml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/urn/schemas_microsoft_com/vml"
)

func TestOfcCT_SkewConstructor(t *testing.T) {
	v := vml.NewOfcCT_Skew()
	if v == nil {
		t.Errorf("vml.NewOfcCT_Skew must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.OfcCT_Skew should validate: %s", err)
	}
}

func TestOfcCT_SkewMarshalUnmarshal(t *testing.T) {
	v := vml.NewOfcCT_Skew()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewOfcCT_Skew()
	xml.Unmarshal(buf, v2)
}
