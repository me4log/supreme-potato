package yml

import "ru.goods/feeds/model"

type ArticleFormatter *func(article *model.Article) *interface{}

type Formatter struct {
	ArticleFormatter ArticleFormatter
}

