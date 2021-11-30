package GetData

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

type Tags struct {
	tags map[string][]string
}
//获取标签信息
func (t *Tags)getTags(doc *goquery.Document){
	doc.Find(".filter-checkbox").Each(func(i int, selection *goquery.Selection) {
		tagType:=selection.Find("td>span").Text()
		fmt.Println(">>",tagType)
		selection.Find(".checkBoxWrapper").Find("label>span").Each(func(i int, selection *goquery.Selection) {
			t.tags[tagType]=append(t.tags[tagType],selection.Text())
			fmt.Println(selection.Text())
		})
		fmt.Println("---")
	})
}

func GetAllTags(doc *goquery.Document)Tags{
	tags:=Tags{tags: make(map[string][]string)}
	tags.getTags(doc)
	return tags
}
