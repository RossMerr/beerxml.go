package beerXML

type Equipment struct {
	Name                   string  `xml:"NAME" json:"name,omitempty"`
	Version                int32   `xml:"VERSION" json:"version,omitempty"`
	Boilsize               float32 `xml:"BOIL_SIZE" json:"boil_size,omitempty"`
	Batchsize              float32 `xml:"BATCH_SIZE" json:"batch_size,omitempty"`
	Tunvolume              float32 `xml:"TUN_VOLUME" json:"tun_volume,omitempty"`
	Tunweight              float32 `xml:"TUN_WEIGHT" json:"tun_weight,omitempty"`
	Tunspecificheat        float32 `xml:"TUN_SPECIFIC_HEAT" json:"tun_specific_heat,omitempty"`
	Topupwater             float32 `xml:"TOP_UP_WATER" json:"top_up_water,omitempty"`
	Trubchillerloss        float32 `xml:"TRUB_CHILLER_LOSS" json:"trub_chiller_loss,omitempty"`
	Evaprate               float32 `xml:"EVAP_RATE" json:"evap_rate,omitempty"`
	Boiltime               float32 `xml:"BOIL_TIME" json:"boil_time,omitempty"`
	Calcboilvolume         bool    `xml:"CALC_BOIL_VOLUME" json:"calc_boil_volume,omitempty"`
	Lauterdeadspace        float32 `xml:"LAUTER_DEADSPACE" json:"lauter_deadspace,omitempty"`
	Topupkettle            float32 `xml:"TOP_UP_KETTLE" json:"top_up_kettle,omitempty"`
	Hoputilization         float32 `xml:"HOP_UTILIZATION" json:"hop_utilization,omitempty"`
	Notes                  string  `xml:"NOTES" json:"notes,omitempty"`
	DisplayBatchSize       string  `xml:"DISPLAY_BATCH_SIZE" json:"DISPLAY_BATCH_SIZE,omitempty"`
	DisplayBoilSize        string  `xml:"DISPLAY_BOIL_SIZE" json:display_boil_size,omitempty"`
	DisplayLauterDeadspace string  `xml:"DISPLAY_LAUTER_DEADSPACE" json:"display_lauter_deadspace,omitempty"`
	DisplayTopUpKettle     string  `xml:"DISPLAY_TOP_UP_KETTLE" json:"display_top_up_kettle,omitempty"`
	DisplayTopUpWater      string  `xml:"DISPLAY_TOP_UP_WATER" json:"display_top_up_water,omitempty"`
	DisplayTrubChillerLoss string  `xml:"DISPLAY_TRUB_CHILLER_LOSS" json:"display_trub_chiller_loss,omitempty"`
	DisplayTunVolume       string  `xml:"DISPLAY_TUN_VOLUME" json:"display_tun_volume,omitempty"`
	DisplayTunWeight       string  `xml:"DISPLAY_TUN_WEIGHT" json:"display_tun_weight,omitempty"`
}
