package document

import (
	"github.com/dhx007/goffice"
	"github.com/dhx007/goffice/measurement"
	"github.com/dhx007/goffice/schema/soo/ofc/sharedTypes"
	"github.com/dhx007/goffice/schema/soo/wml"
)

// RowProperties are the properties for a row within a table
type RowProperties struct {
	x *wml.CT_TrPr
}

// SetHeight allows controlling the height of a row within a table.
func (r RowProperties) SetHeight(ht measurement.Distance, rule wml.ST_HeightRule) {
	if rule == wml.ST_HeightRuleUnset {
		r.x.TrHeight = nil
	} else {
		htv := wml.NewCT_Height()
		htv.HRuleAttr = rule
		htv.ValAttr = &sharedTypes.ST_TwipsMeasure{}
		htv.ValAttr.ST_UnsignedDecimalNumber = goffice.Uint64(uint64(ht / measurement.Twips))
		r.x.TrHeight = []*wml.CT_Height{htv}
	}
}
