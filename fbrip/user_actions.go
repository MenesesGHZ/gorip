package fbrip

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"path"
)

type Action interface {
	execute(*UserRip) bool
	*React | *Publicate | *Comment | *Scrap
}

type React struct {
	Id   string
	Post *Post
}

type Publicate struct {
	Url *url.URL
}

type Comment struct {
	Content string
	Post    *Post
}

type Scrap struct {
	Page             *url.URL
	OutputFolderPath string
	NamePrefix       string
}

func (r *React) execute(user *UserRip) bool {
	transformUrlToBasicFacebook(r.Post.Url)
	response := user.GetRequest(r.Post.Url)
	if response == nil {
		return false
	}
	reactionsPickerUrl := getReactionsPickerUrl(response.Body)
	response.Body.Close()
	response = user.GetRequest(reactionsPickerUrl)
	if response == nil {
		return false
	}
	reactionUrl := getReactionUrl(response.Body, r.Id)
	response.Body.Close()
	user.GetRequest(reactionUrl)
	return true
}

func (r *Publicate) execute(user *UserRip) bool {
	fmt.Println("PUBLICATING...")
	return true
}

func (r *Comment) execute(user *UserRip) bool {
	fmt.Println("COMMENTING...")
	return true
}

func (r *Scrap) execute(user *UserRip) bool {
	fmt.Println("SCRAPPING...")
	return true
}

func Do[A Action](user *UserRip, action A) bool {
	return action.execute(user)
}

type ActionConfig struct {
	GetBasicInfo bool
	Reactions    []React
	Publications []Publicate
	Comments     []Comment
	Scraps       []Scrap
}

type Post struct {
	Url *url.URL
}

func (u *UserRip) Do(actionConfig ActionConfig) {
	if actionConfig.GetBasicInfo {
		u.GetBasicInfo()
	}
	for _, react := range actionConfig.Reactions {
		react.execute(u)
	}
	for _, publicate := range actionConfig.Publications {
		publicate.execute(u)
	}
	for _, comment := range actionConfig.Comments {
		comment.execute(u)
	}
	for _, scrap := range actionConfig.Scraps {
		scrap.execute(u)
	}
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
	} else {
		basicInfoMap := searchBasicInfo(response.Body)
		u.Info.setInfo(basicInfoMap)
	}
}

//func (u *UserRip) DoReaction(Urls []*url.URL, reactions []string) {
//	for i, Url := range Urls {
//Fixing Url & Making GET request in the publication link

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

func NewReaction(id string, postUrl string) *React {
	return &React{
		Id:   id,
		Post: newPost(postUrl),
	}
}

func newPost(rawUrl string) *Post {
	parsedUrl, err := url.Parse(rawUrl)
	if err != nil {
		panic("Error while parsing url")
	}
	return &Post{
		Url: parsedUrl,
	}
}

func NewComment(content string, postUrl string) *Comment {
	var comment *Comment
	comment.Content = content
	comment.Post = newPost(postUrl)
	return comment
}

func (i *UserInfo) setInfo(basicInfo map[string]string) {
	i.Name = basicInfo["Name"]
	i.Birthday = basicInfo["Birthday"]
	i.Gender = basicInfo["Gender"]
}
