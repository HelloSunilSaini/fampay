package apihandlers

import (
	"fampay/client"
	"fampay/handler"
	"net/http"
	"time"
)

type YoutubeHandler struct {
	BaseHandler
}

func (p *YoutubeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := handler.RouteApiCall(p, r)
	response.RenderResponse(w)
}

func (p *YoutubeHandler) Get(r *http.Request) handler.ServiceResponse {
	resp := client.GetYoutubeVideos(time.Now().Add(-time.Minute * 10))
	return handler.Response200OK(resp)
}
