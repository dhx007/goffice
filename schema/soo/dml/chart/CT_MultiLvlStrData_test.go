// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package chart_test

import (
	"encoding/xml"
	"testing"

	"goffice/schema/soo/dml/chart"
)

func TestCT_MultiLvlStrDataConstructor(t *testing.T) {
	v := chart.NewCT_MultiLvlStrData()
	if v == nil {
		t.Errorf("chart.NewCT_MultiLvlStrData must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_MultiLvlStrData should validate: %s", err)
	}
}

func TestCT_MultiLvlStrDataMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_MultiLvlStrData()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_MultiLvlStrData()
	xml.Unmarshal(buf, v2)
}
