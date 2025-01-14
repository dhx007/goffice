package terms_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/purl.org/dc/terms"
)

func TestURIConstructor(t *testing.T) {
	v := terms.NewURI()
	if v == nil {
		t.Errorf("terms.NewURI must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed terms.URI should validate: %s", err)
	}
}

func TestURIMarshalUnmarshal(t *testing.T) {
	v := terms.NewURI()
	buf, _ := xml.Marshal(v)
	v2 := terms.NewURI()
	xml.Unmarshal(buf, v2)
}
