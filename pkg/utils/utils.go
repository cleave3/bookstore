package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type response struct {
	status  bool
	code    int
	message string
	data    interface{}
}

func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll((r.Body)); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

func responseHandler(w http.ResponseWriter, r response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.code)

	response := map[string]interface{}{
		"status":  r.status,
		"code":    r.code,
		"message": r.message,
		"data":    r.data,
	}

	res, _ := json.Marshal(response)

	w.Write(res)
}

func HandleSucess(w http.ResponseWriter, code int, data interface{}) {
	responseHandler(w, response{status: true, code: code, message: "success", data: data})
}

func HandleBadRequest(w http.ResponseWriter, code int, message string) {
	responseHandler(w, response{status: false, code: code, message: message, data: nil})
}
