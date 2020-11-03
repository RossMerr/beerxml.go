package beerXML

import (
	"bytes"
	"encoding/xml"
	"github.com/beerproto/beerxml.go/reader"
	"reflect"
	"strings"
	"testing"
)

func TestMISC_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name    string
		xml string
		wantErr bool
	}{
		{
			name: "Basic",
			wantErr: false,
			xml: `<MISC>
	<NAME>Irish Moss</NAME>
	<VERSION>1</VERSION>
	<TYPE>Fining</TYPE>
	<USE>Boil</USE>
	<TIME>10</TIME>
	<AMOUNT>0.001232</AMOUNT>
	<AMOUNT_IS_WEIGHT>false</AMOUNT_IS_WEIGHT>
	<USE_FOR>Clarity</USE_FOR>
	<NOTES>Fining agent that aids in the post-boil protein break. Reduces protein chill haze and improves beer clarity.</NOTES>
	<DISPLAY_AMOUNT>0.25 tsp</DISPLAY_AMOUNT>
	<INVENTORY>0.00 tsp</INVENTORY>
	<DISPLAY_TIME>10.0 min</DISPLAY_TIME>
</MISC>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := strings.NewReader(tt.xml)

			decoder := xml.NewDecoder(r)
			decoder.CharsetReader = reader.MakeCharsetReader
			dec := xml.NewTokenDecoder(reader.Trimmer{decoder}) // trimming decoder
			misc := &Misc{}
			err := dec.Decode(&misc)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}

			var buf bytes.Buffer
			enc := xml.NewEncoder(&buf)

			enc.Indent("", "\t")

			err = enc.Encode(misc)
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
