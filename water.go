package beerXML

import (
	"encoding/json"
)

type Water struct {
	Name          string  `xml:"NAME" json:"name,omitempty"`
	Version       int32   `xml:"VERSION" json:"version,omitempty"`
	Amount        float32 `xml:"AMOUNT" json:"amount,omitempty"`
	Calcium       float32 `xml:"CALCIUM" json:"calcium,omitempty"`
	Bicarbonate   float32 `xml:"BICARBONATE" json:"bicarbonate,omitempty"`
	Sulfate       float32 `xml:"SULFATE" json:"sulfate,omitempty"`
	Chloride      float32 `xml:"CHLORIDE" json:"chloride,omitempty"`
	Sodium        float32 `xml:"SODIUM" json:"sodium,omitempty"`
	Magnesium     float32 `xml:"MAGNESIUM" json:"magnesium,omitempty"`
	PH            float32 `xml:"PH" json:"ph,omitempty"`
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
