package beerXML

import (
	"encoding/xml"
	"strings"
)

// Encloses a set of one or more Yeast records.
type Yeasts struct {
	Yeast []Yeast `xml:"YEAST" json:"yeast,omitempty"`
}

// The term "yeast" encompasses all yeasts, including dry yeast, liquid yeast and yeast starters.
type Yeast struct {
	// Name of the yeast.
	Name           string  `xml:"NAME" json:"name,omitempty"`
	// Version of the standard.  Should be “1” for this version.
	Version        int32   `xml:"VERSION" json:"version,omitempty"`
	// May be “Ale”, “Lager”, “Wheat”, “Wine” or “Champagne”
	Type           Culture_CultureType  `xml:"TYPE" json:"type,omitempty"`
	// May be “Liquid”, “Dry”, “Slant” or “Culture”
	Form           Culture_FormType  `xml:"FORM" json:"form,omitempty"`
	// The amount of yeast, measured in liters.  For a starter this is the size of the starter.
	// If the flag AMOUNT_IS_WEIGHT is set to TRUE then this measurement is in kilograms and not liters.
	Amount         float64 `xml:"AMOUNT" json:"amount,omitempty"`
	// TRUE if the amount measurement is a weight measurement and FALSE if the amount is a volume measurement.
	// Default value (if not present) is assumed to be FALSE – therefore the yeast measurement is a liquid
	// amount by default.
	AmountIsWeight *bool    `xml:"AMOUNT_IS_WEIGHT,omitempty" json:"amount_is_weight,omitempty"`
	// The name of the laboratory that produced the yeast.
	Laboratory     *string  `xml:"LABORATORY,omitempty" json:"laboratory,omitempty"`
	// The manufacturer’s product ID label or number that identifies this particular strain of yeast.
	ProductID      *string  `xml:"PRODUCT_ID,omitempty" json:"product_id,omitempty"`
	// The minimum recommended temperature for fermenting this yeast strain in degrees Celsius.
	MinTemperature *float64 `xml:"MIN_TEMPERATURE,omitempty" json:"min_temperature,omitempty"`
	// The maximum recommended temperature for fermenting this yeast strain in Celsius.
	MaxTemperature *float64 `xml:"MAX_TEMPERATURE,omitempty" json:"max_temperature,omitempty"`
	// May be “Low”, “Medium”, “High” or “Very High”
	Flocculation   *string  `xml:"FLOCCULATION,omitempty" json:"flocculation,omitempty"`
	// Average attenuation for this yeast strain.
	Attenuation    *float64 `xml:"ATTENUATION,omitempty" json:"attenuation,omitempty"`
	// Notes on this yeast strain.  May be a multiline entry.
	Notes          *string  `xml:"NOTES,omitempty" json:"notes,omitempty"`
	// Styles or types of beer this yeast strain is best suited for.
	BestFor        *string  `xml:"BEST_FOR,omitempty" json:"best_for,omitempty"`
	// Number of times this yeast has been reused as a harvested culture.  This number should be zero if this is a
	// product directly from the manufacturer.
	TimesCultured  *float64 `xml:"TIMES_CULTURED,omitempty" json:"times_cultured,omitempty"`
	// Recommended of times this yeast can be reused (recultured from a previous batch)
	MaxReuse       *float64 `xml:"MAX_REUSE,omitempty" json:"max_reuse,omitempty"`
	// Flag denoting that this yeast was added for a secondary (or later) fermentation as opposed to the primary
	// fermentation.  Useful if one uses two or more yeast strains for a single brew (eg: Lambic).
	// Default value is FALSE.
	AddToSecondary *bool    `xml:"ADD_TO_SECONDARY,omitempty" json:"add_to_secondary,omitempty"`

	// Extensions

	// The amount of yeast or starter in this record along with the units formatted for easy display in the
	// current user defined units.  For example “1.5 oz” or “100 g”.
	DisplayAmount  *string  `xml:"DISPLAY_AMOUNT,omitempty" json:"display_amount,omitempty"`
	// Minimum fermentation temperature converted to current user units along with the units.
	// For example “54.0 F” or “24.2 C”
	DispMinTemp    *string  `xml:"DISP_MIN_TEMP,omitempty" json:"disp_min_temp,omitempty"`
	// Maximum fermentation temperature converted to current user units along with the units.
	// For example “54.0 F” or “24.2 C”
	DispMaxTemp    *string  `xml:"DISP_MAX_TEMP,omitempty" json:"disp_max_temp,omitempty"`
	// Amount in inventory for this hop along with the units – for example “10.0 pkgs”
	Inventory      *string  `xml:"INVENTORY,omitempty" json:"inventory,omitempty"`
	// Date sample was last cultured in a neutral date form such as “10 Dec 04”
	CultureDate    *string  `xml:"CULTURE_DATE,omitempty" json:"culture_date,omitempty"`
}


func (a *Yeast) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Alias Yeast
	aux := &struct {
		*Alias
		Attenuation    *string `xml:"ATTENUATION,omitempty" json:"attenuation,omitempty"`
	}{
		Alias: (*Alias)(a),
	}

	err := d.DecodeElement(aux, &start)
	if err != nil {
		return err
	}

	a.Attenuation = percentToFloat(aux.Attenuation)

	return nil
}

func (a Yeast) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type Alias Yeast
	aux := &struct {
		*Alias
		Attenuation    string `xml:"ATTENUATION,omitempty" json:"attenuation,omitempty"`
	}{
		Alias: (*Alias)(&a),
		Attenuation: floatToPercent(a.Attenuation),
	}

	start.Name.Local = strings.ToUpper(start.Name.Local)

	err := e.EncodeElement(aux, start)
	if err != nil {
		return err
	}

	return nil
}


func (a *Culture_CultureType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var use string

	err := d.DecodeElement(&use, &start)
	if err != nil {
		return err
	}

	if value, ok := Culture_CultureType_value[use]; ok {
		*a = Culture_CultureType(value)
	} else {
		*a = CULTURE_CULTURE_NONE
	}

	return nil
}

func (a Culture_CultureType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if name, ok := Culture_CultureType_name[int32(a)]; ok {
		err := e.EncodeElement(name, start)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *Culture_FormType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var use string

	err := d.DecodeElement(&use, &start)
	if err != nil {
		return err
	}

	if value, ok := Culture_FormType_value[use]; ok {
		*a = Culture_FormType(value)
	} else {
		*a = CULTURE_FORM_NONE
	}

	return nil
}

func (a Culture_FormType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if name, ok := Culture_FormType_name[int32(a)]; ok {
		err := e.EncodeElement(name, start)
		if err != nil {
			return err
		}
	}
	return nil
}

type Culture_CultureType int32

const (
	// None
	CULTURE_CULTURE_NONE Culture_CultureType = 0
	// Ale
	CULTURE_ALE Culture_CultureType = 1
	// Lager
	CULTURE_LAGER Culture_CultureType = 2
	// Wheat
	CULTURE_WHEAT Culture_CultureType = 3
	// Wine
	CULTURE_WINE Culture_CultureType = 4
	// Champagne
	CULTURE_CHAMPAGNE Culture_CultureType = 5
	// Bacteria
	CULTURE_BACTERIA Culture_CultureType = 6
	// Brett
	CULTURE_BRETT Culture_CultureType = 7
	// Kveik
	CULTURE_KVEIK Culture_CultureType = 8
	// Lacto
	CULTURE_LACTO Culture_CultureType = 9
	// Malolactic
	CULTURE_MALOLACTIC Culture_CultureType = 10
	// Mixed-Culture
	CULTURE_MIXED_CULTURE Culture_CultureType = 11
	// Other
	CULTURE_OTHER Culture_CultureType = 12
	// Pedio
	CULTURE_PEDIO Culture_CultureType = 13
	// Spontaneous
	CULTURE_SPONTANEOUS Culture_CultureType = 14
)

var Culture_CultureType_name = map[int32]string{
	0:  "CULTURE_NONE",
	1:  "Ale",
	2:  "Lager",
	3:  "Wheat",
	4:  "Wine",
	5:  "Champagne",
	6:  "Bacteria",
	7:  "Brett",
	8:  "Kveik",
	9:  "Lacto",
	10: "Malolactic",
	11: "Mixed Culture",
	12: "Other",
	13: "Pedio",
	14: "Spontaneous",
}

var Culture_CultureType_value = map[string]int32{
	"CULTURE_NONE":  0,
	"Ale":           1,
	"Lager":         2,
	"Wheat":         3,
	"Wine":          4,
	"Champagne":     5,
	"Bacteria":      6,
	"Brett":         7,
	"Kveik":         8,
	"Lacto":         9,
	"Malolactic":    10,
	"Mixed Culture": 11,
	"Other":         12,
	"Pedio":         13,
	"Spontaneous":   14,
}

type Culture_FormType int32

const (
	// None
	CULTURE_FORM_NONE Culture_FormType = 0
	// Liquid
	CULTURE_LIQUID Culture_FormType = 1
	// Dry
	CULTURE_DRY Culture_FormType = 2
	// Slant
	CULTURE_SLANT Culture_FormType = 3
	// Culture
	CULTURE_CULTURE Culture_FormType = 4
	// Dregs
	CULTURE_DREGS Culture_FormType = 5
)

var Culture_FormType_name = map[int32]string{
	0: "FORM_NONE",
	1: "Liquid",
	2: "Dry",
	3: "Slant",
	4: "Culture",
	5: "Dregs",
}

var Culture_FormType_value = map[string]int32{
	"FORM_NONE": 0,
	"Liquid":    1,
	"Dry":       2,
	"Slant":     3,
	"Culture":   4,
	"Dregs":     5,
}

type Culture_FlocculationType int32

const (
	// None
	CULTURE_FLOCCULATION_NONE Culture_FlocculationType = 0
	// Low
	CULTURE_LOW Culture_FlocculationType = 1
	// Medium
	CULTURE_MEDIUM Culture_FlocculationType = 2
	// High
	CULTURE_HIGH Culture_FlocculationType = 3
	// Very High
	CULTURE_VERY_HIGH Culture_FlocculationType = 4
)

var Culture_FlocculationType_name = map[int32]string{
	0: "FLOCCULATION_NONE",
	1: "Low",
	2: "Medium",
	3: "High",
	4: "Very High",
}

var Culture_FlocculationType_value = map[string]int32{
	"FLOCCULATION_NONE": 0,
	"Low":               1,
	"Medium":            2,
	"High":              3,
	"Very High":         4,
}
