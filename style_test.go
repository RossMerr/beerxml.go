package beerXML_test

import (
	"encoding/xml"
	"testing"

	beerXML "github.com/RossMerr/beerxml.go"
)

func TestStyle_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name    string
		xml     string
		wantErr bool
	}{
		{
			name: "CARB_MAX",
			xml: `<STYLE>
			<NAME>
				English Pale Ale
			</NAME>
			<CARB_MAX>
				2.4 >
			</CARB_MAX>
			</STYLE>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := []byte(tt.xml)
			s := &beerXML.Style{}
			err := xml.Unmarshal(data, s)
			if err != nil {
				t.Error(err)
			}
		})
	}
}
