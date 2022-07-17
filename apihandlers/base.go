package apihandlers

import (
	"fampay/handler"
	"net/http"
)

type BaseHandler struct {
}

func (p *BaseHandler) GetOne(r *http.Request, id string) handler.ServiceResponse {
	return handler.ResponseNotImplemented(nil)
}

func (p *BaseHandler) Get(r *http.Request) handler.ServiceResponse {
	return handler.ResponseNotImplemented(nil)
}

func (p *BaseHandler) Put(r *http.Request) handler.ServiceResponse {
	return handler.ResponseNotImplemented(nil)
}

func (p *BaseHandler) Post(r *http.Request) handler.ServiceResponse {
	return handler.ResponseNotImplemented(nil)
}

func (p *BaseHandler) Delete(r *http.Request) handler.ServiceResponse {
	return handler.ResponseNotImplemented(nil)
}

func (p *BaseHandler) Patch(r *http.Request) handler.ServiceResponse {
	return handler.ResponseNotImplemented(nil)
}

func (p *BaseHandler) Options(r *http.Request) handler.ServiceResponse {
	return handler.OptionsResponseOK("OK")
}
