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
	ep endpoints.Endpoints
	db *gorm.DB
)

//Write response function
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
	router.GET("/v1/user/byID/:userId", getUserByID)
	router.POST("/v1/objects/post/:objectName", addObject)
	router.POST("/v1/users", addUser)

	router.GET("/v1/usersHasObject", userObjectData)
	router.POST("/v1/take", userTakeObject)
	router.POST("/v1/return", userReturnObject)

	router.POST("/v1/plastic", addPlastic)
	router.PUT("/v1/plastic", updatePlastic)
	router.GET("/v1/plastic", getPlastics)
	router.GET("/v1/plastic/:plasticColor", getPlasticByColor)

	router.POST("/v1/command", addCommand)
	router.GET("/v1/command", getCommands)
	router.GET("/v1/command/:customerName", getCommandCustomer)
	router.DELETE("/v1/command/:idCommand", deleteCommand)

	router.DELETE("/v1/objects/:objectToDelete", deleteObject)
	router.DELETE("/v1/users/:userToDelete", deleteUser)

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}
