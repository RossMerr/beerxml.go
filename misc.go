package beerXML

import (
	"encoding/json"
	"strings"
)

// The term "misc" encompasses all non-fermentable miscellaneous ingredients that are not hops or yeast and do not
// significantly change the gravity of the beer.  For example: spices, clarifying agents, water treatments, etc…
type MISC struct {
	// Name of the misc item.
	Name           string  `xml:"NAME" json:"name,omitempty"`
	// Version number of this element.  Should be “1” for this version.
	Version        int32   `xml:"VERSION" json:"version,omitempty"`
	// May be “Spice”, “Fining”, “Water Agent”, “Herb”, “Flavor” or “Other”
	Type           string  `xml:"TYPE" json:"type,omitempty"`
	// May be “Boil”, “Mash”, “Primary”, “Secondary”, “Bottling”
	Use            string  `xml:"USE" json:"use,omitempty"`
	// Amount of time the misc was boiled, steeped, mashed, etc in minutes.
	Time           float64 `xml:"TIME" json:"time,omitempty"`
	// Amount of item used.  The default measurements are by weight, but this may be the measurement in volume units
	// if AMOUNT_IS_WEIGHT is set to TRUE for this record.
	// If a liquid it is in liters,
	// if a solid the weight is measured in kilograms.
	Amount         float64 `xml:"AMOUNT" json:"amount,omitempty"`
	// TRUE if the amount measurement is a weight measurement and FALSE if the amount is a volume measurement.
	// Default value (if not present) is assumed to be FALSE.
	AmountIsWeight *bool    `xml:"AMOUNT_IS_WEIGHT" json:"amount_is_weight,omitempty"`
	// Short description of what the ingredient is used for in text
	UseFor         *string  `xml:"USE_FOR" json:"use_for,omitempty"`
	// Detailed notes on the item including usage.  May be multiline.
	Notes          *string  `xml:"NOTES" json:"notes,omitempty"`

	// Extensions

	// The amount of the item in this record along with the units formatted for easy display in the current user
	// defined units.  For example “1.5 lbs” or “2.1 kg”.
	DisplayAmount  *string  `xml:"DISPLAY_AMOUNT" json:"display_amount,omitempty"`
	// Amount in inventory for this item along with the units – for example “10.0 lb”
	Inventory      *string  `xml:"INVENTORY" json:"inventory,omitempty"`
	// Time in appropriate units along with the units as in “10 min” or “3 days”.
	DisplayTime    *string  `xml:"DISPLAY_TIME" json:"display_time,omitempty"`
}

func (a MISC) MarshalJSON() ([]byte, error) {

	type Alias MISC
	t := func() int32 {
		if t, ok := Misc_MiseType_value[strings.ToUpper(a.Type)]; ok {
			return t
		}
		return int32(Misc_MISE_NONE)
	}()

	use := func() int32 {
		if t, ok := Misc_MiseUseType_value[strings.ToUpper(a.Use)]; ok {
			return t
		}
		return int32(Misc_USE_NONE)
	}()

	return json.Marshal(&struct {
		Type int32 `json:"type,omitempty"`
		Use  int32 `json:"use,omitempty"`
		*Alias
	}{
		Type:  t,
		Use:   use,
		Alias: (*Alias)(&a),
	})
}

func (a *MISC) UnmarshalJSON(b []byte) error {
	return nil
}

type MISCS struct {
	MISC []MISC `xml:"MISC" json:"mise,omitempty"`
}

func (a MISCS) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0)
	b = append(b, []byte("[")...)
	if len(a.MISC) > 0 {
		for _, hop := range a.MISC {
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

func (a *MISCS) UnmarshalJSON(b []byte) error {
	return nil
}

type Misc_MiseType int32

const (
	// None
	Misc_MISE_NONE Misc_MiseType = 0
	// Spice
	Misc_SPICE Misc_MiseType = 1
	// Fining
	Misc_FINING Misc_MiseType = 2
	// Water Agent
	Misc_WATER_AGENT Misc_MiseType = 3
	// Herb
	Misc_HERB Misc_MiseType = 4
	// Flavor
	Misc_FLAVOR Misc_MiseType = 5
	// Other
	Misc_OTHER Misc_MiseType = 6
)

var Misc_MiseType_name = map[int32]string{
	0: "MISE_NONE",
	1: "SPICE",
	2: "FINING",
	3: "WATER_AGENT",
	4: "HERB",
	5: "FLAVOR",
	6: "OTHER",
}

var Misc_MiseType_value = map[string]int32{
	"MISE_NONE":   0,
	"SPICE":       1,
	"FINING":      2,
	"WATER_AGENT": 3,
	"HERB":        4,
	"FLAVOR":      5,
	"OTHER":       6,
}

type Misc_MiseUseType int32

const (
	// None
	Misc_USE_NONE Misc_MiseUseType = 0
	// Boil
	Misc_BOIL Misc_MiseUseType = 1
	// Mash
	Misc_MASH Misc_MiseUseType = 2
	// Primary
	Misc_PRIMARY Misc_MiseUseType = 3
	// Secondary
	Misc_SECONDARY Misc_MiseUseType = 4
	// Bottling
	Misc_BOTTLING Misc_MiseUseType = 5
)

var Misc_MiseUseType_name = map[int32]string{
	0: "USE_NONE",
	1: "BOIL",
	2: "MASH",
	3: "PRIMARY",
	4: "SECONDARY",
	5: "BOTTLING",
}

var Misc_MiseUseType_value = map[string]int32{
	"USE_NONE":  0,
	"BOIL":      1,
	"MASH":      2,
	"PRIMARY":   3,
	"SECONDARY": 4,
	"BOTTLING":  5,
}
