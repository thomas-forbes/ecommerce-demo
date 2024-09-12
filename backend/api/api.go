package api

import (
	"backend/products"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type RouteHandler struct {
	Router         *mux.Router
	ProductService products.ProductProvider
}

func NewRequestHandler(productProvider products.ProductProvider) *RouteHandler {
	return &RouteHandler{
		ProductService: productProvider,
	}
}

func (rh *RouteHandler) Run() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), rh.Router))
}

func (rh *RouteHandler) InitializeRoutes() {
	rh.Router = mux.NewRouter()
	routes := rh.getRoutes()
	for _, route := range routes {
		rh.Router.HandleFunc(route.Path, route.Handler).Methods(route.Method)
	}
}
