package routers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"net/http"

	modelObj "../models"
	DAO "../models/dao"
)

var dao = DAO.RDB_DAO{}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	responses, err := dao.GetAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, responses)
}

func GetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	response, err := dao.GetByID(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Movie ID")
		return
	}
	respondWithJson(w, http.StatusOK, response)
}

func Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var response modelObj.ResponseDataBody_OBJ
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	response.ID = bson.NewObjectId()
	if err := dao.Create(response); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, response)
}

func Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	var response modelObj.ResponseDataBody_OBJ
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Update(params["id"], response); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": response.Request_Url + " atualizado com sucesso!"})
}

func Delete(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	if err := dao.Delete(params["id"]); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}
