package main

import (
	"github.com/fadhlika/tanoshi/internal/handler"
	"github.com/fadhlika/tanoshi/internal/router"
	"github.com/fadhlika/tanoshi/pkg/scraper"
)

func main() {
	sites := make(map[string]scraper.Scraper)
	sites["mangasee"] = scraper.NewMangaseeScraper()
	mangaHandler := handler.NewMangaHandler(sites)

	r := router.NewRouter(mangaHandler)
	r.InitRouter()
	r.Run()
}
