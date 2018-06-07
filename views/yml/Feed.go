package yml

import (
	"encoding/xml"
	"goods.ru/feeds/model"
	"io"
)

type Feed struct {
	XMLName    xml.Name      `xml:"yml_catalog"`
	Categories []interface{} `xml:"categories>category"`
	Articles   []interface{} `xml:"offers>offer"`

}

func CreateFeed(formatter *Formatter, articles []model.Article, categories []model.Category) (*Feed, error) {

	feed := Feed{Articles: make([]interface{}, len(articles))}

	for _, categories := range categories {
		feed.Categories = append(feed.Categories, (*formatter.CategoryFormatter)(&categories))
	}

	for _, article := range articles {
		feed.Articles = append(feed.Articles, (*formatter.ArticleFormatter)(&article))
	}

	return &feed, nil
}

func (feed *Feed) Encode(writer io.Writer) (error) {
	enc := xml.NewEncoder(writer)
	enc.Indent("  ", "    ")
	if err := enc.Encode(&feed); err != nil {
		return err
	}
	return nil
}
