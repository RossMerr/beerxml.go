package beerXML

import (
	"encoding/json"
)

type Water struct {
	Name          string  `xml:"NAME" json:"name,omitempty"`
	Version       int32   `xml:"VERSION" json:"version,omitempty"`
	Amount        float64 `xml:"AMOUNT" json:"amount,omitempty"`
	Calcium       float64 `xml:"CALCIUM" json:"calcium,omitempty"`
	Bicarbonate   float64 `xml:"BICARBONATE" json:"bicarbonate,omitempty"`
	Sulfate       float64 `xml:"SULFATE" json:"sulfate,omitempty"`
	Chloride      float64 `xml:"CHLORIDE" json:"chloride,omitempty"`
	Sodium        float64 `xml:"SODIUM" json:"sodium,omitempty"`
	Magnesium     float64 `xml:"MAGNESIUM" json:"magnesium,omitempty"`
	PH            float64 `xml:"PH" json:"ph,omitempty"`
	Notes         string  `xml:"NOTES" json:"notes,omitempty"`
	DisplayAmount string  `xml:"DISPLAY_AMOUNT" json:"display_amount,omitempty"`
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
