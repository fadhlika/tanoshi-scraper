package router

import (
	"github.com/fadhlika/tanoshi/internal/handler"
	"github.com/gin-gonic/gin"
)

type Router interface {
	InitRouter()
	Run()
}

type router struct {
	router       *gin.Engine
	MangaHandler handler.Handler
}

func NewRouter(mangaHandler handler.Handler) Router {
	r := gin.Default()
	return &router{r, mangaHandler}
}

func (r *router) InitRouter() {
	api := r.router.Group("/api")
	{
		api.GET("/source", r.MangaHandler.GetMangas)
		api.GET("/manga", r.MangaHandler.GetMangaDetail)
		api.GET("/manga/chapter", r.MangaHandler.GetChapters)
		api.GET("/manga/page", r.MangaHandler.GetPages)
	}
}

func (r *router) Run() {
	r.router.Run()
}
