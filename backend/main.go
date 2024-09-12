package main

import (
	"backend/api"
	"backend/database"
	"backend/products"

	"gorm.io/gorm"
)

type Server struct {
	DB           *gorm.DB
	RouteHandler *api.RouteHandler

	ProductProvider products.ProductProvider
}

var s Server

func main() {
	s.DB = database.SetupDatabase()

	s.ProductProvider = products.NewProductManager(s.DB)

	s.initRoutes()
	s.RouteHandler.Run()
}

func (s *Server) initRoutes() {
	rh := api.NewRequestHandler(s.ProductProvider)
	s.RouteHandler = rh
	s.RouteHandler.InitializeRoutes()
}
