package wml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/wml"
)

func TestAG_SectPrAttributesConstructor(t *testing.T) {
	v := wml.NewAG_SectPrAttributes()
	if v == nil {
		t.Errorf("wml.NewAG_SectPrAttributes must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.AG_SectPrAttributes should validate: %s", err)
	}
}

func TestAG_SectPrAttributesMarshalUnmarshal(t *testing.T) {
	v := wml.NewAG_SectPrAttributes()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewAG_SectPrAttributes()
	xml.Unmarshal(buf, v2)
}
