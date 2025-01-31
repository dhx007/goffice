package diagram_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/dml/diagram"
)

func TestCT_CTDescriptionConstructor(t *testing.T) {
	v := diagram.NewCT_CTDescription()
	if v == nil {
		t.Errorf("diagram.NewCT_CTDescription must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_CTDescription should validate: %s", err)
	}
}

func TestCT_CTDescriptionMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_CTDescription()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_CTDescription()
	xml.Unmarshal(buf, v2)
}
