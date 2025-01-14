package dml

import (
	"encoding/xml"
	"fmt"

	"github.com/dhx007/goffice"
)

type CT_EffectContainer struct {
	TypeAttr     ST_EffectContainerType
	NameAttr     *string
	Cont         *CT_EffectContainer
	Effect       *CT_EffectReference
	AlphaBiLevel *CT_AlphaBiLevelEffect
	AlphaCeiling *CT_AlphaCeilingEffect
	AlphaFloor   *CT_AlphaFloorEffect
	AlphaInv     *CT_AlphaInverseEffect
	AlphaMod     *CT_AlphaModulateEffect
	AlphaModFix  *CT_AlphaModulateFixedEffect
	AlphaOutset  *CT_AlphaOutsetEffect
	AlphaRepl    *CT_AlphaReplaceEffect
	BiLevel      *CT_BiLevelEffect
	Blend        *CT_BlendEffect
	Blur         *CT_BlurEffect
	ClrChange    *CT_ColorChangeEffect
	ClrRepl      *CT_ColorReplaceEffect
	Duotone      *CT_DuotoneEffect
	Fill         *CT_FillEffect
	FillOverlay  *CT_FillOverlayEffect
	Glow         *CT_GlowEffect
	Grayscl      *CT_GrayscaleEffect
	Hsl          *CT_HSLEffect
	InnerShdw    *CT_InnerShadowEffect
	Lum          *CT_LuminanceEffect
	OuterShdw    *CT_OuterShadowEffect
	PrstShdw     *CT_PresetShadowEffect
	Reflection   *CT_ReflectionEffect
	RelOff       *CT_RelativeOffsetEffect
	SoftEdge     *CT_SoftEdgesEffect
	Tint         *CT_TintEffect
	Xfrm         *CT_TransformEffect
}

func NewCT_EffectContainer() *CT_EffectContainer {
	ret := &CT_EffectContainer{}
	return ret
}

func (m *CT_EffectContainer) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.TypeAttr != ST_EffectContainerTypeUnset {
		attr, err := m.TypeAttr.MarshalXMLAttr(xml.Name{Local: "type"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.NameAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "name"},
			Value: fmt.Sprintf("%v", *m.NameAttr)})
	}
	e.EncodeToken(start)
	if m.Cont != nil {
		secont := xml.StartElement{Name: xml.Name{Local: "a:cont"}}
		e.EncodeElement(m.Cont, secont)
	}
	if m.Effect != nil {
		seeffect := xml.StartElement{Name: xml.Name{Local: "a:effect"}}
		e.EncodeElement(m.Effect, seeffect)
	}
	if m.AlphaBiLevel != nil {
		sealphaBiLevel := xml.StartElement{Name: xml.Name{Local: "a:alphaBiLevel"}}
		e.EncodeElement(m.AlphaBiLevel, sealphaBiLevel)
	}
	if m.AlphaCeiling != nil {
		sealphaCeiling := xml.StartElement{Name: xml.Name{Local: "a:alphaCeiling"}}
		e.EncodeElement(m.AlphaCeiling, sealphaCeiling)
	}
	if m.AlphaFloor != nil {
		sealphaFloor := xml.StartElement{Name: xml.Name{Local: "a:alphaFloor"}}
		e.EncodeElement(m.AlphaFloor, sealphaFloor)
	}
	if m.AlphaInv != nil {
		sealphaInv := xml.StartElement{Name: xml.Name{Local: "a:alphaInv"}}
		e.EncodeElement(m.AlphaInv, sealphaInv)
	}
	if m.AlphaMod != nil {
		sealphaMod := xml.StartElement{Name: xml.Name{Local: "a:alphaMod"}}
		e.EncodeElement(m.AlphaMod, sealphaMod)
	}
	if m.AlphaModFix != nil {
		sealphaModFix := xml.StartElement{Name: xml.Name{Local: "a:alphaModFix"}}
		e.EncodeElement(m.AlphaModFix, sealphaModFix)
	}
	if m.AlphaOutset != nil {
		sealphaOutset := xml.StartElement{Name: xml.Name{Local: "a:alphaOutset"}}
		e.EncodeElement(m.AlphaOutset, sealphaOutset)
	}
	if m.AlphaRepl != nil {
		sealphaRepl := xml.StartElement{Name: xml.Name{Local: "a:alphaRepl"}}
		e.EncodeElement(m.AlphaRepl, sealphaRepl)
	}
	if m.BiLevel != nil {
		sebiLevel := xml.StartElement{Name: xml.Name{Local: "a:biLevel"}}
		e.EncodeElement(m.BiLevel, sebiLevel)
	}
	if m.Blend != nil {
		seblend := xml.StartElement{Name: xml.Name{Local: "a:blend"}}
		e.EncodeElement(m.Blend, seblend)
	}
	if m.Blur != nil {
		seblur := xml.StartElement{Name: xml.Name{Local: "a:blur"}}
		e.EncodeElement(m.Blur, seblur)
	}
	if m.ClrChange != nil {
		seclrChange := xml.StartElement{Name: xml.Name{Local: "a:clrChange"}}
		e.EncodeElement(m.ClrChange, seclrChange)
	}
	if m.ClrRepl != nil {
		seclrRepl := xml.StartElement{Name: xml.Name{Local: "a:clrRepl"}}
		e.EncodeElement(m.ClrRepl, seclrRepl)
	}
	if m.Duotone != nil {
		seduotone := xml.StartElement{Name: xml.Name{Local: "a:duotone"}}
		e.EncodeElement(m.Duotone, seduotone)
	}
	if m.Fill != nil {
		sefill := xml.StartElement{Name: xml.Name{Local: "a:fill"}}
		e.EncodeElement(m.Fill, sefill)
	}
	if m.FillOverlay != nil {
		sefillOverlay := xml.StartElement{Name: xml.Name{Local: "a:fillOverlay"}}
		e.EncodeElement(m.FillOverlay, sefillOverlay)
	}
	if m.Glow != nil {
		seglow := xml.StartElement{Name: xml.Name{Local: "a:glow"}}
		e.EncodeElement(m.Glow, seglow)
	}
	if m.Grayscl != nil {
		segrayscl := xml.StartElement{Name: xml.Name{Local: "a:grayscl"}}
		e.EncodeElement(m.Grayscl, segrayscl)
	}
	if m.Hsl != nil {
		sehsl := xml.StartElement{Name: xml.Name{Local: "a:hsl"}}
		e.EncodeElement(m.Hsl, sehsl)
	}
	if m.InnerShdw != nil {
		seinnerShdw := xml.StartElement{Name: xml.Name{Local: "a:innerShdw"}}
		e.EncodeElement(m.InnerShdw, seinnerShdw)
	}
	if m.Lum != nil {
		selum := xml.StartElement{Name: xml.Name{Local: "a:lum"}}
		e.EncodeElement(m.Lum, selum)
	}
	if m.OuterShdw != nil {
		seouterShdw := xml.StartElement{Name: xml.Name{Local: "a:outerShdw"}}
		e.EncodeElement(m.OuterShdw, seouterShdw)
	}
	if m.PrstShdw != nil {
		seprstShdw := xml.StartElement{Name: xml.Name{Local: "a:prstShdw"}}
		e.EncodeElement(m.PrstShdw, seprstShdw)
	}
	if m.Reflection != nil {
		sereflection := xml.StartElement{Name: xml.Name{Local: "a:reflection"}}
		e.EncodeElement(m.Reflection, sereflection)
	}
	if m.RelOff != nil {
		serelOff := xml.StartElement{Name: xml.Name{Local: "a:relOff"}}
		e.EncodeElement(m.RelOff, serelOff)
	}
	if m.SoftEdge != nil {
		sesoftEdge := xml.StartElement{Name: xml.Name{Local: "a:softEdge"}}
		e.EncodeElement(m.SoftEdge, sesoftEdge)
	}
	if m.Tint != nil {
		setint := xml.StartElement{Name: xml.Name{Local: "a:tint"}}
		e.EncodeElement(m.Tint, setint)
	}
	if m.Xfrm != nil {
		sexfrm := xml.StartElement{Name: xml.Name{Local: "a:xfrm"}}
		e.EncodeElement(m.Xfrm, sexfrm)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_EffectContainer) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "type" {
			m.TypeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = &parsed
			continue
		}
	}
lCT_EffectContainer:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "cont"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "cont"}:
				m.Cont = NewCT_EffectContainer()
				if err := d.DecodeElement(m.Cont, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "effect"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "effect"}:
				m.Effect = NewCT_EffectReference()
				if err := d.DecodeElement(m.Effect, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "alphaBiLevel"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "alphaBiLevel"}:
				m.AlphaBiLevel = NewCT_AlphaBiLevelEffect()
				if err := d.DecodeElement(m.AlphaBiLevel, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "alphaCeiling"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "alphaCeiling"}:
				m.AlphaCeiling = NewCT_AlphaCeilingEffect()
				if err := d.DecodeElement(m.AlphaCeiling, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "alphaFloor"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "alphaFloor"}:
				m.AlphaFloor = NewCT_AlphaFloorEffect()
				if err := d.DecodeElement(m.AlphaFloor, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "alphaInv"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "alphaInv"}:
				m.AlphaInv = NewCT_AlphaInverseEffect()
				if err := d.DecodeElement(m.AlphaInv, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "alphaMod"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "alphaMod"}:
				m.AlphaMod = NewCT_AlphaModulateEffect()
				if err := d.DecodeElement(m.AlphaMod, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "alphaModFix"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "alphaModFix"}:
				m.AlphaModFix = NewCT_AlphaModulateFixedEffect()
				if err := d.DecodeElement(m.AlphaModFix, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "alphaOutset"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "alphaOutset"}:
				m.AlphaOutset = NewCT_AlphaOutsetEffect()
				if err := d.DecodeElement(m.AlphaOutset, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "alphaRepl"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "alphaRepl"}:
				m.AlphaRepl = NewCT_AlphaReplaceEffect()
				if err := d.DecodeElement(m.AlphaRepl, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "biLevel"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "biLevel"}:
				m.BiLevel = NewCT_BiLevelEffect()
				if err := d.DecodeElement(m.BiLevel, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "blend"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "blend"}:
				m.Blend = NewCT_BlendEffect()
				if err := d.DecodeElement(m.Blend, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "blur"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "blur"}:
				m.Blur = NewCT_BlurEffect()
				if err := d.DecodeElement(m.Blur, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "clrChange"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "clrChange"}:
				m.ClrChange = NewCT_ColorChangeEffect()
				if err := d.DecodeElement(m.ClrChange, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "clrRepl"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "clrRepl"}:
				m.ClrRepl = NewCT_ColorReplaceEffect()
				if err := d.DecodeElement(m.ClrRepl, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "duotone"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "duotone"}:
				m.Duotone = NewCT_DuotoneEffect()
				if err := d.DecodeElement(m.Duotone, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "fill"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "fill"}:
				m.Fill = NewCT_FillEffect()
				if err := d.DecodeElement(m.Fill, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "fillOverlay"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "fillOverlay"}:
				m.FillOverlay = NewCT_FillOverlayEffect()
				if err := d.DecodeElement(m.FillOverlay, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "glow"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "glow"}:
				m.Glow = NewCT_GlowEffect()
				if err := d.DecodeElement(m.Glow, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "grayscl"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "grayscl"}:
				m.Grayscl = NewCT_GrayscaleEffect()
				if err := d.DecodeElement(m.Grayscl, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "hsl"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "hsl"}:
				m.Hsl = NewCT_HSLEffect()
				if err := d.DecodeElement(m.Hsl, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "innerShdw"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "innerShdw"}:
				m.InnerShdw = NewCT_InnerShadowEffect()
				if err := d.DecodeElement(m.InnerShdw, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "lum"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "lum"}:
				m.Lum = NewCT_LuminanceEffect()
				if err := d.DecodeElement(m.Lum, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "outerShdw"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "outerShdw"}:
				m.OuterShdw = NewCT_OuterShadowEffect()
				if err := d.DecodeElement(m.OuterShdw, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "prstShdw"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "prstShdw"}:
				m.PrstShdw = NewCT_PresetShadowEffect()
				if err := d.DecodeElement(m.PrstShdw, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "reflection"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "reflection"}:
				m.Reflection = NewCT_ReflectionEffect()
				if err := d.DecodeElement(m.Reflection, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "relOff"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "relOff"}:
				m.RelOff = NewCT_RelativeOffsetEffect()
				if err := d.DecodeElement(m.RelOff, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "softEdge"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "softEdge"}:
				m.SoftEdge = NewCT_SoftEdgesEffect()
				if err := d.DecodeElement(m.SoftEdge, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "tint"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "tint"}:
				m.Tint = NewCT_TintEffect()
				if err := d.DecodeElement(m.Tint, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "xfrm"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "xfrm"}:
				m.Xfrm = NewCT_TransformEffect()
				if err := d.DecodeElement(m.Xfrm, &el); err != nil {
					return err
				}
			default:
				goffice.Log("skipping unsupported element on CT_EffectContainer %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_EffectContainer
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_EffectContainer and its children
func (m *CT_EffectContainer) Validate() error {
	return m.ValidateWithPath("CT_EffectContainer")
}

// ValidateWithPath validates the CT_EffectContainer and its children, prefixing error messages with path
func (m *CT_EffectContainer) ValidateWithPath(path string) error {
	if err := m.TypeAttr.ValidateWithPath(path + "/TypeAttr"); err != nil {
		return err
	}
	if m.Cont != nil {
		if err := m.Cont.ValidateWithPath(path + "/Cont"); err != nil {
			return err
		}
	}
	if m.Effect != nil {
		if err := m.Effect.ValidateWithPath(path + "/Effect"); err != nil {
			return err
		}
	}
	if m.AlphaBiLevel != nil {
		if err := m.AlphaBiLevel.ValidateWithPath(path + "/AlphaBiLevel"); err != nil {
			return err
		}
	}
	if m.AlphaCeiling != nil {
		if err := m.AlphaCeiling.ValidateWithPath(path + "/AlphaCeiling"); err != nil {
			return err
		}
	}
	if m.AlphaFloor != nil {
		if err := m.AlphaFloor.ValidateWithPath(path + "/AlphaFloor"); err != nil {
			return err
		}
	}
	if m.AlphaInv != nil {
		if err := m.AlphaInv.ValidateWithPath(path + "/AlphaInv"); err != nil {
			return err
		}
	}
	if m.AlphaMod != nil {
		if err := m.AlphaMod.ValidateWithPath(path + "/AlphaMod"); err != nil {
			return err
		}
	}
	if m.AlphaModFix != nil {
		if err := m.AlphaModFix.ValidateWithPath(path + "/AlphaModFix"); err != nil {
			return err
		}
	}
	if m.AlphaOutset != nil {
		if err := m.AlphaOutset.ValidateWithPath(path + "/AlphaOutset"); err != nil {
			return err
		}
	}
	if m.AlphaRepl != nil {
		if err := m.AlphaRepl.ValidateWithPath(path + "/AlphaRepl"); err != nil {
			return err
		}
	}
	if m.BiLevel != nil {
		if err := m.BiLevel.ValidateWithPath(path + "/BiLevel"); err != nil {
			return err
		}
	}
	if m.Blend != nil {
		if err := m.Blend.ValidateWithPath(path + "/Blend"); err != nil {
			return err
		}
	}
	if m.Blur != nil {
		if err := m.Blur.ValidateWithPath(path + "/Blur"); err != nil {
			return err
		}
	}
	if m.ClrChange != nil {
		if err := m.ClrChange.ValidateWithPath(path + "/ClrChange"); err != nil {
			return err
		}
	}
	if m.ClrRepl != nil {
		if err := m.ClrRepl.ValidateWithPath(path + "/ClrRepl"); err != nil {
			return err
		}
	}
	if m.Duotone != nil {
		if err := m.Duotone.ValidateWithPath(path + "/Duotone"); err != nil {
			return err
		}
	}
	if m.Fill != nil {
		if err := m.Fill.ValidateWithPath(path + "/Fill"); err != nil {
			return err
		}
	}
	if m.FillOverlay != nil {
		if err := m.FillOverlay.ValidateWithPath(path + "/FillOverlay"); err != nil {
			return err
		}
	}
	if m.Glow != nil {
		if err := m.Glow.ValidateWithPath(path + "/Glow"); err != nil {
			return err
		}
	}
	if m.Grayscl != nil {
		if err := m.Grayscl.ValidateWithPath(path + "/Grayscl"); err != nil {
			return err
		}
	}
	if m.Hsl != nil {
		if err := m.Hsl.ValidateWithPath(path + "/Hsl"); err != nil {
			return err
		}
	}
	if m.InnerShdw != nil {
		if err := m.InnerShdw.ValidateWithPath(path + "/InnerShdw"); err != nil {
			return err
		}
	}
	if m.Lum != nil {
		if err := m.Lum.ValidateWithPath(path + "/Lum"); err != nil {
			return err
		}
	}
	if m.OuterShdw != nil {
		if err := m.OuterShdw.ValidateWithPath(path + "/OuterShdw"); err != nil {
			return err
		}
	}
	if m.PrstShdw != nil {
		if err := m.PrstShdw.ValidateWithPath(path + "/PrstShdw"); err != nil {
			return err
		}
	}
	if m.Reflection != nil {
		if err := m.Reflection.ValidateWithPath(path + "/Reflection"); err != nil {
			return err
		}
	}
	if m.RelOff != nil {
		if err := m.RelOff.ValidateWithPath(path + "/RelOff"); err != nil {
			return err
		}
	}
	if m.SoftEdge != nil {
		if err := m.SoftEdge.ValidateWithPath(path + "/SoftEdge"); err != nil {
			return err
		}
	}
	if m.Tint != nil {
		if err := m.Tint.ValidateWithPath(path + "/Tint"); err != nil {
			return err
		}
	}
	if m.Xfrm != nil {
		if err := m.Xfrm.ValidateWithPath(path + "/Xfrm"); err != nil {
			return err
		}
	}
	return nil
}
