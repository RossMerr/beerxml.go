package beerXML

import (
	"bytes"
	"encoding/xml"
	"github.com/beerproto/beerxml.go/reader"
	"reflect"
	"strings"
	"testing"
)

func TestHop_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name    string

		xml string
		wantErr bool
	}{
		{
			name: "basic",
			xml: `<HOP>
	<ALPHA>5.5</ALPHA>
	<BETA>6</BETA>
	<HSI>50</HSI>
	<NAME>Cascade</NAME>
	<VERSION>1</VERSION>
	<AMOUNT>0</AMOUNT>
	<USE>Boil</USE>
	<TIME>0</TIME>
	<NOTES>Use For: American ales and lagers Aroma: Strong spicy, floral, grapefriut character Substitutes: Centennial Examples: Sierra Nevade Pale Ale, Anchor Liberty Ale A hops with Northern Brewers Heritage</NOTES>
	<TYPE>Both</TYPE>
	<FORM>Pellet</FORM>
	<ORIGIN>US</ORIGIN>
	<DISPLAY_AMOUNT>0.00 oz</DISPLAY_AMOUNT>
	<INVENTORY>0.00 oz</INVENTORY>
	<DISPLAY_TIME>-</DISPLAY_TIME>
</HOP>`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := strings.NewReader(tt.xml)

			decoder := xml.NewDecoder(r)
			decoder.CharsetReader = reader.MakeCharsetReader
			dec := xml.NewTokenDecoder(reader.Trimmer{decoder}) // trimming decoder
			hop := &Hop{}
			err := dec.Decode(&hop)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}

			var buf bytes.Buffer
			enc := xml.NewEncoder(&buf)

			enc.Indent("", "\t")

			err = enc.Encode(hop)
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
