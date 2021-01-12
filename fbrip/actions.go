package fbrip

import (
	"os"
	"fmt"
	"path"
	"io/ioutil"
	"net/url"
)

type ActionConfig struct{
	GetBasicInfo bool
	React ReactStruct
	Post PostStruct
	Comment CommentStruct
	Scrap ScrapStruct
}

// ACTIONS
func(u *UserRip) GetBasicInfo(){
	// Making GET request
	URL_struct,_ := url.Parse("https://mbasic.facebook.com/profile.php?v=info")
	response := u.GET(URL_struct)
	// Searching for user basic info -> {Name:,Birthday:,Gender:}
	bi := searchBasicInfo(response.Body)
	//Setting basic info for user
	u.Info.setInfo(bi)
}

func(u *UserRip) MakeReaction(Url *url.URL, reaction string){
	//Fixing Url & Making GET request in the publication link
	Url = fixUrl(Url)
	response := u.GET(Url)
	//Searching for Reaction Url (it contains specific Query Parameters) 
	tempUrl := searchReactionPickerUrl(response.Body)
	//Making GET request for the reaction selection link
	response = u.GET(tempUrl)
	//Searching for `ufi/reaction` (it contains specific Query Parameters) 
	tempUrl = searchUfiReactionUrl(response.Body,reaction)
	//Doing reaction
	u.GET(tempUrl)
}

//scrap Urls
func (u *UserRip) Scrap(Urls []*url.URL, folderPath string){
	folderPath = path.Clean(folderPath)
	for _,Url := range Urls{
		//Making filename string
		filename := fmt.Sprintf("%s_%s_%s.html",u.Parameters["email"],Url.Host,Url.Path)
		//fixing filename. Converting "/","." -> "-"
		var rs []rune
		for _,r := range filename{
			if string(r) == "/" || string(r) == "."{
				rs = append(rs,'-')
				continue
			}
			rs = append(rs,r)
		}
		//Making full path to file
		fullpath := path.Join(folderPath,string(rs))
		fmt.Println(fullpath)
		//Getting response from Url
		response := u.GET(Url)
		//Creating folder if it does not exist
		if _, err := os.Stat(folderPath); os.IsNotExist(err) {
			os.Mkdir(folderPath, os.ModeDir)
		}
		err := ioutil.WriteFile(fullpath,bodyToBytes(response.Body), 0755)
		if err != nil {
			fmt.Printf("Unable to write file: %v", err)
		}
	}
}
