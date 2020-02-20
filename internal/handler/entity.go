package handler

import "github.com/fadhlika/tanoshi/pkg/scraper"

type GetMangaByQueryRequest struct {
	Source        string `form:"source"`
	Keyword       string `form:"keyword"`
	SortBy        string `form:"sort_by"`
	SortDirection string `form:"sort_order"`
	PageNo        string `form:"page"`
}

type GetMangaByQueryResponse struct {
	Mangas []scraper.Manga `json:"manga"`
}

type GetChapterRequest struct {
	Source string `form:"source"`
	Path   string `form:"path"`
}

type GetChapterResponse struct {
	Chapters []scraper.Chapter `json:"chapters"`
}

type GetMangaDetailRequest struct {
	GetChapterRequest
	Include string `form:"include"`
}

type GetMangaDetailResponse struct {
	scraper.Manga
}

type GetPageRequest struct {
	GetChapterRequest
}

type GetPageResponse struct {
	Pages []scraper.Page `json:"pages"`
}
