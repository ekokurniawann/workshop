package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"workshop/helper"
	"workshop/models"

	"github.com/gorilla/mux"
)

var ResponseJSON = helper.ResponseJson
var ResponseError = helper.ResponseError

func Test(http.ResponseWriter, *http.Request) {
	fmt.Println("Masuk")
}

func Create(w http.ResponseWriter, r *http.Request) {
	var personData models.Person

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&personData)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, "error decode")
		return
	}
	err = models.DB.Create(&personData).Error
	if err != nil {
		ResponseError(w, http.StatusBadRequest, "error Create")
	}

	ResponseJSON(w, http.StatusOK, personData)
}

func GetPeople(w http.ResponseWriter, r *http.Request) {
	var personData []models.Person

	err := models.DB.Find(&personData).Error
	if err != nil {
		ResponseError(w, http.StatusBadRequest, "data not found")
		return
	}

	ResponseJSON(w, http.StatusOK, personData)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	var personData models.Person

	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, "id invalid")
		return
	}

	err = models.DB.First(&personData, id).Error
	if err != nil {
		ResponseError(w, http.StatusNotFound, "data not found")
	}

	ResponseJSON(w, http.StatusOK, &personData)
	// fmt.Println("person", &personData)
}

func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	var personData models.Person

	decoder := json.NewDecoder(r.Body)

	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, "id invalid")
		return
	}

	err = decoder.Decode(&personData)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, "error decode")
		return
	}

	err = models.DB.Model(&personData).Where("id = ?", id).Updates(&personData).Error
	if err != nil {
		ResponseError(w, http.StatusBadRequest, "error Update")
		return
	}

	ResponseJSON(w, http.StatusOK, &personData)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, "id invalid")
		return
	}

	defer r.Body.Close()

	var personData models.Person

	personData.ID = uint(id)

	err = models.DB.Delete(&personData).Error
	if err != nil {
		ResponseError(w, http.StatusNotFound, "error Delete")
		return
	}

	response := map[string]string{"message": "success deleted id"}
	ResponseJSON(w, http.StatusOK, response)
}
