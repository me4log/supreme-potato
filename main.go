package main

import (
	"goods.ru/feeds/services"
	"goods.ru/feeds/feeds"
	"sync"
)

func main() {

	dataService := &services.DataService{}
	feedService := &services.FeedService{}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go feeds.K50Feed(feedService, dataService, wg)
	wg.Wait()
}

