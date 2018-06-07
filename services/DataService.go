package services

import "goods.ru/feeds/model"

type DataService struct {
}

func (dataService *DataService) getData() ([]model.Article, []model.Category, error) {

	return []model.Article{
		{1, "221233", "1000000232", 234},
		{2, "221236", "1000000231", 6990},
	}, [] model.Category{
		{Id: "01", Name: "Электроника"},
		{Id: "02", Name: "Бытовая техника"},
	}, nil
}
