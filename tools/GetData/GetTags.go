package GetData

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

type Tags struct {
	Tags map[string][]string `json:"tags"`
}
//获取标签信息
func (t *Tags)getTags(doc *goquery.Document){
	doc.Find(".filter-checkbox").Each(func(i int, selection *goquery.Selection) {
		tagType:=selection.Find("td>span").Text()
		fmt.Println(">>",tagType)
		selection.Find(".checkBoxWrapper").Find("label>span").Each(func(i int, selection *goquery.Selection) {
			t.Tags[tagType]=append(t.Tags[tagType],selection.Text())
			fmt.Println(selection.Text())
		})
		fmt.Println("---")
	})
}

func GetAllTags(doc *goquery.Document)Tags{
	tags:=Tags{Tags: make(map[string][]string)}
	tags.getTags(doc)
	return tags
}
