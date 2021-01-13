package fbrip

import (
	"fmt"
	"net/url"
)

type InfoStruct struct{
	Name string
	Birthday string
	Gender string
}

type ReactStruct struct{
	Urls []*url.URL
	Ids []string
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

//Creates
func CreateScrap(path string, urls []string) ScrapStruct{
	parsedUrls := parseUrls(urls)
	return ScrapStruct{
		Urls:parsedUrls,
		FolderPath:path,
	}
}

func CreateReact(ids []string, urls []string) ReactStruct{
	if (len(ids) != len(urls)){
		panic("IDs length != URLs length. Must have the same length")
	}
	parsedUrls := parseUrls(urls)
	return ReactStruct{
		Urls:parsedUrls,
		Ids:ids,
	}
}



//Setters for Structs
func (i *InfoStruct) setInfo(basicInfo map[string]string){
	i.Name = basicInfo["Name"]
	i.Birthday = basicInfo["Birthday"]
	i.Gender = basicInfo["Gender"]
}



func parseUrls(urls[]string) []*url.URL{
	//Slice of parsed urls
	var Urls []*url.URL
	
	//Parsing urls
	for i,Url := range urls{
		tempUrl,err := url.Parse(Url)
		if err!=nil{
			fmt.Printf("Error while parsing #%d -> %s\nIt was omitted.\n",i,Url)
			continue
		}
		Urls = append(Urls,tempUrl)
	}
	return Urls
}
