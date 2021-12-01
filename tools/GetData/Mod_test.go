package GetData

import (
	"fmt"
	"testing"
)

func Test(t *testing.T){
	doc,err:= GetDoc(url)
	if err!=nil{
		t.Error("failed")
	}
	//getTags(dom)
	/*
	rd:=GetRecruitData(doc)
	fmt.Println(rd.TagtoChar)
	 */

	T:=GetAllTags(doc)
	for k,v:=range T.Tags{
		fmt.Println(k,v)
	}

	/*if modules.GetUrl(modules.Char_List)=="http://prts.wiki/index.php?title=%E5%B9%B2%E5%91%98%E4%B8%80%E8%A7%88&mobileaction=toggle_view_mobile"{
		t.Log("ok")
	}else{
		t.Error("failed")
	}
	*/
}
