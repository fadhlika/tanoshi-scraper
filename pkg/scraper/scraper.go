package scraper

type Scraper interface {
	GetMangaByQuery(param Params) ([]Manga, error)
	GetMangaDetail(link string) (Manga, error)
	GetChapterList(path string) ([]Chapter, error)
	GetPageList(path string) ([]Page, error)
}

type Params struct {
	Keyword       string
	SortBy        string
	SortDirection string
	PageNo        string
}

type Manga struct {
	Title        string    `json:"title"`
	Author       string    `json:"author"'`
	Genre        []string  `json:"author"`
	Status       string    `json:"status"`
	Description  string    `json:"description"`
	Path         string    `json:"path"`
	ThumbnailURL string    `json:"thumbnail_url"`
	Chapters     []Chapter `json:"chapters"`
}

type Chapter struct {
	Rank  string `json:"rank"`
	Path  string `json:"link"`
	Pages []Page `json:"pages"`
}

type Page struct {
	Rank int    `json:"rank"`
	Src  string `json:"link"`
}
