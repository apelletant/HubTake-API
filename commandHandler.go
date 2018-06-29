package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func addCommand(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var expectedBody endpoints.addCommand
	if err := readJSONBody(r, &expectedBody); err != nil {
		writeResponse(w, http.StatusNotAcceptable,
			fmt.Sprintf("HubTake-api: %s", err.Error()))
		return
	}
	o, err := ep.AddCommand(db, expectedBody)
	if err != nil {
		writeResponse(w, http.StatusNotAcceptable,
			fmt.Sprintf("HubTake-api: color already exists"))
		return
	}
	writeJSONResponse(w, http.StatusOK, o)
	return
}

func getCommands(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	cmd, err := ep.GetCommands(db)
	if err != nil {
		writeResponse(w, http.StatusBadGateway, fmt.Sprintf("Hubtake-API error: %s", err.Error))
	} else if len(cmd) == 0 {
		writeResponse(w, http.StatusNotFound, "No command found")
	} else {
		rawBody, err := json.Marshal(cmd)
		if err != nil {
			writeResponse(w, http.StatusBadGateway, fmt.Sprintf("Hubtake-API error: %s", err))
		} else {
			writeResponse(w, http.StatusOK, string(rawBody))
		}
	}
}

func getCommandCustomer(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	cmd, err := ep.GetCommands(db)
	if err != nil {
		writeResponse(w, http.StatusBadGateway, fmt.Sprintf("Hubtake-API error: %s", err.Error))
	} else if len(cmd) == 0 {
		writeResponse(w, http.StatusNotFound, "No command found")
	} else {
		rawBody, err := json.Marshal(cmd)
		if err != nil {
			writeResponse(w, http.StatusBadGateway, fmt.Sprintf("Hubtake-API error: %s", err))
		} else {
			writeResponse(w, http.StatusOK, string(rawBody))
		}
	}
}

func deleteCommand(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("idCommand"))
	if err != nil {
		writeResponse(w, http.StatusBadGateway, fmt.Sprintf("Hubtake-API error: %s", err))
	} else {
		err := ep.DeleteCommand(db, id)
		if err != nil {
			writeResponse(w, http.StatusBadGateway, fmt.Sprintf("Hubtake-API error: %s", err))
		}
	}
	writeResponse(w, http.StatusOK, "")
}
