package services

import (
	"ru.goods/feeds/views/yml"
	"ru.goods/feeds/views/yml/K50"
	"io"
)

type FeedService struct {
}

func (FeedService *FeedService) CreateFeed(writer io.Writer, dataService *DataService) (error) {

	articles, err := dataService.getData()
	if err != nil {
		return err
	}

	K50formatter := yml.Formatter{&K50.Format}
	feed, err := yml.CreateFeed(K50formatter, articles)
	if err != nil {
		return err
	}

	if err := feed.Encode(writer); err != nil {
		return err
	}

	return nil
}
