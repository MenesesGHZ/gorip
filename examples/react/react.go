package main

import (
	"fmt"
	"github.com/menesesghz/gorip/fbrip"
)

func main() {
	user := fbrip.NewUserRip("mock@email.com", "superSecretPassword")
	isLogged := user.Rip()
	if isLogged {
		react := fbrip.NewReaction("2", "https://www.facebook.com/RickandMorty/photos/pcb.5282285888534857/5282285578534888/")
		success := fbrip.Do(user, react)
		if success {
			fmt.Println("You have reacted '<3' to a Rick and Morty image !")
		} else {
			fmt.Println("You haven't reacted ;(")
		}
	} else {
		fmt.Println("Unable to login ;(")
	}
}
