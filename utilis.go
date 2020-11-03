package beerXML

import (
	"fmt"
	"strconv"
)

var WithUnits bool = false

func percentToFloat(value *string) *float64 {
	if value == nil {
		return nil
	}
	if *value == "" {
		return nil
	}
	str := string(reg.Find([]byte(*value)))
	if f, err := strconv.ParseFloat(str, 64); err == nil {
		return &f
	}
	return nil
}

func floatToPercent(value *float64) string {
	if value == nil {
		return ""
	}
	if WithUnits {
		return fmt.Sprintf("%v %s", *value, "%")
	} else {
		return fmt.Sprintf("%v", *value)
	}
}

func volumeToFloat(value *string) *float64 {
	if value == nil {
		return nil
	}
	if *value == "" {
		return nil
	}

	str := string(reg.Find([]byte(*value)))
	if f, err := strconv.ParseFloat(str, 64); err == nil {
		return &f
	}
	return nil
}

func floatToVolume(value *float64) string {
	if value == nil {
		return ""
	}
	if WithUnits {
		return fmt.Sprintf("%v %s", *value, "l")
	} else {
		return fmt.Sprintf("%v", *value)
	}
}

func specificGravityToFloat(value *string) *float64 {
	if value == nil {
		return nil
	}
	if *value == "" {
		return nil
	}

	str := string(reg.Find([]byte(*value)))
	if f, err := strconv.ParseFloat(str, 64); err == nil {
		return &f
	}
	return nil
}

func floatToSpecificGravity(value *float64) *string {
	if value == nil {
		return nil
	}

	var result string
	if WithUnits {
		result = fmt.Sprintf("%v %s", *value, "SG")
	} else {
		result =  fmt.Sprintf("%v", *value)
	}

	return &result
}

func ibuToFloat(value *string) *float64 {
	if value == nil {
		return nil
	}
	if *value == "" {
		return nil
	}

	str := string(reg.Find([]byte(*value)))
	if f, err := strconv.ParseFloat(str, 64); err == nil {
		return &f
	}
	return nil
}

func floatToIBU(value *float64) *string {
	if value == nil {
		return nil
	}

	var result string
	if WithUnits {
		result = fmt.Sprintf("%v %s", *value, "SRM")
	} else {
		result =  fmt.Sprintf("%v", *value)
	}

	return &result
}
