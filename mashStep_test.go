package beerXML

import (
	"bytes"
	"encoding/xml"
	"github.com/beerproto/beerxml.go/reader"
	"reflect"
	"strings"
	"testing"
)

func TestMashStep_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
		xml string
	}{
		{
			name: "Basic",
			wantErr: false,
			xml: `<MASHSTEP>
	<NAME>Mash Out</NAME>
	<VERSION>1</VERSION>
	<TYPE>Temperature</TYPE>
	<INFUSE_AMOUNT>0</INFUSE_AMOUNT>
	<STEP_TEMP>75.555556</STEP_TEMP>
	<STEP_TIME>10</STEP_TIME>
	<RAMP_TIME>10</RAMP_TIME>
	<END_TEMP>75.555556</END_TEMP>
	<DESCRIPTION>Heat to 168.0 F over 10 min</DESCRIPTION>
	<WATER_GRAIN_RATIO>1.25</WATER_GRAIN_RATIO>
	<DECOCTION_AMT>0.00 qt</DECOCTION_AMT>
	<INFUSE_TEMP>182.0 F</INFUSE_TEMP>
	<DISPLAY_STEP_TEMP>DISPLAY_STEP_TEMP</DISPLAY_STEP_TEMP>
	<DISPLAY_INFUSE_AMT>0.00 qt</DISPLAY_INFUSE_AMT>
</MASHSTEP>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := strings.NewReader(tt.xml)

			decoder := xml.NewDecoder(r)
			decoder.CharsetReader = reader.MakeCharsetReader
			dec := xml.NewTokenDecoder(reader.Trimmer{decoder}) // trimming decoder
			mashStep := &MashStep{}
			err := dec.Decode(&mashStep)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}

			var buf bytes.Buffer
			enc := xml.NewEncoder(&buf)

			enc.Indent("", "\t")

			err = enc.Encode(mashStep)
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}

			got := buf.String()

			if !reflect.DeepEqual(tt.xml, got) {
				t.Errorf("object not equal \nexpected \n%v \ngot \n%v", tt.xml, got)
			}
		})
	}
}
