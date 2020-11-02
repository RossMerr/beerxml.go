package beerXML

import (
	"encoding/json"
	"encoding/xml"
	"strconv"
	"strings"
)

var _ xml.Unmarshaler = (*Style)(nil)

type Style struct {
	Name           string  `xml:"NAME" json:"name,omitempty"`
	Version        int32   `xml:"VERSION" json:"version,omitempty"`
	Category       string  `xml:"CATEGORY" json:"category,omitempty"`
	Categorynumber string  `xml:"CATEGORY_NUMBER" json:"category_number,omitempty"`
	Styleletter    string  `xml:"STYLE_LETTER" json:"style_letter,omitempty"`
	Styleguide     string  `xml:"STYLE_GUIDE" json:"style_guide,omitempty"`
	Type           string  `xml:"TYPE" json:"type,omitempty"`
	OGMIN          float64 `xml:"OG_MIN" json:"og_min,omitempty"`
	OGMAX          float64 `xml:"OG_MAX" json:"og_max,omitempty"`
	FGMIN          float64 `xml:"FG_MIN" json:"fg_min,omitempty"`
	FGMAX          float64 `xml:"FG_MAX" json:"fg_max,omitempty"`
	IBUMIN         float64 `xml:"IBU_MIN" json:"ibu_min,omitempty"`
	IBUMAX         float64 `xml:"IBU_MAX" json:"ibu_max,omitempty"`
	Colormin       float64 `xml:"COLOR_MIN" json:"color_min,omitempty"`
	Colormax       float64 `xml:"COLOR_MAX" json:"color_max,omitempty"`
	Carbmin        float64 `xml:"CARB_MIN" json:"carb_min,omitempty"`
	Carbmax        float64 `xml:"CARB_MAX" json:"carb_max,omitempty"`
	ABVMAX         float64 `xml:"ABV_MAX" json:"abv_max,omitempty"`
	ABVMIN         float64 `xml:"ABV_MIN" json:"abv_min,omitempty"`
	Notes          string  `xml:"NOTES" json:"notes,omitempty"`
	Profile        string  `xml:"PROFILE" json:"profile,omitempty"`
	Ingredients    string  `xml:"INGREDIENTS" json:"ingredients,omitempty"`
	Examples       string  `xml:"EXAMPLES" json:"examples,omitempty"`

	ABVRange        string `xml:"ABV_RANGE" json:"abv_range,omitempty"`
	CarbRange       string `xml:"CARB_RANGE" json:"carb_range,omitempty"`
	ColorRange      string `xml:"COLOR_RANGE" json:"color_range,omitempty"`
	DisplayColorMax string `xml:"DISPLAY_COLOR_MAX" json:"display_color_max,omitempty"`
	DisplayColorMin string `xml:"DISPLAY_COLOR_MIN" json:"display_color_min,omitempty"`
	DisplayFGMax    string `xml:"DISPLAY_FG_MAX" json:"display_fg_max,omitempty"`
	DisplayFGMin    string `xml:"DISPLAY_FG_MIN" json:"display_fg_min,omitempty"`
	DisplayOGMax    string `xml:"DISPLAY_OG_MAX" json:"display_og_max,omitempty"`
	DisplayOGMin    string `xml:"DISPLAY_OG_MIN" json:"display_og_min,omitempty"`
	FGRange         string `xml:"FG_RANGE" json:"fg_range,omitempty"`
	IBURange        string `xml:"IBU_RANGE" json:"ibu_range,omitempty"`
	OGRange         string `xml:"OG_RANGE" json:"og_range,omitempty"`
}

func (a *Style) MarshalJSON() ([]byte, error) {

	type Alias Style
	t := func() int32 {
		if t, ok := Style_StyleType_value[strings.ToUpper(a.Type)]; ok {
			return t
		}
		return int32(Style_STYLE_NONE)
	}()

	return json.Marshal(&struct {
		Type int32 `json:"type,omitempty"`
		*Alias
	}{
		Type:  t,
		Alias: (*Alias)(a),
	})
}

func (a *Style) UnmarshalJSON(b []byte) error {
	return nil
}

func (a *Style) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Alias Style
	vX := struct {
		Carbmin string `xml:"CARB_MIN"`
		Carbmax string `xml:"CARB_MAX"`
		*Alias
	}{
		Alias: (*Alias)(a),
	}

	err := d.DecodeElement(&vX, &start)

	a.Carbmax = toFloat(vX.Carbmax)
	a.Carbmin = toFloat(vX.Carbmin)

	return err
}

func toFloat(s string) float64 {
	str := reg.FindString(s)
	f, _ := strconv.ParseFloat(str, 64)
	return f
}

type Style_StyleType int32

const (
	// None
	Style_STYLE_NONE Style_StyleType = 0
	// Lager
	Style_LAGER Style_StyleType = 1
	// Ale
	Style_ALE Style_StyleType = 2
	// Mead
	Style_MEAD Style_StyleType = 3
	// Wheat
	Style_WHEAT Style_StyleType = 4
	// Mixed
	Style_MIXED Style_StyleType = 5
	// Cider
	Style_CIDER Style_StyleType = 6
)

var Style_StyleType_name = map[int32]string{
	0: "STYLE_NONE",
	1: "LAGER",
	2: "ALE",
	3: "MEAD",
	4: "WHEAT",
	5: "MIXED",
	6: "CIDER",
}

var Style_StyleType_value = map[string]int32{
	"STYLE_NONE": 0,
	"LAGER":      1,
	"ALE":        2,
	"MEAD":       3,
	"WHEAT":      4,
	"MIXED":      5,
	"CIDER":      6,
}
