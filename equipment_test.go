package beerXML

import (
	"bytes"
	"encoding/xml"
	"github.com/beerproto/beerxml.go/reader"
	"reflect"
	"strings"
	"testing"
)

func TestEquipment_UnmarshalXML(t *testing.T) {

	tests := []struct {
		name    string
		xml string
		wantErr bool
	}{
		{
			name: "Basic",
			wantErr: false,
			xml: `<EQUIPMENT>
	<NAME>Brew Pot (6+gal) and Igloo/Gott Cooler (5 Gal)</NAME>
	<VERSION>1</VERSION>
	<BOIL_SIZE>22.705669</BOIL_SIZE>
	<BATCH_SIZE>18.927168</BATCH_SIZE>
	<TUN_VOLUME>18.927168</TUN_VOLUME>
	<TUN_WEIGHT>1.814368</TUN_WEIGHT>
	<TUN_SPECIFIC_HEAT>0.3</TUN_SPECIFIC_HEAT>
	<TOP_UP_WATER>0</TOP_UP_WATER>
	<TRUB_CHILLER_LOSS>0.946358</TRUB_CHILLER_LOSS>
	<BOIL_TIME>60</BOIL_TIME>
	<CALC_BOIL_VOLUME>true</CALC_BOIL_VOLUME>
	<LAUTER_DEADSPACE>0.946358</LAUTER_DEADSPACE>
	<TOP_UP_KETTLE>0</TOP_UP_KETTLE>
	<HOP_UTILIZATION>100</HOP_UTILIZATION>
	<NOTES>Popular all grain setup. 5 Gallon Gott or Igloo cooler as mash tun with false bottom, and 7-9 gallon brewpot capable of boiling at least 6 gallons of wort. Primarily used for single infusion mashes.</NOTES>
	<DISPLAY_BOIL_SIZE>6.00 gal</DISPLAY_BOIL_SIZE>
	<DISPLAY_BATCH_SIZE>5.00 gal</DISPLAY_BATCH_SIZE>
	<DISPLAY_TUN_VOLUME>5.00 gal</DISPLAY_TUN_VOLUME>
	<DISPLAY_TUN_WEIGHT>4.00 lb</DISPLAY_TUN_WEIGHT>
	<DISPLAY_TOP_UP_WATER>0.00 gal</DISPLAY_TOP_UP_WATER>
	<DISPLAY_TRUB_CHILLER_LOSS>0.25 gal</DISPLAY_TRUB_CHILLER_LOSS>
	<DISPLAY_LAUTER_DEADSPACE>0.25 gal</DISPLAY_LAUTER_DEADSPACE>
	<DISPLAY_TOP_UP_KETTLE>0.00 gal</DISPLAY_TOP_UP_KETTLE>
	<EVAP_RATE>9</EVAP_RATE>
</EQUIPMENT>`,
		},
	}
	for _, tt := range tests {
		r := strings.NewReader(tt.xml)

		decoder := xml.NewDecoder(r)
		decoder.CharsetReader = reader.MakeCharsetReader
		dec := xml.NewTokenDecoder(reader.Trimmer{decoder}) // trimming decoder
		equipment := &Equipment{}
		err := dec.Decode(&equipment)
		if (err != nil) != tt.wantErr {
			t.Errorf("UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
		}

		var buf bytes.Buffer
		enc := xml.NewEncoder(&buf)

		enc.Indent("", "\t")

		err = enc.Encode(equipment)
		if (err != nil) != tt.wantErr {
			t.Errorf("MarshalXML() error = %v, wantErr %v", err, tt.wantErr)
		}

		got := buf.String()

		if !reflect.DeepEqual(tt.xml, got) {
			t.Errorf("object not equal \nexpected \n%v \ngot \n%v", tt.xml, got)
		}
	}
}
