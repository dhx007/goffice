package chartDrawing

import "github.com/dhx007/goffice"

func b2i(b bool) uint8 {
	if b {
		return 1
	}
	return 0
}

// init registers constructor functions for dynamically creating elements based off the XML namespace and name
func init() {
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", "CT_ShapeNonVisual", NewCT_ShapeNonVisual)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", "CT_Shape", NewCT_Shape)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", "CT_ConnectorNonVisual", NewCT_ConnectorNonVisual)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", "CT_Connector", NewCT_Connector)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", "CT_PictureNonVisual", NewCT_PictureNonVisual)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", "CT_Picture", NewCT_Picture)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", "CT_GraphicFrameNonVisual", NewCT_GraphicFrameNonVisual)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", "CT_GraphicFrame", NewCT_GraphicFrame)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", "CT_GroupShapeNonVisual", NewCT_GroupShapeNonVisual)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", "CT_GroupShape", NewCT_GroupShape)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", "CT_Marker", NewCT_Marker)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", "CT_RelSizeAnchor", NewCT_RelSizeAnchor)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", "CT_AbsSizeAnchor", NewCT_AbsSizeAnchor)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", "CT_Drawing", NewCT_Drawing)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", "EG_ObjectChoices", NewEG_ObjectChoices)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", "EG_Anchor", NewEG_Anchor)
}
