package yml

import (
	"encoding/xml"
	"ru.goods/feeds/model"
	"io"
)

type Feed struct {
	XMLName  xml.Name      `xml:"yml_catalog"`
	Articles []interface{} `xml:"offers"`
}

func CreateFeed(formatter Formatter, articles []model.Article) (*Feed, error) {

	feed := Feed{Articles: make([]interface{}, len(articles))}
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
