package handler

import (
	"github.com/fadhlika/tanoshi/pkg/scraper"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	GetSources(c *gin.Context)
	GetMangas(c *gin.Context)
	GetMangaDetail(c *gin.Context)
	GetChapters(c *gin.Context)
	GetPages(c *gin.Context)
}

type mangaHandler struct {
	scrapers map[string]scraper.Scraper
}

func NewMangaHandler(scrapers map[string]scraper.Scraper) Handler {
	return &mangaHandler{scrapers}
}

func (h *mangaHandler) GetSources(c *gin.Context) {
	var sources []string
	for key, _ := range h.scrapers {
		sources = append(sources, key)
	}

	c.JSON(200, sources)
}

func (h *mangaHandler) GetMangas(c *gin.Context) {
	var query GetMangaByQueryRequest
	if err := c.ShouldBind(&query); err != nil {

	}

	mangas, err := h.scrapers[query.Source].GetMangaByQuery(scraper.Params{
		Keyword:       query.Keyword,
		SortBy:        query.SortBy,
		SortDirection: query.SortDirection,
		PageNo:        query.PageNo,
	})
	if err != nil {

	}

	res := GetMangaByQueryResponse{
		Mangas: mangas,
	}
	c.JSON(200, res)
}

func (h *mangaHandler) GetMangaDetail(c *gin.Context) {
	var query GetMangaDetailRequest
	if err := c.ShouldBind(&query); err != nil {

	}

	manga, err := h.scrapers[query.Source].GetMangaDetail(query.Path)
	if err != nil {

	}

	if query.Include == "chapter" {
		manga.Chapters, err = h.scrapers[query.Source].GetChapterList(manga.Path)
		if err != nil {

		}
	}

	res := GetMangaDetailResponse{manga}
	c.JSON(200, res)
}

func (h *mangaHandler) GetChapters(c *gin.Context) {
	var query GetMangaDetailRequest
	if err := c.ShouldBind(&query); err != nil {

	}

	chapters, err := h.scrapers[query.Source].GetChapterList(query.Path)
	if err != nil {

	}

	res := GetChapterResponse{chapters}
	c.JSON(200, res)
}

func (h *mangaHandler) GetPages(c *gin.Context) {
	var query GetPageRequest
	if err := c.ShouldBind(&query); err != nil {

	}

	pages, err := h.scrapers[query.Source].GetPageList(query.Path)
	if err != nil {

	}

	res := GetPageResponse{pages}
	c.JSON(200, res)
}
