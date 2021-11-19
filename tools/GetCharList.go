package tools

import (
	. "ArkToolBackEnd/tools/modules"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/labstack/gommon/log"
	"net/http"
)

var url=GetUrl(Char_List)

const (
	Pos1 = "近战位"
	Pos2 = "远程位"
)

type Character struct {
	cn_name string
	en_name string
	jp_name string
	avatarUrl string
	sex string
	tag map[string]interface{}//用map实现集合
	position int //部署位 
}

func getDoc(url string) (*goquery.Document,error){
	client:=new(http.Client)
	req,err:=http.NewRequest("GET",url,nil)
	req.Header.Add("user-agent","Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.1.2 Safari/605.1.15")
	if err!=nil{
		log.Fatal(err)
		return nil, err
	}
	resp,err:=client.Do(req)
	defer resp.Body.Close()
	if err!=nil{
		log.Warn("Get data failed...",err)
		return nil, err
	}
	dom,err:=goquery.NewDocumentFromReader(resp.Body)
	if err!=nil{
		log.Warn("Data missing...",err)
	}
	return dom,nil
}

func getAllChar(dom *goquery.Document){
	//defer resp.Body.Close()
	Selects:=dom.Find(".smwdata")
	for i:=0;i<Selects.Length();i++{
		for _,attr:=range Selects.Get(i).Attr{
			fmt.Println(attr.Key,attr.Val)
		}
		fmt.Println("------")
	}
}
//获取标签
func getTags(dom *goquery.Document){
	dom.Find(".filter-checkbox").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(">>",selection.Find("td>span").Text())
		selection.Find(".checkBoxWrapper").Find("label>span").Each(func(i int, selection *goquery.Selection) {
			fmt.Println(selection.Text())
		})
		fmt.Println("---")
	})
}

func GetData(){

}

