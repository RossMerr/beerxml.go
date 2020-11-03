package beerXML

import (
	"bytes"
	"encoding/xml"
	"github.com/beerproto/beerxml.go/reader"
	"reflect"
	"strings"
	"testing"
)

func TestFermentable_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name    string
		xml string
		wantErr bool
	}{
		{
			name: "Basic",
			wantErr: false,
			xml: `<FERMENTABLE>
	<NAME>Pale Malt (2 Row) UK</NAME>
	<VERSION>1</VERSION>
	<TYPE>Grain</TYPE>
	<AMOUNT>0</AMOUNT>
	<COLOR>3</COLOR>
	<ADD_AFTER_BOIL>false</ADD_AFTER_BOIL>
	<ORIGIN>United Kingdom</ORIGIN>
	<SUPPLIER></SUPPLIER>
	<NOTES>Base malt for all English beer styles Lower diastatic power than American 2 Row Pale Malt</NOTES>
	<DIASTATIC_POWER>45</DIASTATIC_POWER>
	<RECOMMEND_MASH>true</RECOMMEND_MASH>
	<IBU_GAL_PER_LB>0</IBU_GAL_PER_LB>
	<DISPLAY_AMOUNT>0.00 lb</DISPLAY_AMOUNT>
	<POTENTIAL>1.036</POTENTIAL>
	<INVENTORY>0.00 lb</INVENTORY>
	<DISPLAY_COLOR>3.0 SRM</DISPLAY_COLOR>
	<YIELD>78</YIELD>
	<COARSE_FINE_DIFF>1.5</COARSE_FINE_DIFF>
	<MOISTURE>4</MOISTURE>
	<PROTEIN>10.1</PROTEIN>
	<MAX_IN_BATCH>100</MAX_IN_BATCH>
</FERMENTABLE>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := strings.NewReader(tt.xml)

			decoder := xml.NewDecoder(r)
			decoder.CharsetReader = reader.MakeCharsetReader
			dec := xml.NewTokenDecoder(reader.Trimmer{decoder}) // trimming decoder
			fermentable := &Fermentable{}
			err := dec.Decode(&fermentable)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}

			var buf bytes.Buffer
			enc := xml.NewEncoder(&buf)

			enc.Indent("", "\t")

			err = enc.Encode(fermentable)
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
