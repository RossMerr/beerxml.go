package beerXML

import (
	"encoding/json"
	"strings"
)

type Hop struct {
	Name    string  `xml:"NAME" json:"name,omitempty"`
	Version int32   `xml:"VERSION" json:"version,omitempty"`
	Origin  string  `xml:"ORIGIN" json:"origin,omitempty"`
	Alpha   float64 `xml:"ALPHA" json:"alpha,omitempty"`
	Amount  float64 `xml:"AMOUNT" json:"amount,omitempty"`
	Use     string  `xml:"USE" json:"use,omitempty"`
	Time    float64 `xml:"TIME" json:"time,omitempty"`
	Notes   string  `xml:"NOTES" json:"notes,omitempty"`
	Type    string  `xml:"TYPE" json:"type,omitempty"`
	Form    string  `xml:"FORM" json:"form,omitempty"`
	Beta    float64 `xml:"BETA" json:"beta,omitempty"`
	HSI     float64 `xml:"HSI" json:"hsi,omitempty"`

	DisplayAmount string `xml:"DISPLAY_AMOUNT" json:"display_amount,omitempty"`
	DisplayTime   string `xml:"DISPLAY_TIME" json:"display_time,omitempty"`
	Inventory     string `xml:"INVENTORY" json:"inventory,omitempty"`
}

func (a Hop) MarshalJSON() ([]byte, error) {

	type Alias Hop
	t := func() int32 {
		if t, ok := Hop_HopsType_value[strings.ToUpper(a.Type)]; ok {
			return t
		}
		return int32(Hop_HOPS_NONE)
	}()

	use := func() int32 {
		if t, ok := Hop_HopsUseType_value[strings.ToUpper(a.Use)]; ok {
			return t
		}
		return int32(Hop_HOPS_USE_NONE)
	}()

	form := func() int32 {
		if t, ok := Hop_HopsFormType_value[strings.ToUpper(a.Form)]; ok {
			return t
		}
		return int32(Hop_HOPS_FORM_NONE)
	}()

	return json.Marshal(&struct {
		Type int32 `json:"type,omitempty"`
		Use  int32 `json:"use,omitempty"`
		Form int32 `json:"form,omitempty"`
		*Alias
	}{
		Type:  t,
		Use:   use,
		Form:  form,
		Alias: (*Alias)(&a),
	})
}

func (a *Hop) UnmarshalJSON(b []byte) error {
	return nil
}

type Hops struct {
	Hop []Hop `xml:"HOP" json:"hop,omitempty"`
}

func (a Hops) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0)
	b = append(b, []byte("[")...)
	if len(a.Hop) > 0 {
		for _, hop := range a.Hop {
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

func (a *Hops) UnmarshalJSON(b []byte) error {
	return nil
}

type Hop_HopsType int32

const (
	// None
	Hop_HOPS_NONE Hop_HopsType = 0
	// Bittering
	Hop_BITTERING Hop_HopsType = 1
	// Aroma
	Hop_AROMA Hop_HopsType = 2
	// Both
	Hop_BOTH Hop_HopsType = 3
)

var Hop_HopsType_name = map[int32]string{
	0: "HOPS_NONE",
	1: "BITTERING",
	2: "AROMA",
	3: "BOTH",
}

var Hop_HopsType_value = map[string]int32{
	"HOPS_NONE": 0,
	"BITTERING": 1,
	"AROMA":     2,
	"BOTH":      3,
}

type Hop_HopsUseType int32

const (
	// None
	Hop_HOPS_USE_NONE Hop_HopsUseType = 0
	// Boil
	Hop_BOIL Hop_HopsUseType = 1
	// Dry Hop
	Hop_DRY_HOP Hop_HopsUseType = 2
	// Mash
	Hop_MASH Hop_HopsUseType = 3
	// First Wort
	Hop_FIRST_WORT Hop_HopsUseType = 4
	// Aroma
	Hop_AROMA_OTHER Hop_HopsUseType = 5
)

var Hop_HopsUseType_name = map[int32]string{
	0: "HOPS_USE_NONE",
	1: "BOIL",
	2: "DRY_HOP",
	3: "MASH",
	4: "FIRST_WORT",
	5: "AROMA_OTHER",
}

var Hop_HopsUseType_value = map[string]int32{
	"HOPS_USE_NONE": 0,
	"BOIL":          1,
	"DRY_HOP":       2,
	"MASH":          3,
	"FIRST_WORT":    4,
	"AROMA_OTHER":   5,
}

type Hop_HopsFormType int32

const (
	// None
	Hop_HOPS_FORM_NONE Hop_HopsFormType = 0
	// Pellet
	Hop_PELLET Hop_HopsFormType = 1
	// Plug
	Hop_PLUG Hop_HopsFormType = 2
	// Leaf
	Hop_LEAF Hop_HopsFormType = 3
)

var Hop_HopsFormType_name = map[int32]string{
	0: "HOPS_FORM_NONE",
	1: "PELLET",
	2: "PLUG",
	3: "LEAF",
}

var Hop_HopsFormType_value = map[string]int32{
	"HOPS_FORM_NONE": 0,
	"PELLET":         1,
	"PLUG":           2,
	"LEAF":           3,
}
