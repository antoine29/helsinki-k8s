package server

import (
	"fmt"
	"net/http"
)

func BuildErrorResponse(err error) []byte {
	return []byte(fmt.Sprintf(`{ "error": "%s" }`, err.Error()))
}

func WriteJsonResponse(res http.ResponseWriter, httpCode int, resMessage []byte) {
	res.Header().Set("Content-type", "application/json")
	res.WriteHeader(httpCode)
	if resMessage != nil {
		res.Write(resMessage)
	}
}
