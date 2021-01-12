package main

import(
	"fmt"
	"github.com/menesesghz/gorip/fbrip"
)

func main(){
	// Reading users and action config from JSON
	users,actionConfig := fbrip.ReadRip("./rip.json")
	// Starting sequence 
	for _,u := range users{
		// GET for getting part of Cookies
		u.Sense()
		fmt.Println("* Sense Complete.")
		// POST for get cookie-credentials
		fmt.Println("Ripping...")
		u.Rip()
		fmt.Println("* Done Ripping.")
		// Doing task base in the action config. The action configurations is global for all users.
		u.Do(actionConfig)
		fmt.Printf("\nUser: %s | Gender:[ %s ] Birthday:[ %s ]\n",u.Info.Name,u.Info.Gender,u.Info.Birthday)
		fmt.Println("* Actions Completed.")
	}
}

