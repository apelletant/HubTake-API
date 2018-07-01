package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"./endpoints"

	"github.com/julienschmidt/httprouter"
)

func getUsers(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := ep.GetUsers(db)
	rawBody, err := json.Marshal(u)
	if err != nil {

	}
	writeResponse(w, 200, string(rawBody))
}

func getUserByMail(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := ep.GetUserByMail(db, p.ByName("userEmailGet"))
	rawBody, _ := json.Marshal(u)
	writeResponse(w, 200, string(rawBody))
}

func getUserByID(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u, err := ep.GetUserByID(db, p.ByName("userId"))
	if err != nil {
		writeResponse(w, http.StatusNotAcceptable,
			fmt.Sprintf("HubTake-api: %s", err.Error()))
		return
	}
	rawBody, err := json.Marshal(u)
	if err != nil {
		writeResponse(w, http.StatusNotAcceptable,
			fmt.Sprintf("HubTake-api: %s", err.Error()))
		return
	}
	writeResponse(w, 200, string(rawBody))
}

func addUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var expectedBody endpoints.UserPost
	if err := readJSONBody(r, &expectedBody); err != nil {
		writeResponse(w, http.StatusNotAcceptable,
			fmt.Sprintf("HubTake-api: %s", err.Error()))
		return
	}
	o, err := ep.AddUser(db, expectedBody)
	if err != nil {
		writeResponse(w, http.StatusInternalServerError,
			fmt.Sprintf("HubTake-api: %s", err.Error()))
		return
	}
	writeJSONResponse(w, http.StatusOK, o)
	return
}

func userTakeObject(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var expectedBody endpoints.BorrowReturnData
	if err := readJSONBody(r, &expectedBody); err != nil {
		writeResponse(w, 406,
			fmt.Sprintf("HubTake-api: %s", err.Error()))
		return
	}
	err := ep.UserTakeObject(db, expectedBody)
	if err != nil {
		writeResponse(w, 404,
			fmt.Sprintf("HubTake-api: %s", err.Error()))
		return
	}
	writeJSONResponse(w, 200, expectedBody)
	return
}

func userReturnObject(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var expectedBody endpoints.BorrowReturnData
	if err := readJSONBody(r, &expectedBody); err != nil {
		writeResponse(w, http.StatusNotAcceptable,
			fmt.Sprintf("HubTake-api: %s", err.Error()))
		return
	}
	err := ep.UserReturnObject(db, expectedBody)
	if err != nil {
		writeResponse(w, http.StatusNotFound,
			fmt.Sprintf("HubTake-api: %s", err.Error()))
		return
	}
	writeJSONResponse(w, http.StatusOK, expectedBody)
	return
}

//GET user and object user borrow
func userObjectData(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	data, err := ep.GetUserObjectData(db)
	if err != nil {
		if len(data) == 0 {
			writeResponse(w, 204,
				fmt.Sprintf("HubTake-api: %s", err.Error()))
			return
		}
		writeResponse(w, 404,
			fmt.Sprintf("HubTake-api: %s", err.Error()))
		return
	}
	rawBody, _ := json.Marshal(data)
	writeResponse(w, 200, string(rawBody))
}

func deleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	mail := p.ByName("userToDelete")
	err := ep.DeleteUser(db, mail)
	if err != nil {
		writeResponse(w, 404,
			fmt.Sprintf("User not found"))
		return
	}
	writeResponse(w, http.StatusNoContent,
		fmt.Sprintf("HubTake-api: Remove successfully"))
	return
}
