package beerXML

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

type Measureable struct {
	Value float64 `xml:"VALUE,omitempty" json:"value,omitempty"`
	Unit string `xml:"UNIT,omitempty" json:"unit,omitempty"`
}

func (s Measureable) String() string {
	return fmt.Sprint(s.Value)
}

func NewMeasureable(value *string) *Measureable {
	if value == nil {
		return nil
	}

	if *value == "" {
		return nil
	}

	v := strings.TrimSpace(*value)
	measureable := &Measureable{}
	str := string(reg.Find([]byte(v)))
	if f, err := strconv.ParseFloat(str, 64); err == nil {
		measureable.Value = f
	}

	index := strings.Index(v, " ")
	if index > 0 {
		suffix := v[index:]
		measureable.Unit = strings.TrimSpace(suffix)
	}

	return measureable
}


func (s *Measureable) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var value string

	err := d.DecodeElement(&value, &start)
	if err != nil {
		return err
	}
	measureable := NewMeasureable(&value)
	if measureable != nil {
		s.Unit = measureable.Unit
		s.Value = measureable.Value
	}

	return nil
}

func (s Measureable) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	value := strings.TrimSpace(fmt.Sprintf("%v %s", s.Value, s.Unit))

	err := e.EncodeElement(value, start)
	return err
}
