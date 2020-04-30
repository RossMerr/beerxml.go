package beerXML

import (
	"encoding/json"
	"strings"
)

type RECIPES struct {
	Recipe []Recipe `xml:"RECIPE" json:"recipes,omitempty"`
}

type Recipe struct {
	Name                 string       `xml:"NAME" json:"name,omitempty"`
	Version              int32        `xml:"VERSION" json:"version,omitempty"`
	Type                 string       `xml:"TYPE" json:"type,omitempty"`
	Brewer               string       `xml:"BREWER" json:"author,omitempty"`
	AsstBrewer           string       `xml:"ASST_BREWER" json:"coauthor,omitempty"`
	BatchSize            float32      `xml:"BATCH_SIZE" json:"batch_size,omitempty"`
	BoilSize             float32      `xml:"BOIL_SIZE" json:"boil_size,omitempty"`
	BoilTime             float32      `xml:"BOIL_TIME" json:"boil_time,omitempty"`
	Efficiency           float32      `xml:"EFFICIENCY" json:"efficiency,omitempty"`
	Hops                 *Hops         `xml:"HOPS" json:"hops,omitempty"`
	Fermentables         *Fermentables `xml:"FERMENTABLES" json:"fermentables,omitempty"`
	Miscs                *MISCS        `xml:"MISCS" json:"miscs,omitempty"`
	Yeasts               *Yeasts       `xml:"YEASTS" json:"yeasts,omitempty"`
	Waters               *Waters       `xml:"WATERS" json:"waters,omitempty"`
	Style                *Style        `xml:"STYLE" json:"style,omitempty"`
	Equipment            *Equipment    `xml:"EQUIPMENT" json:"equipment,omitempty"`
	Mash                 *Mash         `xml:"MASH" json:"mash,omitempty"`
	Notes                string       `xml:"NOTES" json:"notes,omitempty"`
	Tastenotes           string       `xml:"TASTE_NOTES" json:"taste_notes,omitempty"`
	Tasterating          float32      `xml:"TASTE_RATING" json:"taste_rating,omitempty"`
	OG                   float32      `xml:"OG" json:"og,omitempty"`
	FG                   float32      `xml:"FG" json:"fg,omitempty"`
	Carbonation          float32      `xml:"CARBONATION" json:"carbonation,omitempty"`
	Fermentationstages   int32        `xml:"FERMENTATION_STAGES" json:"fermentation_stages,omitempty"`
	Primaryage           int32        `xml:"PRIMARY_AGE" json:"primary_age,omitempty"`
	Primarytemp          float64      `xml:"PRIMARY_TEMP" json:"primary_temp,omitempty"`
	Secondaryage         int32        `xml:"SECONDARY_AGE" json:"secondary_age,omitempty"`
	Secondarytemp        float64      `xml:"SECONDARY_TEMP" json:"secondary_temp,omitempty"`
	Tertiaryage          int32        `xml:"TERTIARY_AGE" json:"tertiary_age,omitempty"`
	Tertiarytemp         float64      `xml:"TERTIARY_TEMP" json:"tertiary_temp,omitempty"`
	Age                  float32      `xml:"AGE" json:"age,omitempty"`
	Agetemp              float64      `xml:"AGE_TEMP" json:"age_temp,omitempty"`
	Date                 string       `xml:"DATE" json:"date,omitempty"`
	ABV                  string       `xml:"ABV" json:"abv,omitempty"`
	ActualEfficiency     string       `xml:"ACTUAL_EFFICIENCY" json:"actual_efficiency,omitempty"`
	Calories             string       `xml:"CALORIES" json:"calories,omitempty"`
	CarbonationUsed      string       `xml:"CARBONATION_USED" json:"carbonation_used,omitempty"`
	DisplayAgeTemp       string       `xml:"DISPLAY_AGE_TEMP" json:"display_age_temp,omitempty"`
	DisplayBatchSize     string       `xml:"DISPLAY_BATCH_SIZE" json:"DISPLAY_BATCH_SIZE,omitempty"`
	DisplayBoilSize      string       `xml:"DISPLAY_BOIL_SIZE" json:display_boil_size,omitempty"`
	DisplayFG            string       `xml:"DISPLAY_FG" json:"display_fg,omitempty"`
	DisplayOG            string       `xml:"DISPLAY_OG" json:"display_og,omitempty"`
	DisplayPrimaryTemp   string       `xml:"DISPLAY_PRIMARY_TEMP" json:"display_primary_temp,omitempty"`
	DisplaySecondaryTemp string       `xml:"DISPLAY_SECONDARY_TEMP" json:"display_secondary_temp,omitempty"`
	DisplayTertiaryTemp  string       `xml:"DISPLAY_TERTIARY_TEMP" json:"display_tertiary_temp,omitempty"`
	EstABV               string       `xml:"EST_ABV" json:"est_abv,omitempty"`
	EstColor             string       `xml:"EST_COLOR" json:"est_color,omitempty"`
	EstFF                string       `xml:"EST_FG" json:"est_fg,omitempty"`
	EstOG                string       `xml:"EST_OG" json:"est_og,omitempty"`
	IBU                  string       `xml:"IBU" json:"ibu,omitempty"`
	IBUMethod            string       `xml:"IBU_METHOD" json:"ibu_method,omitempty"`
}

func (a Recipe) MarshalJSON() ([]byte, error) {

	type Alias Recipe
	t := func() int32 {
		if t, ok := Recipe_RecipeType_value[strings.ToUpper(a.Type)]; ok {
			return t
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
