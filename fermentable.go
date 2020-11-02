package beerXML

import (
	"encoding/json"
	"strconv"
	"strings"
)

type Fermentable struct {
	Name           string  `xml:"NAME" json:"name,omitempty"`
	Version        int32   `xml:"VERSION" json:"version,omitempty"`
	Type           string  `xml:"TYPE" json:"type,omitempty"`
	Amount         float64 `xml:"AMOUNT" json:"amount,omitempty"`
	Yield          float64 `xml:"YIELD" json:"yield,omitempty"`
	Color          float64 `xml:"COLOR" json:"color,omitempty"`
	Addafterboil   bool    `xml:"ADD_AFTER_BOIL" json:"add_after_boil,omitempty"`
	Origin         string  `xml:"ORIGIN" json:"origin,omitempty"`
	Supplier       string  `xml:"SUPPLIER" json:"supplier,omitempty"`
	Notes          string  `xml:"NOTES" json:"notes,omitempty"`
	Coarsefinediff string  `xml:"COARSE_FINE_DIFF" json:"coarse_fine_diff,omitempty"`
	Moisture       string  `xml:"MOISTURE" json:"moisture,omitempty"`
	Diastaticpower string  `xml:"DIASTATIC_POWER" json:"diastatic_power,omitempty"`
	Protein        string  `xml:"PROTEIN" json:"protein,omitempty"`
	Maxinbatch     string  `xml:"MAX_IN_BATCH" json:"max_in_batch,omitempty"`
	Recommendmash  bool    `xml:"RECOMMEND_MASH" json:"recommend_mash,omitempty"`
	Ibugalperlb    float64 `xml:"IBU_GAL_PER_LB" json:"ibu_gal_per_lb,omitempty"`
	Potential      string  `xml:"POTENTIAL" json:"potential,omitempty"`

	DisplayAmount string `xml:"DISPLAY_AMOUNT" json:"display_amount,omitempty"`
	DisplayColor  string `xml:"DISPLAY_COLOR" json:"display_color,omitempty"`
	Inventory     string `xml:"INVENTORY" json:"inventory,omitempty"`
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
		Type           int32   `json:"type,omitempty"`
		Coarsefinediff float64 `json:"coarse_fine_diff,omitempty"`
		Moisture       float64 `json:"moisture,omitempty"`
		Diastaticpower float64 `json:"diastatic_power,omitempty"`
		Protein        float64 `json:"protein,omitempty"`
		Maxinbatch     float64 `json:"max_in_batch,omitempty"`
		Potential      float64 `json:"potential,omitempty"`
		*Alias
	}{
		Type: t,
		Coarsefinediff: func() float64 {
			str := string(reg.Find([]byte(a.Coarsefinediff)))
			f, _ := strconv.ParseFloat(str, 64)
			return f
		}(),
		Moisture: func() float64 {
			str := string(reg.Find([]byte(a.Moisture)))
			f, _ := strconv.ParseFloat(str, 64)
			return f
		}(),
		Diastaticpower: func() float64 {
			str := string(reg.Find([]byte(a.Diastaticpower)))
			f, _ := strconv.ParseFloat(str, 64)
			return f
		}(),
		Protein: func() float64 {
			str := string(reg.Find([]byte(a.Protein)))
			f, _ := strconv.ParseFloat(str, 64)
			return f
		}(),
		Maxinbatch: func() float64 {
			str := string(reg.Find([]byte(a.Maxinbatch)))
			f, _ := strconv.ParseFloat(str, 64)
			return f
		}(),
		Potential: func() float64 {
			str := string(reg.Find([]byte(a.Potential)))
			f, _ := strconv.ParseFloat(str, 64)
			return f
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
