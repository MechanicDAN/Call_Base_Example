package server

import (
	"github.com/go-chi/chi"
	"net/http"
)

type Api struct {
	authId     	int
	authSchema 	string
	urlPath 	string
	request *http.Request
}
func (api Api) GetUrlPathParam(param string) string{
	return chi.URLParam(api.request,param)
}

