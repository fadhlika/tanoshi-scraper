package scraper

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
)

type mangasee struct {
	BaseURL string
}

func NewMangaseeScraper() Scraper {
	return &mangasee{"https://mangaseeonline.us"}
}

func (m *mangasee) GetMangaByQuery(params Params) ([]Manga, error) {
	mangas := []Manga{}

	c := colly.NewCollector()
	c.OnHTML(`.requested .row`, func(e *colly.HTMLElement) {
		mangas = append(mangas, Manga{
			Title:        e.ChildText(".resultLink"),
			Author:       "",
			Genre:        nil,
			Status:       "",
			Description:  "",
			Path:         e.ChildAttr(".resultLink", "href"),
			ThumbnailURL: e.ChildAttr("img", "src"),
			Chapters:     nil,
		})
	})

	v := make(map[string][]byte)
	v["keyword"] = []byte(params.Keyword)
	v["page"] = []byte(params.PageNo)
	v["sortOrder"] = []byte(params.SortDirection)
	v["sortBy"] = []byte(params.SortBy)
	c.PostMultipart(fmt.Sprintf("%s/search/request.php", m.BaseURL), v)

	return mangas, nil
}

func (m *mangasee) GetMangaDetail(link string) (Manga, error) {
	c := colly.NewCollector()
	manga := Manga{Path: link}
	c.OnHTML(`.mainWell`, func(e *colly.HTMLElement) {
		manga.Title = e.ChildText(`h1[class="SeriesName"]`)
		manga.Author = e.ChildText(`a[href*=\"author\"]`)
		e.ForEach(`a[href*=\"genre\"]`, func(i int, e *colly.HTMLElement) {
			manga.Genre = append(manga.Genre, e.Text)
		})
		manga.Status = e.ChildAttr(".PublishStatus", "status")
		manga.Description = e.ChildText(`.description`)
		manga.ThumbnailURL = e.ChildAttr(`.leftImage img`, "src")
	})
	c.Visit(fmt.Sprintf("%s%s", m.BaseURL, link))
	return manga, nil
}

func (m *mangasee) GetChapterList(url string) ([]Chapter, error) {
	c := colly.NewCollector()
	chapters := []Chapter{}
	c.OnHTML(`.mainWell .chapter-list a[chapter]`, func(e *colly.HTMLElement) {
		link := e.Attr("href")
		link = strings.Replace(link, "-page-1", "", 1)
		chapters = append(chapters, Chapter{
			Rank:  e.Attr("chapter"),
			Path:  link,
			Pages: nil,
		})

	})
	c.Visit(fmt.Sprintf("%s%s", m.BaseURL, url))
	return chapters, nil
}

func (m *mangasee) GetPageList(url string) ([]Page, error) {
	c := colly.NewCollector()
	pages := []Page{}
	c.OnHTML(`.image-container-manga`, func(e *colly.HTMLElement) {
		e.ForEach(`.fullchapimage img`, func(i int, element *colly.HTMLElement) {
			pages = append(pages, Page{
				Rank: i,
				Src:  element.Attr("src"),
			})
		})
	})
	c.Visit(fmt.Sprintf("%s%s", m.BaseURL, url))
	return pages, nil
}
