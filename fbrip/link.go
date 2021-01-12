package fbrip

import(
	"io"
	"fmt"
	"bytes"
	"net/url"
	"net/http"
)

//Ensuring that URL follows to `mbasic.facebook.com`
func fixUrl(Url *url.URL) *url.URL{
	Url.Scheme = "https"
	Url.Host = "mbasic.facebook.com"
	return Url
}

//It checks if `cookie` is already in `slice of cookies`
func includesCookie(cookies []*http.Cookie, cookie *http.Cookie) bool{
	for _,c := range cookies{
		if c.Name == cookie.Name{
			return true
		}
	}
	return false
}

//Show Body Content in the Console
func showBody(body io.Reader){
	buf := new(bytes.Buffer)
	buf.ReadFrom(body)
	fmt.Println(buf.String())
}
