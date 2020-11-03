package beerXML

import (
	"encoding/json"
	"strconv"
	"strings"
)

// The term "fermentable" encompasses all fermentable items that contribute substantially to the beer including
// extracts, grains, sugars, honey, fruits.
type Fermentable struct {
	// Name of the fermentable.
	Name string `xml:"NAME" json:"name,omitempty"`

	// Should be set to 1 for this version of the XML standard.  May be a higher number for later versions but all later
	// versions shall be backward compatible.
	Version int32 `xml:"VERSION" json:"version,omitempty"`

	// May be "Grain", "Sugar", "Extract", "Dry Extract" or "Adjunct".  Extract refers to liquid extract.
	Type string `xml:"TYPE" json:"type,omitempty"`

	// Weight of the fermentable, extract or sugar in Kilograms.
	Amount float64 `xml:"AMOUNT" json:"amount,omitempty"`
	// Percent dry yield (fine grain) for the grain, or the raw yield by weight if this is an extract adjunct or sugar.
	Yield float64 `xml:"YIELD" json:"yield,omitempty"`
	// The color of the item in Lovibond Units (SRM for liquid extracts).
	Color float64 `xml:"COLOR" json:"color,omitempty"`
	// May be TRUE if this item is normally added after the boil.  The default value is FALSE since most grains are
	// added during the mash or boil.
	AddAfterBoil *bool `xml:"ADD_AFTER_BOIL" json:"add_after_boil,omitempty"`
	// Country or place of origin
	Origin *string `xml:"ORIGIN" json:"origin,omitempty"`
	// Supplier of the grain/extract/sugar
	Supplier *string `xml:"SUPPLIER" json:"supplier,omitempty"`
	// Textual noted describing this ingredient and its use.  May be multiline.
	Notes *string `xml:"NOTES" json:"notes,omitempty"`
	// Percent difference between the coarse grain yield and fine grain yield.  Only appropriate for a "Grain" or
	// "Adjunct" type, otherwise this value is ignored.
	CoarseFineDiff *string `xml:"COARSE_FINE_DIFF" json:"coarse_fine_diff,omitempty"`
	// Percent moisture in the grain.  Only appropriate for a "Grain" or "Adjunct" type, otherwise this value is ignored.
	Moisture *string `xml:"MOISTURE" json:"moisture,omitempty"`
	// The diastatic power of the grain as measured in "Lintner" units. Only appropriate for a "Grain" or "Adjunct"
	// type, otherwise this value is ignored.
	DiastaticPower *string `xml:"DIASTATIC_POWER" json:"diastatic_power,omitempty"`
	// The percent protein in the grain.  Only appropriate for a "Grain" or "Adjunct" type, otherwise this value is
	// ignored.
	Protein *string `xml:"PROTEIN" json:"protein,omitempty"`
	// The recommended maximum percentage (by weight) this ingredient should represent in a batch of beer.
	MaxInBatch *string `xml:"MAX_IN_BATCH" json:"max_in_batch,omitempty"`
	// TRUE if it is recommended the grain be mashed, FALSE if it can be steeped.  A value of TRUE is only appropriate
	// for a "Grain" or "Adjunct" types.  The default value is FALSE.  Note that this does NOT indicate whether
	// the grain is mashed or not – it is only a recommendation used in recipe formulation.
	RecommendMash *bool `xml:"RECOMMEND_MASH" json:"recommend_mash,omitempty"`
	// For hopped extracts only - an estimate of the number of IBUs per pound of extract in a gallon of water.
	// To convert to IBUs we multiply this number by the "AMOUNT" field (in pounds) and divide by the number of
	// gallons in the batch.  Based on a sixty minute boil.
	// Only suitable for use with an "Extract" type, otherwise this value is ignored.
	IIBGalPerlb *float64 `xml:"IBU_GAL_PER_LB" json:"ibu_gal_per_lb,omitempty"`

	// Extensions

	// The amount of fermentables in this record along with the units formatted for easy display in the current
	// user defined units.
	// For example “1.5 lbs” or “2.1 kg”.
	DisplayAmount string `xml:"DISPLAY_AMOUNT" json:"display_amount,omitempty"`
	// The yield of the fermentable converted to specific gravity units for display.
	// For example “1.036” or “1.040” might be valid potentials.
	Potential string `xml:"POTENTIAL" json:"potential,omitempty"`
	// Amount in inventory for this item along with the units – for example “10.0 lb”
	Inventory string `xml:"INVENTORY" json:"inventory,omitempty"`
	// Color in user defined color units along with the unit identified – for example “200L” or “40 ebc”
	DisplayColor string `xml:"DISPLAY_COLOR" json:"display_color,omitempty"`
}

func (a Fermentable) MarshalJSON() ([]byte, error) {

	type Alias Fermentable
	t := func() int32 {
		if t, ok := Fermentable_FermentablesType_value[strings.ToUpper(a.Type)]; ok {
			return t
		}
		return int32(Fermentable_FERMENTABLES_NONE)
	}()

	return json.Marshal(&struct {
		Type           int32    `json:"type,omitempty"`
		CoarseFineDiff *float64 `json:"coarse_fine_diff,omitempty"`
		Moisture       *float64 `json:"moisture,omitempty"`
		DiastaticPower *float64 `json:"diastatic_power,omitempty"`
		Protein        *float64 `json:"protein,omitempty"`
		MaxInBatch     *float64 `json:"max_in_batch,omitempty"`
		Potential      *float64 `json:"potential,omitempty"`
		*Alias
	}{
		Type: t,
		CoarseFineDiff: func() *float64 {
			if a.CoarseFineDiff != nil {
				str := string(reg.Find([]byte(*a.CoarseFineDiff)))
				if f, err := strconv.ParseFloat(str, 64); err == nil {
					return &f
				}
			}
			return nil
		}(),
		Moisture: func() *float64 {
			if a.Moisture != nil {
				str := string(reg.Find([]byte(*a.Moisture)))
				if f, err := strconv.ParseFloat(str, 64); err == nil {
					return &f
				}
			}
			return nil
		}(),
		DiastaticPower: func() *float64 {
			if a.DiastaticPower != nil {
				str := string(reg.Find([]byte(*a.DiastaticPower)))
				if f, err := strconv.ParseFloat(str, 64); err == nil {
					return &f
				}
			}
			return nil
		}(),
		Protein: func() *float64 {
			if a.Protein != nil {
				str := string(reg.Find([]byte(*a.Protein)))
				if f, err := strconv.ParseFloat(str, 64); err == nil {
					return &f
				}
			}
			return nil
		}(),
		MaxInBatch: func() *float64 {
			if a.MaxInBatch != nil {
				str := string(reg.Find([]byte(*a.MaxInBatch)))
				if f, err := strconv.ParseFloat(str, 64); err == nil {
					return &f
				}
			}
			return nil
		}(),
		Potential: func() *float64 {
			if a.Potential != "" {
				str := string(reg.Find([]byte(a.Potential)))
				if f, err := strconv.ParseFloat(str, 64); err == nil {
					return &f
				}
			}
			return nil
		}(),
		Alias: (*Alias)(&a),
	})
}

func (a *Fermentable) UnmarshalJSON(b []byte) error {
	return nil
}

type Fermentables struct {
	Fermentable []Fermentable `xml:"FERMENTABLE" json:"fermentable,omitempty"`
}

func (a Fermentables) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0)
	b = append(b, []byte("[")...)
	if len(a.Fermentable) > 0 {
		for _, hop := range a.Fermentable {
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

func (a *Fermentables) UnmarshalJSON(b []byte) error {
	return nil
}

type Fermentable_FermentablesType int32

const (
	// None
	Fermentable_FERMENTABLES_NONE Fermentable_FermentablesType = 0
	// Grain
	Fermentable_GRAIN Fermentable_FermentablesType = 1
	// Sugar
	Fermentable_SUGAR Fermentable_FermentablesType = 2
	// Extract
	Fermentable_EXTRACT Fermentable_FermentablesType = 3
	// Dry Extract
	Fermentable_DRY_EXTRACT Fermentable_FermentablesType = 4
	// Adjunct
	Fermentable_ADJUNCT Fermentable_FermentablesType = 5
)

var Fermentable_FermentablesType_name = map[int32]string{
	0: "FERMENTABLES_NONE",
	1: "GRAIN",
	2: "SUGAR",
	3: "EXTRACT",
	4: "DRY_EXTRACT",
	5: "ADJUNCT",
}

var Fermentable_FermentablesType_value = map[string]int32{
	"FERMENTABLES_NONE": 0,
	"GRAIN":             1,
	"SUGAR":             2,
	"EXTRACT":           3,
	"DRY_EXTRACT":       4,
	"ADJUNCT":           5,
}
