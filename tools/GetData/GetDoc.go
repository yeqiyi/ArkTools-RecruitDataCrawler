package GetData

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/labstack/gommon/log"
	"net/http"
)

func GetDoc(url string) (*goquery.Document,error){
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
	doc,err:=goquery.NewDocumentFromReader(resp.Body)
	if err!=nil{
		log.Warn("Data missing...",err)
	}
	return doc,nil
}
