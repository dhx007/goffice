package spreadsheet

import (
	"github.com/dhx007/goffice"
	"github.com/dhx007/goffice/schema/soo/sml"
)

type Fills struct {
	x *sml.CT_Fills
}

func NewFills() Fills {
	return Fills{sml.NewCT_Fills()}
}

func (f Fills) X() *sml.CT_Fills {
	return f.x
}

func (f Fills) AddFill() Fill {
	fill := sml.NewCT_Fill()
	f.x.Fill = append(f.x.Fill, fill)
	f.x.CountAttr = goffice.Uint32(uint32(len(f.x.Fill)))
	return Fill{fill, f.x}
}
