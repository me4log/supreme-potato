package services

import "ru.goods/feeds/model"

type DataService struct {
}

func (dataService *DataService) getData() ([]model.Article, error) {

	return []model.Article{
		{1, "221233", "1000000232", 234},
		{2, "221236", "1000000231", 6990},
	}, nil
}
