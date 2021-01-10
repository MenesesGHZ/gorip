package main

import(
	"fbreach"
)


func main(){
	//Creating UserBreach
	userBreach := fbreach.CreateUser("gerry_csm@outlook.com","print(jerry2000)")

	// GET for getting part of Cookies
	userBreach.Sense()

	// POST for get cookie-credentials
	userBreach.Rip()
	
	// Prepare 
	config := fbreach.ActionConfig{
		GetBasicInfo:true,
		MakeReaction:false,
		MakePost:false,
	}
	content := fbreach.ActionContent{}
	
	// DO
	userBreach.Do(content,config)


}

