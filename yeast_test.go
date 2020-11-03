package beerXML
import (
	"bytes"
	"encoding/xml"
	"github.com/beerproto/beerxml.go/reader"
	"reflect"
	"strings"
	"testing"
)

func TestYeast_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name    string
		xml string
		wantErr bool
	}{
		{
			name: "Basic",
			wantErr: false,
			xml: `<YEAST>
	<NAME>European Ale</NAME>
	<VERSION>1</VERSION>
	<TYPE>Ale</TYPE>
	<FORM>Liquid</FORM>
	<AMOUNT>0.035</AMOUNT>
	<AMOUNT_IS_WEIGHT>false</AMOUNT_IS_WEIGHT>
	<LABORATORY>White Labs</LABORATORY>
	<PRODUCT_ID>WLP011</PRODUCT_ID>
	<MIN_TEMPERATURE>18.3</MIN_TEMPERATURE>
	<MAX_TEMPERATURE>21.1</MAX_TEMPERATURE>
	<FLOCCULATION>Medium</FLOCCULATION>
	<NOTES>Malty, Northern European ale yeast. Low ester production, low sulfer, gives a clean profile. Low attenuation contributes to malty taste.</NOTES>
	<BEST_FOR>Alt, Kolsch, malty English Ales, Fruit beers</BEST_FOR>
	<TIMES_CULTURED>0</TIMES_CULTURED>
	<MAX_REUSE>5</MAX_REUSE>
	<ADD_TO_SECONDARY>false</ADD_TO_SECONDARY>
	<DISPLAY_AMOUNT>35 ml</DISPLAY_AMOUNT>
	<DISP_MIN_TEMP>64.9 F</DISP_MIN_TEMP>
	<DISP_MAX_TEMP>70.0 F</DISP_MAX_TEMP>
	<INVENTORY>0 Pkgs</INVENTORY>
	<CULTURE_DATE>8/31/2003</CULTURE_DATE>
	<ATTENUATION>67.5</ATTENUATION>
</YEAST>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := strings.NewReader(tt.xml)

			decoder := xml.NewDecoder(r)
			decoder.CharsetReader = reader.MakeCharsetReader
			dec := xml.NewTokenDecoder(reader.Trimmer{decoder}) // trimming decoder
			yeast := &Yeast{}
			err := dec.Decode(&yeast)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}

			var buf bytes.Buffer
			enc := xml.NewEncoder(&buf)

			enc.Indent("", "\t")

			err = enc.Encode(yeast)
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
