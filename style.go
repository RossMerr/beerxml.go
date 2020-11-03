package beerXML

import (
	"encoding/xml"
	"strings"
)

// Encloses a set of one or more Beer Styles
type Styles struct {
	Style []Style `xml:"STYLE" json:"style,omitempty"`
}


// The term "style" encompasses beer styles.  The beer style may be from the BJCP style guide, Australian, UK or
// local style guides.  Generally a recipe is designed to one style.
type Style struct {
	// Name of the style profile – usually this is the specific name of the style – for example
	// “Scottish Wee Heavy Ale” and not the Category which in this case might be “Scottish Ale”
	Name string `xml:"NAME" json:"name,omitempty"`
	// Category that this style belongs to – usually associated with a group of styles such as “English Ales”
	// or “Amercian Lagers”.
	Category string `xml:"CATEGORY" json:"category,omitempty"`
	// Version of the style record.  Should always be “1” for this version of the XML standard.
	Version int32 `xml:"VERSION" json:"version,omitempty"`
	// Number or identifier associated with this style category.  For example in the BJCP style guide,
	// the “American Lager” category has a category number of “1”.
	CategoryNumber string `xml:"CATEGORY_NUMBER" json:"category_number,omitempty"`
	// The specific style number or subcategory letter associated with this particular style.
	// For example in the BJCP style guide, an American Standard Lager would be style letter “A”
	// under the main category.  Letters should be upper case.
	StyleLetter string `xml:"STYLE_LETTER" json:"style_letter,omitempty"`
	// The name of the style guide that this particular style or category belongs to.
	// For example “BJCP” might denote the BJCP style guide, and “AHA” would be used for the AHA style guide.
	StyleGuide string `xml:"STYLE_GUIDE" json:"style_guide,omitempty"`
	// May be “Lager”, “Ale”, “Mead”, “Wheat”, “Mixed” or “Cider”.  Defines the type of beverage
	// associated with this category.
	Type Style_StyleType `xml:"TYPE" json:"type,omitempty"`
	// The minimum specific gravity as measured relative to water.  For example “1.040” might be a reasonable
	// minimum for a Pale Ale.
	OGMin float64 `xml:"OG_MIN" json:"og_min,omitempty"`
	// The maximum specific gravity as measured relative to water.
	OGMax float64 `xml:"OG_MAX" json:"og_max,omitempty"`
	// The minimum final gravity as measured relative to water.
	FGMin float64 `xml:"FG_MIN" json:"fg_min,omitempty"`
	// The maximum final gravity as measured relative to water.
	FGMax float64 `xml:"FG_MAX" json:"fg_max,omitempty"`
	// The recommended minimum bitterness for this style as measured in International Bitterness Units (IBUs)
	IBUMin float64 `xml:"IBU_MIN" json:"ibu_min,omitempty"`
	// The recommended maximum bitterness for this style as measured in International Bitterness Units (IBUs)
	IBUMax float64 `xml:"IBU_MAX" json:"ibu_max,omitempty"`
	// The minimum recommended color in SRM
	ColorMin float64 `xml:"COLOR_MIN" json:"color_min,omitempty"`
	// The maximum recommended color in SRM.
	ColorMax float64 `xml:"COLOR_MAX" json:"color_max,omitempty"`
	// Minimum recommended carbonation for this style in volumes of CO2
	CarbMin *float64 `xml:"CARB_MIN,omitempty" json:"carb_min,omitempty"`
	// The maximum recommended carbonation for this style in volumes of CO2
	CarbMax *float64 `xml:"CARB_MAX,omitempty" json:"carb_max,omitempty"`
	// The minimum recommended alcohol by volume as a percentage.
	ABVMmin *float64 `xml:"ABV_MIN,omitempty" json:"abv_min,omitempty"`
	// The maximum recommended alcohol by volume as a percentage.
	ABVMax *float64 `xml:"ABV_MAX,omitempty" json:"abv_max,omitempty"`
	// Description of the style, history
	Notes *string `xml:"NOTES,omitempty" json:"notes,omitempty"`
	// Flavor and aroma profile for this style
	Profile *string `xml:"PROFILE,omitempty" json:"profile,omitempty"`
	// Suggested ingredients for this style
	Ingredients *string `xml:"INGREDIENTS,omitempty" json:"ingredients,omitempty"`
	// Example beers of this style.
	Examples *string `xml:"EXAMPLES,omitempty" json:"examples,omitempty"`

	// Extensions

	// Original gravity minimum in user defined units such as “1.036 sg”.
	DisplayOGMin *string `xml:"DISPLAY_OG_MIN" json:"display_og_min,omitempty"`
	// Original gravity max in user defined units such as “1.056 sg”
	DisplayOGMax *string `xml:"DISPLAY_OG_MAX" json:"display_og_max,omitempty"`
	// Final gravity minimum in user defined units such as “1.010 sg”.
	DisplayFGMin *string `xml:"DISPLAY_FG_MIN" json:"display_fg_min,omitempty"`
	// Final gravity maximum in user defined units such as “1.019 sg”.
	DisplayFGMax *string `xml:"DISPLAY_FG_MAX" json:"display_fg_max,omitempty"`
	// Minimum color in user defined units such as “30 srm”.
	DisplayColorMin *string `xml:"DISPLAY_COLOR_MIN" json:"display_color_min,omitempty"`
	// Maximum color in user defined units such as “20 srm”
	DisplayColorMax *string `xml:"DISPLAY_COLOR_MAX" json:"display_color_max,omitempty"`
	// Original gravity range for the style such as “1.030-1.040 sg”
	OGRange *string `xml:"OG_RANGE" json:"og_range,omitempty"`
	// Final gravity range such as “1.010-1.015 sg”
	FGRange *string `xml:"FG_RANGE" json:"fg_range,omitempty"`
	// Bitterness range in IBUs such as “10-20 IBU”
	IBURange *string `xml:"IBU_RANGE" json:"ibu_range,omitempty"`
	// Carbonation range in volumes such as “2.0-2.6 vols”
	CarbRange *string `xml:"CARB_RANGE" json:"carb_range,omitempty"`
	// Color range such as “10-20 SRM”
	ColorRange *string `xml:"COLOR_RANGE" json:"color_range,omitempty"`
	// Color range such as “10-20 SRM”
	ABVRange *string `xml:"ABV_RANGE" json:"abv_range,omitempty"`
}


func (a *Style) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Alias Style
	aux := struct {
		CarbMin *string `xml:"CARB_MIN,omitempty" json:"carb_min,omitempty"`
		CarbMax *string `xml:"CARB_MAX,omitempty" json:"carb_max,omitempty"`
		ABVMmin *string `xml:"ABV_MIN,omitempty" json:"abv_min,omitempty"`
		ABVMax  *string `xml:"ABV_MAX,omitempty" json:"abv_max,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(a),
	}

	err := d.DecodeElement(&aux, &start)

	a.CarbMin = volumeToFloat(aux.CarbMin)
	a.CarbMax = volumeToFloat(aux.CarbMax)
	a.ABVMmin = percentToFloat(aux.ABVMmin)
	a.ABVMax = percentToFloat(aux.ABVMax)

	return err
}

func (a Style) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type Alias Style
	aux := &struct {
		*Alias
		CarbMin string `xml:"CARB_MIN,omitempty" json:"carb_min,omitempty"`
		CarbMax string `xml:"CARB_MAX,omitempty" json:"carb_max,omitempty"`
		ABVMmin string `xml:"ABV_MIN,omitempty" json:"abv_min,omitempty"`
		ABVMax  string `xml:"ABV_MAX,omitempty" json:"abv_max,omitempty"`
	}{
		Alias:   (*Alias)(&a),
		CarbMin: floatToVolume(a.CarbMin),
		CarbMax: floatToVolume(a.CarbMax),
		ABVMmin: floatToPercent(a.ABVMmin),
		ABVMax:  floatToPercent(a.ABVMax),
	}

	start.Name.Local = strings.ToUpper(start.Name.Local)
	err := e.EncodeElement(aux, start)
	if err != nil {
		return err
	}

	return nil
}



func (a *Style_StyleType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var use string

	err := d.DecodeElement(&use, &start)
	if err != nil {
		return err
	}

	if value, ok := Style_StyleType_value[use]; ok {
		*a = Style_StyleType(value)
	} else {
		*a = STYLE_STYLE_NONE
	}

	return nil
}

func (a Style_StyleType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if name, ok := Style_StyleType_name[int32(a)]; ok {
		err := e.EncodeElement(name, start)
		if err != nil {
			return err
		}
	}
	return nil
}

type Style_StyleType int32

const (
	// None
	STYLE_STYLE_NONE Style_StyleType = 0
	// Lager
	STYLE_LAGER Style_StyleType = 1
	// Ale
	STYLE_ALE Style_StyleType = 2
	// Mead
	STYLE_MEAD Style_StyleType = 3
	// Wheat
	STYLE_WHEAT Style_StyleType = 4
	// Mixed
	STYLE_MIXED Style_StyleType = 5
	// Cider
	STYLE_CIDER Style_StyleType = 6
)

var Style_StyleType_name = map[int32]string{
	0: "STYLE_NONE",
	1: "Lager",
	2: "Ale",
	3: "Mead",
	4: "Wheat",
	5: "Mixed",
	6: "Cider",
}

var Style_StyleType_value = map[string]int32{
	"STYLE_NONE": 0,
	"Lager":      1,
	"Ale":        2,
	"Mead":       3,
	"Wheat":      4,
	"Mixed":      5,
	"Cider":      6,
}
