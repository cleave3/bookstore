package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll((r.Body)); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

func HandleSucess(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := make(map[string]interface{})
	response["status"] = true
	response["code"] = http.StatusOK
	response["message"] = "success"
	response["data"] = data

	res, _ := json.Marshal(response)

	w.Write(res)
}

func HandleCreated(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	
	response := make(map[string]interface{})
	response["status"] = true
	response["code"] = http.StatusCreated
	response["message"] = "success"
	response["data"] = data

	res, _ := json.Marshal(response)

	w.Write(res)
}