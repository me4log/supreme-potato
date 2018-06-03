package K50

import (
	"encoding/xml"
	"ru.goods/feeds/model"
	"strconv"
)

type __article struct {
	article     *model.Article
	XMLName     xml.Name `xml:"offer"`
	Id          int64    `xml:"id"`
	OfferId     string   `xml:"offer_id,attr"`
	GoodsId     string   `xml:"goods_id"`
	Price       int32    `xml:"prices>price"`
	Description string   `xml:"description"`
}

var Format = func(article *model.Article) *interface{} {

	var result interface{} = __article{
		article:     article,
		Id:          article.Id,
		OfferId:     article.OfferId,
		GoodsId:     article.GoodsId,
		Price:       article.Price,
		Description: "Best prices from " + strconv.FormatInt(int64(article.Price), 10),
	}

	return &result
}
