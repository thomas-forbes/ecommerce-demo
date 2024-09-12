package api

import (
	"net/http"
)

type Route struct {
	Method  string
	Path    string
	Handler func(w http.ResponseWriter, r *http.Request)
}

func (rh *RouteHandler) getRoutes() []Route {
	return []Route{
		{
			Method:  http.MethodGet,
			Path:    "/products",
			Handler: asGetRequest(rh.ProductService.GetProducts),
		},
		{
			Method:  http.MethodPost,
			Path:    "/products",
			Handler: asPostRequest(rh.ProductService.CreateProduct),
		},
	}
}
