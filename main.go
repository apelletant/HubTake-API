package main

import (
	"fmt"
	"net/http"
	"errors"
	"io/ioutil"
	"encoding/json"

	log "github.com/cihub/seelog"
	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"./endpoints"
)

var (
	ep *endpoints.Endpoints
	db *gorm.DB
)

//GUID: 0f0ea599d683


//Write reponse function
func writeResponse(w http.ResponseWriter, status int, body string) {
	w.Header().Set("Content-Length", fmt.Sprintf("%v", len(body)))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	fmt.Fprintf(w, `%v`, body)
}

func writeJsonResponse(w http.ResponseWriter, status int, body interface{}) {
	rawBody, _ := json.Marshal(body)
	writeResponse(w, status, string(rawBody))
}


//Read request body
func readJsonBody(r *http.Request, expectedBody interface{}) error {
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

	db, err = gorm.Open("sqlite3", "/home/apelletant/go/HubTakeAPI/HubTake.db")
	if err != nil {
		panic("Can't find Database")
	}

	//ep := ep.NewEndpoints(db)
	router := httprouter.New()

	//GET ROUTE FOR OBJECTS
	router.GET("/v1/objects", getObject)
    	router.GET("/v1/objects/isTaken", getTakenObject)
    	router.GET("/v1/objects/notTaken", getNotTakenObject)
    	router.GET("/v1/objects/getByName/:name", getObjectWithName)

    	//GET ROUTE FOR USER
    	router.GET("/v1/users", getUsers)
    	router.GET("/v1/users/:userEmailGet", getUserByMail)
    	router.GET("/v1/usersHasObject", getUserHasObject)

	//POST ROUTE FOR OBJECT
    	router.POST("/v1/objects/post/:objectName", addObject)
    	router.POST("/v1/objects/borrow", updateObjectBorrowState)
    	router.POST("/v1/objects/return", updateObjectReturnState)

    	//POST ROUTE FOR USERS
    	router.POST("/v1/User", addUser)

	panic(http.ListenAndServe(":8080", router))
	db.Close()
}

//GET OBJECT
func getObject(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    o := ep.GetObjects(db)
    rawBody, _ := json.Marshal(o)
    writeResponse(w, 200, string(rawBody))
    return
}

func getObjectWithName(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    o := ep.GetObjectByName(db, "name")
    writeJsonResponse(w, http.StatusOK, o)
    return
}

func getTakenObject(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    o := ep.GetTakenObject(db)
    rawBody, _ := json.Marshal(o)
    writeResponse(w, 200, string(rawBody))
    return
}

func getNotTakenObject(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    o := ep.GetNotTakenObject(db)
    rawBody, _ := json.Marshal(o)
    writeResponse(w, 200, string(rawBody))
    return
}


//POST OBJECT
func addObject(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	expectedBody := struct { Name string }{}
	if err := readJsonBody(r, &expectedBody); err != nil {
		writeResponse(w, http.StatusInternalServerError,
			fmt.Sprintf("HubTake-api: %s", err.Error()))
		return
	}
	o, err := ep.AddObject(db, expectedBody.Name);
	if err != nil {
		writeResponse(w, http.StatusInternalServerError,
			"HubTake-api: add object error: "+err.Error())
		return
	} else {
		writeJsonResponse(w, http.StatusOK, o)
	}
	return
}

func updateObjectBorrowState(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    	fmt.Println("updateBorrowState")
}

func updateObjectReturnState(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    	fmt.Println("getTakenObject")
}

//GET USER
func getUsers(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    	u := ep.GetUsers(db)
    	rawBody, _ := json.Marshal(u)
    	writeResponse(w, 200, string(rawBody))
    	}

func getUserByMail(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    	u := ep.GetUserByMail(db, p.ByName("userEmailGet"))
    	rawBody, _ := json.Marshal(u)
    	writeResponse(w, 200, string(rawBody))
	fmt.Println("getUserByMail")
}

func getUserHasObject(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    fmt.Println("getUserHasObject")
}

//POST USER
func addUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    /*expectedBody := struct { Name string }{}
    if err := readJsonBody(r, &expectedBody); err != nil {
	writeResponse(w, http.StatusInternalServerError,
	    fmt.Sprintf("HubTake-api: %s", err.Error()))
	return
    }
    o, err := ep.AddObject(db, expectedBody.Name);
    if err != nil {
	writeResponse(w, http.StatusInternalServerError,
	    "HubTake-api: add object error: "+err.Error())
	return
    } else {
	writeJsonResponse(w, http.StatusOK, o)
    }
    return*/
}