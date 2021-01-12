package fbrip

import "net/url"

type info struct{
	Name string
	Birthday string
	Gender string
}

type react struct{
	Url *url.URL
	Id string
}

type post struct{
	Url *url.URL
	Content string
}

type comment struct{
	Url *url.URL
	Content string
}

type scrap struct{
	Urls []*url.URL
}

//Setters for Structs
func (i *info) setInfo(basicInfo map[string]string){
	i.Name = basicInfo["Name"]
	i.Birthday = basicInfo["Birthday"]
	i.Gender = basicInfo["Gender"]
}
