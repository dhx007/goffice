package terms_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/purl.org/dc/terms"
)

func TestUDCConstructor(t *testing.T) {
	v := terms.NewUDC()
	if v == nil {
		t.Errorf("terms.NewUDC must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed terms.UDC should validate: %s", err)
	}
}

func TestUDCMarshalUnmarshal(t *testing.T) {
	v := terms.NewUDC()
	buf, _ := xml.Marshal(v)
	v2 := terms.NewUDC()
	xml.Unmarshal(buf, v2)
}
