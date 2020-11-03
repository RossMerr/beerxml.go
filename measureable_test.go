package beerXML

import (
	"reflect"
	"testing"
)

func TestNewMeasureable(t *testing.T) {
	tests := []struct {
		name string
		value string
		want *Measureable
	}{
		{
			name: "Nil",
			value: "",
			want: nil,
		},
		{
			name: "123.5",
			value: "123.5",
			want: &Measureable{
				Value: 123.5,
			},
		},
		{
			name: "SG",
			value: "123.5 SG",
			want: &Measureable{
				Value: 123.5,
				Unit: "SG",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMeasureable(&tt.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMeasureable() = %v, want %v", got, tt.want)
			}
		})
	}
}
