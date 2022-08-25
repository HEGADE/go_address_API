package views

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hegade/go_address_API/service"
	"gorm.io/gorm"
)

type StateView interface {
	Create(ctx *gin.Context)
	Get(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	List(ctx *gin.Context)
	DeleteAll(ctx *gin.Context)
}
type stateview struct {
	service service.StateService
}

func NewStateView(conn *gorm.DB) StateView {
	return &stateview{
		service: service.NewStateService(conn),
	}
}

func (s *stateview) Create(ctx *gin.Context) {
	req := service.StateRequest{}
	ctx.ShouldBind(&req)
	v, err := s.service.Create(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, v)
	return
}

func (s *stateview) List(ctx *gin.Context) {

	// /states?page=1&size=10
	page := ctx.Query("page")
	size := ctx.Query("size")
	v, err := s.service.List(page, size)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusAccepted, v)
	return
}

func (s *stateview) Get(ctx *gin.Context) {
	ID := ctx.Param("id")
	v, err := s.service.Get(ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusAccepted, v)
	return
}

func (s *stateview) Update(ctx *gin.Context) {
	req := service.StateRequest{}
	ctx.ShouldBind(&req)
	ID := ctx.Param("id")
	v, err := s.service.Update(ID, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusAccepted, v)
	return
}

func (s *stateview) Delete(ctx *gin.Context) {
	ID := ctx.Param("id")
	v, err := s.service.Delete(ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusAccepted, v)
	return
}

func (s *stateview) DeleteAll(ctx *gin.Context) {

	v, err := s.service.DeleteAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{"msg": v})
}
