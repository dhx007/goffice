package chart

import (
	"github.com/dhx007/goffice"
	"github.com/dhx007/goffice/drawing"
	"github.com/dhx007/goffice/schema/soo/dml"
	crt "github.com/dhx007/goffice/schema/soo/dml/chart"
)

// SurfaceChart is a 3D surface chart, viewed from the top-down.
type SurfaceChart struct {
	chartBase
	x *crt.CT_SurfaceChart
}

// X returns the inner wrapped XML type.
func (c SurfaceChart) X() *crt.CT_SurfaceChart {
	return c.x
}

func (c SurfaceChart) InitializeDefaults() {
	c.x.Wireframe = crt.NewCT_Boolean()
	c.x.Wireframe.ValAttr = goffice.Bool(false)

	c.x.BandFmts = crt.NewCT_BandFmts()
	for i := 0; i < 15; i++ {
		bfmt := crt.NewCT_BandFmt()
		bfmt.Idx.ValAttr = uint32(i)
		bfmt.SpPr = dml.NewCT_ShapeProperties()

		sp := drawing.MakeShapeProperties(bfmt.SpPr)
		sp.SetSolidFill(c.nextColor(i))
		c.x.BandFmts.BandFmt = append(c.x.BandFmts.BandFmt, bfmt)
	}
}

// AddSeries adds a default series to a Surface chart.
func (c SurfaceChart) AddSeries() SurfaceChartSeries {
	color := c.nextColor(len(c.x.Ser))
	ser := crt.NewCT_SurfaceSer()
	c.x.Ser = append(c.x.Ser, ser)
	ser.Idx.ValAttr = uint32(len(c.x.Ser) - 1)
	ser.Order.ValAttr = uint32(len(c.x.Ser) - 1)

	ls := SurfaceChartSeries{ser}
	ls.InitializeDefaults()
	ls.Properties().LineProperties().SetSolidFill(color)
	return ls
}

// AddAxis adds an axis to a Surface chart.
func (c SurfaceChart) AddAxis(axis Axis) {
	axisID := crt.NewCT_UnsignedInt()
	axisID.ValAttr = axis.AxisID()
	c.x.AxId = append(c.x.AxId, axisID)
}
