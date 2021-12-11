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
	Pub="公开招募"
)

type CharList []Character

type RecruitData struct{
	Charlist map[int]Character
	TagtoChar map[string][]int
}

type Character struct {
	Id int
	Cn_name   string //名称（中）
	En_name   string //名称（英）
	Jp_name   string //名称（日）
	AvatarUrl string //头像url
	Sex       string //干员性别
	Msg       string //干员信息
	More_msg  string //附加信息
	Approach int //获取方式
	Tags []string //标签
	Class string//职阶
	Star int //星级
	Ori_hp int //初始生命力
	Ori_atk int //初始攻击力
	Ori_def int //初始物抗
	Ori_res int //初始法抗
	Ori_dt string //再部署时间
	Cost string//部署费用
	Ori_block string//阻挡
	Ori_cd string//攻击间隔
	Birthplace string//干员出生地
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
					char.Id=id
				}else{
					log.Warn("type transform failed:",err)
				}
			case "data-icon":
				char.AvatarUrl="http:"+attr.Val
			case "data-cn":
				char.Cn_name=attr.Val
			case "data-en":
				char.En_name=attr.Val
			case "data-jp":
				char.Jp_name=attr.Val
			case "data-sex":
				char.Sex=attr.Val
			case "data-des":
				char.Msg=attr.Val
			case "data-moredes":
				char.More_msg=attr.Val
			case "data-tag":
				char.Tags=strings.Split(attr.Val," ")
			case "data-class":
				char.Class=attr.Val
			case "data-rarity":
				if star,err:=strconv.Atoi(attr.Val);err==nil{
					char.Star=star+1
				}else{
					log.Warn("type transform failed:",err)
				}
			case "data-ori-hp":
				if hp,err:=strconv.Atoi(TrimComma(attr.Val));err==nil{
					char.Ori_hp=hp
				}else{
					log.Warn("type transform failed:",err)
				}
			case "data-ori-atk":
				if atk,err:=strconv.Atoi(TrimComma(attr.Val));err==nil{
					char.Ori_atk=atk
				}else{
					log.Warn("type transform failed:",err)
				}
			case "data-ori-def":
				if def,err:=strconv.Atoi(TrimComma(attr.Val));err==nil{
					char.Ori_def=def
				}else{
					log.Warn("type transform failed:",err)
				}
			case "data-ori-res":
				if res,err:=strconv.Atoi(TrimComma(attr.Val));err==nil{
					char.Ori_res=res
				}else{
					log.Warn("type transform failed:",err)
				}
			case "data-ori-dt":
				char.Ori_dt=attr.Val
			case "data-ori-dc":
				char.Cost=attr.Val
			case "data-ori-block":
				char.Ori_block=attr.Val
			case "data-ori-cd":
				char.Ori_cd=attr.Val
			case "data-ori-birthplace":
				char.Birthplace=attr.Val
			case "data-approach":
				if true==strings.Contains(attr.Val,Pub){
					char.Approach=1
				}else{
					char.Approach=0
				}
			}
		}
		rd.Charlist[char.Id]=char
		//fmt.Println("------")
	}
}


func (rd *RecruitData)BuildTagMap(){
	//rd.TagtoChar=make(map[string][]int)
	for _,char :=range rd.Charlist{
		for _,tag:=range char.Tags{
			rd.TagtoChar[tag]=append(rd.TagtoChar[tag],char.Id)
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
	rd:=RecruitData{Charlist:make(map[int]Character),TagtoChar: make(map[string][]int)}
	rd.getAllChar(doc)
	rd.BuildTagMap()
	return rd
}