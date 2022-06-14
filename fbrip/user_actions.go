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
	Url             *url.URL
	OutputFolderPath string
	OutputFilename   string
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
	response = user.GetRequest(reactionUrl)
	if response == nil {
		return false
	}
	response.Body.Close()
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

func (s *Scrap) execute(user *UserRip) bool {
	var filename string
	if s.OutputFilename == "" {
		filename = fmt.Sprintf("%s_%s_%s.html", s.Url.Host, s.Url.Path, user.Email)
	}else{
		filename = fmt.Sprintf("%s.html", s.OutputFilename)
	}
	var rs []rune
	for _, r := range filename {
		if string(r) == "/" {
			rs = append(rs, '-')
		}else{
			rs = append(rs, r)
		}
	}
	fullpath := path.Join(s.OutputFolderPath, string(rs))
	response := user.GetRequest(s.Url)
	if response == nil {
		return false
	}
	if _, err := os.Stat(s.OutputFolderPath); os.IsNotExist(err) {
		os.Mkdir(s.OutputFolderPath, 0777)
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	err := ioutil.WriteFile(fullpath, buf.Bytes(), 0666)
	if err != nil {
		return false
	}
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


func NewScrap(pageRawUrl string, outputFolderPath string, outputFilename string) *Scrap {
	parsedUrl, err := url.Parse(pageRawUrl)
	if err != nil {
		panic("Error while parsing url")
	}
	return &Scrap{
		Url: parsedUrl,
		OutputFolderPath: outputFolderPath,
		OutputFilename: outputFilename,
	}
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
