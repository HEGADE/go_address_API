package views

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hegade/go_address_API/service"
)

type StateView interface {
	Create(ctx *gin.Context)
	Get(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	List(ctx *gin.Context)
}
type stateview struct {
	service service.StateService
}

func NewStateView() StateView {
	return &stateview{
		service: service.NewStateService(),
	}
}
func (s *stateview) Create(ctx *gin.Context) {
	req := service.StateRequest{}
	ctx.ShouldBind(&req)
	v, err := s.service.Create(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Error")
		return
	}

	ctx.JSON(http.StatusCreated, v)
	return
}

func (s *stateview) Get(ctx *gin.Context) {
	ID := ctx.Param("id")
	ctx.JSON(http.StatusAccepted, ID)
	return
}

func (s *stateview) Update(ctx *gin.Context) {
	ID := ctx.Param("id")
	req := stateRequest{}
	ctx.ShouldBind(&req)
	v, err := strconv.ParseInt(ID, 0, 8)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid ID")
		return
	}
	req.ID = v
	ctx.JSON(http.StatusAccepted, req)
	return
}

func (s *stateview) Delete(ctx *gin.Context) {
	ID := ctx.Param("id")
	ctx.JSON(http.StatusAccepted, ID)
	return
}

func (s *stateview) List(ctx *gin.Context) {

	// /states?page=1&size=10
	page := ctx.Query("page")
	size := ctx.Query("size")
	v, err := s.service.List(page, size)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Error")
		return
	}
	ctx.JSON(http.StatusAccepted, v)
	return
}
