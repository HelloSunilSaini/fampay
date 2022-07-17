package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ServiceResponse struct {
	Code     int
	Response interface{}
}

func (s *ServiceResponse) RenderResponse(w http.ResponseWriter) {
	headers := map[string]string{
		"Content-Type":                 "application/json",
		"Access-Control-Allow-Headers": "*",
		"Access-Control-Allow-Origin":  "*",
		"Access-Control-Allow-Methods": "*",
	}
	for h, val := range headers {
		w.Header().Set(h, val)
	}

	data, _ := json.MarshalIndent(s.Response, "", "")
	w.Header().Set("Content-Length", fmt.Sprint(len(data)))
	w.WriteHeader(s.Code)
	fmt.Fprint(w, string(data))
}

func Simple200OK(message string) ServiceResponse {
	return ServiceResponse{http.StatusOK, message}
}
func Simple404Response(message string) ServiceResponse {
	return ServiceResponse{http.StatusNotFound, message}
}

func UnAuthorized(message string) ServiceResponse {
	return ServiceResponse{http.StatusUnauthorized, message}
}

func PreconditionFailed(message string) ServiceResponse {
	return ServiceResponse{http.StatusPreconditionFailed, message}
}

func OptionsResponseOK(message string) ServiceResponse {
	return ServiceResponse{http.StatusOK, message}
}

func SimpleBadRequest(message string) ServiceResponse {
	return ServiceResponse{http.StatusBadRequest, message}
}

func Response200OK(response interface{}) ServiceResponse {
	return ServiceResponse{http.StatusOK, response}
}

func ResponseNotImplemented(response interface{}) ServiceResponse {
	return ServiceResponse{http.StatusNotImplemented, "Method not implementd"}
}

func ReponseInternalError(message string) ServiceResponse {
	return ServiceResponse{http.StatusInternalServerError, message}
}

func ProcessError(err error) ServiceResponse {
	return ReponseInternalError(err.Error())
}
