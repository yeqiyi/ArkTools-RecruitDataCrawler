package modules

import (
	"fmt"
	"net/url"
)

const (
	Home = "首页"
	Char_List="干员一览"
	Item_List="道具一览"
)

func GetUrl(path string)string{
	url:=fmt.Sprintf("http://prts.wiki/index.php?title=%s&mobileaction=toggle_view_mobile",url.QueryEscape(path))
	return url
}

