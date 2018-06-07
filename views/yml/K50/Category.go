package K50

import (
	"encoding/xml"
	"goods.ru/feeds/model"
)

type __category struct {
	category *model.Category
	XMLName  xml.Name `xml:"category"`
	Id       string   `xml:"id,attr"`
	Name     string   `xml:",chardata"`
	Type     string   `xml:"type,attr"`
}

var FormatCategory = func(category *model.Category) *interface{} {

	var result interface{} = __category{
		category: category,
		Id:       category.Id,
		Name:     category.Name,
		Type:     "web",
	}

	return &result
}
