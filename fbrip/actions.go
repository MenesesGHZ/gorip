package fbrip

import "net/url"

type ActionConfig struct{
	GetBasicInfo bool
	React react
	Post post
	Comment comment
}

// ACTIONS
func(u *UserRip) getBasicInfo(){
	// Making GET request
	URL_struct,_ := url.Parse("https://mbasic.facebook.com/profile.php?v=info")
	response := u.GET(URL_struct)
	// Searching for user basic info -> {Name:,Birthday:,Gender:}
	bi := searchBasicInfo(response.Body)
	//Setting basic info for user
	u.Info.setInfo(bi)
}

func(u *UserRip) makeReaction(URL *url.URL, reaction string){
	//Getting Query Parameters from `URL`	
	values := URL.Query()
	//Setting Query Parameters
	URL_struct,_ := url.Parse("https://mbasic.facebook.com/ufi/reaction/")
	rValues := url.Values{}
	if URL.Path == "photo.php"{
		//Missing -> av, ext, hash
		rValues.Set("reaction_type", reaction)
		rValues.Set("basic_origin_uri",URL.String())
		rValues.Set("_ft_",values.Get("_ft_"))
		rValues.Set("ft_ent_identifier",values.Get("set")[3:])// to delete the extra `gm.`
	}
	//Adding Query Parameters to `URL_struct`
	URL_struct.RawQuery = rValues.Encode()
	//Making GET request
	u.GET(URL_struct)
}


//func makeReactionPhase1(){
//	//
//}
