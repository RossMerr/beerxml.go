package beerXML

import (
	"fmt"
	"strconv"
)

func percentToFloat(value string) *float64 {
	if value == "" {
		return nil
	}
	str := string(reg.Find([]byte(value)))
	if f, err := strconv.ParseFloat(str, 64); err == nil {
		return &f
	}
	return nil
}

func floatToPercent(value *float64) string {
	if value == nil {
		return ""
	}
	return fmt.Sprintf("%v %s", *value, "%")
}

func volumeToFloat(value string) *float64 {
	if value == "" {
		return nil
	}
	str := string(reg.Find([]byte(value)))
	if f, err := strconv.ParseFloat(str, 64); err == nil {
		return &f
	}
	return nil
}

func floatToVolume(value *float64) string {
	if value == nil {
		return ""
	}
	return fmt.Sprintf("%v %s", *value, "l")
}
