package main

import (
	"fmt"
	"github.com/menesesghz/gorip/fbrip"
)

func main() {
	user := fbrip.NewUserRip("mock@email.com", "superSecretPassword")
	isLogged := user.Rip()
	if isLogged {
		scrap := fbrip.NewScrap(
			"https://www.facebook.com/profile.php?id=100008137277101",
			"./scraps/",
			"some-name",
		)
		success := fbrip.Do(user, scrap)
		if success {
			fmt.Println("You have reacted 'WOW :O' to a Rick and Morty image !")
		} else {
			fmt.Println("You haven't reacted ;(")
		}
	}
	user.GetBasicInfo()
	fmt.Println(user)
}
