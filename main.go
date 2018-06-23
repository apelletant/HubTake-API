package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"./endpoints"
	log "github.com/cihub/seelog"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/julienschmidt/httprouter"
)

var (
	ep *endpoints.Endpoints
	db *gorm.DB
)

//Write reponse function
func writeResponse(w http.ResponseWriter, status int, body string) {
	w.Header().Set("Content-Length", fmt.Sprintf("%v", len(body)))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	fmt.Fprintf(w, `%v`, body)
}

func writeJSONResponse(w http.ResponseWriter, status int, body interface{}) {
	rawBody, err := json.Marshal(body)
	if err != nil {
		writeResponse(w, http.StatusNotAcceptable, fmt.Sprintf("HubTake api error : %v", err))
	}
	writeResponse(w, status, string(rawBody))
}

//Read request body
func readJSONBody(r *http.Request, expectedBody interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return errors.New("unable to read body")
	}
	if err := json.Unmarshal(body, expectedBody); err != nil {
		log.Errorf(
			"[readJsonBody] invalid request, %s, input was %s",
			err.Error(), string(body))
		return errors.New("invalid json body format")
	}
	return nil
}

func main() {
	var err error

	db, err = gorm.Open("sqlite3", "./HubTake.db")
	if err != nil {
		panic("Can't find Database")
	}
	defer db.Close()
	router := httprouter.New()

	router.GET("/v1/objects", getObject)
	router.GET("/v1/objects/isTaken", getTakenObject)
	router.GET("/v1/objects/notTaken", getNotTakenObject)
	router.GET("/v1/objects/getByName/:name", getObjectWithName)
	router.GET("/v1/users", getUsers)
	router.GET("/v1/users/:userEmailGet", getUserByMail)
	router.POST("/v1/objects/post/:objectName", addObject)
	router.POST("/v1/users", addUser)

	router.GET("/v1/usersHasObject", userObjectData)
	router.POST("/v1/take", userTakeObject)
	router.POST("/v1/return", userReturnObject)

	router.POST("/v1/plastic", addPlastic)
	router.PUT("/v1/plastic", updatePlastic)
	router.GET("/v1/plastic", getPlastics)
	router.GET("/v1/plastic/:plasticColor", getPlasticByColor)

	router.POST("/v1/addCommand", addCommand)
	router.GET("/v1/command", getCommands)
	router.GET("/v1/command/:customerName", getCommandCustomer)
	router.PUT("/v1/command/:idCommand", updateCommand)

	router.DELETE("/v1/objects/:objectToDelete", deleteObject)
	router.DELETE("/v1/users/:userToDelete", deleteUser)

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}

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

func addCommand(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println("function: addCommand")
}

func getCommands(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println("function: getCommands")
}

func getCommandCustomer(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println("function: getCommandCustomer")
}

func updateCommand(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println("function: updateCommand")
}
