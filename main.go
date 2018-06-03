package main

import (
	"ru.goods/feeds/services"
	"os"
)

func main() {

	dataService := services.DataService{}
	feedService := services.FeedService{}

	if err := feedService.CreateFeed(os.Stdout, &dataService); err != nil {
		panic(err)
	}
}
