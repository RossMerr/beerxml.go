package beerXML

import (
	"encoding/json"
	"strings"
)

type Yeast struct {
	Name           string  `xml:"NAME" json:"name,omitempty"`
	Version        int32   `xml:"VERSION" json:"version,omitempty"`
	Type           string  `xml:"TYPE" json:"type,omitempty"`
	Form           string  `xml:"FORM" json:"form,omitempty"`
	Amount         float32 `xml:"AMOUNT" json:"amount,omitempty"`
	Amountisweight bool    `xml:"AMOUNT_IS_WEIGHT" json:"amount_is_weight,omitempty"`
	Laboratory     string  `xml:"LABORATORY" json:"laboratory,omitempty"`
	Productid      string  `xml:"PRODUCT_ID" json:"product_id,omitempty"`
	Mintemperature float64 `xml:"MIN_TEMPERATURE" json:"min_temperature,omitempty"`
	Maxtemperature float64 `xml:"MAX_TEMPERATURE" json:"max_temperature,omitempty"`
	Flocculation   string  `xml:"FLOCCULATION" json:"flocculation,omitempty"`
	Attenuation    float32 `xml:"ATTENUATION" json:"attenuation,omitempty"`
	Notes          string  `xml:"NOTES" json:"notes,omitempty"`
	Bestfor        string  `xml:"BEST_FOR" json:"best_for,omitempty"`
	Maxreuse       float32 `xml:"MAX_REUSE" json:"max_reuse,omitempty"`
	Timescultured  float32 `xml:"TIMES_CULTURED" json:"times_cultured,omitempty"`
	Addtosecondary bool    `xml:"ADD_TO_SECONDARY" json:"add_to_secondary,omitempty"`
	CultureDate    string  `xml:"CULTURE_DATE" json:"culture_date,omitempty"`
	DisplayAmount  string  `xml:"DISPLAY_AMOUNT" json:"display_amount,omitempty"`
	DispMaxTemp    string  `xml:"DISP_MAX_TEMP" json:"disp_max_temp,omitempty"`
	DispMinTemp    string  `xml:"DISP_MIN_TEMP" json:"disp_min_temp,omitempty"`
	Inventory      string  `xml:"INVENTORY" json:"inventory,omitempty"`
}

func (a Yeast) MarshalJSON() ([]byte, error) {

	type Alias Yeast
	t := func() int32 {
		if t, ok := Culture_CultureType_value[strings.ToUpper(a.Type)]; ok {
			return t
		}
		return int32(Culture_CULTURE_NONE)
	}()

	form := func() int32 {
		if t, ok := Culture_FormType_value[strings.ToUpper(a.Form)]; ok {
			return t
		}
		return int32(Culture_FORM_NONE)
	}()

	flocculation := func() int32 {
		if t, ok := Culture_FlocculationType_value[strings.ToUpper(a.Form)]; ok {
			return t
		}
		return int32(Culture_FLOCCULATION_NONE)
	}()

	return json.Marshal(&struct {
		Type         int32 `json:"type,omitempty"`
		Form         int32 `json:"form,omitempty"`
		Flocculation int32 `json:"flocculation,omitempty"`
		*Alias
	}{
		Type:         t,
		Form:         form,
		Flocculation: flocculation,
		Alias:        (*Alias)(&a),
	})
}

func (a *Yeast) UnmarshalJSON(b []byte) error {
	return nil
}

type Yeasts struct {
	Yeast []Yeast `xml:"YEAST" json:"yeast,omitempty"`
}

func (a Yeasts) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0)
	b = append(b, []byte("[")...)
	if len(a.Yeast) > 0 {
		for _, hop := range a.Yeast {
			bb, err := json.Marshal(hop)
			if err != nil {
				return nil, err
			}

			b = append(b, bb...)
			b = append(b, []byte(",")...)
		}

		// remove the trailing ','
		b = b[:len(b)-1]
	}
	b = append(b, []byte("]")...)

	return b, nil
}

func (a *Yeasts) UnmarshalJSON(b []byte) error {
	return nil
}

type Culture_CultureType int32

const (
	// None
	Culture_CULTURE_NONE Culture_CultureType = 0
	// Ale
	Culture_ALE Culture_CultureType = 1
	// Lager
	Culture_LAGER Culture_CultureType = 2
	// Wheat
	Culture_WHEAT Culture_CultureType = 3
	// Wine
	Culture_WINE Culture_CultureType = 4
	// Champagne
	Culture_CHAMPAGNE Culture_CultureType = 5
	// Bacteria
	Culture_BACTERIA Culture_CultureType = 6
	// Brett
	Culture_BRETT Culture_CultureType = 7
	// Kveik
	Culture_KVEIK Culture_CultureType = 8
	// Lacto
	Culture_LACTO Culture_CultureType = 9
	// Malolactic
	Culture_MALOLACTIC Culture_CultureType = 10
	// Mixed-Culture
	Culture_MIXED_CULTURE Culture_CultureType = 11
	// Other
	Culture_OTHER Culture_CultureType = 12
	// Pedio
	Culture_PEDIO Culture_CultureType = 13
	// Spontaneous
	Culture_SPONTANEOUS Culture_CultureType = 14
)

var Culture_CultureType_name = map[int32]string{
	0:  "CULTURE_NONE",
	1:  "ALE",
	2:  "LAGER",
	3:  "WHEAT",
	4:  "WINE",
	5:  "CHAMPAGNE",
	6:  "BACTERIA",
	7:  "BRETT",
	8:  "KVEIK",
	9:  "LACTO",
	10: "MALOLACTIC",
	11: "MIXED_CULTURE",
	12: "OTHER",
	13: "PEDIO",
	14: "SPONTANEOUS",
}

var Culture_CultureType_value = map[string]int32{
	"CULTURE_NONE":  0,
	"ALE":           1,
	"LAGER":         2,
	"WHEAT":         3,
	"WINE":          4,
	"CHAMPAGNE":     5,
	"BACTERIA":      6,
	"BRETT":         7,
	"KVEIK":         8,
	"LACTO":         9,
	"MALOLACTIC":    10,
	"MIXED_CULTURE": 11,
	"OTHER":         12,
	"PEDIO":         13,
	"SPONTANEOUS":   14,
}

type Culture_FormType int32

const (
	// None
	Culture_FORM_NONE Culture_FormType = 0
	// Liquid
	Culture_LIQUID Culture_FormType = 1
	// Dry
	Culture_DRY Culture_FormType = 2
	// Slant
	Culture_SLANT Culture_FormType = 3
	// Culture
	Culture_CULTURE Culture_FormType = 4
	// Dregs
	Culture_DREGS Culture_FormType = 5
)

var Culture_FormType_name = map[int32]string{
	0: "FORM_NONE",
	1: "LIQUID",
	2: "DRY",
	3: "SLANT",
	4: "CULTURE",
	5: "DREGS",
}

var Culture_FormType_value = map[string]int32{
	"FORM_NONE": 0,
	"LIQUID":    1,
	"DRY":       2,
	"SLANT":     3,
	"CULTURE":   4,
	"DREGS":     5,
}

type Culture_FlocculationType int32

const (
	// None
	Culture_FLOCCULATION_NONE Culture_FlocculationType = 0
	// Low
	Culture_LOW Culture_FlocculationType = 1
	// Medium
	Culture_MEDIUM Culture_FlocculationType = 2
	// High
	Culture_HIGH Culture_FlocculationType = 3
	// Very High
	Culture_VERY_HIGH Culture_FlocculationType = 4
)

var Culture_FlocculationType_name = map[int32]string{
	0: "FLOCCULATION_NONE",
	1: "LOW",
	2: "MEDIUM",
	3: "HIGH",
	4: "VERY_HIGH",
}

var Culture_FlocculationType_value = map[string]int32{
	"FLOCCULATION_NONE": 0,
	"LOW":               1,
	"MEDIUM":            2,
	"HIGH":              3,
	"VERY_HIGH":         4,
}
