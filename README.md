**goffice** is a library for creation of Office Open XML documents (.docx, .xlsx
and .pptx).  It's goal is to be the most compatible and highest performance Go
library for creation and editing of docx/xlsx/pptx files.

Requires **go1.8+**, builds are tested with 1.8, 1.9 and tip.


## Status ##

- Documents (docx) [Word]
	- Read/Write/Edit
	- Formatting
	- Images
	- Tables
- Spreadsheets (xlsx) [Excel]
 	- Read/Write/Edit
 	- Cell formatting including conditional formatting
	- Cell validation (drop down combobox, rules, etc.)
    - Retrieve cell values as formatted by Excel (e.g. retrieve a date or number as displayed in Excel)
 	- Formula Evaluation (100+ functions supported currently, more will be added as required)
 	- Embedded Images
 	- All chart types
- PowerPoint (pptx) [PowerPoint]
	- Creation from templates
	- Textboxes/shapes


## Installation ##
    
    go get github.com/unidoc/goffice/
    go build -i github.com/unidoc/goffice/...

## Raw Types ##

The OOXML specification is large and creating a friendly API to cover the entire
specification is a very time consuming endeavor.  This library attempts to
provide an easy to use API for common use cases in creating OOXML documents
while allowing users to fall back to raw document manipulation should the
library's API not cover a specific use case.

The raw XML based types reside in the ```schema/``` directory. These types are
accessible from the wrapper types via a ```X()``` method that returns the raw
type. 

For example, the library currently doesn't have an API for setting a document
background color. However it's easy to do manually via editing the
```CT_Background``` element of the document.

    dox := document.New()
    doc.X().Background = wordprocessingml.NewCT_Background()
	doc.X().Background.ColorAttr = &wordprocessingml.ST_HexColor{}
	doc.X().Background.ColorAttr.ST_HexColorRGB = color.RGB(50, 50, 50).AsRGBString()

