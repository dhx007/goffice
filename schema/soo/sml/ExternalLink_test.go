package sml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/sml"
)

func TestExternalLinkConstructor(t *testing.T) {
	v := sml.NewExternalLink()
	if v == nil {
		t.Errorf("sml.NewExternalLink must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.ExternalLink should validate: %s", err)
	}
}

func TestExternalLinkMarshalUnmarshal(t *testing.T) {
	v := sml.NewExternalLink()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewExternalLink()
	xml.Unmarshal(buf, v2)
}
