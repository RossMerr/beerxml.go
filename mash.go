package beerXML

type Mash struct {
	Name              string    `xml:"NAME" json:"name,omitempty"`
	Version           int32     `xml:"VERSION" json:"version,omitempty"`
	Graintemp         float64   `xml:"GRAIN_TEMP" json:"grain_temp,omitempty"`
	Tuntemp           float64   `xml:"TUN_TEMP" json:"tun_temp,omitempty"`
	Spargetemp        float64   `xml:"SPARGE_TEMP" json:"sparge_temp,omitempty"`
	PH                float32   `xml:"PH" json:"ph,omitempty"`
	Tunweight         float32   `xml:"TUN_WEIGHT" json:"tun_weight,omitempty"`
	Tunspecificheat   float32   `xml:"TUN_SPECIFIC_HEAT" json:"tun_specific_heat,omitempty"`
	Equipadjust       bool      `xml:"EQUIP_ADJUST" json:"equip_adjust,omitempty"`
	Notes             string    `xml:"NOTES" json:"notes,omitempty"`
	Mashsteps         MashSteps `xml:"MASH_STEPS" json:"mash_steps,omitempty"`
	DisplayGrainTemp  string    `xml:"DISPLAY_GRAIN_TEMP" json:"display_grain_temp,omitempty"`
	DisplaySpargeTemp string    `xml:"DISPLAY_SPARGE_TEMP" json:"display_sparge_temp,omitempty"`
	DisplayTunTemp    string    `xml:"DISPLAY_TUN_TEMP" json:"display_tun_temp,omitempty"`
	DisplayTunWeight  string    `xml:"DISPLAY_TUN_WEIGHT" json:"display_tun_weight,omitempty"`
}
