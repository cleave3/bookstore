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


func responseHandler(status bool, code int, message string, data interface{}) interface{} {
	return map[string]interface{}{
		"status":  status,
		"code":    code,
		"message": message,
		"data":    data,
	}
}

func HandleSucess(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := responseHandler(true, http.StatusOK, "success", data)

	res, _ := json.Marshal(response)

	w.Write(res)
}

func HandleCreated(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	response := responseHandler(true, http.StatusCreated, "success", data)

	res, _ := json.Marshal(response)

	w.Write(res)
}
