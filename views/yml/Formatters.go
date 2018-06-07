package yml

import "goods.ru/feeds/model"

type ArticleFormatter *func(article *model.Article) *interface{}
type CategoryFormatter *func(category *model.Category) *interface{}

type Formatter struct {
	ArticleFormatter  ArticleFormatter
	CategoryFormatter CategoryFormatter
}
