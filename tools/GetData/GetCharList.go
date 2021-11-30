package GetData

import (
	. "ArkToolBackEnd/tools/modules"
	"github.com/PuerkitoBio/goquery"
	"github.com/labstack/gommon/log"
	"strconv"
	"strings"
)


var url=GetUrl(Char_List)

const (
	Pos1 = "近战位"
	Pos2 = "远程位"
)

type CharList []Character

type RecruitData struct{
	charlist CharList
	TagtoChar map[string][]int
}

type Character struct {
	id int
	cn_name   string //名称（中）
	en_name   string //名称（英）
	jp_name   string //名称（日）
	avatarUrl string //头像url
	sex       string //干员性别
	msg       string //干员信息
	more_msg  string //附加信息
	approach string //获取方式
	tags []string //标签
	class string//职阶
	star int //星级
	ori_hp int //初始生命力
	ori_atk int //初始攻击力
	ori_def int //初始物抗
	ori_res int //初始法抗
	ori_dt string //再部署时间
	cost string//部署费用
	ori_block string//阻挡
	ori_cd string//攻击间隔
	birthplace string//干员出生地
}
//获取所有干员信息
func (rd *RecruitData)getAllChar(dom *goquery.Document){
	//defer resp.Body.Close()
	Selects:=dom.Find(".smwdata")
	for i:=0;i<Selects.Length();i++{
		char:=Character{}
		for _,attr:=range Selects.Get(i).Attr{
			//fmt.Println(attr.Key,attr.Val)
			switch attr.Key {
			case "data-sort_id":
				if id,err:=strconv.Atoi(TrimComma(attr.Val));err==nil{
					char.id=id
				}else{
					log.Warn("type transform failed:",err)
				}
			case "data-icon":
				char.avatarUrl=attr.Val
			case "data-cn":
				char.cn_name=attr.Val
			case "data-en":
				char.en_name=attr.Val
			case "data-jp":
				char.jp_name=attr.Val
			case "data-sex":
				char.sex=attr.Val
			case "data-des":
				char.msg=attr.Val
			case "data-moredes":
				char.more_msg=attr.Val
			case "data-tag":
				char.tags=strings.Split(attr.Val," ")
			case "data-class":
				char.class=attr.Val
			case "data-rarity":
				if star,err:=strconv.Atoi(attr.Val);err==nil{
					char.star=star+1
				}else{
					log.Warn("type transform failed:",err)
				}
			case "data-ori-hp":
				if hp,err:=strconv.Atoi(TrimComma(attr.Val));err==nil{
					char.ori_hp=hp
				}else{
					log.Warn("type transform failed:",err)
				}
			case "data-ori-atk":
				if atk,err:=strconv.Atoi(TrimComma(attr.Val));err==nil{
					char.ori_atk=atk
				}else{
					log.Warn("type transform failed:",err)
				}
			case "data-ori-def":
				if def,err:=strconv.Atoi(TrimComma(attr.Val));err==nil{
					char.ori_def=def
				}else{
					log.Warn("type transform failed:",err)
				}
			case "data-ori-res":
				if res,err:=strconv.Atoi(TrimComma(attr.Val));err==nil{
					char.ori_res=res
				}else{
					log.Warn("type transform failed:",err)
				}
			case "data-ori-dt":
				char.ori_dt=attr.Val
			case "data-ori-dc":
				char.cost=attr.Val
			case "data-ori-block":
				char.ori_block=attr.Val
			case "data-ori-cd":
				char.ori_cd=attr.Val
			case "data-ori-birthplace":
				char.birthplace=attr.Val
			}
		}
		rd.charlist=append(rd.charlist,char)
		//fmt.Println("------")
	}
}


func (rd *RecruitData)BuildTagMap(){
	//rd.TagtoChar=make(map[string][]int)
	for _,char :=range rd.charlist{
		for _,tag:=range char.tags{
			rd.TagtoChar[tag]=append(rd.TagtoChar[tag],char.id)

		}
	}
}

func TrimComma(str string) string{
	substrs:=strings.Split(str,",")
	var output string
	for _,substr:=range substrs{
		output+=substr
	}
	return output
}

func GetRecruitData(doc *goquery.Document)RecruitData{
	rd:=RecruitData{TagtoChar: make(map[string][]int)}
	rd.getAllChar(doc)
	rd.BuildTagMap()
	return rd
}