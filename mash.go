package beerXML

// A mash profile is a record used either within a recipe or outside the recipe to precisely specify the mash method
// used.  The record consists of some informational items followed by a <MASH_STEPS> record that contains the actual
// mash steps.
type Mash struct {
	// Name of the mash profile.
	Name              string    `xml:"NAME" json:"name,omitempty"`
	// Version of the mash record.  Should always be “1” for this version of the XML standard.
	Version           int32     `xml:"VERSION" json:"version,omitempty"`
	// The temperature of the grain before adding it to the mash in degrees Celsius.
	GrainTemp         float64   `xml:"GRAIN_TEMP" json:"grain_temp,omitempty"`
	// Record set that starts the list of <MASH_STEP> records.  All MASH_STEP records should appear between the
	// <MASH_STEPS> … </MASH_STEPS> pair.
	MashSteps         *MashSteps `xml:"MASH_STEPS" json:"mash_steps,omitempty"`
	// Notes associated with this profile – may be multiline.
	Notes             string    `xml:"NOTES" json:"notes,omitempty"`
	// Grain tun temperature – may be used to adjust the infusion temperature for equipment if the program supports it.
	// Measured in degrees C.
	TunTemp           *float64   `xml:"TUN_TEMP" json:"tun_temp,omitempty"`
	// Temperature of the sparge water used in degrees Celsius.
	SpargeTemp        *float64   `xml:"SPARGE_TEMP" json:"sparge_temp,omitempty"`
	// PH of the sparge.
	PH                *float64   `xml:"PH" json:"ph,omitempty"`
	// Weight of the mash tun in kilograms
	TunWeight         *float64   `xml:"TUN_WEIGHT" json:"tun_weight,omitempty"`
	// Specific heat of the tun material in calories per gram-degree C.
	TunSpecificHeat   *float64   `xml:"TUN_SPECIFIC_HEAT" json:"tun_specific_heat,omitempty"`
	// If TRUE, mash infusion and decoction calculations should take into account the temperature effects of the
	// equipment (tun specific heat and tun weight).  If FALSE, the tun is assumed to be pre-heated.  Default is FALSE.
	EquipAdjust       bool      `xml:"EQUIP_ADJUST" json:"equip_adjust,omitempty"`

	// Extensions

	// Grain temperature in user display units with the units.  For example: “72 F”.
	DisplayGrainTemp  string    `xml:"DISPLAY_GRAIN_TEMP" json:"display_grain_temp,omitempty"`
	// Tun temperature in user display units.  For example “68 F”
	DisplayTunTemp    string    `xml:"DISPLAY_TUN_TEMP" json:"display_tun_temp,omitempty"`
	// Sparge temperature in user defined units.  For example “178 F”
	DisplaySpargeTemp string    `xml:"DISPLAY_SPARGE_TEMP" json:"display_sparge_temp,omitempty"`
	// Tun weight in user defined units – for example “10 lb”
	DisplayTunWeight  string    `xml:"DISPLAY_TUN_WEIGHT" json:"display_tun_weight,omitempty"`
}
