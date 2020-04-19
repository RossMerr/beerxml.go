package beerXML

import (
	"encoding/json"
)

type MashStep struct {
	Name             string  `xml:"NAME" json:"name,omitempty"`
	Version          int32   `xml:"VERSION" json:"version,omitempty"`
	Type             string  `xml:"TYPE" json:"type,omitempty"`
	Infuseamount     float32 `xml:"INFUSE_AMOUNT" json:"infuse_amount,omitempty"`
	Steptime         int64   `xml:"STEP_TIME" json:"step_time,omitempty"`
	Steptemp         float64 `xml:"STEP_TEMP" json:"step_temp,omitempty"`
	Ramptime         int64   `xml:"RAMP_TIME" json:"ramp_time,omitempty"`
	Endtemp          float64 `xml:"END_TEMP" json:"end_temp,omitempty"`
	DecoctionAmt     string  `xml:"DECOCTION_AMT" json:"decoction_amt,omitempty"`
	Description      string  `xml:"DESCRIPTION" json:"description,omitempty"`
	DisplayInfuseAmt string  `xml:"DISPLAY_INFUSE_AMT" json:"display_infuse_amt,omitempty"`
	DisplayStepTemp  string  `xml:"DISPLAY_STEP_TEMP" json:"display_step_temp,omitempty"`
	InfuseTemp       string  `xml:"INFUSE_TEMP" json:"infuse_temp,omitempty"`
	WaterGrainRatio  string  `xml:"WATER_GRAIN_RATIO" json:"water_grain_ratio,omitempty"`
}

type MashSteps struct {
	Mashstep []MashStep `xml:"MASH_STEP" json:"mash_step,omitempty"`
}

func (a MashSteps) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0)
	b = append(b, []byte("[")...)
	if len(a.Mashstep) > 0 {
		for _, hop := range a.Mashstep {
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

func (a *MashSteps) UnmarshalJSON(b []byte) error {
	return nil
}
