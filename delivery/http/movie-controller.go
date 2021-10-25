package http

import (
	"github.com/densus/movie_service/helper"
	external_service "github.com/densus/movie_service/service/external-service"
	internal_service "github.com/densus/movie_service/service/internal-service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type MovieController interface {
	Search(ctx *gin.Context, search, page string)
}

type movieController struct {
	externalService external_service.ExternalService
	internalService internal_service.InternalService
}

func NewMovieController(engine *gin.Engine, service external_service.ExternalService, internalServ internal_service.InternalService) {
	handler := &movieController{externalService: service, internalService: internalServ}

	engine.GET("api/search", handler.Search)
	engine.GET("api/search/:id", handler.GetByImdbID)
}

func (m *movieController) Search(ctx *gin.Context) {
	page := ctx.Query("pagination")
	_page, _ := strconv.Atoi(page)
	movies := m.externalService.Search(ctx.Query("searchword"), _page)
	ctx.JSON(http.StatusOK, movies)
}

func (m *movieController) GetByImdbID(ctx *gin.Context) {
	//fmt.Println("test",ctx.Param("id"))
	movie := m.internalService.GetByImdbID(ctx.Param("id"))
	res := helper.SuccessResponse(movie, "", true)
	ctx.JSON(http.StatusOK, res)
}
