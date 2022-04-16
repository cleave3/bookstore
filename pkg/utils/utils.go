package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Response struct {
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

func responseHandler(w http.ResponseWriter, r Response) {
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

func HandleSucess(w http.ResponseWriter, data interface{}) {
	responseHandler(w, Response{status: true, code: http.StatusOK, message: "success", data: data})
}

func HandleCreated(w http.ResponseWriter, data interface{}) {
	responseHandler(w, Response{status: true, code: http.StatusCreated, message: "success", data: data})
}
