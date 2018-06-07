package feeds

import (
	"goods.ru/feeds/views/yml"
	"goods.ru/feeds/views/yml/K50"
	"os"
	"goods.ru/feeds/services"
	"sync"
)

func K50Feed(feedService *services.FeedService, dataService *services.DataService, wg *sync.WaitGroup) {

	defer wg.Done()

	K50formatter := yml.Formatter{
		ArticleFormatter:  &K50.FormatArticle,
		CategoryFormatter: &K50.FormatCategory,
	}

	if err := feedService.CreateFeed(os.Stdout, dataService, &K50formatter); err != nil {
		panic(err)
	}
}
