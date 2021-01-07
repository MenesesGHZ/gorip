package main

import(
	"fmt"
	"fbreach"
    	"net/url"
)


func main(){
	//Creating UserBreach
	userBreach := fbreach.CreateUser("ghzant.y@gmail.com","password(password)")
	fmt.Println("User:",userBreach)

	// GET for getting part of Cookies
	URL_struct,_ := url.Parse("https://mbasic.facebook.com/")
	userBreach.Sense(URL_struct)
	fmt.Println("User Sensed:",userBreach)

	
		//Defining URL
	URL_struct,_ = url.Parse(URL_struct.String()+"login/device-based/regular/login/")
	userBreach.Rip(URL_struct)
}

