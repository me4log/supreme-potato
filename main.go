package main

import (
	"goods.ru/feeds/services"
	"os"
	"goods.ru/feeds/views/yml"
	"goods.ru/feeds/views/yml/K50"
)

func main() {

	dataService := services.DataService{}
	feedService := services.FeedService{}

	K50formatter := yml.Formatter{
		ArticleFormatter:  &K50.FormatArticle,
		CategoryFormatter: &K50.FormatCategory,
	}

	if err := feedService.CreateFeed(os.Stdout, &dataService, &K50formatter); err != nil {
		panic(err)
	}
}
