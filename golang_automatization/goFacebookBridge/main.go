package main

import(
	"fbreach"
)


func main(){
	//Creating UserBreach
	userBreach := fbreach.CreateUser("ghzant.y@gmail.com","password(password)")

	// GET for getting part of Cookies
	userBreach.Sense()

	// POST for get cookie-credentials
	userBreach.Rip()
}

