package beerXML

import "regexp"

var reg *regexp.Regexp

func init() {
	// Find a number in a string
	reg, _ = regexp.Compile(`\d+\.?\d*`)
}
