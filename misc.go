package beerXML

import (
	"encoding/json"
	"strings"
)

type MISC struct {
	Name           string  `xml:"NAME" json:"name,omitempty"`
	Version        int32   `xml:"VERSION" json:"version,omitempty"`
	Type           string  `xml:"TYPE" json:"type,omitempty"`
	Use            string  `xml:"USE" json:"use,omitempty"`
	Amount         float64 `xml:"AMOUNT" json:"amount,omitempty"`
	Time           float64 `xml:"TIME" json:"time,omitempty"`
	Amountisweight bool    `xml:"AMOUNT_IS_WEIGHT" json:"amount_is_weight,omitempty"`
	Usefor         string  `xml:"USE_FOR" json:"use_for,omitempty"`
	Notes          string  `xml:"NOTES" json:"notes,omitempty"`
	BatchSize      string  `xml:"BATCH_SIZE" json:"batch_size,omitempty"`
	DisplayAmount  string  `xml:"DISPLAY_AMOUNT" json:"display_amount,omitempty"`
	DisplayTime    string  `xml:"DISPLAY_TIME" json:"display_time,omitempty"`
	Inventory      string  `xml:"INVENTORY" json:"inventory,omitempty"`
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
