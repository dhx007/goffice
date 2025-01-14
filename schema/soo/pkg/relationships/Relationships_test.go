package relationships_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/pkg/relationships"
)

func TestRelationshipsConstructor(t *testing.T) {
	v := relationships.NewRelationships()
	if v == nil {
		t.Errorf("relationships.NewRelationships must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed relationships.Relationships should validate: %s", err)
	}
}

func TestRelationshipsMarshalUnmarshal(t *testing.T) {
	v := relationships.NewRelationships()
	buf, _ := xml.Marshal(v)
	v2 := relationships.NewRelationships()
	xml.Unmarshal(buf, v2)
}
