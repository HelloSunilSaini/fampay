package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

type ServiceApiHandler interface {
	GetOne(r *http.Request, id string) ServiceResponse
	Get(r *http.Request) ServiceResponse
	Put(r *http.Request) ServiceResponse
	Post(r *http.Request) ServiceResponse
	Delete(r *http.Request) ServiceResponse
	Patch(r *http.Request) ServiceResponse
	Options(r *http.Request) ServiceResponse
}

func RouteApiCall(sah ServiceApiHandler, r *http.Request) ServiceResponse {
	switch r.Method {
	case "GET":
		params := mux.Vars(r)
		id, present := params["id"]
		if present {
			return sah.GetOne(r, id)
		} else {
			return sah.Get(r)
		}
	case "PUT":
		return sah.Put(r)
	case "POST":
		return sah.Post(r)
	case "PATCH":
		return sah.Patch(r)
	case "DELETE":
		return sah.Delete(r)
	case "OPTIONS":
		return sah.Options(r)
	}
	return ServiceResponse{
		Code:     http.StatusMethodNotAllowed,
		Response: "Method not allowed",
	}

}
