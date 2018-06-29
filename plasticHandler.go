package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"./endpoints"

	"github.com/julienschmidt/httprouter"
)

func addPlastic(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var expectedBody endpoints.Plastic
	if err := readJSONBody(r, &expectedBody); err != nil {
		writeResponse(w, http.StatusNotAcceptable,
			fmt.Sprintf("HubTake-api: %s", err.Error()))
		return
	}
	o, err := ep.AddPlastic(db, expectedBody)
	if err != nil {
		writeResponse(w, http.StatusNotAcceptable,
			fmt.Sprintf("HubTake-api: color already exists"))
		return
	}
	writeJSONResponse(w, http.StatusOK, o)
	return
}

func updatePlastic(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var expectedBody endpoints.Plastic
	if err := readJSONBody(r, &expectedBody); err != nil {
		writeResponse(w, http.StatusNotAcceptable,
			fmt.Sprintf("HubTake-api: %s", fmt.Errorf("Hubtake-API error: %v", err.Error())))
		return
	}
	o, err := ep.UpdatePlastic(db, expectedBody)
	if err != nil {
		writeResponse(w, http.StatusNotAcceptable,
			fmt.Sprintf("HubTake-api: color already exists"))
		return
	}
	writeJSONResponse(w, http.StatusOK, o)
	return
}

func getPlastics(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u, err := ep.GetPlastics(db)
	if err != nil {
		writeResponse(w, http.StatusBadGateway, "Internal server error")
		return
	}
	rawBodyByte, err := json.Marshal(u)
	if err != nil {
		writeResponse(w, http.StatusBadGateway, fmt.Sprintf("Internal server error: %s", err.Error))
		return
	}
	if len(u) == 0 {
		writeResponse(w, 204, "No content")
		return
	}
	writeResponse(w, http.StatusOK, string(rawBodyByte))
}

func getPlasticByColor(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u, err := ep.GetPlasticByColor(db, p.ByName("plasticColor"))
	if err != nil {
		writeResponse(w, http.StatusNotFound, "Plastic not found")
		return
	}
	rawBodyByte, err := json.Marshal(u)
	if err != nil {
		writeResponse(w, http.StatusBadGateway, fmt.Sprintf("Hubtake-API error: %s", err.Error))
		return
	}
	writeResponse(w, http.StatusOK, string(rawBodyByte))
}
