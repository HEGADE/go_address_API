package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hegade/go_address_API/views"
)

func main() {
	router := gin.Default()
	s := views.NewStateView()
	router.POST("/states", s.Create)
	router.GET("/states", s.List)
	router.GET("/states/:id", s.Get)
	router.PUT("/states/:id", s.Update)
	router.DELETE("/states/:id", s.Delete)
	router.Run()
}
