package fbrip

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"path"
	"bytes"
)

type Action interface {
	React | Publicate | Comment | Scrap
}

type React struct { // ReactStruct
	Id string
	Post *Post
}

type Publicate struct { // PostStruct
	Url *url.URL
}

type Comment struct { // CommentStruct
	Content string
	Post *Post
}

type Post struct {
	Url *url.URL
}

type Scrap struct { // ScrapStruct
	Page *url.URL
	OutputFolderPath string
	NamePrefix string
}

func (u *UserRip) Do(action Action){
	fmt.Println(action)
}

//type ActionConfig struct {
//	GetBasicInfo bool
//	React        ReactStruct
//	Publicate         PostStruct
//	Comment      CommentStruct
//	Scrap        ScrapStruct
//}

type ActionConfig struct {
	GetBasicInfo bool
	Reactions []React
	Publications []Publicate
	Comments []Comment
	Scraps []Scrap
}

type UserInfo struct {
	Name     string
	Birthday string
	Gender   string
}

func (u *UserRip) GetBasicInfo() {
	profileUrl, _ := url.Parse("https://mbasic.facebook.com/profile.php?v=info")
	response := u.GetRequest(profileUrl)
	if response == nil {
		fmt.Printf("** Error while making GET request to: %s | %s\n", profileUrl.String(), u.Email)
		return
	}
	basicInfoMap := searchBasicInfo(response.Body)
	u.Info.setInfo(basicInfoMap)
}

//func (u *UserRip) DoReaction(Urls []*url.URL, reactions []string) {
//	for i, Url := range Urls {
//		//Fixing Url & Making GET request in the publication link
//		transformUrlToBasicFacebook(Url)
//		response := u.GetRequest(Url)
//		//Handling error. NEED TO BE IMPROVED
//		if response == nil {
//			fmt.Printf("** Error while making GET request to: %s | %s", Url.String(), u.Email)
//			return
//		}
//		//Searching for React Url (it contains specific Query Parameters)
//		tempUrl := searchReactionPickerUrl(response.Body)
//		//Making GET request for the reaction selection link
//		response = u.GetRequest(tempUrl)
//		//Handling error. NEED TO BE IMPROVED
//		if response == nil {
//			fmt.Printf("** Error while making GET request to: %s | %s\n", Url.String(), u.Email)
//			return
//		}
//		//Searching for `ufi/reaction` (it contains specific Query Parameters)
//		tempUrl = searchUfiReactionUrl(response.Body, reactions[i])
//		//Doing reaction
//		u.GetRequest(tempUrl)
//	}
//}

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
		response := u.GetRequest(Url)
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


func NewScrap(pageRawUrl string, outputFolderPath string, namePrefix string) *Scrap {
	var scrap *Scrap
	parsedUrl, err := url.Parse(pageRawUrl)
	if err != nil {
		panic("Error while parsing url")
	}
	scrap.Page = parsedUrl
	scrap.OutputFolderPath = outputFolderPath
	scrap.NamePrefix = namePrefix
	return scrap
}

func NewReaction(id string, post *Post) *React {
	var react *React
	react.Id = id
	react.Post = post
	return react
}

func NewPost(rawUrl string) *Post {
	var post *Post
	parsedUrl, err := url.Parse(rawUrl)
	if err != nil {
		panic("Error while parsing url")
	}
	post.Url = parsedUrl
	return post
}

func NewComment(content string, post *Post) *Comment{
	var comment *Comment
	comment.Content = content
	comment.Post = post
	return comment
}

func (i *UserInfo) setInfo(basicInfo map[string]string) {
	i.Name = basicInfo["Name"]
	i.Birthday = basicInfo["Birthday"]
	i.Gender = basicInfo["Gender"]
}
