package beerXML

import (
	"bytes"
	"encoding/xml"
	"github.com/beerproto/beerxml.go/reader"
	"reflect"
	"strings"
	"testing"
)

func TestStyle_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name    string
		xml     string
		wantErr bool
	}{
		{
			name: "Basic",
			xml: `<STYLE>
	<NAME>Bohemian Pilsner</NAME>
	<CATEGORY>European Pale Lager</CATEGORY>
	<VERSION>1</VERSION>
	<CATEGORY_NUMBER>2</CATEGORY_NUMBER>
	<STYLE_LETTER>A</STYLE_LETTER>
	<STYLE_GUIDE>BJCP 1999</STYLE_GUIDE>
	<TYPE>Lager</TYPE>
	<OG_MIN>1.044</OG_MIN>
	<OG_MAX>1.056</OG_MAX>
	<FG_MIN>1.013</FG_MIN>
	<FG_MAX>1.017</FG_MAX>
	<IBU_MIN>35</IBU_MIN>
	<IBU_MAX>45</IBU_MAX>
	<COLOR_MIN>3</COLOR_MIN>
	<COLOR_MAX>5</COLOR_MAX>
	<NOTES>Famous beer from Pilsen, Czech Republic. Brewed with very soft water and high hop rates.</NOTES>
	<PROFILE>Light to medium body with some sweetness. Saaz hop flavor and aroma, but no lingering bitterness. Clean flavor, low diacetyls. Hoppy and malty with no aftertaste.</PROFILE>
	<INGREDIENTS>Saaz hops, light pilsner malt, pilsner yeast, soft water profile.</INGREDIENTS>
	<EXAMPLES>Budvar, Pilsner Urquell, Gambrinus</EXAMPLES>
	<DISPLAY_OG_MIN>1.044 SG</DISPLAY_OG_MIN>
	<DISPLAY_OG_MAX>1.056 SG</DISPLAY_OG_MAX>
	<DISPLAY_FG_MIN>1.013 SG</DISPLAY_FG_MIN>
	<DISPLAY_FG_MAX>1.017 SG</DISPLAY_FG_MAX>
	<DISPLAY_COLOR_MIN>3.0 SRM</DISPLAY_COLOR_MIN>
	<DISPLAY_COLOR_MAX>5.0 SRM</DISPLAY_COLOR_MAX>
	<OG_RANGE>1.044-1.056 SG</OG_RANGE>
	<FG_RANGE>1.013-1.017 SG</FG_RANGE>
	<IBU_RANGE>35.0-45.0 IBU</IBU_RANGE>
	<CARB_RANGE>2.3-2.6 vols</CARB_RANGE>
	<COLOR_RANGE>3.0-5.0 SRM</COLOR_RANGE>
	<ABV_RANGE>4.0-5.3 %</ABV_RANGE>
	<CARB_MIN>2.3</CARB_MIN>
	<CARB_MAX>2.6</CARB_MAX>
	<ABV_MIN>4</ABV_MIN>
	<ABV_MAX>5.3</ABV_MAX>
</STYLE>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := strings.NewReader(tt.xml)

			decoder := xml.NewDecoder(r)
			decoder.CharsetReader = reader.MakeCharsetReader
			dec := xml.NewTokenDecoder(reader.Trimmer{decoder}) // trimming decoder
			style := &Style{}
			err := dec.Decode(&style)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}

			var buf bytes.Buffer
			enc := xml.NewEncoder(&buf)

			enc.Indent("", "\t")

			err = enc.Encode(style)
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
