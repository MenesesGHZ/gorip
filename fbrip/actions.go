package fbrip

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"path"
	"bytes"
)

type ActionConfig struct {
	GetBasicInfo bool
	React        ReactStruct
	Post         PostStruct
	Comment      CommentStruct
	Scrap        ScrapStruct
}

// ACTIONS
func (u *UserRip) GetBasicInfo() {
	// Making GET request
	Url, _ := url.Parse("https://mbasic.facebook.com/profile.php?v=info")
	response := u.GET(Url)
	//Handling error. NEED TO BE IMPROVED
	if response == nil {
		fmt.Printf("** Error while making GET request to: %s | %s\n", Url.String(), u.Email)
		return
	}
	// Searching for user basic info -> {Name:,Birthday:,Gender:}
	bi := searchBasicInfo(response.Body)
	//Setting basic info for user
	u.Info.setInfo(bi)
}

func (u *UserRip) MakeReactions(Urls []*url.URL, reactions []string) {
	for i, Url := range Urls {
		//Fixing Url & Making GET request in the publication link
		Url = fixUrl(Url)
		response := u.GET(Url)
		//Handling error. NEED TO BE IMPROVED
		if response == nil {
			fmt.Printf("** Error while making GET request to: %s | %s", Url.String(), u.Email)
			return
		}
		//Searching for Reaction Url (it contains specific Query Parameters)
		tempUrl := searchReactionPickerUrl(response.Body)
		//Making GET request for the reaction selection link
		response = u.GET(tempUrl)
		//Handling error. NEED TO BE IMPROVED
		if response == nil {
			fmt.Printf("** Error while making GET request to: %s | %s\n", Url.String(), u.Email)
			return
		}
		//Searching for `ufi/reaction` (it contains specific Query Parameters)
		tempUrl = searchUfiReactionUrl(response.Body, reactions[i])
		//Doing reaction
		u.GET(tempUrl)
	}
}

//scrap Urls
func (u *UserRip) Scrap(Urls []*url.URL, folderPath string) {
	folderPath = path.Clean(folderPath)
	for i, Url := range Urls {
		//Making filename string
		filename := fmt.Sprintf("%s_%s_%s.html", Url.Host, Url.Path, u.Email)
		//fixing filename. Converting "/" -> "-"
		var rs []rune
		for _, r := range filename {
			if string(r) == "/" {
				rs = append(rs, '-')
				continue
			}
			rs = append(rs, r)
		}
		//Making full path to file
		fullpath := path.Join(folderPath, string(rs))
		//Getting response from Url
		response := u.GET(Url)
		if response == nil {
			fmt.Printf("** Error while making GET request to: Scrap > Urls[%d]  |  %s\n", i, u.Email)
			continue
		}
		//Creating folder if it does not exist
		if _, err := os.Stat(folderPath); os.IsNotExist(err) {
			os.Mkdir(folderPath, 0777)
		}

		//Writing in location
		buf := new(bytes.Buffer)
		buf.ReadFrom(response.Body)
		err := ioutil.WriteFile(fullpath, buf.Bytes(), 0666)
		if err != nil {
			fmt.Printf("Unable to write file: %v", err)
		}
	}
}


type InfoStruct struct {
	Name     string
	Birthday string
	Gender   string
}

type ReactStruct struct {
	Urls []*url.URL
	Ids  []string
}

type PostStruct struct {
	Url     *url.URL
	Content string
}

type CommentStruct struct {
	Url     *url.URL
	Content string
}

type ScrapStruct struct {
	Urls       []*url.URL
	FolderPath string
}

//Creates
func CreateScrap(path string, urls []string) *ScrapStruct {
	parsedUrls := parseUrls(urls)
	return &ScrapStruct{
		Urls:       parsedUrls,
		FolderPath: path,
	}
}

func CreateReact(ids []string, urls []string) ReactStruct {
	if len(ids) != len(urls) {
		panic("IDs length != URLs length. Must have the same length")
	}
	parsedUrls := parseUrls(urls)
	return ReactStruct{
		Urls: parsedUrls,
		Ids:  ids,
	}
}

//Setters for Structs
func (i *InfoStruct) setInfo(basicInfo map[string]string) {
	i.Name = basicInfo["Name"]
	i.Birthday = basicInfo["Birthday"]
	i.Gender = basicInfo["Gender"]
}

//Checks
func (r *ReactStruct) Checks() bool {
	boolOut := (len(r.Urls) > 0 && len(r.Ids) > 0)
	boolOut = (len(r.Urls) == len(r.Urls)) && boolOut

	for _, Url := range r.Urls {
		boolOut = (Url.String() != "" && boolOut)
	}

	for id := range r.Ids {
		boolOut = (fmt.Sprint(id) != "" && boolOut)
	}

	return boolOut
}

func parseUrls(urls []string) []*url.URL {
	//Slice of parsed urls
	var Urls []*url.URL

	//Parsing urls
	for i, Url := range urls {
		tempUrl, err := url.Parse(Url)
		if err != nil {
			fmt.Printf("Error while parsing #%d -> %s\nIt was omitted.\n", i, Url)
			continue
		}
		Urls = append(Urls, tempUrl)
	}
	return Urls
}
