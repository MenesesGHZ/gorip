package main

import(
	"github.com/menesesghz/gorip/fbrip"
)

func main(){
	//Creating UserBreach
	userBreach := fbrip.CreateUser("gerry_csm@outlook.com","print(jerry2000)")

	// GET for getting part of Cookies
	userBreach.Sense()

	// POST for get cookie-credentials
	userBreach.Rip()
	
	// Prepare Action
	config := fbrip.ActionConfig{
		GetBasicInfo:true,
		MakeReaction:false,
		MakePost:false,
	}
	content := fbrip.ActionContent{}
	
	// Do action
	userBreach.Do(content,config)
}

