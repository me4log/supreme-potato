package services

import (
	"goods.ru/feeds/views/yml"
	"io"
)

type FeedService struct {
}

func (FeedService *FeedService) CreateFeed(writer io.Writer, dataService *DataService, formatter *yml.Formatter) (error) {

	articles, categories, err := dataService.getData()
	if err != nil {
		return err
	}
	feed, err := yml.CreateFeed(formatter, articles, categories)
	if err != nil {
		return err
	}

	if err := feed.Encode(writer); err != nil {
		return err
	}

	return nil
}
