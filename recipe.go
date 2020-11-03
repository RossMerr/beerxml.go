package beerXML

import (
	"encoding/json"
	"encoding/xml"
	"strings"
)

// Encloses one or more recipe records.
type RECIPES struct {
	Recipe []Recipe `xml:"RECIPE" json:"recipes,omitempty"`
}

// A recipe record denotes a single recipe.  A recipe record may use records from any of the previously described
// record formats to specify ingredients and other data.
type Recipe struct {
	// Name of the recipe.
	Name string `xml:"NAME" json:"name,omitempty"`
	// Version of the recipe record.  Should always be “1” for this version of the XML standard.
	Version int32 `xml:"VERSION" json:"version,omitempty"`
	// May be one of “Extract”, “Partial Mash” or “All Grain”
	Type string `xml:"TYPE" json:"type,omitempty"`
	// The style of the beer this recipe is associated with.  All of the required items for a valid style should be
	// between the <STYLE>…</STYLE> tags.
	Style *Style `xml:"STYLE" json:"style,omitempty"`
	// An equipment record is optional.  If included the BATCH_SIZE and BOIL_SIZE in the equipment record must match
	// the values in this recipe record.
	Equipment *Equipment `xml:"EQUIPMENT" json:"equipment,omitempty"`
	// Name of the brewer
	Brewer string `xml:"BREWER" json:"author,omitempty"`
	// Optional name of the assistant brewer
	AsstBrewer *string `xml:"ASST_BREWER" json:"coauthor,omitempty"`
	// Target size of the finished batch in liters.
	BatchSize float64 `xml:"BATCH_SIZE" json:"batch_size,omitempty"`
	// Starting size for the main boil of the wort in liters.
	BoilSize float64 `xml:"BOIL_SIZE" json:"boil_size,omitempty"`
	// The total time to boil the wort in minutes.
	BoilTime float64 `xml:"BOIL_TIME" json:"boil_time,omitempty"`
	// The percent brewhouse efficiency to be used for estimating the starting gravity of the beer.
	// Not required for “Extract” recipes, but is required for “Partial Mash” and “All Grain” recipes.
	Efficiency *float64 `xml:"EFFICIENCY" json:"efficiency,omitempty"`
	// Zero or more HOP ingredient records may appear between the <HOPS>…</HOPS> tags.
	Hops *Hops `xml:"HOPS" json:"hops,omitempty"`
	// Zero or more FERMENTABLE ingredients may appear between the <FERMENTABLES> … </FERMENTABLES> tags.
	Fermentables *Fermentables `xml:"FERMENTABLES" json:"fermentables,omitempty"`
	// Zero or more MISC records may appear between <MISCS> … </MISCS>
	Miscs *MISCS `xml:"MISCS" json:"miscs,omitempty"`
	// Zero or more YEAST records may appear between <YEASTS> … </YEASTS>
	Yeasts *Yeasts `xml:"YEASTS" json:"yeasts,omitempty"`
	// Zero or more WATER records may appear between <WATERS> … </WATERS>
	Waters *Waters `xml:"WATERS" json:"waters,omitempty"`
	// A MASH profile record containing one or more MASH_STEPs.  NOTE: No Mash record is needed for “Extract” type
	// brews.
	Mash *Mash `xml:"MASH" json:"mash,omitempty"`
	// Notes associated with this recipe – may be multiline.
	Notes *string `xml:"NOTES" json:"notes,omitempty"`
	// Tasting notes – may be multiline.
	TasteNotes *string `xml:"TASTE_NOTES" json:"taste_notes,omitempty"`
	// Number between zero and 50.0 denoting the taste rating – corresponds to the 50 point BJCP rating system.
	TasteRating *float64 `xml:"TASTE_RATING" json:"taste_rating,omitempty"`
	// The measured original (pre-fermentation) specific gravity of the beer.
	OG *float64 `xml:"OG" json:"og,omitempty"`
	// The measured final gravity of the finished beer.
	FG *float64 `xml:"FG" json:"fg,omitempty"`
	// The number of fermentation stages used – typically a number between one and three
	FermentationStages *int32 `xml:"FERMENTATION_STAGES" json:"fermentation_stages,omitempty"`
	// Time spent in the primary in days
	PrimaryAge *int32 `xml:"PRIMARY_AGE" json:"primary_age,omitempty"`
	// Temperature in degrees Celsius for the primary fermentation.
	PrimaryTemp *float64 `xml:"PRIMARY_TEMP" json:"primary_temp,omitempty"`
	// Time spent in the secondary in days.
	SecondaryAge *int32 `xml:"SECONDARY_AGE" json:"secondary_age,omitempty"`
	// Temperature in degrees Celsius for the secondary fermentation.
	SecondaryTemp *float64 `xml:"SECONDARY_TEMP" json:"secondary_temp,omitempty"`
	// Time spent in the third fermenter in days.
	TertiaryAge *int32 `xml:"TERTIARY_AGE" json:"tertiary_age,omitempty"`
	// Temperature in the tertiary fermenter.
	TertiaryTemp *float64 `xml:"TERTIARY_TEMP" json:"tertiary_temp,omitempty"`
	// The time to age the beer in days after bottling.
	Age *float32 `xml:"AGE" json:"age,omitempty"`
	// Temperature for aging the beer after bottling.
	AgeTemp *float64 `xml:"AGE_TEMP" json:"age_temp,omitempty"`
	// Date brewed in a easily recognizable format such as “3 Dec 04”.
	Date *string `xml:"DATE" json:"date,omitempty"`
	// Floating point value corresponding to the target volumes of CO2 used to carbonate this beer.
	Carbonation *float64 `xml:"CARBONATION" json:"carbonation,omitempty"`
	// TRUE if the batch was force carbonated using CO2 pressure,
	// FALSE if the batch was carbonated using a priming agent.  Default is FALSE
	ForcedCarbonation *bool `xml:"FORCED_CARBONATION" json:"forced_carbonation,omitempty"`
	// Text describing the priming agent such as “Honey” or “Corn Sugar” – used only if this is not a forced
	// carbonation
	PrimingSugarName *string `xml:"PRIMING_SUGAR_NAME" json:"priming_sugar_name,omitempty"`
	// The temperature for either bottling or forced carbonation.
	CarbonationTemp *float64 `xml:"CARBONATION_TEMP" json:"carbonation_temp,omitempty"`
	// Factor used to convert this priming agent to an equivalent amount of corn sugar for a bottled scenario.
	// For example, “Dry Malt Extract” would have a value of 1.4 because it requires 1.4 times as much DME as
	// corn sugar to carbonate.  To calculate the amount of DME needed, the program can calculate the amount of
	// corn sugar needed and then multiply by this factor.
	PrimingSugarEquiv *float64 `xml:"PRIMING_SUGAR_EQUIV" json:"priming_sugar_equiv,omitempty"`
	// Used to factor in the smaller amount of sugar needed for large containers.
	// For example, this might be 0.5 for a typical 5 gallon keg since naturally priming a keg requires about 50%
	// as much sugar as priming bottles.
	KegPrimingFactor *float64 `xml:"KEG_PRIMING_FACTOR" json:"keg_priming_factor,omitempty"`

	// Extensions

	// Calculated estimate of the original gravity for this recipe along with the units.
	EstOG string `xml:"EST_OG" json:"est_og,omitempty"`
	// Calculated estimate for the final specific gravity of this recipe along with the units as in “1.015 sg”
	EstFF string `xml:"EST_FG" json:"est_fg,omitempty"`
	// The estimated color of the beer in user defined color units.
	EstColor string `xml:"EST_COLOR" json:"est_color,omitempty"`
	// The estimated bitterness level of the beer in IBUs
	IBU string `xml:"IBU" json:"ibu,omitempty"`
	// May be “Rager”, “Tinseth” or “Garetz” corresponding to the method/equation used to estimate IBUs for this recipe.
	IBUMethod string `xml:"IBU_METHOD" json:"ibu_method,omitempty"`
	// Estimated percent alcohol by volume for this recipe.
	EstABV *float64 `xml:"EST_ABV" json:"est_abv,omitempty"`
	// Actual alcohol by volume calculated from the OG and FG measured.
	ABV *float64 `xml:"ABV" json:"abv,omitempty"`
	// The actual efficiency as calculated using the measured original and final gravity.
	ActualEfficiency *float64 `xml:"ACTUAL_EFFICIENCY" json:"actual_efficiency,omitempty"`
	// Calorie estimate based on the measured starting and ending gravity.  Note that calories should be quoted in
	// “Cal” or kilocalories which is the normal dietary measure
	// (i.e. a beer is usually in the range of 100-250 calories per 12 oz).  Examples “180 Cal/pint”,
	Calories string `xml:"CALORIES" json:"calories,omitempty"`
	// Batch size in user defined units along with the units as in “5.0 gal”
	DisplayBatchSize string `xml:"DISPLAY_BATCH_SIZE" json:"DISPLAY_BATCH_SIZE,omitempty"`
	// Boil size with user defined units as in “6.3 gal”
	DisplayBoilSize string `xml:"DISPLAY_BOIL_SIZE" json:display_boil_size,omitempty"`
	// Measured original gravity in user defined units as in “6.4 plato”
	DisplayOG string `xml:"DISPLAY_OG" json:"display_og,omitempty"`
	// Measured final gravity in user defined units as in “1.035 sg”
	DisplayFG string `xml:"DISPLAY_FG" json:"display_fg,omitempty"`
	// Primary fermentation temperature in user defined units such as “64 F”
	DisplayPrimaryTemp string `xml:"DISPLAY_PRIMARY_TEMP" json:"display_primary_temp,omitempty"`
	// Secondary fermentation temperature in user defined units such as “56 F”
	DisplaySecondaryTemp string `xml:"DISPLAY_SECONDARY_TEMP" json:"display_secondary_temp,omitempty"`
	// Tertiary temperature in user defined units such as “20 C”
	DisplayTertiaryTemp string `xml:"DISPLAY_TERTIARY_TEMP" json:"display_tertiary_temp,omitempty"`
	// Temperature to use when aging the beer in user units such as “55 F”
	DisplayAgeTemp string `xml:"DISPLAY_AGE_TEMP" json:"display_age_temp,omitempty"`
	// Text description of the carbonation used such as “50g corn sugar” or “Kegged at 20psi”
	CarbonationUsed string `xml:"CARBONATION_USED" json:"carbonation_used,omitempty"`
	// Carbonation/Bottling temperature in appropriate units such as “40F” or “32 C”
	DisplayCarbTemp *string `xml:"DISPLAY_CARB_TEMP" json:"display_carb_temp,omitempty"`
}

func (a *Recipe) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Alias Recipe
	aux := &struct {
		EstABV           string `xml:"EST_ABV" json:"est_abv,omitempty"`
		ABV              string `xml:"ABV" json:"abv,omitempty"`
		ActualEfficiency string `xml:"ACTUAL_EFFICIENCY" json:"actual_efficiency,omitempty"`
		Efficiency       string `xml:"EFFICIENCY" json:"efficiency,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(a),
	}

	err := d.DecodeElement(aux, &start)
	if err != nil {
		return err
	}

	a.EstABV = percentToFloat(aux.EstABV)
	a.ABV = percentToFloat(aux.ABV)
	a.ActualEfficiency = percentToFloat(aux.ActualEfficiency)
	a.Efficiency = percentToFloat(aux.Efficiency)

	return nil
}

func (a Recipe) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type Alias Recipe
	aux := &struct {
		EstABV           string `xml:"EST_ABV" json:"est_abv,omitempty"`
		ABV              string `xml:"ABV" json:"abv,omitempty"`
		ActualEfficiency string `xml:"ACTUAL_EFFICIENCY" json:"actual_efficiency,omitempty"`
		Efficiency       string `xml:"EFFICIENCY" json:"efficiency,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(&a),
		EstABV: floatToPercent(a.EstABV),
		ABV: floatToPercent(a.ABV),
		ActualEfficiency: floatToPercent(a.ActualEfficiency),
		Efficiency: floatToPercent(a.Efficiency),
	}

	err := e.EncodeElement(aux, start)
	if err != nil {
		return err
	}

	return nil
}

func (a Recipe) MarshalJSON() ([]byte, error) {

	type Alias Recipe
	t := func() int32 {
		if t, ok := Recipe_RecipeType_value[strings.ToUpper(a.Type)]; ok {
			return  t
		}
		return int32(Recipe_RECIPE_NONE)
	}()

	return json.Marshal(&struct {
		Type int32 `json:"type,omitempty"`
		*Alias
	}{
		Type:  t,
		Alias: (*Alias)(&a),
	})
}

func (a *Recipe) UnmarshalJSON(b []byte) error {

	type Alias Recipe
	aux := &struct {
		EstABV           string `xml:"EST_ABV" json:"est_abv,omitempty"`
		ABV              string `xml:"ABV" json:"abv,omitempty"`
		ActualEfficiency string `xml:"ACTUAL_EFFICIENCY" json:"actual_efficiency,omitempty"`
		Efficiency       string `xml:"EFFICIENCY" json:"efficiency,omitempty"`

		*Alias
	}{
		Alias: (*Alias)(a),
	}
	if err := json.Unmarshal(b, &aux); err != nil {
		return err
	}

	a.EstABV = percentToFloat(aux.EstABV)
	a.ABV = percentToFloat(aux.ABV)
	a.ActualEfficiency = percentToFloat(aux.ActualEfficiency)
	a.Efficiency = percentToFloat(aux.Efficiency)

	return nil
}


type Recipe_RecipeType int32

const (
	// None
	Recipe_RECIPE_NONE Recipe_RecipeType = 0
	// Extract
	Recipe_EXTRACT Recipe_RecipeType = 1
	// Partial Mash
	Recipe_PARTIAL_MASH Recipe_RecipeType = 2
	// All Grain
	Recipe_ALL_GRAIN Recipe_RecipeType = 3
	// Cider
	Recipe_CIDER Recipe_RecipeType = 4
	// Kombucha
	Recipe_KOMBUCHA Recipe_RecipeType = 5
	// Soda
	Recipe_SODA Recipe_RecipeType = 6
	// Other
	Recipe_OTHER Recipe_RecipeType = 7
	// Mead
	Recipe_MEAD Recipe_RecipeType = 8
	// Wine
	Recipe_WINE Recipe_RecipeType = 9
)

var Recipe_RecipeType_name = map[int32]string{
	0: "RECIPE_NONE",
	1: "EXTRACT",
	2: "PARTIAL_MASH",
	3: "ALL_GRAIN",
	4: "CIDER",
	5: "KOMBUCHA",
	6: "SODA",
	7: "OTHER",
	8: "MEAD",
	9: "WINE",
}

var Recipe_RecipeType_value = map[string]int32{
	"RECIPE_NONE":  0,
	"EXTRACT":      1,
	"PARTIAL_MASH": 2,
	"ALL_GRAIN":    3,
	"CIDER":        4,
	"KOMBUCHA":     5,
	"SODA":         6,
	"OTHER":        7,
	"MEAD":         8,
	"WINE":         9,
}
