package beerXML

import (
	"bytes"
	"encoding/xml"
	"github.com/beerproto/beerxml.go/reader"
	"reflect"
	"strings"
	"testing"
)

func TestWater_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name    string
		xml string
		wantErr bool
	}{
		{
			name: "Basic",
			wantErr: false,
			xml: `<WATER>
	<NAME>Edinburg, Scotland</NAME>
	<VERSION>1</VERSION>
	<AMOUNT>18.927168</AMOUNT>
	<CALCIUM>120</CALCIUM>
	<BICARBONATE>225</BICARBONATE>
	<SULFATE>140</SULFATE>
	<CHLORIDE>20</CHLORIDE>
	<SODIUM>55</SODIUM>
	<MAGNESIUM>25</MAGNESIUM>
	<PH>8</PH>
	<NOTES>Used for dark, malty Brown ales with low bitterness.</NOTES>
	<DISPLAY_AMOUNT>5.00 gal</DISPLAY_AMOUNT>
</WATER>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := strings.NewReader(tt.xml)

			decoder := xml.NewDecoder(r)
			decoder.CharsetReader = reader.MakeCharsetReader
			dec := xml.NewTokenDecoder(reader.Trimmer{decoder}) // trimming decoder
			water := &Water{}
			err := dec.Decode(&water)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}

			var buf bytes.Buffer
			enc := xml.NewEncoder(&buf)

			enc.Indent("", "\t")

			err = enc.Encode(water)
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
