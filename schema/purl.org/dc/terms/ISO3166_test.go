// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package terms_test

import (
	"encoding/xml"
	"testing"

	"goffice/schema/purl.org/dc/terms"
)

func TestISO3166Constructor(t *testing.T) {
	v := terms.NewISO3166()
	if v == nil {
		t.Errorf("terms.NewISO3166 must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed terms.ISO3166 should validate: %s", err)
	}
}

func TestISO3166MarshalUnmarshal(t *testing.T) {
	v := terms.NewISO3166()
	buf, _ := xml.Marshal(v)
	v2 := terms.NewISO3166()
	xml.Unmarshal(buf, v2)
}
