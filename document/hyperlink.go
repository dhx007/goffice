package document

import (
	"github.com/dhx007/goffice"
	"github.com/dhx007/goffice/common"
	"github.com/dhx007/goffice/schema/soo/wml"
)

// HyperLink is a link within a document.
type HyperLink struct {
	d *Document
	x *wml.CT_Hyperlink
}

// X returns the inner wrapped XML type.
func (h HyperLink) X() *wml.CT_Hyperlink {
	return h.x
}

// SetTargetByRef sets the URL target of the hyperlink and is more efficient if a link
// destination will be used many times.
func (h HyperLink) SetTargetByRef(link common.Hyperlink) {
	h.x.IdAttr = goffice.String(common.Relationship(link).ID())
	h.x.AnchorAttr = nil
}

// SetTarget sets the URL target of the hyperlink.
func (h HyperLink) SetTarget(url string) {
	rel := h.d.AddHyperlink(url)
	h.x.IdAttr = goffice.String(common.Relationship(rel).ID())
	h.x.AnchorAttr = nil
}

// SetTargetBookmark sets the bookmark target of the hyperlink.
func (h HyperLink) SetTargetBookmark(bm Bookmark) {
	h.x.AnchorAttr = goffice.String(bm.Name())
	h.x.IdAttr = nil
}

// SetToolTip sets the tooltip text for a hyperlink.
func (h HyperLink) SetToolTip(text string) {
	if text == "" {
		h.x.TooltipAttr = nil
	} else {
		h.x.TooltipAttr = goffice.String(text)
	}
}

// AddRun adds a run of text to a hyperlink. This is the text that will be linked.
func (h HyperLink) AddRun() Run {
	rc := wml.NewEG_ContentRunContent()
	h.x.EG_ContentRunContent = append(h.x.EG_ContentRunContent, rc)
	r := wml.NewCT_R()
	rc.R = r
	return Run{h.d, r}
}
