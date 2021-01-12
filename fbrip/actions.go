package fbrip

import (
	"net/url"
)

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

func(u *UserRip) makeReaction(Url *url.URL, reaction string){
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
