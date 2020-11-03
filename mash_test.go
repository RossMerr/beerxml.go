package beerXML

import (
	"bytes"
	"encoding/xml"
	"github.com/beerproto/beerxml.go/reader"
	"reflect"
	"strings"
	"testing"
)

func TestMash_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name    string
		xml string
		wantErr bool
	}{
		{
			name: "Basic",
			wantErr: false,
			xml: `<MASH>
	<NAME>Single Infusion, Light Body, No Mash Out</NAME>
	<VERSION>1</VERSION>
	<GRAIN_TEMP>22.222222</GRAIN_TEMP>
	<MASH_STEPS>
		<MASH_STEP>
			<NAME>Mash In</NAME>
			<VERSION>1</VERSION>
			<TYPE>Infusion</TYPE>
			<INFUSE_AMOUNT>11.82948</INFUSE_AMOUNT>
			<STEP_TEMP>65.555556</STEP_TEMP>
			<STEP_TIME>75</STEP_TIME>
			<RAMP_TIME>2</RAMP_TIME>
			<END_TEMP>65.555556</END_TEMP>
			<DESCRIPTION>Add 12.50 qt of water at 161.4 F</DESCRIPTION>
			<WATER_GRAIN_RATIO>1.25</WATER_GRAIN_RATIO>
			<DECOCTION_AMT>0.00 qt</DECOCTION_AMT>
			<INFUSE_TEMP>161.4 F</INFUSE_TEMP>
			<DISPLAY_STEP_TEMP>DISPLAY_STEP_TEMP</DISPLAY_STEP_TEMP>
			<DISPLAY_INFUSE_AMT>12.50 qt</DISPLAY_INFUSE_AMT>
		</MASH_STEP>
	</MASH_STEPS>
	<NOTES>Simple single infusion mash for use with most modern well modified grains (about 95% of the time).</NOTES>
	<TUN_TEMP>22.222222</TUN_TEMP>
	<SPARGE_TEMP>75.555556</SPARGE_TEMP>
	<PH>5.4</PH>
	<TUN_WEIGHT>0</TUN_WEIGHT>
	<TUN_SPECIFIC_HEAT>0.12</TUN_SPECIFIC_HEAT>
	<EQUIP_ADJUST>false</EQUIP_ADJUST>
	<DISPLAY_GRAIN_TEMP>72.0 F</DISPLAY_GRAIN_TEMP>
	<DISPLAY_TUN_TEMP>72.0</DISPLAY_TUN_TEMP>
	<DISPLAY_SPARGE_TEMP>168.0 F</DISPLAY_SPARGE_TEMP>
	<DISPLAY_TUN_WEIGHT>0.00 lb</DISPLAY_TUN_WEIGHT>
</MASH>`,
		},
	}
	for _, tt := range tests {
		r := strings.NewReader(tt.xml)

		decoder := xml.NewDecoder(r)
		decoder.CharsetReader = reader.MakeCharsetReader
		dec := xml.NewTokenDecoder(reader.Trimmer{decoder}) // trimming decoder
		mash := &Mash{}
		err := dec.Decode(&mash)
		if (err != nil) != tt.wantErr {
			t.Errorf("UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
		}

		var buf bytes.Buffer
		enc := xml.NewEncoder(&buf)

		enc.Indent("", "\t")

		err = enc.Encode(mash)
		if (err != nil) != tt.wantErr {
			t.Errorf("MarshalXML() error = %v, wantErr %v", err, tt.wantErr)
		}

		got := buf.String()

		if !reflect.DeepEqual(tt.xml, got) {
			t.Errorf("object not equal \nexpected \n%v \ngot \n%v", tt.xml, got)
		}
	}
}
