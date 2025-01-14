package elements_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/purl.org/dc/elements"
)

func TestElementContainerConstructor(t *testing.T) {
	v := elements.NewElementContainer()
	if v == nil {
		t.Errorf("elements.NewElementContainer must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed elements.ElementContainer should validate: %s", err)
	}
}

func TestElementContainerMarshalUnmarshal(t *testing.T) {
	v := elements.NewElementContainer()
	buf, _ := xml.Marshal(v)
	v2 := elements.NewElementContainer()
	xml.Unmarshal(buf, v2)
}
