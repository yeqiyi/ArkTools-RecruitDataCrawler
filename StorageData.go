package main

import (
	. "ArkToolBackEnd/tools/GetData"
	"ArkToolBackEnd/tools/modules"
	"encoding/json"
	"github.com/labstack/gommon/log"
	"io/ioutil"
)
var url=modules.GetUrl(modules.Char_List)

func TagtoJson(tags Tags)error{
	data,err:=json.Marshal(tags)
	if err!=nil{
		return err
	}else{
		err:=ioutil.WriteFile("Tags.json",data,0644)
		if err!=nil{
			return err
		}
	}
	return nil
}
func RDtoJson(rd RecruitData)error{
	c_data,err:=json.Marshal(rd.Charlist)
	if err!=nil{
		return err
	}else{
		err:=ioutil.WriteFile("CharacterData.json",c_data,0644)
		if err!=nil{
			return err
		}
	}
	t_data,err:=json.Marshal(rd.TagtoChar)
	if err!=nil{
		return err
	}else{
		err:=ioutil.WriteFile("TagtoChar.json",t_data,0644)
		if err!=nil{
			return err
		}
	}
	return nil
}

func main() {
	doc, err := GetDoc(url)
	if err != nil {
		log.Error(err)
	}
	/*
	var tags Tags = GetAllTags(doc)
	err = TagtoJson(tags)
	if err != nil {
		log.Error(err)
	}*/
	rd:=GetRecruitData(doc)
	err = RDtoJson(rd)
	if err!=nil{
		log.Error(err)
	}
	return
}

