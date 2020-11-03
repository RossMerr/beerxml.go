package beerXML

import (
	"encoding/xml"
	"strings"
)


// Encloses a set of one or more Hop records.
type Hops struct {
	Hop []Hop `xml:"HOP" json:"hop,omitempty"`
}


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
	Use     Hop_HopsUseType  `xml:"USE" json:"use,omitempty"`
	// The time as measured in minutes.  Meaning is dependent on the “USE” field.
	// For “Boil” this is the boil time.
	// For “Mash” this is the mash time.
	// For “First Wort” this is the boil time.
	// For “Aroma” this is the steep time.
	// For “Dry Hop” this is the amount of time to dry hop.
	Time    float64 `xml:"TIME" json:"time,omitempty"`
	// Textual notes about the hops, usage, substitutes.  May be a multiline entry.
	Notes   *string  `xml:"NOTES,omitempty" json:"notes,omitempty"`
	// May be "Bittering", "Aroma" or "Both"
	Type    Hop_HopsType  `xml:"TYPE" json:"type,omitempty"`
	// May be "Pellet", "Plug" or "Leaf"
	Form    Hop_HopsFormType  `xml:"FORM" json:"form,omitempty"`
	// Hop beta percentage - for example "4.4" denotes 4.4 % beta
	Beta    *float64 `xml:"BETA,omitempty" json:"beta,omitempty"`
	// Hop Stability Index - defined as the percentage of hop alpha lost in 6 months of storage
	HSI     *float64 `xml:"HSI,omitempty" json:"hsi,omitempty"`
	// Place of origin for the hops
	Origin  *string  `xml:"ORIGIN,omitempty" json:"origin,omitempty"`
	// Substitutes that can be used for this hops
	Substitutes *string `xml:"SUBSTITUTES,omitempty" json:"substitutes,omitempty"`
	// Humulene level in percent.
	Humulene *float64 `xml:"HUMULENE,omitempty" json:"humulene,omitempty"`
	// Caryophyllene level in percent.
	Caryophyllene *float64 `xml:"CARYOPHYLLENE,omitempty" json:"caryophyllene,omitempty"`
	// Cohumulone level in percent
	Cohumulone *float64 `xml:"COHUMULONE,omitempty" json:"cohumulone,omitempty"`
	// Myrcene level in percent
	Myrcene *float64  `xml:"MYRCENE,omitempty" json:"myrcene,omitempty"`

	// Extensions

	// The amount of hops in this record along with the units formatted for easy display in the current user defined units.  For example “100 g” or “1.5 oz”.
	DisplayAmount *string `xml:"DISPLAY_AMOUNT,omitempty" json:"display_amount,omitempty"`
	// Amount in inventory for this item along with the units – for example “10.0 oz”
	Inventory     *string `xml:"INVENTORY,omitempty" json:"inventory,omitempty"`
	// Time displayed in minutes for all uses except for the dry hop which is in days.  For example “60 min”, “3 days”.
	DisplayTime   *string `xml:"DISPLAY_TIME,omitempty" json:"display_time,omitempty"`
}

func (a *Hop) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Alias Hop
	aux := &struct {
		Alpha   *string `xml:"ALPHA" json:"alpha,omitempty"`
		Beta    *string `xml:"BETA" json:"beta,omitempty"`
		HSI     *string `xml:"HSI" json:"hsi,omitempty"`
		Humulene *string `xml:"HUMULENE" json:"humulene,omitempty"`
		Caryophyllene *string `xml:"CARYOPHYLLENE" json:"caryophyllene,omitempty"`
		Cohumulone *string `xml:"COHUMULONE" json:"cohumulone,omitempty"`
		Myrcene *string  `xml:"MYRCENE" json:"myrcene,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(a),
	}

	err := d.DecodeElement(aux, &start)
	if err != nil {
		return err
	}

	a.Alpha = *percentToFloat(aux.Alpha)
	a.Beta = percentToFloat(aux.Beta)
	a.HSI = percentToFloat(aux.HSI)
	a.Beta = percentToFloat(aux.Beta)
	a.Humulene = percentToFloat(aux.Humulene)
	a.Caryophyllene = percentToFloat(aux.Caryophyllene)
	a.Cohumulone = percentToFloat(aux.Cohumulone)
	a.Myrcene = percentToFloat(aux.Myrcene)
	return nil
}

func (a Hop) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type Alias Hop
	aux := &struct {
		Alpha   string `xml:"ALPHA" json:"alpha,omitempty"`
		Beta    string `xml:"BETA" json:"beta,omitempty"`
		HSI     string `xml:"HSI,omitempty" json:"hsi,omitempty"`
		Humulene string `xml:"HUMULENE,omitempty" json:"humulene,omitempty"`
		Caryophyllene string `xml:"CARYOPHYLLENE,omitempty" json:"caryophyllene,omitempty"`
		Cohumulone string `xml:"COHUMULONE,omitempty" json:"cohumulone,omitempty"`
		Myrcene string  `xml:"MYRCENE,omitempty" json:"myrcene,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(&a),
		Alpha: floatToPercent(&a.Alpha),
		Beta: floatToPercent(a.Beta),
		HSI: floatToPercent(a.HSI),
		Humulene: floatToPercent(a.Humulene),
		Caryophyllene: floatToPercent(a.Caryophyllene),
		Cohumulone: floatToPercent(a.Cohumulone),
		Myrcene: floatToPercent(a.Myrcene),
	}

	start.Name.Local = strings.ToUpper(start.Name.Local)

	err := e.EncodeElement(aux, start)
	if err != nil {
		return err
	}

	return nil
}


func (a *Hop_HopsUseType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var use string

	err := d.DecodeElement(&use, &start)
	if err != nil {
		return err
	}

	if value, ok := Hop_HopsUseType_value[use]; ok {
		*a = Hop_HopsUseType(value)
	} else {
		*a = HOP_HOPS_USE_NONE
	}

	return nil
}

func (a Hop_HopsUseType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if name, ok := Hop_HopsUseType_name[int32(a)]; ok {
		err := e.EncodeElement(name, start)
		if err != nil {
			return err
		}
	}
	return nil
}



func (a *Hop_HopsType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var use string

	err := d.DecodeElement(&use, &start)
	if err != nil {
		return err
	}

	if value, ok := Hop_HopsType_value[use]; ok {
		*a = Hop_HopsType(value)
	} else {
		*a = HOP_HOPS_NONE
	}

	return nil
}

func (a Hop_HopsType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if name, ok := Hop_HopsType_name[int32(a)]; ok {
		err := e.EncodeElement(name, start)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *Hop_HopsFormType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var use string

	err := d.DecodeElement(&use, &start)
	if err != nil {
		return err
	}

	if value, ok := Hop_HopsFormType_value[use]; ok {
		*a = Hop_HopsFormType(value)
	} else {
		*a = Hop_HOPS_FORM_NONE
	}

	return nil
}

func (a Hop_HopsFormType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if name, ok := Hop_HopsFormType_name[int32(a)]; ok {
		err := e.EncodeElement(name, start)
		if err != nil {
			return err
		}
	}
	return nil
}


type Hop_HopsType int32

const (
	// None
	HOP_HOPS_NONE Hop_HopsType = 0
	// Bittering
	HOP_BITTERING Hop_HopsType = 1
	// Aroma
	HOP_AROMA Hop_HopsType = 2
	// Both
	HOP_BOTH Hop_HopsType = 3
)

var Hop_HopsType_name = map[int32]string{
	0: "HOPS_NONE",
	1: "Bittering",
	2: "Aroma",
	3: "Both",
}

var Hop_HopsType_value = map[string]int32{
	"HOPS_NONE": 0,
	"Bittering": 1,
	"Aroma":     2,
	"Both":      3,
}

type Hop_HopsUseType int32

const (
	// None
	HOP_HOPS_USE_NONE Hop_HopsUseType = 0
	// Boil
	HOP_BOIL Hop_HopsUseType = 1
	// Dry Hop
	HOP_DRY_HOP Hop_HopsUseType = 2
	// Mash
	HOP_MASH Hop_HopsUseType = 3
	// First Wort
	HOP_FIRST_WORT Hop_HopsUseType = 4
	// Aroma
	HOP_AROMA_OTHER Hop_HopsUseType = 5
)

var Hop_HopsUseType_name = map[int32]string{
	0: "HOPS_USE_NONE",
	1: "Boil",
	2: "Dry Hop",
	3: "Mash",
	4: "First Wort",
	5: "Aroma",
}

var Hop_HopsUseType_value = map[string]int32{
	"HOPS_USE_NONE": 0,
	"Boil":          1,
	"Dry Hop":       2,
	"Mash":          3,
	"First Wort":    4,
	"Aroma":   5,
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
	1: "Pellet",
	2: "Plug",
	3: "Leaf",
}

var Hop_HopsFormType_value = map[string]int32{
	"HOPS_FORM_NONE": 0,
	"Pellet":         1,
	"Plug":           2,
	"Leaf":           3,
}
