package beerXML

import (
	"encoding/xml"
	"strings"
)

// Though an equipment record is optional, when used it in a recipe or on its own it provides details needed to
// calculate total water usage as well as water needed for each step.  It also contains information about the thermal
// parameters of the mash tun and large batch hop utilization factors.
type Equipment struct {
	// Name of the equipment profile – usually a text description of the brewing setup.
	Name string `xml:"NAME" json:"name,omitempty"`
	// Version of the equipment record.  Should always be “1” for this version of the XML standard.
	Version int32 `xml:"VERSION" json:"version,omitempty"`
	// The pre-boil volume used in this particular instance for this equipment setup.
	// Note that this may be a calculated value depending on the CALC_BOIL_VOLUME parameter.
	BoilSize float64 `xml:"BOIL_SIZE" json:"boil_size,omitempty"`
	// The target volume of the batch at the start of fermentation.
	BatchSize float64 `xml:"BATCH_SIZE" json:"batch_size,omitempty"`
	// Volume of the mash tun in liters.  This parameter can be used to calculate if a particular mash and grain
	// profile will fit in the mash tun.  It may also be used for thermal calculations
	// in the case of a partially full mash tun.
	TunVolume *float64 `xml:"TUN_VOLUME,omitempty" json:"tun_volume,omitempty"`
	// Weight of the mash tun in kilograms.
	// Used primarily to calculate the thermal parameters of the mash tun – in conjunction with the volume and specific
	// heat.
	TunWeight *float64 `xml:"TUN_WEIGHT,omitempty" json:"tun_weight,omitempty"`
	// The specific heat of the mash tun which is usually a function of the material it is made of.
	// Typical ranges are 0.1-0.25 for metal and 0.2-0.5 for plastic materials.
	TunSpecificHeat *float64 `xml:"TUN_SPECIFIC_HEAT,omitempty" json:"tun_specific_heat,omitempty"`
	// The amount of top up water normally added just prior to starting fermentation.  Usually used for extract brewing.
	TopUpWater *float64 `xml:"TOP_UP_WATER,omitempty" json:"top_up_water,omitempty"`
	// The amount of wort normally lost during transition from the boiler to the fermentation vessel.
	// Includes both unusable wort due to trub and wort lost to the chiller and transfer systems.
	TrubChillerLoss *float64 `xml:"TRUB_CHILLER_LOSS,omitempty" json:"trub_chiller_loss,omitempty"`
	// The percentage of wort lost to evaporation per hour of the boil.
	EvapRate *float64 `xml:"EVAP_RATE,omitempty" json:"evap_rate,omitempty"`
	// The normal amount of time one boils for this equipment setup.  This can be used with the evaporation rate to
	// calculate the evaporation loss.
	BoilTime *float64 `xml:"BOIL_TIME,omitempty" json:"boil_time,omitempty"`
	// Flag denoting that the program should calculate the boil size.  Flag may be TRUE or FALSE.
	// If TRUE, then BOIL_SIZE = (BATCH_SIZE – TOP_UP_WATER – TRUB_CHILLER_LOSS) * (1+BOIL_TIME * EVAP_RATE )
	// If set then the boil size should match this value.
	CalcBoilVolume *bool `xml:"CALC_BOIL_VOLUME,omitempty" json:"calc_boil_volume,omitempty"`
	// Amount lost to the lauter tun and equipment associated with the lautering process.
	LauterDeadspace *float64 `xml:"LAUTER_DEADSPACE,omitempty" json:"lauter_deadspace,omitempty"`
	// Amount normally added to the boil kettle before the boil.
	TopUpKettle *float64 `xml:"TOP_UP_KETTLE,omitempty" json:"top_up_kettle,omitempty"`
	// Large batch hop utilization.  This value should be 100% for batches less than 20 gallons,
	// but may be higher (200% or more) for very large batch equipment.
	HopUtilization *float64 `xml:"HOP_UTILIZATION,omitempty" json:"hop_utilization,omitempty"`
	// Notes associated with the equipment.  May be a multiline entry.
	Notes *string `xml:"NOTES,omitempty" json:"notes,omitempty"`

	// Extensions

	// The pre-boil volume normally used for a batch of this size shown in display volume units such as “5.5 gal”
	DisplayBoilSize *string `xml:"DISPLAY_BOIL_SIZE,omitempty" json:display_boil_size,omitempty"`
	// The target volume of the batch at the start of fermentation in display volume units such as “5.0 gal”
	DisplayBatchSize *string `xml:"DISPLAY_BATCH_SIZE,omitempty" json:"DISPLAY_BATCH_SIZE,omitempty"`
	// Volume of the mash tun in display units such as “10.0 gal” or “20.0 l”
	DisplayTunVolume *string `xml:"DISPLAY_TUN_VOLUME,omitempty" json:"display_tun_volume,omitempty"`
	// Weight of the mash tun in display units such as “3.0 kg” or “6.0 lb”
	DisplayTunWeight *string `xml:"DISPLAY_TUN_WEIGHT,omitempty" json:"display_tun_weight,omitempty"`
	// The amount of top up water normally added just prior to starting fermentation in display volume such as “1.0 gal”
	DisplayTopUpWater *string `xml:"DISPLAY_TOP_UP_WATER,omitempty" json:"display_top_up_water,omitempty"`
	// The amount of wort normally lost during transition from the boiler to the fermentation vessel.
	// Includes both unusable wort due to trub and wort lost to the chiller and transfer systems.
	// Expressed in user units - Ex: “1.5 qt”
	DisplayTrubChillerLoss *string `xml:"DISPLAY_TRUB_CHILLER_LOSS,omitempty" json:"display_trub_chiller_loss,omitempty"`
	// Amount lost to the lauter tun and equipment associated with the lautering process. Ex: “2.0 gal” or “1.0 l”
	DisplayLauterDeadspace *string `xml:"DISPLAY_LAUTER_DEADSPACE,omitempty" json:"display_lauter_deadspace,omitempty"`
	// Amount normally added to the boil kettle before the boil. Ex: “1.0 gal”
	DisplayTopUpKettle *string `xml:"DISPLAY_TOP_UP_KETTLE,omitempty" json:"display_top_up_kettle,omitempty"`
}

func (a *Equipment) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Alias Equipment
	aux := &struct {
		*Alias
		EvapRate *string `xml:"EVAP_RATE,omitempty" json:"evap_rate,omitempty"`
	}{
		Alias: (*Alias)(a),
	}

	err := d.DecodeElement(aux, &start)
	if err != nil {
		return err
	}

	a.EvapRate = percentToFloat(aux.EvapRate)

	return nil
}

func (a Equipment) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type Alias Equipment
	aux := &struct {
		*Alias
		EvapRate string `xml:"EVAP_RATE,omitempty" json:"evap_rate,omitempty"`
	}{
		Alias:    (*Alias)(&a),
		EvapRate: floatToPercent(a.EvapRate),
	}

	start.Name.Local = strings.ToUpper(start.Name.Local)

	err := e.EncodeElement(aux, start)
	if err != nil {
		return err
	}

	return nil
}
