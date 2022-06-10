package fbrip

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"path"
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
		fmt.Printf("** Error while making GET request to: %s | %s\n", Url.String(), u.Parameters["email"])
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
			fmt.Printf("** Error while making GET request to: %s | %s", Url.String(), u.Parameters["email"])
			return
		}
		//Searching for Reaction Url (it contains specific Query Parameters)
		tempUrl := searchReactionPickerUrl(response.Body)
		//Making GET request for the reaction selection link
		response = u.GET(tempUrl)
		//Handling error. NEED TO BE IMPROVED
		if response == nil {
			fmt.Printf("** Error while making GET request to: %s | %s\n", Url.String(), u.Parameters["email"])
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
		filename := fmt.Sprintf("%s_%s_%s.html", Url.Host, Url.Path, u.Parameters["email"])
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
			fmt.Printf("** Error while making GET request to: Scrap > Urls[%d]  |  %s\n", i, u.Parameters["email"])
			continue
		}
		//Creating folder if it does not exist
		if _, err := os.Stat(folderPath); os.IsNotExist(err) {
			os.Mkdir(folderPath, 0777)
		}
		//Writing in location
		err := ioutil.WriteFile(fullpath, bodyToBytes(response.Body), 0666)
		if err != nil {
			fmt.Printf("Unable to write file: %v", err)
		}
	}
}
