package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func getObject(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	o, err := ep.GetObjects(db)
	if len(o) == 0 {
		writeResponse(w, http.StatusNoContent, "")
		return
	}
	if err != nil {
		writeResponse(w, http.StatusInternalServerError, fmt.Sprintf("Hubtake api error; %v", err))
		return
	}
	rawBody, err := json.Marshal(o)
	if err != nil {
		writeResponse(w, http.StatusInternalServerError, fmt.Sprintf("Hubtake api error; %v", err))
		return
	}
	writeResponse(w, http.StatusOK, string(rawBody))
	return
}

func getObjectWithName(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	o := ep.GetObjectByName(db, "name")
	writeJSONResponse(w, 200, o)
	return
}

func getTakenObject(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	o := ep.GetTakenObject(db)
	rawBody, err := json.Marshal(o)
	if err != nil {
		writeResponse(w, 404, "not found")
	}
	writeResponse(w, 200, string(rawBody))
	return
}

func getNotTakenObject(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	o := ep.GetNotTakenObject(db)
	rawBody, _ := json.Marshal(o)
	writeResponse(w, 200, string(rawBody))
	return
}

func addObject(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	_, err := ep.AddObject(db, p.ByName("objectName"))
	if err != nil {
		writeResponse(w, http.StatusInternalServerError,
			"HubTake-api: add object error: "+err.Error())
		return
	}
	writeResponse(w, 204, "object added correctly")
	return
}

func deleteObject(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var status bool
	name := p.ByName("objectToDelete")
	fmt.Println(name)
	status = ep.DeleteObject(db, name)
	if status {
		writeResponse(w, 200,
			fmt.Sprintf("Remove successfully"))
		return
	}
	writeResponse(w, 404,
		fmt.Sprintf("Object not found"))
	return
}
