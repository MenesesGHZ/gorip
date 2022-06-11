package scrap

import (
	"fmt"
	"github.com/menesesghz/gorip/fbrip"
)

func main() {
	// Reading Users and Action configuration from JSON
	users, actionConfig := fbrip.ReadRip("./rip.json")

	// Main Loop
	for _, user := range users {
		// Login sequence
		user.Sense()
		isLogged := user.Rip()
		if isLogged {
			user.Do(actionConfig)
			fmt.Printf("\n> User: %s | Gender:[ %s ] Birthday:[ %s ]\n", user.Info.Name, user.Info.Gender, user.Info.Birthday)
			fmt.Printf("* Actions Completed for -> %s\n\n", user.Email)
		}
	}
}
