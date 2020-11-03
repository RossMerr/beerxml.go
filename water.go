package beerXML

import (
	"encoding/json"
)

// The term "water" encompasses water profiles.  Though not strictly required for recipes, the water record allows
// supporting programs to record the water profile used for brewing a particular batch.
type Water struct {
	// Name of the water profile – usually the city and country of the water profile.
	Name          string  `xml:"NAME" json:"name,omitempty"`
	// Version of the water record.  Should always be “1” for this version of the XML standard.
	Version       int32   `xml:"VERSION" json:"version,omitempty"`
	// Volume of water to use in a recipe in liters.
	Amount        float64 `xml:"AMOUNT" json:"amount,omitempty"`
	// The amount of calcium (Ca) in parts per million.
	Calcium       float64 `xml:"CALCIUM" json:"calcium,omitempty"`
	// The amount of bicarbonate (HCO3) in parts per million.
	Bicarbonate   float64 `xml:"BICARBONATE" json:"bicarbonate,omitempty"`
	// The amount of Sulfate (SO4) in parts per million.
	Sulfate       float64 `xml:"SULFATE" json:"sulfate,omitempty"`
	// The amount of Chloride (Cl) in parts per million.
	Chloride      float64 `xml:"CHLORIDE" json:"chloride,omitempty"`
	// The amount of Sodium (Na) in parts per million.
	Sodium        float64 `xml:"SODIUM" json:"sodium,omitempty"`
	// The amount of Magnesium (Mg) in parts per million.
	Magnesium     float64 `xml:"MAGNESIUM" json:"magnesium,omitempty"`
	// The PH of the water.
	PH            *float64 `xml:"PH" json:"ph,omitempty"`
	Notes         *string  `xml:"NOTES" json:"notes,omitempty"`
	// Notes about the water profile.  May be multiline.

	// Extensions

	// The amount of water in this record along with the units formatted for easy display in the current
	// user defined units.  For example “5.0 gal” or “20.0 l”.
	DisplayAmount *string  `xml:"DISPLAY_AMOUNT" json:"display_amount,omitempty"`
}

type Waters struct {
	Water []Water `xml:"WATER" json:"water,omitempty"`
}

func (a Waters) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0)

	b = append(b, []byte("[")...)
	if len(a.Water) > 0 {
		for _, hop := range a.Water {
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

func (a *Waters) UnmarshalJSON(b []byte) error {
	return nil
}
