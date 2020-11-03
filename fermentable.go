package beerXML

import (
	"encoding/xml"
	"strings"
)

// Encloses a set of one or more Fermentable records.
type Fermentables struct {
	Fermentable []Fermentable `xml:"FERMENTABLE" json:"fermentable,omitempty"`
}

// The term "fermentable" encompasses all fermentable items that contribute substantially to the beer including
// extracts, grains, sugars, honey, fruits.
type Fermentable struct {
	// Name of the fermentable.
	Name string `xml:"NAME" json:"name,omitempty"`

	// Should be set to 1 for this version of the XML standard.  May be a higher number for later versions but all later
	// versions shall be backward compatible.
	Version int32 `xml:"VERSION" json:"version,omitempty"`

	// May be "Grain", "Sugar", "Extract", "Dry Extract" or "Adjunct".  Extract refers to liquid extract.
	Type Fermentable_FermentablesType `xml:"TYPE" json:"type,omitempty"`

	// Weight of the fermentable, extract or sugar in Kilograms.
	Amount float64 `xml:"AMOUNT" json:"amount,omitempty"`
	// Percent dry yield (fine grain) for the grain, or the raw yield by weight if this is an extract adjunct or sugar.
	Yield float64 `xml:"YIELD" json:"yield,omitempty"`
	// The color of the item in Lovibond Units (SRM for liquid extracts).
	Color float64 `xml:"COLOR" json:"color,omitempty"`
	// May be TRUE if this item is normally added after the boil.  The default value is FALSE since most grains are
	// added during the mash or boil.
	AddAfterBoil *bool `xml:"ADD_AFTER_BOIL,omitempty" json:"add_after_boil,omitempty"`
	// Country or place of origin
	Origin *string `xml:"ORIGIN,omitempty" json:"origin,omitempty"`
	// Supplier of the grain/extract/sugar
	Supplier *string `xml:"SUPPLIER,omitempty" json:"supplier,omitempty"`
	// Textual noted describing this ingredient and its use.  May be multiline.
	Notes *string `xml:"NOTES,omitempty" json:"notes,omitempty"`
	// Percent difference between the coarse grain yield and fine grain yield.  Only appropriate for a "Grain" or
	// "Adjunct" type, otherwise this value is ignored.
	CoarseFineDiff *float64 `xml:"COARSE_FINE_DIFF,omitempty" json:"coarse_fine_diff,omitempty"`
	// Percent moisture in the grain.  Only appropriate for a "Grain" or "Adjunct" type, otherwise this value is ignored.
	Moisture *float64 `xml:"MOISTURE,omitempty" json:"moisture,omitempty"`
	// The diastatic power of the grain as measured in "Lintner" units. Only appropriate for a "Grain" or "Adjunct"
	// type, otherwise this value is ignored.
	DiastaticPower *float64 `xml:"DIASTATIC_POWER,omitempty" json:"diastatic_power,omitempty"`
	// The percent protein in the grain.  Only appropriate for a "Grain" or "Adjunct" type, otherwise this value is
	// ignored.
	Protein *float64 `xml:"PROTEIN,omitempty" json:"protein,omitempty"`
	// The recommended maximum percentage (by weight) this ingredient should represent in a batch of beer.
	MaxInBatch *float64 `xml:"MAX_IN_BATCH,omitempty" json:"max_in_batch,omitempty"`
	// TRUE if it is recommended the grain be mashed, FALSE if it can be steeped.  A value of TRUE is only appropriate
	// for a "Grain" or "Adjunct" types.  The default value is FALSE.  Note that this does NOT indicate whether
	// the grain is mashed or not – it is only a recommendation used in recipe formulation.
	RecommendMash *bool `xml:"RECOMMEND_MASH,omitempty" json:"recommend_mash,omitempty"`
	// For hopped extracts only - an estimate of the number of IBUs per pound of extract in a gallon of water.
	// To convert to IBUs we multiply this number by the "AMOUNT" field (in pounds) and divide by the number of
	// gallons in the batch.  Based on a sixty minute boil.
	// Only suitable for use with an "Extract" type, otherwise this value is ignored.
	IIBGalPerlb *float64 `xml:"IBU_GAL_PER_LB,omitempty" json:"ibu_gal_per_lb,omitempty"`

	// Extensions

	// The amount of fermentables in this record along with the units formatted for easy display in the current
	// user defined units.
	// For example “1.5 lbs” or “2.1 kg”.
	DisplayAmount *string `xml:"DISPLAY_AMOUNT,omitempty" json:"display_amount,omitempty"`
	// The yield of the fermentable converted to specific gravity units for display.
	// For example “1.036” or “1.040” might be valid potentials.
	Potential *string `xml:"POTENTIAL,omitempty" json:"potential,omitempty"`
	// Amount in inventory for this item along with the units – for example “10.0 lb”
	Inventory *string `xml:"INVENTORY,omitempty" json:"inventory,omitempty"`
	// Color in user defined color units along with the unit identified – for example “200L” or “40 ebc”
	DisplayColor *string `xml:"DISPLAY_COLOR,omitempty" json:"display_color,omitempty"`
}


func (a *Fermentable) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Alias Fermentable
	aux := &struct {
		*Alias
		Yield *string `xml:"YIELD" json:"yield,omitempty"`
		CoarseFineDiff *string `xml:"COARSE_FINE_DIFF,omitempty" json:"coarse_fine_diff,omitempty"`
		Moisture *string `xml:"MOISTURE,omitempty" json:"moisture,omitempty"`
		Protein *string `xml:"PROTEIN,omitempty" json:"protein,omitempty"`
		MaxInBatch *string `xml:"MAX_IN_BATCH,omitempty" json:"max_in_batch,omitempty"`
	}{
		Alias: (*Alias)(a),
	}

	err := d.DecodeElement(aux, &start)
	if err != nil {
		return err
	}

	a.Yield = *percentToFloat(aux.Yield)
	a.CoarseFineDiff = percentToFloat(aux.CoarseFineDiff)
	a.Moisture = percentToFloat(aux.Moisture)
	a.Protein = percentToFloat(aux.Protein)
	a.MaxInBatch = percentToFloat(aux.MaxInBatch)

	return nil
}

func (a Fermentable) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type Alias Fermentable
	aux := &struct {
		*Alias
		Yield string `xml:"YIELD" json:"yield,omitempty"`
		CoarseFineDiff string `xml:"COARSE_FINE_DIFF,omitempty" json:"coarse_fine_diff,omitempty"`
		Moisture string `xml:"MOISTURE,omitempty" json:"moisture,omitempty"`
		Protein string `xml:"PROTEIN,omitempty" json:"protein,omitempty"`
		MaxInBatch string `xml:"MAX_IN_BATCH,omitempty" json:"max_in_batch,omitempty"`
	}{
		Alias: (*Alias)(&a),
		Yield: floatToPercent(&a.Yield),
		CoarseFineDiff: floatToPercent(a.CoarseFineDiff),
		Moisture: floatToPercent(a.Moisture),
		Protein: floatToPercent(a.Protein),
		MaxInBatch: floatToPercent(a.MaxInBatch),
	}
	start.Name.Local = strings.ToUpper(start.Name.Local)

	err := e.EncodeElement(aux, start)
	if err != nil {
		return err
	}

	return nil
}

func (a *Fermentable_FermentablesType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var use string

	err := d.DecodeElement(&use, &start)
	if err != nil {
		return err
	}

	if value, ok := Fermentable_FermentablesType_value[use]; ok {
		*a = Fermentable_FermentablesType(value)
	} else {
		*a = FERMENTABLE_FERMENTABLES_NONE
	}

	return nil
}

func (a Fermentable_FermentablesType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if name, ok := Fermentable_FermentablesType_name[int32(a)]; ok {
		err := e.EncodeElement(name, start)
		if err != nil {
			return err
		}
	}
	return nil
}


type Fermentable_FermentablesType int32

const (
	// None
	FERMENTABLE_FERMENTABLES_NONE Fermentable_FermentablesType = 0
	// Grain
	FERMENTABLE_GRAIN Fermentable_FermentablesType = 1
	// Sugar
	FERMENTABLE_SUGAR Fermentable_FermentablesType = 2
	// Extract
	FERMENTABLE_EXTRACT Fermentable_FermentablesType = 3
	// Dry Extract
	FERMENTABLE_DRY_EXTRACT Fermentable_FermentablesType = 4
	// Adjunct
	FERMENTABLE_ADJUNCT Fermentable_FermentablesType = 5
)

var Fermentable_FermentablesType_name = map[int32]string{
	0: "FERMENTABLES_NONE",
	1: "Grain",
	2: "Sugar",
	3: "Extract",
	4: "Dry Extract",
	5: "Adjunct",
}

var Fermentable_FermentablesType_value = map[string]int32{
	"FERMENTABLES_NONE": 0,
	"Grain":             1,
	"Sugar":             2,
	"Extract":           3,
	"Dry Extract":       4,
	"Adjunct":           5,
}
