package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/hegade/go_address_API/utils"
	"github.com/hegade/go_address_API/views"
)

func main() {
	conn, err := utils.GormConnection(context.Background())
	if err != nil {
		panic(err)
	}
	router := gin.Default()
	s := views.NewStateView(conn)
	router.POST("/states", s.Create)
	router.GET("/states", s.List)
	router.GET("/states/:id", s.Get)
	router.PUT("/states/:id", s.Update)
	router.DELETE("/states/:id", s.Delete)
	// router.DELETE("states/all", s.DeleteAll)
	router.Run()
}
