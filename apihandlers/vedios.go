package apihandlers

import (
	"fampay/dbrepository"
	"fampay/handler"
	"net/http"
	"strconv"
)

type VediosHandler struct {
	BaseHandler
	DBRepo dbrepository.IRepo
}

func (v *VediosHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := handler.RouteApiCall(v, r)
	response.RenderResponse(w)
}

func (v *VediosHandler) Get(r *http.Request) handler.ServiceResponse {
	values := r.URL.Query()
	searchTerm := values.Get("searchterm")
	if searchTerm == "" {
		return handler.SimpleBadRequest("search Term Not Provided.")
	}
	offset_str := values.Get("offset")
	offset, _ := strconv.Atoi(offset_str)
	size_str := values.Get("pagesize")
	size, err := strconv.Atoi(size_str)
	if err != nil {
		return handler.SimpleBadRequest("Error parsing pagesize : " + err.Error())
	}
	resp, err := v.DBRepo.GetVedioDetailsBySearchTerm(searchTerm, offset, size)
	if err != nil {
		return handler.SimpleBadRequest(err.Error())
	}
	return handler.Response200OK(resp)
}
