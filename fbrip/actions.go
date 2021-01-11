package fbrip

import(
	"fmt"
	"net/url"
	"golang.org/x/net/html"
)

type ActionConfig struct{
	GetBasicInfo bool
	React react
	Post post
	Comment comment
}

// ACTIONS
func(u *UserRip) getBasicInfo(){
	// Making GET request
	URL_struct,_ := url.Parse("https://mbasic.facebook.com/profile.php?v=info")
	response := u.GET(URL_struct)

	//TEMPORAL IMPLEMENTATION. Just getting the name (for now)
	z := html.NewTokenizer(response.Body)
	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			break
		}
		if string(z.Raw()) == "<title>"{
			tt = z.Next()

			// TEMPORAL SOLUTION. It needs to be determine from the rip's phases whether the user could login or not.
			out := string(z.Raw())
			if out == "Page Not Found" || out == "Pagina No Encontrada"{
				panic("> FAILED AT LOGIN ;(")
				break
			}
			fmt.Println("> Welcome ->",out)
			u.Info.Name = out
			break
		}
	}
}

func(u *UserRip) makeReaction(URL *url.URL, reaction string){
	//Getting Query Parameters from `URL`	
	values := URL.Query()
	//Setting Query Parameters
	URL_struct,_ := URL.Parse("https://mbasic.facebook.com/ufi/reaction/")
	rValues := url.Values{}
	if URL.Path == "photo.php"{
		//Missing -> av, ext, hash
		rValues.Set("reaction_type", reaction)
		rValues.Set("basic_origin_uri",URL.String())
		rValues.Set("_ft_",values.Get("_ft_"))
		rValues.Set("ft_ent_identifier",values.Get("set")[3:])// to delete the extra `gm.`
	}
	//Adding Query Parameters to `URL_struct`
	URL_struct.RawQuery = rValues.Encode()
	fmt.Println("RawQuery",URL_struct.RawQuery)
	//Making GET request
	u.GET(URL_struct)	
}


//func makeReactionPhase1(){
//	//
//}
