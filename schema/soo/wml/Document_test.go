package wml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/wml"
)

func TestDocumentConstructor(t *testing.T) {
	v := wml.NewDocument()
	if v == nil {
		t.Errorf("wml.NewDocument must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.Document should validate: %s", err)
	}
}

func TestDocumentMarshalUnmarshal(t *testing.T) {
	v := wml.NewDocument()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewDocument()
	xml.Unmarshal(buf, v2)
}
