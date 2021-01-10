package main

import(
	"github.com/menesesghz/gofbreach"
)

func main(){
	//Creating UserBreach
	userBreach := gofbreach.CreateUser("gerry_csm@outlook.com","print(jerry2000)")

	// GET for getting part of Cookies
	userBreach.Sense()

	// POST for get cookie-credentials
	userBreach.Rip()
	
	// Prepare Action
	config := gofbreach.ActionConfig{
		GetBasicInfo:true,
		MakeReaction:false,
		MakePost:false,
	}
	content := gofbreach.ActionContent{}
	
	// Do action
	userBreach.Do(content,config)


}

