package fbrip

import(
	"net/url"
	"io/ioutil"
	"encoding/json"
)

func ReadRip(path string) ([]*UserRip, *ActionConfig){
	content,err := ioutil.ReadFile(path)
	if err!=nil{
		panic(err)
	}
	//Decoding JSON file
	var i interface{}
	err = json.Unmarshal(content,&i)
	if err!=nil{
		panic(err)
	}
	//Reading contentn interface
	m := i.(map[string]interface{})
	us := readUsers(m["Users"])
	ac := readActionConfig(m["ActionConfig"])
	return us,ac
}

func readUsers(users interface{}) []*UserRip{
	var us []*UserRip
	for _,user := range users.([]interface{}){
		us = append(us,readUser(user))
	}
	return us
}

func readUser(user interface{}) *UserRip{
	u := user.(map[string]interface{})
	p := readUserParameters(u["Parameters"])
	//c := readUserCookies(u["Cookies"])
	//i := readUserInfo(u["Info"])
	return &UserRip{
		Parameters:p,
	}
}

func readUserParameters(pI interface{}) map[string]string {
	p := make(map[string]string)
	for k, v := range pI.(map[string]interface{}){
		switch vv := v.(type) {
		case string:
			p[k] = vv
		default:
			panic("Error while reading `Users` in JSON")
		}
	}
	return p
}

//func readUserCookies(cI interface{}) []*http.Cookie{
//	//To Do
//}
//
//func readUserInfo(iI interface{}) *info{
//	//To Do
//}

func readActionConfig(acI interface{}) *ActionConfig{
	ac := acI.(map[string]interface{})
	r := readActionConfigReact(ac["React"])
//	p := readActionConfigPost(ac["Post"])
//	c := readActionConfigComment(ac["Comment"])
	s := readActionConfigScrap(ac["Scrap"])
	return &ActionConfig{
		GetBasicInfo:true, // For now default is `true` in order to get the cookies. Need to parse cookies from string properly in order to change this. 
		React:r,
//		Post:p,
//		Comment:c,
		Scrap:s,
	}
}

func readActionConfigPost(pI interface{}) post{
	p := make(map[string]string)
	pI = pI.(map[string]interface{})
	for k, v := range pI.(map[string]interface{}){
		switch vv := v.(type) {
		case string:
			p[k] = vv
		default:
			panic("Error while reading `ActionConfig > Post` in JSON")
		}
	}
	url,_ := url.Parse(p["Url"])
	return post{
		Url:url,
		Content:p["Content"],
	}
}

func readActionConfigReact(rI interface{}) react{
	r  := make(map[string]string)
	rI = rI.(map[string]interface{})
	for k, v := range rI.(map[string]interface{}){
		switch vv := v.(type) {
		case string:
			r[k] = vv
		default:
			panic("Error while reading `ActionConfig > React` in JSON")
		}
	}
	url,_ := url.Parse(r["Url"])
	return react{
		Url:url,
		Id:r["Id"],
	}
}

func readActionConfigComment(cI interface{}) comment{
	c := make(map[string]string)
	cI = cI.(map[string]interface{})
	for k, v := range cI.(map[string]interface{}){
		switch vv := v.(type) {
		case string:
			c[k] = vv
		default:
			panic("Error while reading `ActionConfig > Comment` in JSON")
		}
	}
	url,_ := url.Parse(c["Url"])
	return comment{
		Url:url,
		Content:c["Content"],
	}
}

func readActionConfigScrap(sI interface{}) scrap{
	var s []*url.URL
	sI = sI.([]interface{})
	for _,v := range sI.([]interface{}){
		switch vv := v.(type) {
		case string:
			tempUrl,err := url.Parse(vv)
			if err!=nil{
				panic("Error while reading `ActionConfig > Scrap` in JSON")
			}
			s = append(s,tempUrl)
		default:
			panic("Error while reading `ActionConfig > Comment` in JSON")
		}
	}
	return scrap{
		Urls:s,
	}
}
