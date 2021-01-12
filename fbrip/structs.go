package fbrip

import "net/url"

type InfoStruct struct{
	Name string
	Birthday string
	Gender string
}

type ReactStruct struct{
	Url *url.URL
	Id string
}

type PostStruct struct{
	Url *url.URL
	Content string
}

type CommentStruct struct{
	Url *url.URL
	Content string
}

type ScrapStruct struct{
	Urls []*url.URL
	FolderPath string
}

//Setters for Structs
func (i *InfoStruct) setInfo(basicInfo map[string]string){
	i.Name = basicInfo["Name"]
	i.Birthday = basicInfo["Birthday"]
	i.Gender = basicInfo["Gender"]
}
