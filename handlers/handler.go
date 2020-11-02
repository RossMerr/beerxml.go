package handlers

import (
	"encoding/xml"
	"net/http"

	beerXML "github.com/beerproto/beerxml.go"
	"github.com/beerproto/beerxml.go/reader"
)

func BeerXML(w http.ResponseWriter, r *http.Request, recipes *beerXML.RECIPES) {
	defer r.Body.Close()

	decoder := xml.NewDecoder(r.Body)
	decoder.CharsetReader = reader.MakeCharsetReader

	dec := xml.NewTokenDecoder(reader.Trimmer{decoder}) // trimming decoder

	err := dec.Decode(&recipes)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
