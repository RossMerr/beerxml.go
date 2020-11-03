package beerXML

import (
	"bytes"
	"encoding/xml"
	"github.com/beerproto/beerxml.go/reader"
	"reflect"
	"strings"
	"testing"
)

func TestRecipe_UnmarshalXML(t *testing.T) {

	tests := []struct {
		name    string
		xml string
		wantErr bool
	}{
		{
			name: "Basic",
			wantErr: false,
			xml: `<RECIPE>
	<EST_ABV>5.3</EST_ABV>
	<ABV>5.1</ABV>
	<ACTUAL_EFFICIENCY>69.7</ACTUAL_EFFICIENCY>
	<EFFICIENCY>72</EFFICIENCY>
	<IBU>32.4</IBU>
	<NAME>Burton Ale</NAME>
	<VERSION>1</VERSION>
	<TYPE>All Grain</TYPE>
	<BREWER>Brad Smith</BREWER>
	<ASST_BREWER></ASST_BREWER>
	<BATCH_SIZE>18.927168</BATCH_SIZE>
	<BOIL_SIZE>20.819885</BOIL_SIZE>
	<BOIL_TIME>60</BOIL_TIME>
	<HOPS></HOPS>
	<FERMENTABLES></FERMENTABLES>
	<MISCS></MISCS>
	<YEASTS></YEASTS>
	<WATERS></WATERS>
	<NOTES></NOTES>
	<TASTE_NOTES>A smooth tasting pale ale -- full in body and a great long lasting head. Another great beer to stock -- favorite with guests. Make&#39;s a great black and tan when combined with light bodied Irish stout.</TASTE_NOTES>
	<TASTE_RATING>38</TASTE_RATING>
	<OG>1.054</OG>
	<FG>1.015</FG>
	<FERMENTATION_STAGES>2</FERMENTATION_STAGES>
	<PRIMARY_AGE>4</PRIMARY_AGE>
	<PRIMARY_TEMP>20</PRIMARY_TEMP>
	<SECONDARY_AGE>7</SECONDARY_AGE>
	<SECONDARY_TEMP>20</SECONDARY_TEMP>
	<TERTIARY_AGE>0</TERTIARY_AGE>
	<AGE>14</AGE>
	<AGE_TEMP>11.111</AGE_TEMP>
	<DATE>4/6/2003</DATE>
	<CARBONATION>2.4</CARBONATION>
	<EST_OG>1.056 SG</EST_OG>
	<EST_FG>1.015 SG</EST_FG>
	<EST_COLOR>7 SRM</EST_COLOR>
	<IBU_METHOD>Tinseth</IBU_METHOD>
	<CALORIES>242 cal/pint</CALORIES>
	<DISPLAY_BATCH_SIZE>5.00 gal</DISPLAY_BATCH_SIZE>
	<DISPLAY_BOIL_SIZE>5.50 gal</DISPLAY_BOIL_SIZE>
	<DISPLAY_OG>1.054 SG</DISPLAY_OG>
	<DISPLAY_FG>1.015 SG</DISPLAY_FG>
	<DISPLAY_PRIMARY_TEMP>68.0 F</DISPLAY_PRIMARY_TEMP>
	<DISPLAY_SECONDARY_TEMP>68.0 F</DISPLAY_SECONDARY_TEMP>
	<DISPLAY_TERTIARY_TEMP>68.0 F</DISPLAY_TERTIARY_TEMP>
	<DISPLAY_AGE_TEMP>52.0 F</DISPLAY_AGE_TEMP>
	<CARBONATION_USED>12 PSI</CARBONATION_USED>
</RECIPE>`,
		},
	}
	for _, tt := range tests {

		r := strings.NewReader(tt.xml)

		decoder := xml.NewDecoder(r)
		decoder.CharsetReader = reader.MakeCharsetReader
		dec := xml.NewTokenDecoder(reader.Trimmer{decoder}) // trimming decoder
		recipe := &Recipe{}
		err := dec.Decode(&recipe)
		if (err != nil) != tt.wantErr {
			t.Errorf("UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
		}

		var buf bytes.Buffer
		enc := xml.NewEncoder(&buf)

		enc.Indent("", "\t")

		err = enc.Encode(recipe)
		if (err != nil) != tt.wantErr {
			t.Errorf("MarshalXML() error = %v, wantErr %v", err, tt.wantErr)
		}

		got := buf.String()

		if !reflect.DeepEqual(tt.xml, got) {
			t.Errorf("object not equal \nexpected \n%v \ngot \n%v", tt.xml, got)
		}
	}
}
