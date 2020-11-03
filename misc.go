package beerXML

import (
	"encoding/xml"
	"strings"
)

// Encloses a set of one or more Misc records
type MISCS struct {
	MISC []Misc `xml:"MISC" json:"mise,omitempty"`
}

// The term "misc" encompasses all non-fermentable miscellaneous ingredients that are not hops or yeast and do not
// significantly change the gravity of the beer.  For example: spices, clarifying agents, water treatments, etc…
type Misc struct {
	// Name of the misc item.
	Name           string  `xml:"NAME" json:"name,omitempty"`
	// Version number of this element.  Should be “1” for this version.
	Version        int32   `xml:"VERSION" json:"version,omitempty"`
	// May be “Spice”, “Fining”, “Water Agent”, “Herb”, “Flavor” or “Other”
	Type           Misc_MiseType  `xml:"TYPE" json:"type,omitempty"`
	// May be “Boil”, “Mash”, “Primary”, “Secondary”, “Bottling”
	Use            Misc_MiseUseType  `xml:"USE" json:"use,omitempty"`
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


func (a *Misc) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Alias Misc
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(a),
	}

	err := d.DecodeElement(aux, &start)
	if err != nil {
		return err
	}

	return nil
}

func (a Misc) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type Alias Misc
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(&a),
	}
	start.Name.Local = strings.ToUpper(start.Name.Local)

	err := e.EncodeElement(aux, start)
	if err != nil {
		return err
	}

	return nil
}


func (a *Misc_MiseType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var use string

	err := d.DecodeElement(&use, &start)
	if err != nil {
		return err
	}

	if value, ok := Misc_MiseType_value[use]; ok {
		*a = Misc_MiseType(value)
	} else {
		*a = MISC_MISE_NONE
	}

	return nil
}

func (a Misc_MiseType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if name, ok := Misc_MiseType_name[int32(a)]; ok {
		err := e.EncodeElement(name, start)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *Misc_MiseUseType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var use string

	err := d.DecodeElement(&use, &start)
	if err != nil {
		return err
	}

	if value, ok := Misc_MiseUseType_value[use]; ok {
		*a = Misc_MiseUseType(value)
	} else {
		*a = MISC_USE_NONE
	}

	return nil
}

func (a Misc_MiseUseType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if name, ok := Misc_MiseUseType_name[int32(a)]; ok {
		err := e.EncodeElement(name, start)
		if err != nil {
			return err
		}
	}
	return nil
}

type Misc_MiseType int32

const (
	// None
	MISC_MISE_NONE Misc_MiseType = 0
	// Spice
	MISC_SPICE Misc_MiseType = 1
	// Fining
	MISC_FINING Misc_MiseType = 2
	// Water Agent
	MISC_WATER_AGENT Misc_MiseType = 3
	// Herb
	MISC_HERB Misc_MiseType = 4
	// Flavor
	MISC_FLAVOR Misc_MiseType = 5
	// Other
	MISC_OTHER Misc_MiseType = 6
)

var Misc_MiseType_name = map[int32]string{
	0: "MISE_NONE",
	1: "Spice",
	2: "Fining",
	3: "Water Agent",
	4: "Herb",
	5: "Flavor",
	6: "Other",
}

var Misc_MiseType_value = map[string]int32{
	"MISE_NONE":   0,
	"Spice":       1,
	"Fining":      2,
	"Water Agent": 3,
	"Herb":        4,
	"Flavor":      5,
	"Other":       6,
}

type Misc_MiseUseType int32

const (
	// None
	MISC_USE_NONE Misc_MiseUseType = 0
	// Boil
	MISC_BOIL Misc_MiseUseType = 1
	// Mash
	MISC_MASH Misc_MiseUseType = 2
	// Primary
	MISC_PRIMARY Misc_MiseUseType = 3
	// Secondary
	MISC_SECONDARY Misc_MiseUseType = 4
	// Bottling
	MISC_BOTTLING Misc_MiseUseType = 5
)

var Misc_MiseUseType_name = map[int32]string{
	0: "USE_NONE",
	1: "Boil",
	2: "Mash",
	3: "Primary",
	4: "Secondary",
	5: "Bottling",
}

var Misc_MiseUseType_value = map[string]int32{
	"USE_NONE":  0,
	"Boil":      1,
	"Mash":      2,
	"Primary":   3,
	"Secondary": 4,
	"Bottling":  5,
}
