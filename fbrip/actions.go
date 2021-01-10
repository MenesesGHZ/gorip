package fbrip

import(
	"fmt"
	"net/url"
	"golang.org/x/net/html"
)

type ActionConfig struct {
	GetBasicInfo bool
	MakeReaction bool
	MakePost bool
}

type ActionContent struct {
	url *url.URL
	reaction int
	comment string
}

//
// ACTIONS
//
func(u *UserBreach) getBasicInfo() bool{

	// Making GET request
	URL_struct,_ := url.Parse("https://mbasic.facebook.com/profile.php?v=info")
	response := u.GET(URL_struct)

	// Just getting the name (for now)
	z := html.NewTokenizer(response.Body)
	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			// ...
			return false
		}
		if string(z.Raw()) == "<title>"{
			tt = z.Next()
			fmt.Println("> Welcome ->",z.Token())	
			u.name = string(z.Raw())
		}
	}
	return true
}

func(u *UserBreach) makeReaction(url *url.URL, reaction int) bool{
	success:=false
	return success
}


