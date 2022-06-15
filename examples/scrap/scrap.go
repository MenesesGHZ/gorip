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
			"some-name-facebook-original",
			false,
		)
		success := fbrip.Do(user, scrap)
		if success {
			fmt.Println("You have scrapped the www.facebook profile of the developer !")
		} else {
			fmt.Println("Unable to scrap ;(")
		}

		scrap = fbrip.NewScrap(
			"https://www.facebook.com/profile.php?id=100008137277101",
			"./scraps/",
			"some-name-mbasic-facebook",
			true,
		)
		success = fbrip.Do(user, scrap)
		if success {
			fmt.Println("You have scrapped the mbasic.facebook profile of the developer !")
		} else {
			fmt.Println("Unable to scrap ;(")
		}
	}else{
		fmt.Println("Unable to login ;(")
	}
}
