package beerXML

import (
	"encoding/json"
	"strconv"
	"strings"
)

type Style struct {
	Name           string  `xml:"NAME" json:"name,omitempty"`
	Version        int32   `xml:"VERSION" json:"version,omitempty"`
	Category       string  `xml:"CATEGORY" json:"category,omitempty"`
	Categorynumber string  `xml:"CATEGORY_NUMBER" json:"category_number,omitempty"`
	Styleletter    string  `xml:"STYLE_LETTER" json:"style_letter,omitempty"`
	Styleguide     string  `xml:"STYLE_GUIDE" json:"style_guide,omitempty"`
	Type           string  `xml:"TYPE" json:"type,omitempty"`
	OGMIN          float32 `xml:"OG_MIN" json:"og_min,omitempty"`
	OGMAX          float32 `xml:"OG_MAX" json:"og_max,omitempty"`
	FGMIN          float32 `xml:"FG_MIN" json:"fg_min,omitempty"`
	FGMAX          float32 `xml:"FG_MAX" json:"fg_max,omitempty"`
	IBUMIN         float32 `xml:"IBU_MIN" json:"ibu_min,omitempty"`
	IBUMAX         float32 `xml:"IBU_MAX" json:"ibu_max,omitempty"`
	Colormin       float32 `xml:"COLOR_MIN" json:"color_min,omitempty"`
	Colormax       float32 `xml:"COLOR_MAX" json:"color_max,omitempty"`
	Carbmin        string  `xml:"CARB_MIN" json:"carb_min,omitempty"`
	Carbmax        string  `xml:"CARB_MAX" json:"carb_max,omitempty"`
	ABVMAX         float32 `xml:"ABV_MAX" json:"abv_max,omitempty"`
	ABVMIN         float32 `xml:"ABV_MIN" json:"abv_min,omitempty"`
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

func (a Style) MarshalJSON() ([]byte, error) {

	type Alias Style
	t := func() int32 {
		if t, ok := Style_StyleType_value[strings.ToUpper(a.Type)]; ok {
			return t
		}
		return int32(Style_STYLE_NONE)
	}()

	return json.Marshal(&struct {
		Type    int32   `json:"type,omitempty"`
		CarbMin float32 `json:"carb_min,omitempty"`
		CarbMax float32 `json:"carb_max,omitempty"`
		*Alias
	}{
		Type: t,
		CarbMin: func() float32 {
			str := string(reg.Find([]byte(a.Carbmax)))
			f, _ := strconv.ParseFloat(str, 64)
			return float32(f)
		}(),
		CarbMax: func() float32 {
			str := string(reg.Find([]byte(a.Carbmax)))
			f, _ := strconv.ParseFloat(str, 64)
			return float32(f)
		}(),
		Alias: (*Alias)(&a),
	})
}

func (a *Style) UnmarshalJSON(b []byte) error {
	return nil
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
