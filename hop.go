package beerXML

import (
	"encoding/json"
	"strings"
)

// The “HOP” identifier is used to define all varieties of hops.
type Hop struct {
	// Name of the hops
	Name    string  `xml:"NAME" json:"name,omitempty"`
	// Should be set to 1 for this version of the XML standard.  May be a higher number for later versions but all
	// later versions shall be backward compatible.
	Version int32   `xml:"VERSION" json:"version,omitempty"`
	// Percent alpha of hops - for example "5.5" represents 5.5% alpha
	Alpha   float64 `xml:"ALPHA" json:"alpha,omitempty"`
	// Weight in Kilograms of the hops used in the recipe.
	Amount  float64 `xml:"AMOUNT" json:"amount,omitempty"`
	// May be "Boil", "Dry Hop", "Mash", "First Wort" or "Aroma".  Note that "Aroma" and "Dry Hop" do not contribute
	// to the bitterness of the beer while the others do.  Aroma hops are added after the boil and do not contribute
	// substantially to beer bitterness.
	Use     string  `xml:"USE" json:"use,omitempty"`
	// The time as measured in minutes.  Meaning is dependent on the “USE” field.
	// For “Boil” this is the boil time.
	// For “Mash” this is the mash time.
	// For “First Wort” this is the boil time.
	// For “Aroma” this is the steep time.
	// For “Dry Hop” this is the amount of time to dry hop.
	Time    float64 `xml:"TIME" json:"time,omitempty"`
	// Textual notes about the hops, usage, substitutes.  May be a multiline entry.
	Notes   *string  `xml:"NOTES" json:"notes,omitempty"`
	// May be "Bittering", "Aroma" or "Both"
	Type    *string  `xml:"TYPE" json:"type,omitempty"`
	// May be "Pellet", "Plug" or "Leaf"
	Form    *string  `xml:"FORM" json:"form,omitempty"`
	// Hop beta percentage - for example "4.4" denotes 4.4 % beta
	Beta    *float64 `xml:"BETA" json:"beta,omitempty"`
	// Hop Stability Index - defined as the percentage of hop alpha lost in 6 months of storage
	HSI     *float64 `xml:"HSI" json:"hsi,omitempty"`
	// Place of origin for the hops
	Origin  *string  `xml:"ORIGIN" json:"origin,omitempty"`
	// Substitutes that can be used for this hops
	Substitutes *string `xml:"SUBSTITUTES" json:"substitutes,omitempty"`
	// Humulene level in percent.
	Humulene *string `xml:"HUMULENE" json:"humulene,omitempty"`
	// Caryophyllene level in percent.
	Caryophyllene *string `xml:"CARYOPHYLLENE" json:"caryophyllene,omitempty"`
	// Cohumulone level in percent
	Cohumulone *string `xml:"COHUMULONE" json:"cohumulone,omitempty"`
	// Myrcene level in percent
	Myrcene *string  `xml:"MYRCENE" json:"myrcene,omitempty"`

	// Extensions

	// The amount of hops in this record along with the units formatted for easy display in the current user defined units.  For example “100 g” or “1.5 oz”.
	DisplayAmount *string `xml:"DISPLAY_AMOUNT" json:"display_amount,omitempty"`
	// Amount in inventory for this item along with the units – for example “10.0 oz”
	Inventory     *string `xml:"INVENTORY" json:"inventory,omitempty"`
	// Time displayed in minutes for all uses except for the dry hop which is in days.  For example “60 min”, “3 days”.
	DisplayTime   *string `xml:"DISPLAY_TIME" json:"display_time,omitempty"`
}

func (a Hop) MarshalJSON() ([]byte, error) {

	type Alias Hop
	t := func() int32 {
		if a.Type != nil {
			if t, ok := Hop_HopsType_value[strings.ToUpper(*a.Type)]; ok {
				return t
			}
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
		if a.Form != nil {
			if t, ok := Hop_HopsFormType_value[strings.ToUpper(*a.Form)]; ok {
				return t
			}
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
