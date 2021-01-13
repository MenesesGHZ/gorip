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
		fmt.Println("* Sense Completed.")
		// POST for get cookie-credentials
		fmt.Println("Ripping...")
		isLogged := u.Rip()
		if isLogged {
			fmt.Println("* Done Ripping.")
			// Doing task(s) base in ActionConfig. The action configuration is global for all users. FUTURE WORK: Make independent user config from rip.json
			u.Do(actionConfig)
			fmt.Printf("\n> User: %s | Gender:[ %s ] Birthday:[ %s ]\n",u.Info.Name,u.Info.Gender,u.Info.Birthday)
			fmt.Printf("* Actions Completed for -> %s\n\n",u.Parameters["email"])
		}
	}
}

